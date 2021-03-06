// Code generated by go-swagger; DO NOT EDIT.

package identities

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

	"lwebco.de/cosmic-apis-spec/identities/models"
)

// NewIdentitiesUpdateAccountDetailsParams creates a new IdentitiesUpdateAccountDetailsParams object
// with the default values initialized.
func NewIdentitiesUpdateAccountDetailsParams() *IdentitiesUpdateAccountDetailsParams {
	var ()
	return &IdentitiesUpdateAccountDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIdentitiesUpdateAccountDetailsParamsWithTimeout creates a new IdentitiesUpdateAccountDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIdentitiesUpdateAccountDetailsParamsWithTimeout(timeout time.Duration) *IdentitiesUpdateAccountDetailsParams {
	var ()
	return &IdentitiesUpdateAccountDetailsParams{

		timeout: timeout,
	}
}

// NewIdentitiesUpdateAccountDetailsParamsWithContext creates a new IdentitiesUpdateAccountDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewIdentitiesUpdateAccountDetailsParamsWithContext(ctx context.Context) *IdentitiesUpdateAccountDetailsParams {
	var ()
	return &IdentitiesUpdateAccountDetailsParams{

		Context: ctx,
	}
}

// NewIdentitiesUpdateAccountDetailsParamsWithHTTPClient creates a new IdentitiesUpdateAccountDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIdentitiesUpdateAccountDetailsParamsWithHTTPClient(client *http.Client) *IdentitiesUpdateAccountDetailsParams {
	var ()
	return &IdentitiesUpdateAccountDetailsParams{
		HTTPClient: client,
	}
}

/*IdentitiesUpdateAccountDetailsParams contains all the parameters to send to the API endpoint
for the identities update account details operation typically these are written to a http.Request
*/
type IdentitiesUpdateAccountDetailsParams struct {

	/*AccountID*/
	AccountID string
	/*Body*/
	Body *models.V1UpdateAccountDetailsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the identities update account details params
func (o *IdentitiesUpdateAccountDetailsParams) WithTimeout(timeout time.Duration) *IdentitiesUpdateAccountDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identities update account details params
func (o *IdentitiesUpdateAccountDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identities update account details params
func (o *IdentitiesUpdateAccountDetailsParams) WithContext(ctx context.Context) *IdentitiesUpdateAccountDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identities update account details params
func (o *IdentitiesUpdateAccountDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identities update account details params
func (o *IdentitiesUpdateAccountDetailsParams) WithHTTPClient(client *http.Client) *IdentitiesUpdateAccountDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identities update account details params
func (o *IdentitiesUpdateAccountDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the identities update account details params
func (o *IdentitiesUpdateAccountDetailsParams) WithAccountID(accountID string) *IdentitiesUpdateAccountDetailsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the identities update account details params
func (o *IdentitiesUpdateAccountDetailsParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithBody adds the body to the identities update account details params
func (o *IdentitiesUpdateAccountDetailsParams) WithBody(body *models.V1UpdateAccountDetailsRequest) *IdentitiesUpdateAccountDetailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identities update account details params
func (o *IdentitiesUpdateAccountDetailsParams) SetBody(body *models.V1UpdateAccountDetailsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IdentitiesUpdateAccountDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_id
	if err := r.SetPathParam("account_id", o.AccountID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
