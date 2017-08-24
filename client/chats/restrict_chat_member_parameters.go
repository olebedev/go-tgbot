// Code generated by go-swagger; DO NOT EDIT.

package chats

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

// NewRestrictChatMemberParams creates a new RestrictChatMemberParams object
// with the default values initialized.
func NewRestrictChatMemberParams() *RestrictChatMemberParams {
	var ()
	return &RestrictChatMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRestrictChatMemberParamsWithTimeout creates a new RestrictChatMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRestrictChatMemberParamsWithTimeout(timeout time.Duration) *RestrictChatMemberParams {
	var ()
	return &RestrictChatMemberParams{

		timeout: timeout,
	}
}

// NewRestrictChatMemberParamsWithContext creates a new RestrictChatMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewRestrictChatMemberParamsWithContext(ctx context.Context) *RestrictChatMemberParams {
	var ()
	return &RestrictChatMemberParams{

		Context: ctx,
	}
}

// NewRestrictChatMemberParamsWithHTTPClient creates a new RestrictChatMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRestrictChatMemberParamsWithHTTPClient(client *http.Client) *RestrictChatMemberParams {
	var ()
	return &RestrictChatMemberParams{
		HTTPClient: client,
	}
}

/*RestrictChatMemberParams contains all the parameters to send to the API endpoint
for the restrict chat member operation typically these are written to a http.Request
*/
type RestrictChatMemberParams struct {

	/*Body*/
	Body *models.RestrictChatMemberBody
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the restrict chat member params
func (o *RestrictChatMemberParams) WithTimeout(timeout time.Duration) *RestrictChatMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restrict chat member params
func (o *RestrictChatMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restrict chat member params
func (o *RestrictChatMemberParams) WithContext(ctx context.Context) *RestrictChatMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restrict chat member params
func (o *RestrictChatMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restrict chat member params
func (o *RestrictChatMemberParams) WithHTTPClient(client *http.Client) *RestrictChatMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restrict chat member params
func (o *RestrictChatMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the restrict chat member params
func (o *RestrictChatMemberParams) WithBody(body *models.RestrictChatMemberBody) *RestrictChatMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the restrict chat member params
func (o *RestrictChatMemberParams) SetBody(body *models.RestrictChatMemberBody) {
	o.Body = body
}

// WithToken adds the token to the restrict chat member params
func (o *RestrictChatMemberParams) WithToken(token *string) *RestrictChatMemberParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the restrict chat member params
func (o *RestrictChatMemberParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *RestrictChatMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.RestrictChatMemberBody)
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
