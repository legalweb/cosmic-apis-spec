// Code generated by go-swagger; DO NOT EDIT.

package identities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"lwebco.de/cosmic-apis-spec/identities/models"
)

// IdentitiesCreateAccountReader is a Reader for the IdentitiesCreateAccount structure.
type IdentitiesCreateAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentitiesCreateAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentitiesCreateAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIdentitiesCreateAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentitiesCreateAccountOK creates a IdentitiesCreateAccountOK with default headers values
func NewIdentitiesCreateAccountOK() *IdentitiesCreateAccountOK {
	return &IdentitiesCreateAccountOK{}
}

/*IdentitiesCreateAccountOK handles this case with default header values.

A successful response.
*/
type IdentitiesCreateAccountOK struct {
	Payload models.V1CreateAccountResponse
}

func (o *IdentitiesCreateAccountOK) Error() string {
	return fmt.Sprintf("[POST /v1/accounts][%d] identitiesCreateAccountOK  %+v", 200, o.Payload)
}

func (o *IdentitiesCreateAccountOK) GetPayload() models.V1CreateAccountResponse {
	return o.Payload
}

func (o *IdentitiesCreateAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentitiesCreateAccountDefault creates a IdentitiesCreateAccountDefault with default headers values
func NewIdentitiesCreateAccountDefault(code int) *IdentitiesCreateAccountDefault {
	return &IdentitiesCreateAccountDefault{
		_statusCode: code,
	}
}

/*IdentitiesCreateAccountDefault handles this case with default header values.

An unexpected error response
*/
type IdentitiesCreateAccountDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the identities create account default response
func (o *IdentitiesCreateAccountDefault) Code() int {
	return o._statusCode
}

func (o *IdentitiesCreateAccountDefault) Error() string {
	return fmt.Sprintf("[POST /v1/accounts][%d] Identities_CreateAccount default  %+v", o._statusCode, o.Payload)
}

func (o *IdentitiesCreateAccountDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IdentitiesCreateAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
