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

// GetAccountReader is a Reader for the GetAccount structure.
type GetAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountOK creates a GetAccountOK with default headers values
func NewGetAccountOK() *GetAccountOK {
	return &GetAccountOK{}
}

/*GetAccountOK handles this case with default header values.

A successful response.
*/
type GetAccountOK struct {
	Payload *models.V1GetAccountResponse
}

func (o *GetAccountOK) Error() string {
	return fmt.Sprintf("[GET /v1/accounts/{account_id}][%d] getAccountOK  %+v", 200, o.Payload)
}

func (o *GetAccountOK) GetPayload() *models.V1GetAccountResponse {
	return o.Payload
}

func (o *GetAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}