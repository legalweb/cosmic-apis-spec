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

// PensionPensionType pension pension type
// swagger:model PensionPensionType
type PensionPensionType string

const (

	// PensionPensionTypeSIPP captures enum value "SIPP"
	PensionPensionTypeSIPP PensionPensionType = "SIPP"

	// PensionPensionTypeSSAS captures enum value "SSAS"
	PensionPensionTypeSSAS PensionPensionType = "SSAS"

	// PensionPensionTypeWORKPLACEPENSION captures enum value "WORKPLACE_PENSION"
	PensionPensionTypeWORKPLACEPENSION PensionPensionType = "WORKPLACE_PENSION"
)

// for schema
var pensionPensionTypeEnum []interface{}

func init() {
	var res []PensionPensionType
	if err := json.Unmarshal([]byte(`["SIPP","SSAS","WORKPLACE_PENSION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pensionPensionTypeEnum = append(pensionPensionTypeEnum, v)
	}
}

func (m PensionPensionType) validatePensionPensionTypeEnum(path, location string, value PensionPensionType) error {
	if err := validate.Enum(path, location, value, pensionPensionTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this pension pension type
func (m PensionPensionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePensionPensionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
