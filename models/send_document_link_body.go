// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SendDocumentLinkBody send document link body
// swagger:model SendDocumentLinkBody

type SendDocumentLinkBody struct {

	// caption
	Caption string `json:"caption,omitempty"`

	// chat id
	// Required: true
	ChatID interface{} `json:"chat_id"`

	// disable notification
	DisableNotification bool `json:"disable_notification,omitempty"`

	// document
	// Required: true
	Document *string `json:"document"`

	// reply markup
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`

	// reply to message id
	ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`
}

/* polymorph SendDocumentLinkBody caption false */

/* polymorph SendDocumentLinkBody chat_id false */

/* polymorph SendDocumentLinkBody disable_notification false */

/* polymorph SendDocumentLinkBody document false */

/* polymorph SendDocumentLinkBody reply_markup false */

/* polymorph SendDocumentLinkBody reply_to_message_id false */

// Validate validates this send document link body
func (m *SendDocumentLinkBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChatID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDocument(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendDocumentLinkBody) validateChatID(formats strfmt.Registry) error {

	return nil
}

func (m *SendDocumentLinkBody) validateDocument(formats strfmt.Registry) error {

	if err := validate.Required("document", "body", m.Document); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendDocumentLinkBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendDocumentLinkBody) UnmarshalBinary(b []byte) error {
	var res SendDocumentLinkBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
