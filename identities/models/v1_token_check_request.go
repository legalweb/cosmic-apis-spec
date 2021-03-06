// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1TokenCheckRequest v1 token check request
// swagger:model v1TokenCheckRequest
type V1TokenCheckRequest struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this v1 token check request
func (m *V1TokenCheckRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1TokenCheckRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TokenCheckRequest) UnmarshalBinary(b []byte) error {
	var res V1TokenCheckRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
