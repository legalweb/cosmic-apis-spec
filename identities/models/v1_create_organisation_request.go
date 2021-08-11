// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CreateOrganisationRequest v1 create organisation request
// swagger:model v1CreateOrganisationRequest
type V1CreateOrganisationRequest struct {

	// address
	Address *InternalwktAddress `json:"address,omitempty"`

	// administrator id
	AdministratorID string `json:"administrator_id,omitempty"`

	// email address
	EmailAddress string `json:"email_address,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// organisation id
	OrganisationID string `json:"organisation_id,omitempty"`

	// Contact details
	PhoneNumber string `json:"phone_number,omitempty"`

	// website
	Website string `json:"website,omitempty"`
}

// Validate validates this v1 create organisation request
func (m *V1CreateOrganisationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateOrganisationRequest) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateOrganisationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateOrganisationRequest) UnmarshalBinary(b []byte) error {
	var res V1CreateOrganisationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}