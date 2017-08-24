// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Voice voice
// swagger:model Voice
type Voice struct {

	// duration
	Duration int64 `json:"duration,omitempty"`

	// file id
	FileID string `json:"file_id,omitempty"`

	// file size
	FileSize int64 `json:"file_size,omitempty"`

	// mime type
	MimeType string `json:"mime_type,omitempty"`
}

// Validate validates this voice
func (m *Voice) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Voice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Voice) UnmarshalBinary(b []byte) error {
	var res Voice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
