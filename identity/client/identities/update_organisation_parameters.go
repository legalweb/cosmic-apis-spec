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

// NewUpdateOrganisationParams creates a new UpdateOrganisationParams object
// with the default values initialized.
func NewUpdateOrganisationParams() *UpdateOrganisationParams {
	var ()
	return &UpdateOrganisationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganisationParamsWithTimeout creates a new UpdateOrganisationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOrganisationParamsWithTimeout(timeout time.Duration) *UpdateOrganisationParams {
	var ()
	return &UpdateOrganisationParams{

		timeout: timeout,
	}
}

// NewUpdateOrganisationParamsWithContext creates a new UpdateOrganisationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOrganisationParamsWithContext(ctx context.Context) *UpdateOrganisationParams {
	var ()
	return &UpdateOrganisationParams{

		Context: ctx,
	}
}

// NewUpdateOrganisationParamsWithHTTPClient creates a new UpdateOrganisationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOrganisationParamsWithHTTPClient(client *http.Client) *UpdateOrganisationParams {
	var ()
	return &UpdateOrganisationParams{
		HTTPClient: client,
	}
}

/*UpdateOrganisationParams contains all the parameters to send to the API endpoint
for the update organisation operation typically these are written to a http.Request
*/
type UpdateOrganisationParams struct {

	/*Body*/
	Body *models.V1UpdateOrganisationRequest
	/*OrganisationID*/
	OrganisationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update organisation params
func (o *UpdateOrganisationParams) WithTimeout(timeout time.Duration) *UpdateOrganisationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organisation params
func (o *UpdateOrganisationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organisation params
func (o *UpdateOrganisationParams) WithContext(ctx context.Context) *UpdateOrganisationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organisation params
func (o *UpdateOrganisationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organisation params
func (o *UpdateOrganisationParams) WithHTTPClient(client *http.Client) *UpdateOrganisationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organisation params
func (o *UpdateOrganisationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update organisation params
func (o *UpdateOrganisationParams) WithBody(body *models.V1UpdateOrganisationRequest) *UpdateOrganisationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update organisation params
func (o *UpdateOrganisationParams) SetBody(body *models.V1UpdateOrganisationRequest) {
	o.Body = body
}

// WithOrganisationID adds the organisationID to the update organisation params
func (o *UpdateOrganisationParams) WithOrganisationID(organisationID string) *UpdateOrganisationParams {
	o.SetOrganisationID(organisationID)
	return o
}

// SetOrganisationID adds the organisationId to the update organisation params
func (o *UpdateOrganisationParams) SetOrganisationID(organisationID string) {
	o.OrganisationID = organisationID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganisationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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