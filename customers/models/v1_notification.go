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

// V1Notification v1 notification
// swagger:model v1Notification
type V1Notification struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// The name of the notification.
	Name string `json:"name,omitempty"`

	// The type of notification.
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 notification
func (m *V1Notification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Notification) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Notification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Notification) UnmarshalBinary(b []byte) error {
	var res V1Notification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
