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

// PromoteChatMemberBody promote chat member body
// swagger:model PromoteChatMemberBody

type PromoteChatMemberBody struct {

	// can change info
	CanChangeInfo *bool `json:"can_change_info,omitempty"`

	// can delete messages
	CanDeleteMessages *bool `json:"can_delete_messages,omitempty"`

	// can edit messages
	CanEditMessages *bool `json:"can_edit_messages,omitempty"`

	// can invite users
	CanInviteUsers *bool `json:"can_invite_users,omitempty"`

	// can pin messages
	CanPinMessages *bool `json:"can_pin_messages,omitempty"`

	// can post messages
	CanPostMessages *bool `json:"can_post_messages,omitempty"`

	// can promote members
	CanPromoteMembers *bool `json:"can_promote_members,omitempty"`

	// can restrict members
	CanRestrictMembers *bool `json:"can_restrict_members,omitempty"`

	// chat id
	// Required: true
	ChatID interface{} `json:"chat_id"`

	// user id
	// Required: true
	UserID *int64 `json:"user_id"`
}

/* polymorph PromoteChatMemberBody can_change_info false */

/* polymorph PromoteChatMemberBody can_delete_messages false */

/* polymorph PromoteChatMemberBody can_edit_messages false */

/* polymorph PromoteChatMemberBody can_invite_users false */

/* polymorph PromoteChatMemberBody can_pin_messages false */

/* polymorph PromoteChatMemberBody can_post_messages false */

/* polymorph PromoteChatMemberBody can_promote_members false */

/* polymorph PromoteChatMemberBody can_restrict_members false */

/* polymorph PromoteChatMemberBody chat_id false */

/* polymorph PromoteChatMemberBody user_id false */

// Validate validates this promote chat member body
func (m *PromoteChatMemberBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChatID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PromoteChatMemberBody) validateChatID(formats strfmt.Registry) error {

	return nil
}

func (m *PromoteChatMemberBody) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PromoteChatMemberBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PromoteChatMemberBody) UnmarshalBinary(b []byte) error {
	var res PromoteChatMemberBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
