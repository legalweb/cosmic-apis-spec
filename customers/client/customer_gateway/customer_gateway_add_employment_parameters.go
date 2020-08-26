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

// NewCustomerGatewayAddEmploymentParams creates a new CustomerGatewayAddEmploymentParams object
// with the default values initialized.
func NewCustomerGatewayAddEmploymentParams() *CustomerGatewayAddEmploymentParams {

	return &CustomerGatewayAddEmploymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerGatewayAddEmploymentParamsWithTimeout creates a new CustomerGatewayAddEmploymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerGatewayAddEmploymentParamsWithTimeout(timeout time.Duration) *CustomerGatewayAddEmploymentParams {

	return &CustomerGatewayAddEmploymentParams{

		timeout: timeout,
	}
}

// NewCustomerGatewayAddEmploymentParamsWithContext creates a new CustomerGatewayAddEmploymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerGatewayAddEmploymentParamsWithContext(ctx context.Context) *CustomerGatewayAddEmploymentParams {

	return &CustomerGatewayAddEmploymentParams{

		Context: ctx,
	}
}

// NewCustomerGatewayAddEmploymentParamsWithHTTPClient creates a new CustomerGatewayAddEmploymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerGatewayAddEmploymentParamsWithHTTPClient(client *http.Client) *CustomerGatewayAddEmploymentParams {

	return &CustomerGatewayAddEmploymentParams{
		HTTPClient: client,
	}
}

/*CustomerGatewayAddEmploymentParams contains all the parameters to send to the API endpoint
for the customer gateway add employment operation typically these are written to a http.Request
*/
type CustomerGatewayAddEmploymentParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer gateway add employment params
func (o *CustomerGatewayAddEmploymentParams) WithTimeout(timeout time.Duration) *CustomerGatewayAddEmploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer gateway add employment params
func (o *CustomerGatewayAddEmploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer gateway add employment params
func (o *CustomerGatewayAddEmploymentParams) WithContext(ctx context.Context) *CustomerGatewayAddEmploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer gateway add employment params
func (o *CustomerGatewayAddEmploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer gateway add employment params
func (o *CustomerGatewayAddEmploymentParams) WithHTTPClient(client *http.Client) *CustomerGatewayAddEmploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer gateway add employment params
func (o *CustomerGatewayAddEmploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerGatewayAddEmploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}