# go-tgbot [![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/olebedev/go-tgbot)

> Pure Golang telegram bot API wrapper generated from swagger definition, session-based routing and middlewares.

### Usage benefits

1. **No need to learn any other library API.** You will use methods with payload exactly like it presented on telegram bot API description page. With only couple trade-offs, b/c of telegram bot API is generics a bit.
2. **All models and methods are being supported.** The models and methods were generated from `swagger.yaml` description file. So, new entities/methods could be added by describing in the YAML swagger file. This approach allows validating the description, avoid typos and develop fast.
3. `easyjson` is plugged. So, **it's fast**.
4. **`context.Context` based** HTTP client
5. **Session-based routing**, not only message text based.

### Client

Client package could be used as regular `go-swagger` client library without using Router. There are the only two additional features over `go-swagger`, a possibility to setup token by default(It solved as an [interface checking](https://github.com/olebedev/go-tgbot/blob/master/client.go#L36)) and [API throttling](https://github.com/olebedev/go-tgbot/blob/master/client.go#L26) for 30 calls per second([see more info](https://core.telegram.org/bots/faq#my-bot-is-hitting-limits-how-do-i-avoid-this)). 

Example:

```go
package main

import (
	"context"
	"flag"
	"log"
	"time"

	tgbot "github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/client/users"
)

var token *string

func main() {
	token = flag.String("token", "", "telegram bot token")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	api := tgbot.NewClient(ctx, *token)

	log.Println(api.Users.GetMe(nil))

	// Also, every calls could be done with given context
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err := api.Users.GetMe(
		users.NewGetMeParams().
			WithContext(ctx),
	)

	if err != nil {
		// users.NewGetMeBadRequest()
		if e, ok := err.(*users.GetMeBadRequest); ok {
			log.Println(e.Payload.ErrorCode, e.Payload.Description)
		}
	}
}
```

Since swagger covers many other platforms/technologies the same libraries could be generated for them too. See the source here - [`swagger.yaml`](https://github.com/olebedev/go-tgbot/blob/master/swagger.yaml).  
Also, have a look at [Swagger UI for telegram API](http://petstore.swagger.io/?url=https://raw.githubusercontent.com/olebedev/go-tgbot/master/swagger.yaml#/)(`master` branch).

### Router

The Router allows binding between kinds of updates and handlers, which are being checked via regexp. Router includes client API library as embedded struct. Example:

```go
package main

import (
	"context"
	"flag"
	"log"
	"os"

	tgbot "github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/olebedev/go-tgbot/models"
)

var token *string

func main() {
	token = flag.String("token", "", "telegram bot token")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := tgbot.New(ctx, *token)

	// setup global middleware
	r.Use(tgbot.Recover)
	r.Use(tgbot.Logger(os.Stdout))

	// modify path to be able to match user's commands via router
	r.Use(func(c *tgbot.Context) error {
		c.Path = c.Path + c.Text
		return nil
	})

	// bind handler
	r.Bind(`^/message/(?:.*)/text/start(?:\s(.*))?$`, func(c *tgbot.Context) error {
		log.Println(c.Capture)             // - ^ from path
		log.Println(c.Update.Message.Text) // or c.Text

		// send greeting message back
		message := "hi there what's up"
		resp, err := r.Messages.SendMessage(
			messages.NewSendMessageParams().WithBody(&models.SendMessageBody{
				Text:   &message,
				ChatID: c.Update.Message.Chat.ID,
			}),
		)
		if err != nil {
			return err
		}
		if resp != nil {
			log.Println(resp.Payload.Result.MessageID)
		}
		return nil
	})

	if err := r.Poll(ctx, models.AllowedUpdateMessage); err != nil {
		log.Fatal(err)
	}
}
```

Default string representation of any kind of an update could be found here - [`router.go`](https://github.com/olebedev/go-tgbot/blob/master/router.go#L96-L226).

Router implements `http.Handler` interface to be able to serve HTTP as well. But, it's not recommended because webhooks are much much slower than polling.

More examples can be found at [godoc](https://godoc.org/github.com/olebedev/go-tgbot).

---

See [the documentation](https://godoc.org/github.com/olebedev/go-tgbot) for more details.

### LICENSE

http://www.apache.org/licenses/LICENSE-2.0
