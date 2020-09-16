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

// NewCustomerGatewayGetProfileParams creates a new CustomerGatewayGetProfileParams object
// with the default values initialized.
func NewCustomerGatewayGetProfileParams() *CustomerGatewayGetProfileParams {

	return &CustomerGatewayGetProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerGatewayGetProfileParamsWithTimeout creates a new CustomerGatewayGetProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerGatewayGetProfileParamsWithTimeout(timeout time.Duration) *CustomerGatewayGetProfileParams {

	return &CustomerGatewayGetProfileParams{

		timeout: timeout,
	}
}

// NewCustomerGatewayGetProfileParamsWithContext creates a new CustomerGatewayGetProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerGatewayGetProfileParamsWithContext(ctx context.Context) *CustomerGatewayGetProfileParams {

	return &CustomerGatewayGetProfileParams{

		Context: ctx,
	}
}

// NewCustomerGatewayGetProfileParamsWithHTTPClient creates a new CustomerGatewayGetProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerGatewayGetProfileParamsWithHTTPClient(client *http.Client) *CustomerGatewayGetProfileParams {

	return &CustomerGatewayGetProfileParams{
		HTTPClient: client,
	}
}

/*CustomerGatewayGetProfileParams contains all the parameters to send to the API endpoint
for the customer gateway get profile operation typically these are written to a http.Request
*/
type CustomerGatewayGetProfileParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer gateway get profile params
func (o *CustomerGatewayGetProfileParams) WithTimeout(timeout time.Duration) *CustomerGatewayGetProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer gateway get profile params
func (o *CustomerGatewayGetProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer gateway get profile params
func (o *CustomerGatewayGetProfileParams) WithContext(ctx context.Context) *CustomerGatewayGetProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer gateway get profile params
func (o *CustomerGatewayGetProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer gateway get profile params
func (o *CustomerGatewayGetProfileParams) WithHTTPClient(client *http.Client) *CustomerGatewayGetProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer gateway get profile params
func (o *CustomerGatewayGetProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerGatewayGetProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
