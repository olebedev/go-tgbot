package tgbot_test

import (
	"context"
	"flag"
	"log"
	"time"

	tgbot "github.com/olebedev/go-tgbot"
	"github.com/olebedev/go-tgbot/client/users"
)

var token *string

func ExampleNewClient() {
	token = flag.String("token", "", "telegram bot token")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	api := tgbot.NewClient(ctx, *token)

	log.Println(api.Users.GetMe(nil))

	// Also, every calls could be done with given token and/or context
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	t := "<overwrite default token>"
	_, err := api.Users.GetMe(
		users.NewGetMeParams().
			WithContext(ctx).
			WithToken(&t),
	)

	if err != nil {
		// users.NewGetMeBadRequest()
		if e, ok := err.(*users.GetMeBadRequest); ok {
			log.Println(e.Payload.ErrorCode, e.Payload.Description)
		}
	}
}
