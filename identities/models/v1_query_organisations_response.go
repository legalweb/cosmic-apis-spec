// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1QueryOrganisationsResponse v1 query organisations response
// swagger:model v1QueryOrganisationsResponse
type V1QueryOrganisationsResponse struct {

	// organisations
	Organisations []*V1Organisation `json:"organisations"`

	// pagination
	Pagination *InternalwktPaginationResponse `json:"pagination,omitempty"`
}

// Validate validates this v1 query organisations response
func (m *V1QueryOrganisationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganisations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1QueryOrganisationsResponse) validateOrganisations(formats strfmt.Registry) error {

	if swag.IsZero(m.Organisations) { // not required
		return nil
	}

	for i := 0; i < len(m.Organisations); i++ {
		if swag.IsZero(m.Organisations[i]) { // not required
			continue
		}

		if m.Organisations[i] != nil {
			if err := m.Organisations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organisations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1QueryOrganisationsResponse) validatePagination(formats strfmt.Registry) error {

	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1QueryOrganisationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1QueryOrganisationsResponse) UnmarshalBinary(b []byte) error {
	var res V1QueryOrganisationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
