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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCustomerGatewayListDocumentsParams creates a new CustomerGatewayListDocumentsParams object
// with the default values initialized.
func NewCustomerGatewayListDocumentsParams() *CustomerGatewayListDocumentsParams {
	var ()
	return &CustomerGatewayListDocumentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerGatewayListDocumentsParamsWithTimeout creates a new CustomerGatewayListDocumentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerGatewayListDocumentsParamsWithTimeout(timeout time.Duration) *CustomerGatewayListDocumentsParams {
	var ()
	return &CustomerGatewayListDocumentsParams{

		timeout: timeout,
	}
}

// NewCustomerGatewayListDocumentsParamsWithContext creates a new CustomerGatewayListDocumentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerGatewayListDocumentsParamsWithContext(ctx context.Context) *CustomerGatewayListDocumentsParams {
	var ()
	return &CustomerGatewayListDocumentsParams{

		Context: ctx,
	}
}

// NewCustomerGatewayListDocumentsParamsWithHTTPClient creates a new CustomerGatewayListDocumentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerGatewayListDocumentsParamsWithHTTPClient(client *http.Client) *CustomerGatewayListDocumentsParams {
	var ()
	return &CustomerGatewayListDocumentsParams{
		HTTPClient: client,
	}
}

/*CustomerGatewayListDocumentsParams contains all the parameters to send to the API endpoint
for the customer gateway list documents operation typically these are written to a http.Request
*/
type CustomerGatewayListDocumentsParams struct {

	/*ApplicationName
	  Filter down to document folders only relevant for the application
	name provided.

	*/
	ApplicationName []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer gateway list documents params
func (o *CustomerGatewayListDocumentsParams) WithTimeout(timeout time.Duration) *CustomerGatewayListDocumentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer gateway list documents params
func (o *CustomerGatewayListDocumentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer gateway list documents params
func (o *CustomerGatewayListDocumentsParams) WithContext(ctx context.Context) *CustomerGatewayListDocumentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer gateway list documents params
func (o *CustomerGatewayListDocumentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer gateway list documents params
func (o *CustomerGatewayListDocumentsParams) WithHTTPClient(client *http.Client) *CustomerGatewayListDocumentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer gateway list documents params
func (o *CustomerGatewayListDocumentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationName adds the applicationName to the customer gateway list documents params
func (o *CustomerGatewayListDocumentsParams) WithApplicationName(applicationName []string) *CustomerGatewayListDocumentsParams {
	o.SetApplicationName(applicationName)
	return o
}

// SetApplicationName adds the applicationName to the customer gateway list documents params
func (o *CustomerGatewayListDocumentsParams) SetApplicationName(applicationName []string) {
	o.ApplicationName = applicationName
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerGatewayListDocumentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesApplicationName := o.ApplicationName

	joinedApplicationName := swag.JoinByFormat(valuesApplicationName, "multi")
	// query array param application_name
	if err := r.SetQueryParam("application_name", joinedApplicationName...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
