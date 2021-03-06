// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InternalwktPaginationResponse internalwkt pagination response
// swagger:model internalwktPaginationResponse
type InternalwktPaginationResponse struct {

	// current page
	CurrentPage int64 `json:"currentPage,omitempty"`

	// limit
	Limit int64 `json:"limit,omitempty"`

	// total pages
	TotalPages int64 `json:"totalPages,omitempty"`
}

// Validate validates this internalwkt pagination response
func (m *InternalwktPaginationResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InternalwktPaginationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InternalwktPaginationResponse) UnmarshalBinary(b []byte) error {
	var res InternalwktPaginationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
