package tgbot_test

import (
	"context"
	"flag"
	"log"

	tgbot "github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/olebedev/go-tgbot/models"
)

var token *string

func ExampleNew() {
	token = flag.String("token", "", "telegram bot token")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := tgbot.New(ctx, *token)

	// setup global middleware
	r.Use(tgbot.Recover)

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

	if err := r.Poll(ctx, []models.AllowedUpdate{models.AllowedUpdateMessage}); err != nil {
		log.Fatal(err)
	}
}
