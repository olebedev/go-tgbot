package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ResponseParameters response parameters
// swagger:model ResponseParameters
type ResponseParameters struct {

	// migrate to chat id
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`

	// retry after
	RetryAfter int64 `json:"retry_after,omitempty"`
}

// Validate validates this response parameters
func (m *ResponseParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
