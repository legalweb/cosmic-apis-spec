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

// TokenRevokeReader is a Reader for the TokenRevoke structure.
type TokenRevokeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokenRevokeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTokenRevokeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTokenRevokeOK creates a TokenRevokeOK with default headers values
func NewTokenRevokeOK() *TokenRevokeOK {
	return &TokenRevokeOK{}
}

/*TokenRevokeOK handles this case with default header values.

A successful response.
*/
type TokenRevokeOK struct {
	Payload models.V1TokenRevokeResponse
}

func (o *TokenRevokeOK) Error() string {
	return fmt.Sprintf("[POST /v1/token/revoke][%d] tokenRevokeOK  %+v", 200, o.Payload)
}

func (o *TokenRevokeOK) GetPayload() models.V1TokenRevokeResponse {
	return o.Payload
}

func (o *TokenRevokeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}