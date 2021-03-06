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

// IdentitiesUpdateAccountPasswordReader is a Reader for the IdentitiesUpdateAccountPassword structure.
type IdentitiesUpdateAccountPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentitiesUpdateAccountPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentitiesUpdateAccountPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIdentitiesUpdateAccountPasswordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentitiesUpdateAccountPasswordOK creates a IdentitiesUpdateAccountPasswordOK with default headers values
func NewIdentitiesUpdateAccountPasswordOK() *IdentitiesUpdateAccountPasswordOK {
	return &IdentitiesUpdateAccountPasswordOK{}
}

/*IdentitiesUpdateAccountPasswordOK handles this case with default header values.

A successful response.
*/
type IdentitiesUpdateAccountPasswordOK struct {
	Payload models.V1UpdateAccountPasswordResponse
}

func (o *IdentitiesUpdateAccountPasswordOK) Error() string {
	return fmt.Sprintf("[POST /v1/accounts/{account_id}/change_password][%d] identitiesUpdateAccountPasswordOK  %+v", 200, o.Payload)
}

func (o *IdentitiesUpdateAccountPasswordOK) GetPayload() models.V1UpdateAccountPasswordResponse {
	return o.Payload
}

func (o *IdentitiesUpdateAccountPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentitiesUpdateAccountPasswordDefault creates a IdentitiesUpdateAccountPasswordDefault with default headers values
func NewIdentitiesUpdateAccountPasswordDefault(code int) *IdentitiesUpdateAccountPasswordDefault {
	return &IdentitiesUpdateAccountPasswordDefault{
		_statusCode: code,
	}
}

/*IdentitiesUpdateAccountPasswordDefault handles this case with default header values.

An unexpected error response
*/
type IdentitiesUpdateAccountPasswordDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the identities update account password default response
func (o *IdentitiesUpdateAccountPasswordDefault) Code() int {
	return o._statusCode
}

func (o *IdentitiesUpdateAccountPasswordDefault) Error() string {
	return fmt.Sprintf("[POST /v1/accounts/{account_id}/change_password][%d] Identities_UpdateAccountPassword default  %+v", o._statusCode, o.Payload)
}

func (o *IdentitiesUpdateAccountPasswordDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IdentitiesUpdateAccountPasswordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
