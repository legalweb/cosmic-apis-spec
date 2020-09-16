// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CreateUploadDocumentLinkRequest v1 create upload document link request
// swagger:model v1CreateUploadDocumentLinkRequest
type V1CreateUploadDocumentLinkRequest struct {

	// collection id
	CollectionID string `json:"collection_id,omitempty"`

	// label
	Label string `json:"label,omitempty"`
}

// Validate validates this v1 create upload document link request
func (m *V1CreateUploadDocumentLinkRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateUploadDocumentLinkRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateUploadDocumentLinkRequest) UnmarshalBinary(b []byte) error {
	var res V1CreateUploadDocumentLinkRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}