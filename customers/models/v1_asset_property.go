// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AssetProperty v1 asset property
// swagger:model v1AssetProperty
type V1AssetProperty struct {

	// address
	Address *InternalwktAddress `json:"address,omitempty"`

	// beds
	Beds int32 `json:"beds,omitempty"`

	// When true the property type MUST be PROPERTY_TYPE_RESIDENTIAL.
	IsOwnHome bool `json:"is_own_home,omitempty"`

	// subtype
	Subtype string `json:"subtype,omitempty"`

	// type
	Type V1AssetPropertyPropertyType `json:"type,omitempty"`
}

// Validate validates this v1 asset property
func (m *V1AssetProperty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AssetProperty) validateAddress(formats strfmt.Registry) error {

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

func (m *V1AssetProperty) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AssetProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AssetProperty) UnmarshalBinary(b []byte) error {
	var res V1AssetProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
