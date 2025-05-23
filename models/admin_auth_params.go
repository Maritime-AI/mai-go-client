// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdminAuthParams admin auth params
//
// swagger:model AdminAuthParams
type AdminAuthParams struct {

	// email
	Email string `json:"email,omitempty"`

	// password
	Password string `json:"password,omitempty"`
}

// Validate validates this admin auth params
func (m *AdminAuthParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this admin auth params based on context it is used
func (m *AdminAuthParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdminAuthParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminAuthParams) UnmarshalBinary(b []byte) error {
	var res AdminAuthParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
