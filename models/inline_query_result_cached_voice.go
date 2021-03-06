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

// InlineQueryResultCachedVoice inline query result cached voice
// swagger:model InlineQueryResultCachedVoice
type InlineQueryResultCachedVoice struct {

	// caption
	Caption string `json:"caption,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// input message content
	InputMessageContent interface{} `json:"input_message_content,omitempty"`

	// parse mode
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// reply markup
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	// Required: true
	Type InlineType `json:"type"`

	// voice file id
	// Required: true
	VoiceFileID *string `json:"voice_file_id"`
}

// Validate validates this inline query result cached voice
func (m *InlineQueryResultCachedVoice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParseMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplyMarkup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVoiceFileID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineQueryResultCachedVoice) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultCachedVoice) validateParseMode(formats strfmt.Registry) error {

	if swag.IsZero(m.ParseMode) { // not required
		return nil
	}

	if err := m.ParseMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("parse_mode")
		}
		return err
	}

	return nil
}

func (m *InlineQueryResultCachedVoice) validateReplyMarkup(formats strfmt.Registry) error {

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

func (m *InlineQueryResultCachedVoice) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *InlineQueryResultCachedVoice) validateVoiceFileID(formats strfmt.Registry) error {

	if err := validate.Required("voice_file_id", "body", m.VoiceFileID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineQueryResultCachedVoice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineQueryResultCachedVoice) UnmarshalBinary(b []byte) error {
	var res InlineQueryResultCachedVoice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
