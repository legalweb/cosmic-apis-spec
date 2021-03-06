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

// NewCustomerGatewayListOwnershipsParams creates a new CustomerGatewayListOwnershipsParams object
// with the default values initialized.
func NewCustomerGatewayListOwnershipsParams() *CustomerGatewayListOwnershipsParams {
	var ()
	return &CustomerGatewayListOwnershipsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerGatewayListOwnershipsParamsWithTimeout creates a new CustomerGatewayListOwnershipsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerGatewayListOwnershipsParamsWithTimeout(timeout time.Duration) *CustomerGatewayListOwnershipsParams {
	var ()
	return &CustomerGatewayListOwnershipsParams{

		timeout: timeout,
	}
}

// NewCustomerGatewayListOwnershipsParamsWithContext creates a new CustomerGatewayListOwnershipsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerGatewayListOwnershipsParamsWithContext(ctx context.Context) *CustomerGatewayListOwnershipsParams {
	var ()
	return &CustomerGatewayListOwnershipsParams{

		Context: ctx,
	}
}

// NewCustomerGatewayListOwnershipsParamsWithHTTPClient creates a new CustomerGatewayListOwnershipsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerGatewayListOwnershipsParamsWithHTTPClient(client *http.Client) *CustomerGatewayListOwnershipsParams {
	var ()
	return &CustomerGatewayListOwnershipsParams{
		HTTPClient: client,
	}
}

/*CustomerGatewayListOwnershipsParams contains all the parameters to send to the API endpoint
for the customer gateway list ownerships operation typically these are written to a http.Request
*/
type CustomerGatewayListOwnershipsParams struct {

	/*BelongsTo*/
	BelongsTo []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer gateway list ownerships params
func (o *CustomerGatewayListOwnershipsParams) WithTimeout(timeout time.Duration) *CustomerGatewayListOwnershipsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer gateway list ownerships params
func (o *CustomerGatewayListOwnershipsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer gateway list ownerships params
func (o *CustomerGatewayListOwnershipsParams) WithContext(ctx context.Context) *CustomerGatewayListOwnershipsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer gateway list ownerships params
func (o *CustomerGatewayListOwnershipsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer gateway list ownerships params
func (o *CustomerGatewayListOwnershipsParams) WithHTTPClient(client *http.Client) *CustomerGatewayListOwnershipsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer gateway list ownerships params
func (o *CustomerGatewayListOwnershipsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBelongsTo adds the belongsTo to the customer gateway list ownerships params
func (o *CustomerGatewayListOwnershipsParams) WithBelongsTo(belongsTo []string) *CustomerGatewayListOwnershipsParams {
	o.SetBelongsTo(belongsTo)
	return o
}

// SetBelongsTo adds the belongsTo to the customer gateway list ownerships params
func (o *CustomerGatewayListOwnershipsParams) SetBelongsTo(belongsTo []string) {
	o.BelongsTo = belongsTo
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerGatewayListOwnershipsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesBelongsTo := o.BelongsTo

	joinedBelongsTo := swag.JoinByFormat(valuesBelongsTo, "multi")
	// query array param belongs_to
	if err := r.SetQueryParam("belongs_to", joinedBelongsTo...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
