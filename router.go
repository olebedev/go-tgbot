package tgbot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/olebedev/go-tgbot/client"
	"github.com/olebedev/go-tgbot/client/updates"
	"github.com/olebedev/go-tgbot/models"
	"github.com/pkg/errors"
)

const DefaultPollTimeout int64 = 29

// Router ...
type Router struct {
	*client.TelegramBot
	PollTimeout int64
	middlewares []func(*Context) error
	handlers    []func(*Context) error
	routes      []struct {
		key   string
		re    *regexp.Regexp
		value []func(*Context) error
	}
}

type Error struct {
	Message string
	Code    uint16
}

func (e Error) Error() string {
	if e.Code != 0 {
		return fmt.Sprintf("%d: %s", e.Code, e.Message)
	}
	return e.Message
}

// New returns a router.
func New(ctx context.Context, token string) *Router {
	r := &Router{
		TelegramBot: NewClient(ctx, token),
		middlewares: make([]func(*Context) error, 0),
		handlers:    make([]func(*Context) error, 1),
		routes: make([]struct {
			key   string
			re    *regexp.Regexp
			value []func(*Context) error
		}, 0),
		PollTimeout: DefaultPollTimeout,
	}
	r.handlers[0] = r.routeMiddleware
	return r
}

// Use appends given middlewares into call chain.
func (r *Router) Use(middlewares ...func(*Context) error) {
	r.middlewares = append(r.middlewares, middlewares...)
	r.handlers = make([]func(*Context) error, len(r.middlewares)+1)
	copy(r.handlers, r.middlewares)
	r.handlers[len(r.middlewares)] = r.routeMiddleware
}

// Bind binds an update to the handler(s).
func (r *Router) Bind(route string, handlers ...func(*Context) error) error {
	re, err := regexp.Compile(route)
	if err != nil {
		return err
	}

	r.routes = append(r.routes, struct {
		key   string
		re    *regexp.Regexp
		value []func(*Context) error
	}{
		key:   route,
		re:    re,
		value: handlers,
	})

	return nil
}

// Route routes given updates into the bound handlers
// through defined middlewares.
func (r *Router) Route(u *models.Update) error {
	return errors.Wrap(r.buildContext(u).Next(), "call context chain")
}

