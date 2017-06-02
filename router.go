package tgbot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/go-tgbot/client"
	"github.com/olebedev/go-tgbot/client/updates"
	"github.com/olebedev/go-tgbot/models"
	"github.com/pkg/errors"
)

//go:generate stringer -type=rkind
type rkind int

const (
	// Kind is const to route different updates
	KindMessage rkind = 1 << iota
	KindEditedMessage
	KindChannelPost
	KindEditedChannelPost
	KindInlineQuery
	KindCallbackQuery
	KindChosenInlineResult
	KindShippingQuery
	KindPreCheckoutQuery
	KindAll rkind = (1 << iota) - 1

	DefaultPollTimeout int64 = 29
)

// Options ... All field are optional.
type Options struct {
	Context    context.Context
	Token      string
	GetSession func(*models.Update) (fmt.Stringer, error)
}

// Router ...
type Router struct {
	*client.TelegramBot
	middlewares []func(*Context) error
	routes      []struct {
		kind  rkind
		key   string
		re    *regexp.Regexp
		value []func(*Context) error
	}
	getSession func(*models.Update) (fmt.Stringer, error)
}

// New returns a router.
func New(o *Options) *Router {
	if o == nil {
		o = &Options{}
	}
	r := &Router{
		TelegramBot: NewClient(o.Context, o.Token),
		middlewares: make([]func(*Context) error, 0),
		routes: make([]struct {
			kind  rkind
			key   string
			re    *regexp.Regexp
			value []func(*Context) error
		}, 0),
	}

	if o.GetSession != nil {
		r.getSession = o.GetSession
	} else {
		r.getSession = GetSession
	}
	return r
}

// Use appends given middlewares into call chain.
func (r *Router) Use(middlewares ...func(*Context) error) {
	r.middlewares = append(r.middlewares, middlewares...)
}

// All binds all kind of updates and the route through handler/handlers.
func (r *Router) All(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindAll, route, handlers...)
}

// Message binds message updates and the route through handler/handlers.
func (r *Router) Message(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindMessage, route, handlers...)
}

// EditedMessage binss message updates and the route through handler/handlers.
func (r *Router) EditedMessage(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindEditedMessage, route, handlers...)
}

// ChannelPost binds message updates and the route through handler/handlers.
func (r *Router) ChannelPost(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindChannelPost, route, handlers...)
}

// EditedChannelPost binds message updates and the route through handler/handlers.
func (r *Router) EditedChannelPost(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindEditedChannelPost, route, handlers...)
}

// InlineQuery binds inline queries and the route through handler/handlers.
func (r *Router) InlineQuery(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindInlineQuery, route, handlers...)
}

// CallbackQuery binds callback updates and the route through handler/handlers.
func (r *Router) CallbackQuery(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindCallbackQuery, route, handlers...)
}

// ChosenInlineResult binds chosen inline result updates and the route through
// handler/handlers.
func (r *Router) ChosenInlineResult(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindChosenInlineResult, route, handlers...)
}

// ShippingQuery binds shipping result updates and the route through
// handler/handlers.
func (r *Router) ShippingQuery(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindShippingQuery, route, handlers...)
}

// PreCheckout binds pre-checkout result updates and the route through
// handler/handlers.
func (r *Router) PreCheckout(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindPreCheckoutQuery, route, handlers...)
}

// Handle binds specific kinds of an update to the handler/handlers.
func (r *Router) Handle(kind rkind, routeRegExp string, handlers ...func(*Context) error) error {
	re, err := regexp.Compile(routeRegExp)
	if err != nil {
		return err
	}

	r.routes = append(r.routes, struct {
		kind  rkind
		key   string
		re    *regexp.Regexp
		value []func(*Context) error
	}{
		kind:  kind,
		key:   routeRegExp,
		re:    re,
		value: handlers,
	})

	return nil
}

// Route routes given updates into the bound handlers
// through defined middlewares.
func (r *Router) Route(u *models.Update) error {

	sess, err := r.getSession(u)
	if err != nil {
		return errors.Wrap(err, "getting session")
	}

	var kind rkind
	switch true {
	case u.Message != nil:
		kind = KindMessage
	case u.EditedMessage != nil:
		kind = KindEditedMessage
	case u.ChannelPost != nil:
		kind = KindChannelPost
	case u.EditedChannelPost != nil:
		kind = KindEditedChannelPost
	case u.InlineQuery != nil:
		kind = KindInlineQuery
	case u.CallbackQuery != nil:
		kind = KindCallbackQuery
	case u.ChosenInlineResult != nil:
		kind = KindChosenInlineResult
	case u.ShippingQuery != nil:
		kind = KindShippingQuery
	case u.PreCheckoutQuery != nil:
		kind = KindPreCheckoutQuery
	default:
		return errors.New("unknown update")
	}

	text := sess.String()

	for _, ro := range r.routes {
		if (ro.kind|kind) == ro.kind && ro.re.MatchString(text) {
			ctx := &Context{
				Session:  sess,
				Update:   u,
				Params:   ro.re.FindStringSubmatch(text)[1:],
				Kind:     kind,
				handlers: make([]func(*Context) error, len(r.middlewares)+len(ro.value)),
			}

			copy(ctx.handlers, r.middlewares)
			copy(ctx.handlers[len(r.middlewares):], ro.value)

			err := ctx.Next()
			if err != nil {
				return errors.Wrap(err, "ctx.Next()")
			}

			if ctx.isSkipped {
				ctx.isSkipped = false
				continue
			} else {
				break
			}
		}
	}
	return nil
}

// Poll does a polling of API endpoints and routes consumed updates. It returns an error
// if any of handlers return the error.
func (r *Router) Poll(ctx context.Context, allowed []models.AllowedUpdate) error {
	var a []string
	for _, item := range allowed {
		a = append(a, string(item))
	}

	p := updates.NewGetUpdatesParams().
		WithContext(ctx).
		WithTimeout(pointer.ToInt64(DefaultPollTimeout))

	if len(a) > 0 {
		p.SetAllowedUpdates(a)
	}

	for {
		u, err := r.Updates.GetUpdates(p)
		if err != nil {
			return errors.Wrap(err, "polling updates")
		}
		for _, update := range u.Payload.Result {
			p.SetOffset(pointer.ToInt64(update.UpdateID + 1))
			err = r.Route(update)
			if err != nil {
				return errors.Wrap(err, "route update")
			}
		}
	}

	return nil
}

// ServeHTTP implements http.ServeHTTP interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var u *models.Update
	if err := json.NewDecoder(req.Body).Decode(u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	defer req.Body.Close()

	if err := r.Route(u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
