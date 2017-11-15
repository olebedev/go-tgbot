package tgbot_test

import (
	"context"
	"flag"
	"log"

	tgbot "github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/olebedev/go-tgbot/models"

	"app" // you application
)

type Session struct {
	User    *app.User
	Billing *app.UserBillingInfo
}

func ExampleNew_session() {
	token := flag.String("token", "", "telegram bot token")
	flag.Parse()

	r := tgbot.New(context.Background(), *token)

	// setup global middleware
	r.Use(tgbot.Recover)

	r.Use(func(c *tgbot.Context) error {
		s, err := app.GetSession(c.From)
		if err != nil {
			return nil, err
		}
		c.Set("session", s)

		if !s.Billing.Active {
			c.Path = "/pay_first~" + c.Path
		}

		err = c.Next()
		if err != nil {
			return err
		}
		// finalize session if needed
	})

	// Bind handlers

	// middleware for users who have not paid
	r.Bind(`^/pay_first~.*`, func(c *tgbot.Context) error {
		s := c.MustGet("session").(*Session)
		// business logic
		return r.Route(c.Update)
	})

	r.Bind(`^/message/-/private/text/start\sstart$`, func(c *tgbot.Context) error {
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
