// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StickerSet sticker set
// swagger:model StickerSet
type StickerSet struct {

	// is masks
	IsMasks bool `json:"is_masks,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// stickers
	Stickers []*Sticker `json:"stickers"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this sticker set
func (m *StickerSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStickers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StickerSet) validateStickers(formats strfmt.Registry) error {

	if swag.IsZero(m.Stickers) { // not required
		return nil
	}

	for i := 0; i < len(m.Stickers); i++ {

		if swag.IsZero(m.Stickers[i]) { // not required
			continue
		}

		if m.Stickers[i] != nil {

			if err := m.Stickers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stickers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StickerSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StickerSet) UnmarshalBinary(b []byte) error {
	var res StickerSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
