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

// CustomerGatewayUpdateOwnershipReader is a Reader for the CustomerGatewayUpdateOwnership structure.
type CustomerGatewayUpdateOwnershipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerGatewayUpdateOwnershipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerGatewayUpdateOwnershipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerGatewayUpdateOwnershipDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerGatewayUpdateOwnershipOK creates a CustomerGatewayUpdateOwnershipOK with default headers values
func NewCustomerGatewayUpdateOwnershipOK() *CustomerGatewayUpdateOwnershipOK {
	return &CustomerGatewayUpdateOwnershipOK{}
}

/*CustomerGatewayUpdateOwnershipOK handles this case with default header values.

A successful response.
*/
type CustomerGatewayUpdateOwnershipOK struct {
	Payload models.V1UpdateOwnershipResponse
}

func (o *CustomerGatewayUpdateOwnershipOK) Error() string {
	return fmt.Sprintf("[PUT /v1/ownerships/{ownership_id}][%d] customerGatewayUpdateOwnershipOK  %+v", 200, o.Payload)
}

func (o *CustomerGatewayUpdateOwnershipOK) GetPayload() models.V1UpdateOwnershipResponse {
	return o.Payload
}

func (o *CustomerGatewayUpdateOwnershipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerGatewayUpdateOwnershipDefault creates a CustomerGatewayUpdateOwnershipDefault with default headers values
func NewCustomerGatewayUpdateOwnershipDefault(code int) *CustomerGatewayUpdateOwnershipDefault {
	return &CustomerGatewayUpdateOwnershipDefault{
		_statusCode: code,
	}
}

/*CustomerGatewayUpdateOwnershipDefault handles this case with default header values.

An unexpected error response
*/
type CustomerGatewayUpdateOwnershipDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the customer gateway update ownership default response
func (o *CustomerGatewayUpdateOwnershipDefault) Code() int {
	return o._statusCode
}

func (o *CustomerGatewayUpdateOwnershipDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/ownerships/{ownership_id}][%d] CustomerGateway_UpdateOwnership default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerGatewayUpdateOwnershipDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CustomerGatewayUpdateOwnershipDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
