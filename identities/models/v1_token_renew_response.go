// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1TokenRenewResponse v1 token renew response
// swagger:model v1TokenRenewResponse
type V1TokenRenewResponse struct {

	// renew token
	RenewToken string `json:"renew_token,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this v1 token renew response
func (m *V1TokenRenewResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1TokenRenewResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TokenRenewResponse) UnmarshalBinary(b []byte) error {
	var res V1TokenRenewResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
