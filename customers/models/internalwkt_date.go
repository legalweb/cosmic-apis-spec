// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InternalwktDate internalwkt date
// swagger:model internalwktDate
type InternalwktDate struct {

	// day
	Day int32 `json:"day,omitempty"`

	// month
	Month int32 `json:"month,omitempty"`

	// year
	Year int32 `json:"year,omitempty"`
}

// Validate validates this internalwkt date
func (m *InternalwktDate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InternalwktDate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InternalwktDate) UnmarshalBinary(b []byte) error {
	var res InternalwktDate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}