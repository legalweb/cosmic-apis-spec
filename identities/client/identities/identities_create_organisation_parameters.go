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

// NewIdentitiesCreateOrganisationParams creates a new IdentitiesCreateOrganisationParams object
// with the default values initialized.
func NewIdentitiesCreateOrganisationParams() *IdentitiesCreateOrganisationParams {
	var ()
	return &IdentitiesCreateOrganisationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIdentitiesCreateOrganisationParamsWithTimeout creates a new IdentitiesCreateOrganisationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIdentitiesCreateOrganisationParamsWithTimeout(timeout time.Duration) *IdentitiesCreateOrganisationParams {
	var ()
	return &IdentitiesCreateOrganisationParams{

		timeout: timeout,
	}
}

// NewIdentitiesCreateOrganisationParamsWithContext creates a new IdentitiesCreateOrganisationParams object
// with the default values initialized, and the ability to set a context for a request
func NewIdentitiesCreateOrganisationParamsWithContext(ctx context.Context) *IdentitiesCreateOrganisationParams {
	var ()
	return &IdentitiesCreateOrganisationParams{

		Context: ctx,
	}
}

// NewIdentitiesCreateOrganisationParamsWithHTTPClient creates a new IdentitiesCreateOrganisationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIdentitiesCreateOrganisationParamsWithHTTPClient(client *http.Client) *IdentitiesCreateOrganisationParams {
	var ()
	return &IdentitiesCreateOrganisationParams{
		HTTPClient: client,
	}
}

/*IdentitiesCreateOrganisationParams contains all the parameters to send to the API endpoint
for the identities create organisation operation typically these are written to a http.Request
*/
type IdentitiesCreateOrganisationParams struct {

	/*Body*/
	Body *models.V1CreateOrganisationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the identities create organisation params
func (o *IdentitiesCreateOrganisationParams) WithTimeout(timeout time.Duration) *IdentitiesCreateOrganisationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identities create organisation params
func (o *IdentitiesCreateOrganisationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identities create organisation params
func (o *IdentitiesCreateOrganisationParams) WithContext(ctx context.Context) *IdentitiesCreateOrganisationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identities create organisation params
func (o *IdentitiesCreateOrganisationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identities create organisation params
func (o *IdentitiesCreateOrganisationParams) WithHTTPClient(client *http.Client) *IdentitiesCreateOrganisationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identities create organisation params
func (o *IdentitiesCreateOrganisationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the identities create organisation params
func (o *IdentitiesCreateOrganisationParams) WithBody(body *models.V1CreateOrganisationRequest) *IdentitiesCreateOrganisationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identities create organisation params
func (o *IdentitiesCreateOrganisationParams) SetBody(body *models.V1CreateOrganisationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IdentitiesCreateOrganisationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
