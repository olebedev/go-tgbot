package tgbot

import (
	"context"
	"fmt"
	"regexp"

	"github.com/olebedev/go-tgbot/client"
	"github.com/olebedev/go-tgbot/models"
	"github.com/pkg/errors"
)

type rkind int

const (
	KindMessage rkind = 1 << iota
	KindEditedMessage
	KindChannelPost
	KindEditedChannelPost
	KindInlineQuery
	KindCallbackQuery
	KindChosenInlineResult
	KindAll

	allKind = KindMessage |
		KindEditedMessage |
		KindChannelPost |
		KindEditedChannelPost |
		KindInlineQuery |
		KindCallbackQuery |
		KindChosenInlineResult
)

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

func New(ctx context.Context, token string, repo func(*models.Update) (fmt.Stringer, error)) *Router {
	r := &Router{
		TelegramBot: NewClient(ctx, token),
		middlewares: make([]func(*Context) error, 0),
		routes: make([]struct {
			kind  rkind
			key   string
			re    *regexp.Regexp
			value []func(*Context) error
		}, 0),
	}

	if repo != nil {
		r.getSession = repo
	} else {
		r.getSession = GetSession
	}
	return r
}

// Use appends given middlewares into call chain.
func (r *Router) Use(middlewares ...func(*Context) error) {
	r.middlewares = append(r.middlewares, middlewares...)
}

// All bind all kind of updates and the route through handler/handlers.
func (r *Router) All(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindAll, route, handlers...)
}

// Message bind message updates and the route through handler/handlers.
func (r *Router) Message(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindMessage, route, handlers...)
}

// EditedMessage bind message updates and the route through handler/handlers.
func (r *Router) EditedMessage(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindEditedMessage, route, handlers...)
}

// ChannelPost bind message updates and the route through handler/handlers.
func (r *Router) ChannelPost(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindChannelPost, route, handlers...)
}

// EditedChannelPost bind message updates and the route through handler/handlers.
func (r *Router) EditedChannelPost(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindEditedChannelPost, route, handlers...)
}

// InlineQuery bind inline queries and the route through handler/handlers.
func (r *Router) InlineQuery(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindInlineQuery, route, handlers...)
}

// CallbackQuery bind callback updates and the route through handler/handlers.
func (r *Router) CallbackQuery(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindCallbackQuery, route, handlers...)
}

// ChosenInlineResult bind chosen inline result updates and the route through
// handler/handlers.
func (r *Router) ChosenInlineResult(route string, handlers ...func(*Context) error) error {
	return r.Handle(KindChosenInlineResult, route, handlers...)
}

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
