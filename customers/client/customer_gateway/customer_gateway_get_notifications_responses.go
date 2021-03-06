// Code generated by go-swagger; DO NOT EDIT.

package customer_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"lwebco.de/cosmic-apis-spec/customers/models"
)

// CustomerGatewayGetNotificationsReader is a Reader for the CustomerGatewayGetNotifications structure.
type CustomerGatewayGetNotificationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerGatewayGetNotificationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerGatewayGetNotificationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerGatewayGetNotificationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerGatewayGetNotificationsOK creates a CustomerGatewayGetNotificationsOK with default headers values
func NewCustomerGatewayGetNotificationsOK() *CustomerGatewayGetNotificationsOK {
	return &CustomerGatewayGetNotificationsOK{}
}

/*CustomerGatewayGetNotificationsOK handles this case with default header values.

A successful response.(streaming responses)
*/
type CustomerGatewayGetNotificationsOK struct {
	Payload *CustomerGatewayGetNotificationsOKBody
}

func (o *CustomerGatewayGetNotificationsOK) Error() string {
	return fmt.Sprintf("[GET /v1/notifications][%d] customerGatewayGetNotificationsOK  %+v", 200, o.Payload)
}

func (o *CustomerGatewayGetNotificationsOK) GetPayload() *CustomerGatewayGetNotificationsOKBody {
	return o.Payload
}

func (o *CustomerGatewayGetNotificationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CustomerGatewayGetNotificationsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerGatewayGetNotificationsDefault creates a CustomerGatewayGetNotificationsDefault with default headers values
func NewCustomerGatewayGetNotificationsDefault(code int) *CustomerGatewayGetNotificationsDefault {
	return &CustomerGatewayGetNotificationsDefault{
		_statusCode: code,
	}
}

/*CustomerGatewayGetNotificationsDefault handles this case with default header values.

An unexpected error response
*/
type CustomerGatewayGetNotificationsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the customer gateway get notifications default response
func (o *CustomerGatewayGetNotificationsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerGatewayGetNotificationsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/notifications][%d] CustomerGateway_GetNotifications default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerGatewayGetNotificationsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CustomerGatewayGetNotificationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CustomerGatewayGetNotificationsOKBody Stream result of v1Notification
swagger:model CustomerGatewayGetNotificationsOKBody
*/
type CustomerGatewayGetNotificationsOKBody struct {

	// error
	Error *models.RuntimeStreamError `json:"error,omitempty"`

	// result
	Result *models.V1Notification `json:"result,omitempty"`
}

// Validate validates this customer gateway get notifications o k body
func (o *CustomerGatewayGetNotificationsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CustomerGatewayGetNotificationsOKBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerGatewayGetNotificationsOK" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *CustomerGatewayGetNotificationsOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerGatewayGetNotificationsOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CustomerGatewayGetNotificationsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CustomerGatewayGetNotificationsOKBody) UnmarshalBinary(b []byte) error {
	var res CustomerGatewayGetNotificationsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
