package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestrictChatMemberBody restrict chat member body
// swagger:model RestrictChatMemberBody
type RestrictChatMemberBody struct {

	// can add web page previews
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

	// can send media messages
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`

	// can send messages
	CanSendMessages bool `json:"can_send_messages,omitempty"`

	// can send other messages
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

	// chat id
	// Required: true
	ChatID interface{} `json:"chat_id"`

	// until date
	UntilDate int64 `json:"until_date,omitempty"`

	// user id
	// Required: true
	UserID *int64 `json:"user_id"`
}

// Validate validates this restrict chat member body
func (m *RestrictChatMemberBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChatID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestrictChatMemberBody) validateChatID(formats strfmt.Registry) error {

	return nil
}

func (m *RestrictChatMemberBody) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestrictChatMemberBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestrictChatMemberBody) UnmarshalBinary(b []byte) error {
	var res RestrictChatMemberBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}