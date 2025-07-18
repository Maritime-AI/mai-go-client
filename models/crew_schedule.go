// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CrewSchedule crew schedule
//
// swagger:model CrewSchedule
type CrewSchedule struct {

	// company name
	CompanyName *string `json:"company_name,omitempty"`

	// The unique external id of the crew schedule record
	ContextRefID string `json:"context_ref_id,omitempty"`

	// The date when the crew member was last off the vessel
	// Format: date
	CrewedOff *strfmt.Date `json:"crewed_off,omitempty"`

	// The date when the crew member was last assigned to the vessel
	// Format: date
	CrewedOn strfmt.Date `json:"crewed_on,omitempty"`

	// The total number of days the mariner has been at sea before shift in hours is applied
	NumDays *float64 `json:"num_days,omitempty"`

	// The position of the crew member on the vessel
	Position string `json:"position,omitempty"`

	// The credentials required for the crew member's position
	RequiredCredentials []*CredentialParams `json:"required_credentials"`

	// shift in hours
	ShiftInHours *int64 `json:"shift_in_hours,omitempty"`

	// The current status of the crew member (e.g., active, inactive)
	Status string `json:"status,omitempty"`

	// Additional identifier for the vessel (e.g., call sign)
	VesselAdditionalIdentifier *string `json:"vessel_additional_identifier,omitempty"`

	// vessel capacity gt
	VesselCapacityGt *int64 `json:"vessel_capacity_gt,omitempty"`

	// The unique reference id of the vessel
	VesselExternalID string `json:"vessel_external_id,omitempty"`

	// vessel flag
	VesselFlag *string `json:"vessel_flag,omitempty"`

	// vessel horse power
	VesselHorsePower *float64 `json:"vessel_horse_power,omitempty"`

	// vessel imo number
	VesselImoNumber *string `json:"vessel_imo_number,omitempty"`

	// vessel mmsi number
	VesselMmsiNumber *string `json:"vessel_mmsi_number,omitempty"`

	// The name of the vessel
	VesselName string `json:"vessel_name,omitempty"`

	// vessel propulsion type
	VesselPropulsionType *string `json:"vessel_propulsion_type,omitempty"`

	// vessel type
	VesselType *string `json:"vessel_type,omitempty"`

	// water way
	WaterWay *string `json:"water_way,omitempty"`
}

// Validate validates this crew schedule
func (m *CrewSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrewedOff(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrewedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrewSchedule) validateCrewedOff(formats strfmt.Registry) error {
	if swag.IsZero(m.CrewedOff) { // not required
		return nil
	}

	if err := validate.FormatOf("crewed_off", "body", "date", m.CrewedOff.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CrewSchedule) validateCrewedOn(formats strfmt.Registry) error {
	if swag.IsZero(m.CrewedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("crewed_on", "body", "date", m.CrewedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CrewSchedule) validateRequiredCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.RequiredCredentials) { // not required
		return nil
	}

	for i := 0; i < len(m.RequiredCredentials); i++ {
		if swag.IsZero(m.RequiredCredentials[i]) { // not required
			continue
		}

		if m.RequiredCredentials[i] != nil {
			if err := m.RequiredCredentials[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("required_credentials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("required_credentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this crew schedule based on the context it is used
func (m *CrewSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequiredCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrewSchedule) contextValidateRequiredCredentials(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RequiredCredentials); i++ {

		if m.RequiredCredentials[i] != nil {

			if swag.IsZero(m.RequiredCredentials[i]) { // not required
				return nil
			}

			if err := m.RequiredCredentials[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("required_credentials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("required_credentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CrewSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrewSchedule) UnmarshalBinary(b []byte) error {
	var res CrewSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
