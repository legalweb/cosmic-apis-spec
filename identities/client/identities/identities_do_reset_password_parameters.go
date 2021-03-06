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

// NewIdentitiesDoResetPasswordParams creates a new IdentitiesDoResetPasswordParams object
// with the default values initialized.
func NewIdentitiesDoResetPasswordParams() *IdentitiesDoResetPasswordParams {
	var ()
	return &IdentitiesDoResetPasswordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIdentitiesDoResetPasswordParamsWithTimeout creates a new IdentitiesDoResetPasswordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIdentitiesDoResetPasswordParamsWithTimeout(timeout time.Duration) *IdentitiesDoResetPasswordParams {
	var ()
	return &IdentitiesDoResetPasswordParams{

		timeout: timeout,
	}
}

// NewIdentitiesDoResetPasswordParamsWithContext creates a new IdentitiesDoResetPasswordParams object
// with the default values initialized, and the ability to set a context for a request
func NewIdentitiesDoResetPasswordParamsWithContext(ctx context.Context) *IdentitiesDoResetPasswordParams {
	var ()
	return &IdentitiesDoResetPasswordParams{

		Context: ctx,
	}
}

// NewIdentitiesDoResetPasswordParamsWithHTTPClient creates a new IdentitiesDoResetPasswordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIdentitiesDoResetPasswordParamsWithHTTPClient(client *http.Client) *IdentitiesDoResetPasswordParams {
	var ()
	return &IdentitiesDoResetPasswordParams{
		HTTPClient: client,
	}
}

/*IdentitiesDoResetPasswordParams contains all the parameters to send to the API endpoint
for the identities do reset password operation typically these are written to a http.Request
*/
type IdentitiesDoResetPasswordParams struct {

	/*Body*/
	Body *models.V1DoResetPasswordRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the identities do reset password params
func (o *IdentitiesDoResetPasswordParams) WithTimeout(timeout time.Duration) *IdentitiesDoResetPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identities do reset password params
func (o *IdentitiesDoResetPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identities do reset password params
func (o *IdentitiesDoResetPasswordParams) WithContext(ctx context.Context) *IdentitiesDoResetPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identities do reset password params
func (o *IdentitiesDoResetPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identities do reset password params
func (o *IdentitiesDoResetPasswordParams) WithHTTPClient(client *http.Client) *IdentitiesDoResetPasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identities do reset password params
func (o *IdentitiesDoResetPasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the identities do reset password params
func (o *IdentitiesDoResetPasswordParams) WithBody(body *models.V1DoResetPasswordRequest) *IdentitiesDoResetPasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identities do reset password params
func (o *IdentitiesDoResetPasswordParams) SetBody(body *models.V1DoResetPasswordRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IdentitiesDoResetPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
