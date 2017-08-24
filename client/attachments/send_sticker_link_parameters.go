// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// NewSendStickerLinkParams creates a new SendStickerLinkParams object
// with the default values initialized.
func NewSendStickerLinkParams() *SendStickerLinkParams {
	var ()
	return &SendStickerLinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendStickerLinkParamsWithTimeout creates a new SendStickerLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendStickerLinkParamsWithTimeout(timeout time.Duration) *SendStickerLinkParams {
	var ()
	return &SendStickerLinkParams{

		timeout: timeout,
	}
}

// NewSendStickerLinkParamsWithContext creates a new SendStickerLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendStickerLinkParamsWithContext(ctx context.Context) *SendStickerLinkParams {
	var ()
	return &SendStickerLinkParams{

		Context: ctx,
	}
}

// NewSendStickerLinkParamsWithHTTPClient creates a new SendStickerLinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendStickerLinkParamsWithHTTPClient(client *http.Client) *SendStickerLinkParams {
	var ()
	return &SendStickerLinkParams{
		HTTPClient: client,
	}
}

/*SendStickerLinkParams contains all the parameters to send to the API endpoint
for the send sticker link operation typically these are written to a http.Request
*/
type SendStickerLinkParams struct {

	/*Body*/
	Body *models.SendStickerLinkBody
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send sticker link params
func (o *SendStickerLinkParams) WithTimeout(timeout time.Duration) *SendStickerLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send sticker link params
func (o *SendStickerLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send sticker link params
func (o *SendStickerLinkParams) WithContext(ctx context.Context) *SendStickerLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send sticker link params
func (o *SendStickerLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send sticker link params
func (o *SendStickerLinkParams) WithHTTPClient(client *http.Client) *SendStickerLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send sticker link params
func (o *SendStickerLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send sticker link params
func (o *SendStickerLinkParams) WithBody(body *models.SendStickerLinkBody) *SendStickerLinkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send sticker link params
func (o *SendStickerLinkParams) SetBody(body *models.SendStickerLinkBody) {
	o.Body = body
}

// WithToken adds the token to the send sticker link params
func (o *SendStickerLinkParams) WithToken(token *string) *SendStickerLinkParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the send sticker link params
func (o *SendStickerLinkParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *SendStickerLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.SendStickerLinkBody)
	}

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
