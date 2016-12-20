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

// NewSendPhotoLinkParams creates a new SendPhotoLinkParams object
// with the default values initialized.
func NewSendPhotoLinkParams() *SendPhotoLinkParams {
	var ()
	return &SendPhotoLinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendPhotoLinkParamsWithTimeout creates a new SendPhotoLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendPhotoLinkParamsWithTimeout(timeout time.Duration) *SendPhotoLinkParams {
	var ()
	return &SendPhotoLinkParams{

		timeout: timeout,
	}
}

// NewSendPhotoLinkParamsWithContext creates a new SendPhotoLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendPhotoLinkParamsWithContext(ctx context.Context) *SendPhotoLinkParams {
	var ()
	return &SendPhotoLinkParams{

		Context: ctx,
	}
}

/*SendPhotoLinkParams contains all the parameters to send to the API endpoint
for the send photo link operation typically these are written to a http.Request
*/
type SendPhotoLinkParams struct {

	/*Body*/
	Body *models.SendPhotoLinkBody
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send photo link params
func (o *SendPhotoLinkParams) WithTimeout(timeout time.Duration) *SendPhotoLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send photo link params
func (o *SendPhotoLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send photo link params
func (o *SendPhotoLinkParams) WithContext(ctx context.Context) *SendPhotoLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send photo link params
func (o *SendPhotoLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the send photo link params
func (o *SendPhotoLinkParams) WithBody(body *models.SendPhotoLinkBody) *SendPhotoLinkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send photo link params
func (o *SendPhotoLinkParams) SetBody(body *models.SendPhotoLinkBody) {
	o.Body = body
}

// WithToken adds the token to the send photo link params
func (o *SendPhotoLinkParams) WithToken(token *string) *SendPhotoLinkParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the send photo link params
func (o *SendPhotoLinkParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *SendPhotoLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.SendPhotoLinkBody)
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
