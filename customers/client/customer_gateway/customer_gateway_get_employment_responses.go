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

// CustomerGatewayGetEmploymentReader is a Reader for the CustomerGatewayGetEmployment structure.
type CustomerGatewayGetEmploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerGatewayGetEmploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerGatewayGetEmploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerGatewayGetEmploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerGatewayGetEmploymentOK creates a CustomerGatewayGetEmploymentOK with default headers values
func NewCustomerGatewayGetEmploymentOK() *CustomerGatewayGetEmploymentOK {
	return &CustomerGatewayGetEmploymentOK{}
}

/*CustomerGatewayGetEmploymentOK handles this case with default header values.

A successful response.
*/
type CustomerGatewayGetEmploymentOK struct {
	Payload *models.V1GetEmploymentResponse
}

func (o *CustomerGatewayGetEmploymentOK) Error() string {
	return fmt.Sprintf("[GET /v1/employment][%d] customerGatewayGetEmploymentOK  %+v", 200, o.Payload)
}

func (o *CustomerGatewayGetEmploymentOK) GetPayload() *models.V1GetEmploymentResponse {
	return o.Payload
}

func (o *CustomerGatewayGetEmploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetEmploymentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerGatewayGetEmploymentDefault creates a CustomerGatewayGetEmploymentDefault with default headers values
func NewCustomerGatewayGetEmploymentDefault(code int) *CustomerGatewayGetEmploymentDefault {
	return &CustomerGatewayGetEmploymentDefault{
		_statusCode: code,
	}
}

/*CustomerGatewayGetEmploymentDefault handles this case with default header values.

An unexpected error response
*/
type CustomerGatewayGetEmploymentDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the customer gateway get employment default response
func (o *CustomerGatewayGetEmploymentDefault) Code() int {
	return o._statusCode
}

func (o *CustomerGatewayGetEmploymentDefault) Error() string {
	return fmt.Sprintf("[GET /v1/employment][%d] CustomerGateway_GetEmployment default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerGatewayGetEmploymentDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CustomerGatewayGetEmploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
