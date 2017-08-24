// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SuccessfulPayment successful payment
// swagger:model SuccessfulPayment
type SuccessfulPayment struct {

	// currency
	Currency string `json:"currency,omitempty"`

	// invoice payload
	InvoicePayload string `json:"invoice_payload,omitempty"`

	// order info
	OrderInfo *OrderInfo `json:"order_info,omitempty"`

	// provider payment charge id
	ProviderPaymentChargeID string `json:"provider_payment_charge_id,omitempty"`

	// shipping option id
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// telegram payment charge id
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id,omitempty"`

	// total amount
	TotalAmount int64 `json:"total_amount,omitempty"`
}

// Validate validates this successful payment
func (m *SuccessfulPayment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SuccessfulPayment) validateOrderInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderInfo) { // not required
		return nil
	}

	if m.OrderInfo != nil {

		if err := m.OrderInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("order_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SuccessfulPayment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuccessfulPayment) UnmarshalBinary(b []byte) error {
	var res SuccessfulPayment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
