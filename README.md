# go-tgbot [![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/olebedev/go-tgbot)

> Pure Golang telegram bot API wrapper generated from swagger definition, session-based routing and middlewares.

### Usage benefits

1. **No need to learn any other library API.** You will use methods with payload exactly like it presented on telegram bot API description page. With only couple trade-offs, b/c of telegram bot API is generics a bit.
2. **All models and methods are being supported.** The models and methods were generated from `swagger.yaml` description file. So, new entities/methods could be added by describing in the YAML swagger file. This approach allows validating the description, avoid typos and develop fast.
3. `ffjson` is plugged. So, **it's pretty fast**.
4. **`context.Context` based** HTTP client
5. **Session-based routing**, not only message text based.

### Client

Client package could be used as regular `go-swagger` client library without using Router. There is the only single additional feature over `go-swagger` - a possibility to setup token by default. It solved as a [custom HTTP transport](https://github.com/olebedev/go-tgbot/blob/master/client.go#L13-L20). Example:

```go
package main

import (
	"log"
	"context"
	tgbot "github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/models"
	"github.com/olebedev/go-tgbot/client/users"
)

func main() {
	token := flag.String("token", "", "telegram bot token")
	flag.Parse()
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	api := tgbot.NewClient(ctx, *token)
	
	log.Println(api.Users.GetMe(nil))
	
	// Also, every calls could be done with given token and/or context
	ctx, cancel = context.WithTimeout(ctx, 10 * time.Second)
	defer cancel()
	
	r, err := api.Users.GetMe(
		users.NewGetMeParams().
			WithContext(ctx).
			WithToken("<overwrite default token>"),
	)
	
	if err != nil && {
		if e, ok := err.(*models.Error); ok {
			log.Println(e.Payload. ErrorCode, e.Payload.Description)
		}
	}
}
```

Since swagger covers many other platforms/technologies the same libraries could be generated for them too. See the source here - [`swagger.yaml`](https://github.com/olebedev/go-tgbot/blob/master/swagger.yaml).

### Router

The Router allows binding between kinds of updates and handlers, which are being checked via regexp. router include client API library as embedded struct. Example:

```go
package main

import (
	"log"
	"context"
	tgbot "github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/models"
	"github.com/olebedev/go-tgbot/client/messages"
)

func main() {
	token := flag.String("token", "", "telegram bot token")
	flag.Parse()
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	r := tgbot.New(&tgbot.Options{
		Context: ctx,
		Token: *token,
	})
	
	// setup global middleware
	r.Use(tgbot.Recover)
	
	// Bind handler
	r.Message("^/start\\sstart$", func(c *tgbot.Context) error {
		log.Println(c.Update.Message.Text)
		// send greeting message back
		message := "hi there what's up"
		resp, err := r.Messages.SendMessage(
			messages.NewSendMessageParams().WithBody(messages.SendMessageBody{
				Text:   &message,
				ChatID: c.Update.Message.Chat.ID,
			}),
		)
		if err != nil {
			retuirn err
		}
		if resp != nil {
			log.Println(resp.Payload.MessageID)
		}
		return nil
	})
	
	if err := r.Poll(ctx, []string{"message"}); err != nil {
		log.Fatal(err)
	}
}
```

Default string representation of any kind of an update could be found here - [`session.go`](https://github.com/olebedev/go-tgbot/blob/master/session.go#L19-L43).

Router implements `http.Handler` interface to be able to serve HTTP as well. But, it's not recommended because webhooks are much much slower than polling.

### Session-based routing

After many bots were developed, the one principal thing could be marked - routing need to be matched with session instead of received message text. So,  we should be able to wrap the update into a string representation, to be matched with a handler. For this purpose, Router accepts a [`GetSession `](https://github.com/olebedev/go-tgbot/blob/master/router.go#L37) optional argument. It's a function which returns `(fmt.Stringer, error)`, the `fmt.Stringer` instance will be placed as [`c.Session`](https://github.com/olebedev/go-tgbot/blob/master/context.go#L13) during handlers chain call. Example:

```go
package main

import (
	"log"
	"fmt"
	"context"
	tgbot "github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/models"
	"github.com/olebedev/go-tgbot/client/messages"
	
	"app"
)

type Session struct {
	User    *app.User
	Billing *app.UserBillingInfo
	// ... etc
	Route string
}

func (s Session) String() string {
	return s.Route
}

func GetSessionFunc(u *models.Update) (fmt.Stringer, error) {
	sess, err := tgbot.GetSession(u)
	if err != nil {
		return nil, err
	}
	
	s := &Session{
		Route: "~" + sess.String(),
	}
	
	switch true {
	case u.Message != nil:
		s.User, err := app.GetUserByTgID(u.Message.From.ID)
		if err != nil {
			return err
		}
		s.Billing, err := app.GetBillingByID(s.User.ID)
		if err != nil {
			return err
		}
		
		if !s.Billing.Active {
			s.Route = "/pay" + s.Route
		}
	}
	
	return s, nil
}

func main() {
	token := flag.String("token", "", "telegram bot token")
	flag.Parse()
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	r := tgbot.New(&tgbot.Options{
		Context: ctx,
		Token: *token,
		GetSession: GetSessionFunc,
	})
	
	// setup global middleware
	r.Use(tgbot.Recover)
	
	// Bind handlers
	r.Message("^/pay~.*", func(c *tgbot.Context) error {
		s := c.Session.(*Session)
		// TODO handle payment here before say hello
		return r.Route(c.Update)
	})
	
	r.Message("^~/start\\sstart$", func(c *tgbot.Context) error {
		log.Println(c.Update.Message.Text)
		// send greeting message back
		message := "hi there what's up"
		resp, err := r.Messages.SendMessage(
			messages.NewSendMessageParams().WithBody(messages.SendMessageBody{
				Text:   &message,
				ChatID: c.Update.Message.Chat.ID,
			}),
		)
		if err != nil {
			retuirn err
		}
		if resp != nil {
			log.Println(resp.Payload.MessageID)
		}
		return nil
	})
	
	if err := r.Poll(ctx, []string{"message"}); err != nil {
		log.Fatal(err)
	}
}
```

---

See [the documentation](https://godoc.org/github.com/olebedev/go-tgbot) for more details.

### LICENSE

http://www.apache.org/licenses/LICENSE-2.0