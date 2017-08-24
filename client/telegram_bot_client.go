// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/client/attachments"
	"github.com/olebedev/go-tgbot/client/callbacks"
	"github.com/olebedev/go-tgbot/client/chats"
	"github.com/olebedev/go-tgbot/client/games"
	"github.com/olebedev/go-tgbot/client/inline"
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/olebedev/go-tgbot/client/payments"
	"github.com/olebedev/go-tgbot/client/stickers"
	"github.com/olebedev/go-tgbot/client/updates"
	"github.com/olebedev/go-tgbot/client/users"
)

// Default telegram bot HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.telegram.org"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new telegram bot HTTP client.
func NewHTTPClient(formats strfmt.Registry) *TelegramBot {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new telegram bot HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *TelegramBot {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new telegram bot client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *TelegramBot {
	cli := new(TelegramBot)
	cli.Transport = transport

	cli.Attachments = attachments.New(transport, formats)

	cli.Callbacks = callbacks.New(transport, formats)

	cli.Chats = chats.New(transport, formats)

	cli.Games = games.New(transport, formats)

	cli.Inline = inline.New(transport, formats)

	cli.Messages = messages.New(transport, formats)

	cli.Payments = payments.New(transport, formats)

	cli.Stickers = stickers.New(transport, formats)

	cli.Updates = updates.New(transport, formats)

	cli.Users = users.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// TelegramBot is a client for telegram bot
type TelegramBot struct {
	Attachments *attachments.Client

	Callbacks *callbacks.Client

	Chats *chats.Client

	Games *games.Client

	Inline *inline.Client

	Messages *messages.Client

	Payments *payments.Client

	Stickers *stickers.Client

	Updates *updates.Client

	Users *users.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *TelegramBot) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Attachments.SetTransport(transport)

	c.Callbacks.SetTransport(transport)

	c.Chats.SetTransport(transport)

	c.Games.SetTransport(transport)

	c.Inline.SetTransport(transport)

	c.Messages.SetTransport(transport)

	c.Payments.SetTransport(transport)

	c.Stickers.SetTransport(transport)

	c.Updates.SetTransport(transport)

	c.Users.SetTransport(transport)

}
