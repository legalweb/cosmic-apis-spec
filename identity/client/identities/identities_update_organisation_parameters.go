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

	"lwebco.de/cosmic-apis-spec/identity/models"
)

// NewIdentitiesUpdateOrganisationParams creates a new IdentitiesUpdateOrganisationParams object
// with the default values initialized.
func NewIdentitiesUpdateOrganisationParams() *IdentitiesUpdateOrganisationParams {
	var ()
	return &IdentitiesUpdateOrganisationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIdentitiesUpdateOrganisationParamsWithTimeout creates a new IdentitiesUpdateOrganisationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIdentitiesUpdateOrganisationParamsWithTimeout(timeout time.Duration) *IdentitiesUpdateOrganisationParams {
	var ()
	return &IdentitiesUpdateOrganisationParams{

		timeout: timeout,
	}
}

// NewIdentitiesUpdateOrganisationParamsWithContext creates a new IdentitiesUpdateOrganisationParams object
// with the default values initialized, and the ability to set a context for a request
func NewIdentitiesUpdateOrganisationParamsWithContext(ctx context.Context) *IdentitiesUpdateOrganisationParams {
	var ()
	return &IdentitiesUpdateOrganisationParams{

		Context: ctx,
	}
}

// NewIdentitiesUpdateOrganisationParamsWithHTTPClient creates a new IdentitiesUpdateOrganisationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIdentitiesUpdateOrganisationParamsWithHTTPClient(client *http.Client) *IdentitiesUpdateOrganisationParams {
	var ()
	return &IdentitiesUpdateOrganisationParams{
		HTTPClient: client,
	}
}

/*IdentitiesUpdateOrganisationParams contains all the parameters to send to the API endpoint
for the identities update organisation operation typically these are written to a http.Request
*/
type IdentitiesUpdateOrganisationParams struct {

	/*Body*/
	Body *models.V1UpdateOrganisationRequest
	/*OrganisationID*/
	OrganisationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the identities update organisation params
func (o *IdentitiesUpdateOrganisationParams) WithTimeout(timeout time.Duration) *IdentitiesUpdateOrganisationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identities update organisation params
func (o *IdentitiesUpdateOrganisationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identities update organisation params
func (o *IdentitiesUpdateOrganisationParams) WithContext(ctx context.Context) *IdentitiesUpdateOrganisationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identities update organisation params
func (o *IdentitiesUpdateOrganisationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identities update organisation params
func (o *IdentitiesUpdateOrganisationParams) WithHTTPClient(client *http.Client) *IdentitiesUpdateOrganisationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identities update organisation params
func (o *IdentitiesUpdateOrganisationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the identities update organisation params
func (o *IdentitiesUpdateOrganisationParams) WithBody(body *models.V1UpdateOrganisationRequest) *IdentitiesUpdateOrganisationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identities update organisation params
func (o *IdentitiesUpdateOrganisationParams) SetBody(body *models.V1UpdateOrganisationRequest) {
	o.Body = body
}

// WithOrganisationID adds the organisationID to the identities update organisation params
func (o *IdentitiesUpdateOrganisationParams) WithOrganisationID(organisationID string) *IdentitiesUpdateOrganisationParams {
	o.SetOrganisationID(organisationID)
	return o
}

// SetOrganisationID adds the organisationId to the identities update organisation params
func (o *IdentitiesUpdateOrganisationParams) SetOrganisationID(organisationID string) {
	o.OrganisationID = organisationID
}

// WriteToRequest writes these params to a swagger request
func (o *IdentitiesUpdateOrganisationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param organisation_id
	if err := r.SetPathParam("organisation_id", o.OrganisationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}