// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new messages API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for messages API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteMessage delete message API
*/
func (a *Client) DeleteMessage(params *DeleteMessageParams) (*DeleteMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMessageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMessage",
		Method:             "GET",
		PathPattern:        "/bot{token}/deleteMessage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMessageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMessageOK), nil

}

/*
EditMessageCaption edit message caption API
*/
func (a *Client) EditMessageCaption(params *EditMessageCaptionParams) (*EditMessageCaptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditMessageCaptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editMessageCaption",
		Method:             "POST",
		PathPattern:        "/bot{token}/editMessageCaption",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EditMessageCaptionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EditMessageCaptionOK), nil

}

/*
EditMessageReplyMarkup edit message reply markup API
*/
func (a *Client) EditMessageReplyMarkup(params *EditMessageReplyMarkupParams) (*EditMessageReplyMarkupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditMessageReplyMarkupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editMessageReplyMarkup",
		Method:             "POST",
		PathPattern:        "/bot{token}/editMessageReplyMarkup",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EditMessageReplyMarkupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EditMessageReplyMarkupOK), nil

}

/*
EditMessageText edit message text API
*/
func (a *Client) EditMessageText(params *EditMessageTextParams) (*EditMessageTextOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditMessageTextParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editMessageText",
		Method:             "POST",
		PathPattern:        "/bot{token}/editMessageText",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EditMessageTextReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EditMessageTextOK), nil

}

/*
ForwardMessage forward message API
*/
func (a *Client) ForwardMessage(params *ForwardMessageParams) (*ForwardMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForwardMessageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "forwardMessage",
		Method:             "POST",
		PathPattern:        "/bot{token}/forwardMessage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForwardMessageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ForwardMessageOK), nil

}

/*
SendMessage send message API
*/
func (a *Client) SendMessage(params *SendMessageParams) (*SendMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendMessageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendMessage",
		Method:             "POST",
		PathPattern:        "/bot{token}/sendMessage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendMessageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SendMessageOK), nil

}

/*
SendMessageBytes send message bytes API
*/
func (a *Client) SendMessageBytes(params *SendMessageBytesParams) (*SendMessageBytesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendMessageBytesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendMessageBytes",
		Method:             "POST",
		PathPattern:        "/bot{token}/sendMessage#bytes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendMessageBytesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SendMessageBytesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
