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
	if ctx == nil {
		ctx = context.Background()
	}
	transport.Context = ctx
	transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetPathParam("token", token)
	})
	return client.New(transport, nil)
}
