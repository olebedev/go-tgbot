// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSendMessageBytesParams creates a new SendMessageBytesParams object
// with the default values initialized.
func NewSendMessageBytesParams() *SendMessageBytesParams {
	var ()
	return &SendMessageBytesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendMessageBytesParamsWithTimeout creates a new SendMessageBytesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendMessageBytesParamsWithTimeout(timeout time.Duration) *SendMessageBytesParams {
	var ()
	return &SendMessageBytesParams{

		timeout: timeout,
	}
}

// NewSendMessageBytesParamsWithContext creates a new SendMessageBytesParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendMessageBytesParamsWithContext(ctx context.Context) *SendMessageBytesParams {
	var ()
	return &SendMessageBytesParams{

		Context: ctx,
	}
}

// NewSendMessageBytesParamsWithHTTPClient creates a new SendMessageBytesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendMessageBytesParamsWithHTTPClient(client *http.Client) *SendMessageBytesParams {
	var ()
	return &SendMessageBytesParams{
		HTTPClient: client,
	}
}

/*SendMessageBytesParams contains all the parameters to send to the API endpoint
for the send message bytes operation typically these are written to a http.Request
*/
type SendMessageBytesParams struct {

	/*Body*/
	Body io.ReadCloser
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send message bytes params
func (o *SendMessageBytesParams) WithTimeout(timeout time.Duration) *SendMessageBytesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send message bytes params
func (o *SendMessageBytesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send message bytes params
func (o *SendMessageBytesParams) WithContext(ctx context.Context) *SendMessageBytesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send message bytes params
func (o *SendMessageBytesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send message bytes params
func (o *SendMessageBytesParams) WithHTTPClient(client *http.Client) *SendMessageBytesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send message bytes params
func (o *SendMessageBytesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send message bytes params
func (o *SendMessageBytesParams) WithBody(body io.ReadCloser) *SendMessageBytesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send message bytes params
func (o *SendMessageBytesParams) SetBody(body io.ReadCloser) {
	o.Body = body
}

// WithToken adds the token to the send message bytes params
func (o *SendMessageBytesParams) WithToken(token *string) *SendMessageBytesParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the send message bytes params
func (o *SendMessageBytesParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *SendMessageBytesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.Token != nil {

		// path param token
		if err := r.SetPathParam("token", *o.Token); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
