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

// IdentitiesSendResetPasswordReader is a Reader for the IdentitiesSendResetPassword structure.
type IdentitiesSendResetPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentitiesSendResetPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentitiesSendResetPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIdentitiesSendResetPasswordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentitiesSendResetPasswordOK creates a IdentitiesSendResetPasswordOK with default headers values
func NewIdentitiesSendResetPasswordOK() *IdentitiesSendResetPasswordOK {
	return &IdentitiesSendResetPasswordOK{}
}

/*IdentitiesSendResetPasswordOK handles this case with default header values.

A successful response.
*/
type IdentitiesSendResetPasswordOK struct {
	Payload models.V1SendResetPasswordResponse
}

func (o *IdentitiesSendResetPasswordOK) Error() string {
	return fmt.Sprintf("[POST /v1/users/send-reset][%d] identitiesSendResetPasswordOK  %+v", 200, o.Payload)
}

func (o *IdentitiesSendResetPasswordOK) GetPayload() models.V1SendResetPasswordResponse {
	return o.Payload
}

func (o *IdentitiesSendResetPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentitiesSendResetPasswordDefault creates a IdentitiesSendResetPasswordDefault with default headers values
func NewIdentitiesSendResetPasswordDefault(code int) *IdentitiesSendResetPasswordDefault {
	return &IdentitiesSendResetPasswordDefault{
		_statusCode: code,
	}
}

/*IdentitiesSendResetPasswordDefault handles this case with default header values.

An unexpected error response
*/
type IdentitiesSendResetPasswordDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the identities send reset password default response
func (o *IdentitiesSendResetPasswordDefault) Code() int {
	return o._statusCode
}

func (o *IdentitiesSendResetPasswordDefault) Error() string {
	return fmt.Sprintf("[POST /v1/users/send-reset][%d] Identities_SendResetPassword default  %+v", o._statusCode, o.Payload)
}

func (o *IdentitiesSendResetPasswordDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *IdentitiesSendResetPasswordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
