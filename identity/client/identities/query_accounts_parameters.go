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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewQueryAccountsParams creates a new QueryAccountsParams object
// with the default values initialized.
func NewQueryAccountsParams() *QueryAccountsParams {
	var ()
	return &QueryAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryAccountsParamsWithTimeout creates a new QueryAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryAccountsParamsWithTimeout(timeout time.Duration) *QueryAccountsParams {
	var ()
	return &QueryAccountsParams{

		timeout: timeout,
	}
}

// NewQueryAccountsParamsWithContext creates a new QueryAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryAccountsParamsWithContext(ctx context.Context) *QueryAccountsParams {
	var ()
	return &QueryAccountsParams{

		Context: ctx,
	}
}

// NewQueryAccountsParamsWithHTTPClient creates a new QueryAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryAccountsParamsWithHTTPClient(client *http.Client) *QueryAccountsParams {
	var ()
	return &QueryAccountsParams{
		HTTPClient: client,
	}
}

/*QueryAccountsParams contains all the parameters to send to the API endpoint
for the query accounts operation typically these are written to a http.Request
*/
type QueryAccountsParams struct {

	/*AccountID*/
	AccountID []string
	/*EmailAddress*/
	EmailAddress []string
	/*PaginationLimit*/
	PaginationLimit *int64
	/*PaginationPage*/
	PaginationPage *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query accounts params
func (o *QueryAccountsParams) WithTimeout(timeout time.Duration) *QueryAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query accounts params
func (o *QueryAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query accounts params
func (o *QueryAccountsParams) WithContext(ctx context.Context) *QueryAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query accounts params
func (o *QueryAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query accounts params
func (o *QueryAccountsParams) WithHTTPClient(client *http.Client) *QueryAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query accounts params
func (o *QueryAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the query accounts params
func (o *QueryAccountsParams) WithAccountID(accountID []string) *QueryAccountsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the query accounts params
func (o *QueryAccountsParams) SetAccountID(accountID []string) {
	o.AccountID = accountID
}

// WithEmailAddress adds the emailAddress to the query accounts params
func (o *QueryAccountsParams) WithEmailAddress(emailAddress []string) *QueryAccountsParams {
	o.SetEmailAddress(emailAddress)
	return o
}

// SetEmailAddress adds the emailAddress to the query accounts params
func (o *QueryAccountsParams) SetEmailAddress(emailAddress []string) {
	o.EmailAddress = emailAddress
}

// WithPaginationLimit adds the paginationLimit to the query accounts params
func (o *QueryAccountsParams) WithPaginationLimit(paginationLimit *int64) *QueryAccountsParams {
	o.SetPaginationLimit(paginationLimit)
	return o
}

// SetPaginationLimit adds the paginationLimit to the query accounts params
func (o *QueryAccountsParams) SetPaginationLimit(paginationLimit *int64) {
	o.PaginationLimit = paginationLimit
}

// WithPaginationPage adds the paginationPage to the query accounts params
func (o *QueryAccountsParams) WithPaginationPage(paginationPage *int64) *QueryAccountsParams {
	o.SetPaginationPage(paginationPage)
	return o
}

// SetPaginationPage adds the paginationPage to the query accounts params
func (o *QueryAccountsParams) SetPaginationPage(paginationPage *int64) {
	o.PaginationPage = paginationPage
}

// WriteToRequest writes these params to a swagger request
func (o *QueryAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesAccountID := o.AccountID

	joinedAccountID := swag.JoinByFormat(valuesAccountID, "multi")
	// query array param account_id
	if err := r.SetQueryParam("account_id", joinedAccountID...); err != nil {
		return err
	}

	valuesEmailAddress := o.EmailAddress

	joinedEmailAddress := swag.JoinByFormat(valuesEmailAddress, "multi")
	// query array param email_address
	if err := r.SetQueryParam("email_address", joinedEmailAddress...); err != nil {
		return err
	}

	if o.PaginationLimit != nil {

		// query param pagination.limit
		var qrPaginationLimit int64
		if o.PaginationLimit != nil {
			qrPaginationLimit = *o.PaginationLimit
		}
		qPaginationLimit := swag.FormatInt64(qrPaginationLimit)
		if qPaginationLimit != "" {
			if err := r.SetQueryParam("pagination.limit", qPaginationLimit); err != nil {
				return err
			}
		}

	}

	if o.PaginationPage != nil {

		// query param pagination.page
		var qrPaginationPage int64
		if o.PaginationPage != nil {
			qrPaginationPage = *o.PaginationPage
		}
		qPaginationPage := swag.FormatInt64(qrPaginationPage)
		if qPaginationPage != "" {
			if err := r.SetQueryParam("pagination.page", qPaginationPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}