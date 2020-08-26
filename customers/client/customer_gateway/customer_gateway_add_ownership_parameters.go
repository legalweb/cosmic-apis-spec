// Code generated by go-swagger; DO NOT EDIT.

package customer_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCustomerGatewayAddOwnershipParams creates a new CustomerGatewayAddOwnershipParams object
// with the default values initialized.
func NewCustomerGatewayAddOwnershipParams() *CustomerGatewayAddOwnershipParams {

	return &CustomerGatewayAddOwnershipParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerGatewayAddOwnershipParamsWithTimeout creates a new CustomerGatewayAddOwnershipParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerGatewayAddOwnershipParamsWithTimeout(timeout time.Duration) *CustomerGatewayAddOwnershipParams {

	return &CustomerGatewayAddOwnershipParams{

		timeout: timeout,
	}
}

// NewCustomerGatewayAddOwnershipParamsWithContext creates a new CustomerGatewayAddOwnershipParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerGatewayAddOwnershipParamsWithContext(ctx context.Context) *CustomerGatewayAddOwnershipParams {

	return &CustomerGatewayAddOwnershipParams{

		Context: ctx,
	}
}

// NewCustomerGatewayAddOwnershipParamsWithHTTPClient creates a new CustomerGatewayAddOwnershipParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerGatewayAddOwnershipParamsWithHTTPClient(client *http.Client) *CustomerGatewayAddOwnershipParams {

	return &CustomerGatewayAddOwnershipParams{
		HTTPClient: client,
	}
}

/*CustomerGatewayAddOwnershipParams contains all the parameters to send to the API endpoint
for the customer gateway add ownership operation typically these are written to a http.Request
*/
type CustomerGatewayAddOwnershipParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer gateway add ownership params
func (o *CustomerGatewayAddOwnershipParams) WithTimeout(timeout time.Duration) *CustomerGatewayAddOwnershipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer gateway add ownership params
func (o *CustomerGatewayAddOwnershipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer gateway add ownership params
func (o *CustomerGatewayAddOwnershipParams) WithContext(ctx context.Context) *CustomerGatewayAddOwnershipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer gateway add ownership params
func (o *CustomerGatewayAddOwnershipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer gateway add ownership params
func (o *CustomerGatewayAddOwnershipParams) WithHTTPClient(client *http.Client) *CustomerGatewayAddOwnershipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer gateway add ownership params
func (o *CustomerGatewayAddOwnershipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerGatewayAddOwnershipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}