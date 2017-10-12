// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Sticker sticker
// swagger:model Sticker

type Sticker struct {

	// emoji
	Emoji string `json:"emoji,omitempty"`

	// file id
	FileID string `json:"file_id,omitempty"`

	// file size
	FileSize int64 `json:"file_size,omitempty"`

	// height
	Height int64 `json:"height,omitempty"`

	// mask position
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// set name
	SetName string `json:"set_name,omitempty"`

	// thumb
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// width
	Width int64 `json:"width,omitempty"`
}

/* polymorph Sticker emoji false */

/* polymorph Sticker file_id false */

/* polymorph Sticker file_size false */

/* polymorph Sticker height false */

/* polymorph Sticker mask_position false */

/* polymorph Sticker set_name false */

/* polymorph Sticker thumb false */

/* polymorph Sticker width false */

// Validate validates this sticker
func (m *Sticker) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaskPosition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateThumb(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Sticker) validateMaskPosition(formats strfmt.Registry) error {

	if swag.IsZero(m.MaskPosition) { // not required
		return nil
	}

	if m.MaskPosition != nil {

		if err := m.MaskPosition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mask_position")
			}
			return err
		}
	}

	return nil
}

func (m *Sticker) validateThumb(formats strfmt.Registry) error {

	if swag.IsZero(m.Thumb) { // not required
		return nil
	}

	if m.Thumb != nil {

		if err := m.Thumb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thumb")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Sticker) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Sticker) UnmarshalBinary(b []byte) error {
	var res Sticker
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
