// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OtherAssetOtherType other asset other type
// swagger:model OtherAssetOtherType
type OtherAssetOtherType string

const (

	// OtherAssetOtherTypePERSONALEFFECTS captures enum value "PERSONAL_EFFECTS"
	OtherAssetOtherTypePERSONALEFFECTS OtherAssetOtherType = "PERSONAL_EFFECTS"

	// OtherAssetOtherTypeART captures enum value "ART"
	OtherAssetOtherTypeART OtherAssetOtherType = "ART"
)

// for schema
var otherAssetOtherTypeEnum []interface{}

func init() {
	var res []OtherAssetOtherType
	if err := json.Unmarshal([]byte(`["PERSONAL_EFFECTS","ART"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		otherAssetOtherTypeEnum = append(otherAssetOtherTypeEnum, v)
	}
}

func (m OtherAssetOtherType) validateOtherAssetOtherTypeEnum(path, location string, value OtherAssetOtherType) error {
	if err := validate.Enum(path, location, value, otherAssetOtherTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this other asset other type
func (m OtherAssetOtherType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOtherAssetOtherTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
