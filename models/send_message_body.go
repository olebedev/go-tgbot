package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// SendMessageBody send message body
// swagger:model SendMessageBody
type SendMessageBody struct {

	// chat id
	// Required: true
	ChatID interface{} `json:"chat_id"`

	// disable notification
	DisableNotification bool `json:"disable_notification,omitempty"`

	// disable web page preview
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

	// parse mode
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// reply markup
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`

	// reply to message id
	ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

	// text
	// Required: true
	Text *string `json:"text"`
}

// Validate validates this send message body
func (m *SendMessageBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChatID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParseMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendMessageBody) validateChatID(formats strfmt.Registry) error {

	return nil
}

func (m *SendMessageBody) validateParseMode(formats strfmt.Registry) error {

	if swag.IsZero(m.ParseMode) { // not required
		return nil
	}

	if err := m.ParseMode.Validate(formats); err != nil {
		return err
	}

	return nil
}

func (m *SendMessageBody) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}
