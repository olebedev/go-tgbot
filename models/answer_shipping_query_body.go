// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AnswerShippingQueryBody answer shipping query body
// swagger:model AnswerShippingQueryBody
type AnswerShippingQueryBody struct {

	// error message
	ErrorMessage string `json:"error_message,omitempty"`

	// ok
	// Required: true
	Ok *bool `json:"ok"`

	// shipping options
	ShippingOptions []*ShippingOption `json:"shipping_options"`

	// shipping query id
	// Required: true
	ShippingQueryID *string `json:"shipping_query_id"`
}

// Validate validates this answer shipping query body
func (m *AnswerShippingQueryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOk(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShippingOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShippingQueryID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnswerShippingQueryBody) validateOk(formats strfmt.Registry) error {

	if err := validate.Required("ok", "body", m.Ok); err != nil {
		return err
	}

	return nil
}

func (m *AnswerShippingQueryBody) validateShippingOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.ShippingOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.ShippingOptions); i++ {

		if swag.IsZero(m.ShippingOptions[i]) { // not required
			continue
		}

		if m.ShippingOptions[i] != nil {

			if err := m.ShippingOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shipping_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AnswerShippingQueryBody) validateShippingQueryID(formats strfmt.Registry) error {

	if err := validate.Required("shipping_query_id", "body", m.ShippingQueryID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnswerShippingQueryBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnswerShippingQueryBody) UnmarshalBinary(b []byte) error {
	var res AnswerShippingQueryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
