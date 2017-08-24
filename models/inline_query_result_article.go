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

// InlineQueryResultArticle inline query result article
// swagger:model InlineQueryResultArticle
type InlineQueryResultArticle struct {

	// description
	Description string `json:"description,omitempty"`

	// hide url
	HideURL bool `json:"hide_url,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// input message content
	// Required: true
	InputMessageContent interface{} `json:"input_message_content"`

	// reply markup
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// thumb height
	ThumbHeight int64 `json:"thumb_height,omitempty"`

	// thumb url
	ThumbURL string `json:"thumb_url,omitempty"`

	// thumb width
	ThumbWidth int64 `json:"thumb_width,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`

	// type
	// Required: true
	Type InlineType `json:"type"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this inline query result article
func (m *InlineQueryResultArticle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInputMessageContent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplyMarkup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineQueryResultArticle) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultArticle) validateInputMessageContent(formats strfmt.Registry) error {

	return nil
}

func (m *InlineQueryResultArticle) validateReplyMarkup(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyMarkup) { // not required
		return nil
	}

	if m.ReplyMarkup != nil {

		if err := m.ReplyMarkup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reply_markup")
			}
			return err
		}
	}

	return nil
}

func (m *InlineQueryResultArticle) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultArticle) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineQueryResultArticle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineQueryResultArticle) UnmarshalBinary(b []byte) error {
	var res InlineQueryResultArticle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
