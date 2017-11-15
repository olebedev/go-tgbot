package tgbot

import (
	"context"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/olebedev/go-tgbot/client"
	"golang.org/x/time/rate"
)

// Add wrapper for API throttling
type wrapper struct {
	*rate.Limiter
	*httptransport.Runtime
	context.Context
}

// Submit wraps httpclient.Submit for throttling
func (w *wrapper) Submit(op *runtime.ClientOperation) (interface{}, error) {
	if op.Context != nil {
		w.Wait(op.Context)
	} else if w.Context != nil {
		w.Wait(w.Context)
	} else {
		w.Wait(context.Background())
	}

	return w.Runtime.Submit(op)
}

// NewClient returns new plain API client
func NewClient(ctx context.Context, token string) *client.TelegramBot {
	transport := httptransport.New("api.telegram.org", "/", []string{"https"})
	if ctx == nil {
		ctx = context.Background()
	}
	transport.Context = ctx
	transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetPathParam("token", token)
	})

	return client.New(&wrapper{
		Limiter: rate.NewLimiter(30, 1),
		Runtime: transport,
		Context: ctx,
	}, nil)
}
