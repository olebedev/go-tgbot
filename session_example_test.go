package tgbot_test

import (
	"context"
	"flag"
	"fmt"
	"log"

	tgbot "github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/olebedev/go-tgbot/models"

	"app" // you application
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

	s.User, err = app.GetUserByTgID(u.Message.From.ID)
	if err != nil {
		return err
	}
	s.Billing, err = app.GetBillingByID(s.User.ID)
	if err != nil {
		return err
	}

	if !s.Billing.Active {
		s.Route = "/pay" + s.Route
	}

	return s, nil
}

func ExampleNew_session() {
	token := flag.String("token", "", "telegram bot token")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := tgbot.New(&tgbot.Options{
		Context:    ctx,
		Token:      *token,
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
			messages.NewSendMessageParams().WithBody(&models.SendMessageBody{
				Text:   &message,
				ChatID: c.Update.Message.Chat.ID,
			}),
		)
		if err != nil {
			return err
		}
		if resp != nil {
			log.Println(resp.Payload.MessageID)
		}
		return nil
	})

	if err := r.Poll(ctx, []models.AllowedUpdate{models.AllowedUpdateMessage}); err != nil {
		log.Fatal(err)
	}
}
