package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// Venue venue
// swagger:model Venue
type Venue struct {

	// address
	Address string `json:"address,omitempty"`

	// foursquare id
	FoursquareID string `json:"foursquare_id,omitempty"`

	// location
	Location *Location `json:"location,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this venue
func (m *Venue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Venue) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {

		if err := m.Location.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}