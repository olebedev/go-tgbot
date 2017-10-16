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

	// transport.Consumers[runtime.JSONMime] = runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
	// 	dec := json.NewDecoder(reader)
	// 	dec.UseNumber() // preserve number formats
	// 	return dec.Decode(data)
	// })
	//
	// transport.Producers[runtime.JSONMime] = runtime.ProducerFunc(func(writer io.Writer, data interface{}) error {
	// 	enc := json.NewEncoder(writer)
	// 	return enc.Encode(data)
	// })

	return client.New(transport, nil)
}