func (r *Router) buildContext(up *models.Update) *Context {
	var from *models.User
	var chat *models.Chat
	var u, text string
	switch true {
	case up.Message != nil || up.EditedMessage != nil || up.ChannelPost != nil || up.EditedChannelPost != nil:
		var msg *models.Message
		switch true {
		case up.Message != nil:
			u += "/message"
			msg = up.Message
		case up.EditedMessage != nil:
			u += "/edited_message"
			msg = up.EditedMessage
		case up.ChannelPost != nil:
			u += "/channel_post"
			msg = up.ChannelPost
		case up.EditedChannelPost != nil:
			u += "/edited_channel_post"
			msg = up.EditedChannelPost
		}
		text = msg.Text
		from = msg.From
		chat = msg.Chat

		// reply/forward/direct
		switch true {
		case msg.ReplyToMessage != nil:
			u += "/reply"
		case msg.ForwardFrom != nil:
			u += "/forward"
		default:
			u += "/-"
		}

		// chat type
		u += "/" + *msg.Chat.Type

		// message type
		switch true {
		case msg.Text != "":
			u += "/text"
		case msg.Audio != nil:
			u += "/audio"
		case msg.Document != nil:
			u += "/document"
		case msg.Game != nil:
			u += "/game"
		case msg.Photo != nil:
			u += "/photo"
		case msg.Sticker != nil:
			u += "/sticker"
		case msg.Video != nil:
			u += "/video"
		case msg.Voice != nil:
			u += "/voice"
		case msg.VideoNote != nil:
			u += "/video_note"
		case msg.Contact != nil:
			u += "/contact"
		case msg.Location != nil:
			u += "/location"
		case msg.Venue != nil:
			u += "/venue"
		case msg.NewChatMembers != nil:
			u += "/new_chat_members"
		case msg.LeftChatMember != nil:
			u += "/left_chat_member"
		case msg.NewChatTitle != "":
			u += "/new_chat_title"
		case msg.NewChatPhoto != nil:
			u += "/new_chat_photo"
		case msg.DeleteChatPhoto:
			u += "/delete_chat_photo"
		case msg.GroupChatCreated:
			u += "/group_chat_created"
		case msg.SupergroupChatCreated:
			u += "/supergroup_chat_created"
		case msg.ChannelChatCreated:
			u += "/channel_chat_created"
		case msg.PinnedMessage != nil:
			u += "/pinned_message"
		case msg.Invoice != nil:
			u += "/invoice"
		case msg.SuccessfulPayment != nil:
			u += "/successful_payment"
		default:
			u += "/text"
		}
	case up.InlineQuery != nil:
		u += "/inline_query"
		text = up.InlineQuery.Query
		from = up.InlineQuery.From
		switch true {
		case up.InlineQuery.Location != nil:
			u += "/location"
		default:
			u += "/-"
		}
	case up.CallbackQuery != nil:
		u += "/callback_query"
		text = up.CallbackQuery.Data
		from = up.CallbackQuery.From
		suffix := "-"
		if up.CallbackQuery.Message != nil {
			chat = up.CallbackQuery.Message.Chat
			if chat != nil && chat.Type != nil {
				suffix = *chat.Type
			}
		}
		u += "/" + suffix
	case up.ChosenInlineResult != nil:
		u += "/chosen_inline_result"
		switch true {
		case up.ChosenInlineResult.Location != nil:
			u += "/location"
		default:
			u += "/-"
		}
		text = *up.ChosenInlineResult.ResultID
		from = up.ChosenInlineResult.From
	case up.ShippingQuery != nil:
		u += "/shipping_query"
		text = up.ShippingQuery.InvoicePayload
		from = up.ShippingQuery.From
	case up.PreCheckoutQuery != nil:
		u += "/pre_checkout_query"
		text = up.PreCheckoutQuery.InvoicePayload
		from = up.PreCheckoutQuery.From
	default:
		b, _ := json.MarshalIndent(u, "", "  ")
		panic("unknown update type: " + string(b))
	}
	c := &Context{
		Path:     u,
		From:     from,
		Text:     text,
		Chat:     chat,
		Update:   up,
		handlers: r.handlers,
	}
	return c
}

// https://play.golang.org/p/L-x3h11Bak
func (r *Router) routeMiddleware(c *Context) error {
	var m []string
	for _, ro := range r.routes {
		if m = ro.re.FindStringSubmatch(c.Path); len(m) != 0 {
			c.Capture = m[1:]
			c.handlers = ro.value
			c.index = 0
			err := c.Next()
			if err != nil {
				return err
			}

			if c.isSkipped {
				c.isSkipped = false
				continue
			} else {
				return nil
			}
		}
	}

	return &Error{
		Message: fmt.Sprintf("missed handler for '%s'", c.Path),
		Code:    404,
	}
}

// Poll does a polling of API endpoints and routes consumed updates. It returns an error
// if any of handlers return the error.
func (r *Router) Poll(ctx context.Context, allowed ...models.AllowedUpdate) error {
	var a []string
	for _, item := range allowed {
		a = append(a, string(item))
	}

	p := updates.NewGetUpdatesParams().
		WithContext(ctx).WithBody(&models.GetUpdatesBody{
		Timeout:        r.PollTimeout,
		AllowedUpdates: allowed,
	})

	for {
		u, err := r.Updates.GetUpdates(p)
		if err != nil {
			return errors.Wrap(err, "polling updates")
		}
		for _, update := range u.Payload.Result {
			p.Body.Offset = update.UpdateID + 1
			go r.Route(update)
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
