package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new attachments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for attachments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SendPhoto send photo API
*/
func (a *Client) SendPhoto(params *SendPhotoParams) (*SendPhotoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendPhotoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendPhoto",
		Method:             "POST",
		PathPattern:        "/bot{token}/sendPhoto",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendPhotoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SendPhotoOK), nil

}

/*
SendPhotoLink send photo link API
*/
func (a *Client) SendPhotoLink(params *SendPhotoLinkParams) (*SendPhotoLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendPhotoLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendPhotoLink",
		Method:             "POST",
		PathPattern:        "/bot{token}/sendPhoto#link",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendPhotoLinkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SendPhotoLinkOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
