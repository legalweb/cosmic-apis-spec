// Code generated by go-swagger; DO NOT EDIT.

package customer_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"lwebco.de/cosmic-apis-spec/customers/models"
)

// CustomerGatewayGetProfileReader is a Reader for the CustomerGatewayGetProfile structure.
type CustomerGatewayGetProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerGatewayGetProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerGatewayGetProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerGatewayGetProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerGatewayGetProfileOK creates a CustomerGatewayGetProfileOK with default headers values
func NewCustomerGatewayGetProfileOK() *CustomerGatewayGetProfileOK {
	return &CustomerGatewayGetProfileOK{}
}

/*CustomerGatewayGetProfileOK handles this case with default header values.

A successful response.
*/
type CustomerGatewayGetProfileOK struct {
	Payload *models.V1GetProfileResponse
}

func (o *CustomerGatewayGetProfileOK) Error() string {
	return fmt.Sprintf("[GET /v1/profile][%d] customerGatewayGetProfileOK  %+v", 200, o.Payload)
}

func (o *CustomerGatewayGetProfileOK) GetPayload() *models.V1GetProfileResponse {
	return o.Payload
}

func (o *CustomerGatewayGetProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetProfileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerGatewayGetProfileDefault creates a CustomerGatewayGetProfileDefault with default headers values
func NewCustomerGatewayGetProfileDefault(code int) *CustomerGatewayGetProfileDefault {
	return &CustomerGatewayGetProfileDefault{
		_statusCode: code,
	}
}

/*CustomerGatewayGetProfileDefault handles this case with default header values.

An unexpected error response
*/
type CustomerGatewayGetProfileDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the customer gateway get profile default response
func (o *CustomerGatewayGetProfileDefault) Code() int {
	return o._statusCode
}

func (o *CustomerGatewayGetProfileDefault) Error() string {
	return fmt.Sprintf("[GET /v1/profile][%d] CustomerGateway_GetProfile default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerGatewayGetProfileDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CustomerGatewayGetProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
