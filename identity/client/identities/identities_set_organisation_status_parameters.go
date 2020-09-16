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

// NewIdentitiesSetOrganisationStatusParams creates a new IdentitiesSetOrganisationStatusParams object
// with the default values initialized.
func NewIdentitiesSetOrganisationStatusParams() *IdentitiesSetOrganisationStatusParams {
	var ()
	return &IdentitiesSetOrganisationStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIdentitiesSetOrganisationStatusParamsWithTimeout creates a new IdentitiesSetOrganisationStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIdentitiesSetOrganisationStatusParamsWithTimeout(timeout time.Duration) *IdentitiesSetOrganisationStatusParams {
	var ()
	return &IdentitiesSetOrganisationStatusParams{

		timeout: timeout,
	}
}

// NewIdentitiesSetOrganisationStatusParamsWithContext creates a new IdentitiesSetOrganisationStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewIdentitiesSetOrganisationStatusParamsWithContext(ctx context.Context) *IdentitiesSetOrganisationStatusParams {
	var ()
	return &IdentitiesSetOrganisationStatusParams{

		Context: ctx,
	}
}

// NewIdentitiesSetOrganisationStatusParamsWithHTTPClient creates a new IdentitiesSetOrganisationStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIdentitiesSetOrganisationStatusParamsWithHTTPClient(client *http.Client) *IdentitiesSetOrganisationStatusParams {
	var ()
	return &IdentitiesSetOrganisationStatusParams{
		HTTPClient: client,
	}
}

/*IdentitiesSetOrganisationStatusParams contains all the parameters to send to the API endpoint
for the identities set organisation status operation typically these are written to a http.Request
*/
type IdentitiesSetOrganisationStatusParams struct {

	/*Body*/
	Body *models.V1SetOrganisationStatusRequest
	/*OrganisationID*/
	OrganisationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the identities set organisation status params
func (o *IdentitiesSetOrganisationStatusParams) WithTimeout(timeout time.Duration) *IdentitiesSetOrganisationStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identities set organisation status params
func (o *IdentitiesSetOrganisationStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identities set organisation status params
func (o *IdentitiesSetOrganisationStatusParams) WithContext(ctx context.Context) *IdentitiesSetOrganisationStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identities set organisation status params
func (o *IdentitiesSetOrganisationStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identities set organisation status params
func (o *IdentitiesSetOrganisationStatusParams) WithHTTPClient(client *http.Client) *IdentitiesSetOrganisationStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identities set organisation status params
func (o *IdentitiesSetOrganisationStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the identities set organisation status params
func (o *IdentitiesSetOrganisationStatusParams) WithBody(body *models.V1SetOrganisationStatusRequest) *IdentitiesSetOrganisationStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identities set organisation status params
func (o *IdentitiesSetOrganisationStatusParams) SetBody(body *models.V1SetOrganisationStatusRequest) {
	o.Body = body
}

// WithOrganisationID adds the organisationID to the identities set organisation status params
func (o *IdentitiesSetOrganisationStatusParams) WithOrganisationID(organisationID string) *IdentitiesSetOrganisationStatusParams {
	o.SetOrganisationID(organisationID)
	return o
}

// SetOrganisationID adds the organisationId to the identities set organisation status params
func (o *IdentitiesSetOrganisationStatusParams) SetOrganisationID(organisationID string) {
	o.OrganisationID = organisationID
}

// WriteToRequest writes these params to a swagger request
func (o *IdentitiesSetOrganisationStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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