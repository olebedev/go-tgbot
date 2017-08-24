// Code generated by go-swagger; DO NOT EDIT.

package stickers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSetStickerPositionInSetParams creates a new SetStickerPositionInSetParams object
// with the default values initialized.
func NewSetStickerPositionInSetParams() *SetStickerPositionInSetParams {
	var ()
	return &SetStickerPositionInSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetStickerPositionInSetParamsWithTimeout creates a new SetStickerPositionInSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetStickerPositionInSetParamsWithTimeout(timeout time.Duration) *SetStickerPositionInSetParams {
	var ()
	return &SetStickerPositionInSetParams{

		timeout: timeout,
	}
}

// NewSetStickerPositionInSetParamsWithContext creates a new SetStickerPositionInSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetStickerPositionInSetParamsWithContext(ctx context.Context) *SetStickerPositionInSetParams {
	var ()
	return &SetStickerPositionInSetParams{

		Context: ctx,
	}
}

// NewSetStickerPositionInSetParamsWithHTTPClient creates a new SetStickerPositionInSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetStickerPositionInSetParamsWithHTTPClient(client *http.Client) *SetStickerPositionInSetParams {
	var ()
	return &SetStickerPositionInSetParams{
		HTTPClient: client,
	}
}

/*SetStickerPositionInSetParams contains all the parameters to send to the API endpoint
for the set sticker position in set operation typically these are written to a http.Request
*/
type SetStickerPositionInSetParams struct {

	/*Position*/
	Position int64
	/*Sticker*/
	Sticker string
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set sticker position in set params
func (o *SetStickerPositionInSetParams) WithTimeout(timeout time.Duration) *SetStickerPositionInSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set sticker position in set params
func (o *SetStickerPositionInSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set sticker position in set params
func (o *SetStickerPositionInSetParams) WithContext(ctx context.Context) *SetStickerPositionInSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set sticker position in set params
func (o *SetStickerPositionInSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set sticker position in set params
func (o *SetStickerPositionInSetParams) WithHTTPClient(client *http.Client) *SetStickerPositionInSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set sticker position in set params
func (o *SetStickerPositionInSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPosition adds the position to the set sticker position in set params
func (o *SetStickerPositionInSetParams) WithPosition(position int64) *SetStickerPositionInSetParams {
	o.SetPosition(position)
	return o
}

// SetPosition adds the position to the set sticker position in set params
func (o *SetStickerPositionInSetParams) SetPosition(position int64) {
	o.Position = position
}

// WithSticker adds the sticker to the set sticker position in set params
func (o *SetStickerPositionInSetParams) WithSticker(sticker string) *SetStickerPositionInSetParams {
	o.SetSticker(sticker)
	return o
}

// SetSticker adds the sticker to the set sticker position in set params
func (o *SetStickerPositionInSetParams) SetSticker(sticker string) {
	o.Sticker = sticker
}

// WithToken adds the token to the set sticker position in set params
func (o *SetStickerPositionInSetParams) WithToken(token *string) *SetStickerPositionInSetParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the set sticker position in set params
func (o *SetStickerPositionInSetParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *SetStickerPositionInSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param position
	frPosition := o.Position
	fPosition := swag.FormatInt64(frPosition)
	if fPosition != "" {
		if err := r.SetFormParam("position", fPosition); err != nil {
			return err
		}
	}

	// form param sticker
	frSticker := o.Sticker
	fSticker := frSticker
	if fSticker != "" {
		if err := r.SetFormParam("sticker", fSticker); err != nil {
			return err
		}
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
