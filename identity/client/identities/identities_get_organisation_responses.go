// Code generated by go-swagger; DO NOT EDIT.

package identities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"lwebco.de/cosmic-apis-spec/identity/models"
)

// IdentitiesGetOrganisationReader is a Reader for the IdentitiesGetOrganisation structure.
type IdentitiesGetOrganisationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentitiesGetOrganisationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentitiesGetOrganisationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIdentitiesGetOrganisationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentitiesGetOrganisationOK creates a IdentitiesGetOrganisationOK with default headers values
func NewIdentitiesGetOrganisationOK() *IdentitiesGetOrganisationOK {
	return &IdentitiesGetOrganisationOK{}
}

/*IdentitiesGetOrganisationOK handles this case with default header values.

A successful response.
*/
type IdentitiesGetOrganisationOK struct {
	Payload *models.V1GetOrganisationResponse
}

func (o *IdentitiesGetOrganisationOK) Error() string {
	return fmt.Sprintf("[GET /v1/organisations/{organisation_id}][%d] identitiesGetOrganisationOK  %+v", 200, o.Payload)
}

func (o *IdentitiesGetOrganisationOK) GetPayload() *models.V1GetOrganisationResponse {
	return o.Payload
}

func (o *IdentitiesGetOrganisationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetOrganisationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentitiesGetOrganisationDefault creates a IdentitiesGetOrganisationDefault with default headers values
func NewIdentitiesGetOrganisationDefault(code int) *IdentitiesGetOrganisationDefault {
	return &IdentitiesGetOrganisationDefault{
		_statusCode: code,
	}
}

/*IdentitiesGetOrganisationDefault handles this case with default header values.

An unexpected error response
*/
type IdentitiesGetOrganisationDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the identities get organisation default response
func (o *IdentitiesGetOrganisationDefault) Code() int {
	return o._statusCode
}

func (o *IdentitiesGetOrganisationDefault) Error() string {
	return fmt.Sprintf("[GET /v1/organisations/{organisation_id}][%d] Identities_GetOrganisation default  %+v", o._statusCode, o.Payload)
}

func (o *IdentitiesGetOrganisationDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IdentitiesGetOrganisationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
