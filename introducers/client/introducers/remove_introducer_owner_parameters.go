// Code generated by go-swagger; DO NOT EDIT.

package introducers

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

// NewRemoveIntroducerOwnerParams creates a new RemoveIntroducerOwnerParams object
// with the default values initialized.
func NewRemoveIntroducerOwnerParams() *RemoveIntroducerOwnerParams {
	var ()
	return &RemoveIntroducerOwnerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveIntroducerOwnerParamsWithTimeout creates a new RemoveIntroducerOwnerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveIntroducerOwnerParamsWithTimeout(timeout time.Duration) *RemoveIntroducerOwnerParams {
	var ()
	return &RemoveIntroducerOwnerParams{

		timeout: timeout,
	}
}

// NewRemoveIntroducerOwnerParamsWithContext creates a new RemoveIntroducerOwnerParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveIntroducerOwnerParamsWithContext(ctx context.Context) *RemoveIntroducerOwnerParams {
	var ()
	return &RemoveIntroducerOwnerParams{

		Context: ctx,
	}
}

// NewRemoveIntroducerOwnerParamsWithHTTPClient creates a new RemoveIntroducerOwnerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveIntroducerOwnerParamsWithHTTPClient(client *http.Client) *RemoveIntroducerOwnerParams {
	var ()
	return &RemoveIntroducerOwnerParams{
		HTTPClient: client,
	}
}

/*RemoveIntroducerOwnerParams contains all the parameters to send to the API endpoint
for the remove introducer owner operation typically these are written to a http.Request
*/
type RemoveIntroducerOwnerParams struct {

	/*AccountID*/
	AccountID string
	/*IntroducerID*/
	IntroducerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove introducer owner params
func (o *RemoveIntroducerOwnerParams) WithTimeout(timeout time.Duration) *RemoveIntroducerOwnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove introducer owner params
func (o *RemoveIntroducerOwnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove introducer owner params
func (o *RemoveIntroducerOwnerParams) WithContext(ctx context.Context) *RemoveIntroducerOwnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove introducer owner params
func (o *RemoveIntroducerOwnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove introducer owner params
func (o *RemoveIntroducerOwnerParams) WithHTTPClient(client *http.Client) *RemoveIntroducerOwnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove introducer owner params
func (o *RemoveIntroducerOwnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the remove introducer owner params
func (o *RemoveIntroducerOwnerParams) WithAccountID(accountID string) *RemoveIntroducerOwnerParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the remove introducer owner params
func (o *RemoveIntroducerOwnerParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithIntroducerID adds the introducerID to the remove introducer owner params
func (o *RemoveIntroducerOwnerParams) WithIntroducerID(introducerID string) *RemoveIntroducerOwnerParams {
	o.SetIntroducerID(introducerID)
	return o
}

// SetIntroducerID adds the introducerId to the remove introducer owner params
func (o *RemoveIntroducerOwnerParams) SetIntroducerID(introducerID string) {
	o.IntroducerID = introducerID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveIntroducerOwnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_id
	if err := r.SetPathParam("account_id", o.AccountID); err != nil {
		return err
	}

	// path param introducer_id
	if err := r.SetPathParam("introducer_id", o.IntroducerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
