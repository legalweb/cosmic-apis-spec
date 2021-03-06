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

// CustomerGatewaySetEmploymentReader is a Reader for the CustomerGatewaySetEmployment structure.
type CustomerGatewaySetEmploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerGatewaySetEmploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerGatewaySetEmploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerGatewaySetEmploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerGatewaySetEmploymentOK creates a CustomerGatewaySetEmploymentOK with default headers values
func NewCustomerGatewaySetEmploymentOK() *CustomerGatewaySetEmploymentOK {
	return &CustomerGatewaySetEmploymentOK{}
}

/*CustomerGatewaySetEmploymentOK handles this case with default header values.

A successful response.
*/
type CustomerGatewaySetEmploymentOK struct {
	Payload models.V1SetEmploymentResponse
}

func (o *CustomerGatewaySetEmploymentOK) Error() string {
	return fmt.Sprintf("[PUT /v1/employment][%d] customerGatewaySetEmploymentOK  %+v", 200, o.Payload)
}

func (o *CustomerGatewaySetEmploymentOK) GetPayload() models.V1SetEmploymentResponse {
	return o.Payload
}

func (o *CustomerGatewaySetEmploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerGatewaySetEmploymentDefault creates a CustomerGatewaySetEmploymentDefault with default headers values
func NewCustomerGatewaySetEmploymentDefault(code int) *CustomerGatewaySetEmploymentDefault {
	return &CustomerGatewaySetEmploymentDefault{
		_statusCode: code,
	}
}

/*CustomerGatewaySetEmploymentDefault handles this case with default header values.

An unexpected error response
*/
type CustomerGatewaySetEmploymentDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the customer gateway set employment default response
func (o *CustomerGatewaySetEmploymentDefault) Code() int {
	return o._statusCode
}

func (o *CustomerGatewaySetEmploymentDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/employment][%d] CustomerGateway_SetEmployment default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerGatewaySetEmploymentDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CustomerGatewaySetEmploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
