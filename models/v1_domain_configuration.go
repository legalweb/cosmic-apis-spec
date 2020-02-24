// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1DomainConfiguration v1 domain configuration
// swagger:model v1DomainConfiguration
type V1DomainConfiguration struct {

	// The domain name which the configuration relates too.
	DomainName string `json:"domain_name,omitempty"`

	// The label is the organisation customer friendly name for example
	// "Commercial Finance Group".
	Label string `json:"label,omitempty"`

	// A URI to the logo the customers will see for the organisation.
	LogoURI string `json:"logo_uri,omitempty"`

	// The organisation ID related to the domain configuration.
	// This should be used on subsequent requests to authentication and any
	// other service that requires the organisation id.
	OrganisationID string `json:"organisation_id,omitempty"`

	// An e-mail address that can be used for customers to get support.
	SupportEmail []string `json:"support_email"`

	// A telephone number that can be used for customers to get support.
	SupportTelephone []string `json:"support_telephone"`
}

// Validate validates this v1 domain configuration
func (m *V1DomainConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DomainConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DomainConfiguration) UnmarshalBinary(b []byte) error {
	var res V1DomainConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}