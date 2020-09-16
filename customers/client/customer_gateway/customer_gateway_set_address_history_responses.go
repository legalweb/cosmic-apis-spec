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

// CustomerGatewaySetAddressHistoryReader is a Reader for the CustomerGatewaySetAddressHistory structure.
type CustomerGatewaySetAddressHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerGatewaySetAddressHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerGatewaySetAddressHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerGatewaySetAddressHistoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerGatewaySetAddressHistoryOK creates a CustomerGatewaySetAddressHistoryOK with default headers values
func NewCustomerGatewaySetAddressHistoryOK() *CustomerGatewaySetAddressHistoryOK {
	return &CustomerGatewaySetAddressHistoryOK{}
}

/*CustomerGatewaySetAddressHistoryOK handles this case with default header values.

A successful response.
*/
type CustomerGatewaySetAddressHistoryOK struct {
	Payload models.V1SetAddressHistoryResponse
}

func (o *CustomerGatewaySetAddressHistoryOK) Error() string {
	return fmt.Sprintf("[POST /v1/addresses][%d] customerGatewaySetAddressHistoryOK  %+v", 200, o.Payload)
}

func (o *CustomerGatewaySetAddressHistoryOK) GetPayload() models.V1SetAddressHistoryResponse {
	return o.Payload
}

func (o *CustomerGatewaySetAddressHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerGatewaySetAddressHistoryDefault creates a CustomerGatewaySetAddressHistoryDefault with default headers values
func NewCustomerGatewaySetAddressHistoryDefault(code int) *CustomerGatewaySetAddressHistoryDefault {
	return &CustomerGatewaySetAddressHistoryDefault{
		_statusCode: code,
	}
}

/*CustomerGatewaySetAddressHistoryDefault handles this case with default header values.

An unexpected error response
*/
type CustomerGatewaySetAddressHistoryDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the customer gateway set address history default response
func (o *CustomerGatewaySetAddressHistoryDefault) Code() int {
	return o._statusCode
}

func (o *CustomerGatewaySetAddressHistoryDefault) Error() string {
	return fmt.Sprintf("[POST /v1/addresses][%d] CustomerGateway_SetAddressHistory default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerGatewaySetAddressHistoryDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CustomerGatewaySetAddressHistoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}