// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Account v1 account
// swagger:model v1Account
type V1Account struct {

	// account id
	AccountID string `json:"account_id,omitempty"`

	// active
	Active bool `json:"active,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// email address
	EmailAddress string `json:"email_address,omitempty"`

	// forename
	Forename string `json:"forename,omitempty"`

	// surname
	Surname string `json:"surname,omitempty"`

	// telephone
	Telephone string `json:"telephone,omitempty"`
}

// Validate validates this v1 account
func (m *V1Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Account) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Account) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Account) UnmarshalBinary(b []byte) error {
	var res V1Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
