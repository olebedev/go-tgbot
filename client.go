package tgbot

import (
	"context"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/olebedev/go-tgbot/client"
)

func NewClient(ctx context.Context, token string) *client.TelegramBot {
	transport := httptransport.New("api.telegram.org", "/", []string{"https"})
	transport.Context = ctx
	transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetPathParam("token", token)
		return nil
	})
	return client.New(transport, nil)
}
