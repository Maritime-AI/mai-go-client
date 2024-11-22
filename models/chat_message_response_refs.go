// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChatMessageResponseRefs chat message response refs
//
// swagger:model ChatMessageResponseRefs
type ChatMessageResponseRefs struct {

	// crew external ids
	CrewExternalIds []string `json:"crew_external_ids"`
}

// Validate validates this chat message response refs
func (m *ChatMessageResponseRefs) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this chat message response refs based on context it is used
func (m *ChatMessageResponseRefs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChatMessageResponseRefs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChatMessageResponseRefs) UnmarshalBinary(b []byte) error {
	var res ChatMessageResponseRefs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}