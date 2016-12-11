package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// PhotoSize photo size
// swagger:model PhotoSize
type PhotoSize struct {

	// file id
	FileID string `json:"file_id,omitempty"`

	// file size
	FileSize int64 `json:"file_size,omitempty"`

	// height
	Height int64 `json:"height,omitempty"`

	// width
	Width int64 `json:"width,omitempty"`
}

// Validate validates this photo size
func (m *PhotoSize) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
