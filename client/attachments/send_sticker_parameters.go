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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSendStickerParams creates a new SendStickerParams object
// with the default values initialized.
func NewSendStickerParams() *SendStickerParams {
	var ()
	return &SendStickerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendStickerParamsWithTimeout creates a new SendStickerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendStickerParamsWithTimeout(timeout time.Duration) *SendStickerParams {
	var ()
	return &SendStickerParams{

		timeout: timeout,
	}
}

// NewSendStickerParamsWithContext creates a new SendStickerParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendStickerParamsWithContext(ctx context.Context) *SendStickerParams {
	var ()
	return &SendStickerParams{

		Context: ctx,
	}
}

/*SendStickerParams contains all the parameters to send to the API endpoint
for the send sticker operation typically these are written to a http.Request
*/
type SendStickerParams struct {

	/*ChatID*/
	ChatID string
	/*DisableNotification*/
	DisableNotification *bool
	/*ReplyMarkup
	  json string of reply_markup object

	*/
	ReplyMarkup *string
	/*ReplyToMessageID*/
	ReplyToMessageID *int64
	/*Sticker*/
	Sticker runtime.NamedReadCloser
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send sticker params
func (o *SendStickerParams) WithTimeout(timeout time.Duration) *SendStickerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send sticker params
func (o *SendStickerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send sticker params
func (o *SendStickerParams) WithContext(ctx context.Context) *SendStickerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send sticker params
func (o *SendStickerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithChatID adds the chatID to the send sticker params
func (o *SendStickerParams) WithChatID(chatID string) *SendStickerParams {
	o.SetChatID(chatID)
	return o
}

// SetChatID adds the chatId to the send sticker params
func (o *SendStickerParams) SetChatID(chatID string) {
	o.ChatID = chatID
}

// WithDisableNotification adds the disableNotification to the send sticker params
func (o *SendStickerParams) WithDisableNotification(disableNotification *bool) *SendStickerParams {
	o.SetDisableNotification(disableNotification)
	return o
}

// SetDisableNotification adds the disableNotification to the send sticker params
func (o *SendStickerParams) SetDisableNotification(disableNotification *bool) {
	o.DisableNotification = disableNotification
}

// WithReplyMarkup adds the replyMarkup to the send sticker params
func (o *SendStickerParams) WithReplyMarkup(replyMarkup *string) *SendStickerParams {
	o.SetReplyMarkup(replyMarkup)
	return o
}

// SetReplyMarkup adds the replyMarkup to the send sticker params
func (o *SendStickerParams) SetReplyMarkup(replyMarkup *string) {
	o.ReplyMarkup = replyMarkup
}

// WithReplyToMessageID adds the replyToMessageID to the send sticker params
func (o *SendStickerParams) WithReplyToMessageID(replyToMessageID *int64) *SendStickerParams {
	o.SetReplyToMessageID(replyToMessageID)
	return o
}

// SetReplyToMessageID adds the replyToMessageId to the send sticker params
func (o *SendStickerParams) SetReplyToMessageID(replyToMessageID *int64) {
	o.ReplyToMessageID = replyToMessageID
}

// WithSticker adds the sticker to the send sticker params
func (o *SendStickerParams) WithSticker(sticker runtime.NamedReadCloser) *SendStickerParams {
	o.SetSticker(sticker)
	return o
}

// SetSticker adds the sticker to the send sticker params
func (o *SendStickerParams) SetSticker(sticker runtime.NamedReadCloser) {
	o.Sticker = sticker
}

// WithToken adds the token to the send sticker params
func (o *SendStickerParams) WithToken(token *string) *SendStickerParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the send sticker params
func (o *SendStickerParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *SendStickerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// form param chat_id
	frChatID := o.ChatID
	fChatID := frChatID
	if fChatID != "" {
		if err := r.SetFormParam("chat_id", fChatID); err != nil {
			return err
		}
	}

	if o.DisableNotification != nil {

		// form param disable_notification
		var frDisableNotification bool
		if o.DisableNotification != nil {
			frDisableNotification = *o.DisableNotification
		}
		fDisableNotification := swag.FormatBool(frDisableNotification)
		if fDisableNotification != "" {
			if err := r.SetFormParam("disable_notification", fDisableNotification); err != nil {
				return err
			}
		}

	}

	if o.ReplyMarkup != nil {

		// form param reply_markup
		var frReplyMarkup string
		if o.ReplyMarkup != nil {
			frReplyMarkup = *o.ReplyMarkup
		}
		fReplyMarkup := frReplyMarkup
		if fReplyMarkup != "" {
			if err := r.SetFormParam("reply_markup", fReplyMarkup); err != nil {
				return err
			}
		}

	}

	if o.ReplyToMessageID != nil {

		// form param reply_to_message_id
		var frReplyToMessageID int64
		if o.ReplyToMessageID != nil {
			frReplyToMessageID = *o.ReplyToMessageID
		}
		fReplyToMessageID := swag.FormatInt64(frReplyToMessageID)
		if fReplyToMessageID != "" {
			if err := r.SetFormParam("reply_to_message_id", fReplyToMessageID); err != nil {
				return err
			}
		}

	}

	// form file param sticker
	if err := r.SetFileParam("sticker", o.Sticker); err != nil {
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
