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

// CustomerGatewayAddEmploymentReader is a Reader for the CustomerGatewayAddEmployment structure.
type CustomerGatewayAddEmploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerGatewayAddEmploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerGatewayAddEmploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerGatewayAddEmploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerGatewayAddEmploymentOK creates a CustomerGatewayAddEmploymentOK with default headers values
func NewCustomerGatewayAddEmploymentOK() *CustomerGatewayAddEmploymentOK {
	return &CustomerGatewayAddEmploymentOK{}
}

/*CustomerGatewayAddEmploymentOK handles this case with default header values.

A successful response.
*/
type CustomerGatewayAddEmploymentOK struct {
	Payload *models.V1AddEmploymentResponse
}

func (o *CustomerGatewayAddEmploymentOK) Error() string {
	return fmt.Sprintf("[POST /v1/employment][%d] customerGatewayAddEmploymentOK  %+v", 200, o.Payload)
}

func (o *CustomerGatewayAddEmploymentOK) GetPayload() *models.V1AddEmploymentResponse {
	return o.Payload
}

func (o *CustomerGatewayAddEmploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AddEmploymentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerGatewayAddEmploymentDefault creates a CustomerGatewayAddEmploymentDefault with default headers values
func NewCustomerGatewayAddEmploymentDefault(code int) *CustomerGatewayAddEmploymentDefault {
	return &CustomerGatewayAddEmploymentDefault{
		_statusCode: code,
	}
}

/*CustomerGatewayAddEmploymentDefault handles this case with default header values.

An unexpected error response
*/
type CustomerGatewayAddEmploymentDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the customer gateway add employment default response
func (o *CustomerGatewayAddEmploymentDefault) Code() int {
	return o._statusCode
}

func (o *CustomerGatewayAddEmploymentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/employment][%d] CustomerGateway_AddEmployment default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerGatewayAddEmploymentDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CustomerGatewayAddEmploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}