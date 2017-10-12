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

// InlineQueryResultAudio inline query result audio
// swagger:model InlineQueryResultAudio

type InlineQueryResultAudio struct {

	// audio duration
	AudioDuration int64 `json:"audio_duration,omitempty"`

	// audio url
	// Required: true
	AudioURL *string `json:"audio_url"`

	// caption
	Caption string `json:"caption,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// input message content
	InputMessageContent interface{} `json:"input_message_content,omitempty"`

	// performer
	Performer string `json:"performer,omitempty"`

	// reply markup
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`

	// type
	// Required: true
	Type InlineType `json:"type"`
}

/* polymorph InlineQueryResultAudio audio_duration false */

/* polymorph InlineQueryResultAudio audio_url false */

/* polymorph InlineQueryResultAudio caption false */

/* polymorph InlineQueryResultAudio id false */

/* polymorph InlineQueryResultAudio input_message_content false */

/* polymorph InlineQueryResultAudio performer false */

/* polymorph InlineQueryResultAudio reply_markup false */

/* polymorph InlineQueryResultAudio title false */

/* polymorph InlineQueryResultAudio type false */

// Validate validates this inline query result audio
func (m *InlineQueryResultAudio) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudioURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

func (m *InlineQueryResultAudio) validateAudioURL(formats strfmt.Registry) error {

	if err := validate.Required("audio_url", "body", m.AudioURL); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultAudio) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultAudio) validateReplyMarkup(formats strfmt.Registry) error {

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

func (m *InlineQueryResultAudio) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultAudio) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineQueryResultAudio) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineQueryResultAudio) UnmarshalBinary(b []byte) error {
	var res InlineQueryResultAudio
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
