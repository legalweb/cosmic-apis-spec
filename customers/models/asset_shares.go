// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AssetShares asset shares
// swagger:model AssetShares
type AssetShares struct {

	// company name
	CompanyName string `json:"company_name,omitempty"`

	// Is this is false we assume is it a public limited company
	// otherwise it is a private limited company.
	IsLimitedCompanyPrivate bool `json:"is_limited_company_private,omitempty"`
}

// Validate validates this asset shares
func (m *AssetShares) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AssetShares) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetShares) UnmarshalBinary(b []byte) error {
	var res AssetShares
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}