package tgbot

import (
	"context"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/olebedev/go-tgbot/client"
	"golang.org/x/time/rate"
)

// Add wrapper for API throttling
type wrapper struct {
	token string
	*rate.Limiter
	*httptransport.Runtime
	context.Context
}

type ist interface {
	SetToken(*string)
}

// Submit wraps httpclient.Submit for throttling
func (w *wrapper) Submit(op *runtime.ClientOperation) (interface{}, error) {
	if op.Context != nil {
		w.Wait(op.Context)
	} else if w.Context != nil {
		w.Wait(w.Context)
	} else if w.Runtime.Context != nil {
		w.Wait(w.Runtime.Context)
	} else {
		w.Wait(context.Background())
	}

	if w.token != "" {
		if st, ok := op.Params.(ist); ok {
			st.SetToken(&w.token)
		}
	}

	return w.Runtime.Submit(op)
}

// NewClient returns new plain API client
func NewClient(ctx context.Context, token string) *client.TelegramBot {
	transport := httptransport.New(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes)
	if ctx == nil {
		ctx = context.Background()
	}
	transport.Context = ctx

	return client.New(&wrapper{
		token:   token,
		Limiter: rate.NewLimiter(30, 1),
		Runtime: transport,
		Context: ctx,
	}, nil)
}
