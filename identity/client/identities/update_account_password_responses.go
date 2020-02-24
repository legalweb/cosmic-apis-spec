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

// UpdateAccountPasswordReader is a Reader for the UpdateAccountPassword structure.
type UpdateAccountPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAccountPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAccountPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAccountPasswordOK creates a UpdateAccountPasswordOK with default headers values
func NewUpdateAccountPasswordOK() *UpdateAccountPasswordOK {
	return &UpdateAccountPasswordOK{}
}

/*UpdateAccountPasswordOK handles this case with default header values.

A successful response.
*/
type UpdateAccountPasswordOK struct {
	Payload models.V1UpdateAccountPasswordResponse
}

func (o *UpdateAccountPasswordOK) Error() string {
	return fmt.Sprintf("[POST /v1/accounts/{account_id}/change_password][%d] updateAccountPasswordOK  %+v", 200, o.Payload)
}

func (o *UpdateAccountPasswordOK) GetPayload() models.V1UpdateAccountPasswordResponse {
	return o.Payload
}

func (o *UpdateAccountPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}