// Code generated by go-swagger; DO NOT EDIT.

package customer_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new customer gateway API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer gateway API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CustomerGatewayAddCompany(params *CustomerGatewayAddCompanyParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayAddCompanyOK, error)

	CustomerGatewayAddEmployment(params *CustomerGatewayAddEmploymentParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayAddEmploymentOK, error)

	CustomerGatewayAddOwnership(params *CustomerGatewayAddOwnershipParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayAddOwnershipOK, error)

	CustomerGatewayCreateDownloadDocumentLink(params *CustomerGatewayCreateDownloadDocumentLinkParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayCreateDownloadDocumentLinkOK, error)

	CustomerGatewayCreateUploadDocumentLink(params *CustomerGatewayCreateUploadDocumentLinkParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayCreateUploadDocumentLinkOK, error)

	CustomerGatewayDeleteDocument(params *CustomerGatewayDeleteDocumentParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayDeleteDocumentOK, error)

	CustomerGatewayGetAddressHistory(params *CustomerGatewayGetAddressHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayGetAddressHistoryOK, error)

	CustomerGatewayGetApplication(params *CustomerGatewayGetApplicationParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayGetApplicationOK, error)

	CustomerGatewayGetNotifications(params *CustomerGatewayGetNotificationsParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayGetNotificationsOK, error)

	CustomerGatewayGetProfile(params *CustomerGatewayGetProfileParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayGetProfileOK, error)

	CustomerGatewayGetSummary(params *CustomerGatewayGetSummaryParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayGetSummaryOK, error)

	CustomerGatewayListApplications(params *CustomerGatewayListApplicationsParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayListApplicationsOK, error)

	CustomerGatewayListCompanies(params *CustomerGatewayListCompaniesParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayListCompaniesOK, error)

	CustomerGatewayListDocuments(params *CustomerGatewayListDocumentsParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayListDocumentsOK, error)

	CustomerGatewayListEmployment(params *CustomerGatewayListEmploymentParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayListEmploymentOK, error)

	CustomerGatewayListNotificationPreferences(params *CustomerGatewayListNotificationPreferencesParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayListNotificationPreferencesOK, error)

	CustomerGatewayListOwnerships(params *CustomerGatewayListOwnershipsParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayListOwnershipsOK, error)

	CustomerGatewayRemoveEmployment(params *CustomerGatewayRemoveEmploymentParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayRemoveEmploymentOK, error)

	CustomerGatewayRemoveOwnership(params *CustomerGatewayRemoveOwnershipParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayRemoveOwnershipOK, error)

	CustomerGatewaySetAddressHistory(params *CustomerGatewaySetAddressHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewaySetAddressHistoryOK, error)

	CustomerGatewaySetNotificationPreferences(params *CustomerGatewaySetNotificationPreferencesParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewaySetNotificationPreferencesOK, error)

	CustomerGatewayUpdateEmployment(params *CustomerGatewayUpdateEmploymentParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayUpdateEmploymentOK, error)

	CustomerGatewayUpdateOwnership(params *CustomerGatewayUpdateOwnershipParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayUpdateOwnershipOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CustomerGatewayAddCompany customer gateway add company API
*/
func (a *Client) CustomerGatewayAddCompany(params *CustomerGatewayAddCompanyParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayAddCompanyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayAddCompanyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_AddCompany",
		Method:             "POST",
		PathPattern:        "/v1/companies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayAddCompanyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayAddCompanyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayAddCompanyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayAddEmployment customer gateway add employment API
*/
func (a *Client) CustomerGatewayAddEmployment(params *CustomerGatewayAddEmploymentParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayAddEmploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayAddEmploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_AddEmployment",
		Method:             "POST",
		PathPattern:        "/v1/employment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayAddEmploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayAddEmploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayAddEmploymentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayAddOwnership customer gateway add ownership API
*/
func (a *Client) CustomerGatewayAddOwnership(params *CustomerGatewayAddOwnershipParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayAddOwnershipOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayAddOwnershipParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_AddOwnership",
		Method:             "POST",
		PathPattern:        "/v1/ownerships",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayAddOwnershipReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayAddOwnershipOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayAddOwnershipDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayCreateDownloadDocumentLink customer gateway create download document link API
*/
func (a *Client) CustomerGatewayCreateDownloadDocumentLink(params *CustomerGatewayCreateDownloadDocumentLinkParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayCreateDownloadDocumentLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayCreateDownloadDocumentLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_CreateDownloadDocumentLink",
		Method:             "POST",
		PathPattern:        "/v1/documents/{collection_id}/{document_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayCreateDownloadDocumentLinkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayCreateDownloadDocumentLinkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayCreateDownloadDocumentLinkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayCreateUploadDocumentLink customer gateway create upload document link API
*/
func (a *Client) CustomerGatewayCreateUploadDocumentLink(params *CustomerGatewayCreateUploadDocumentLinkParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayCreateUploadDocumentLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayCreateUploadDocumentLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_CreateUploadDocumentLink",
		Method:             "POST",
		PathPattern:        "/v1/documents/{collection_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayCreateUploadDocumentLinkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayCreateUploadDocumentLinkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayCreateUploadDocumentLinkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayDeleteDocument customer gateway delete document API
*/
func (a *Client) CustomerGatewayDeleteDocument(params *CustomerGatewayDeleteDocumentParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayDeleteDocumentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayDeleteDocumentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_DeleteDocument",
		Method:             "DELETE",
		PathPattern:        "/v1/documents/{collection_id}/{document_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayDeleteDocumentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayDeleteDocumentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayDeleteDocumentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayGetAddressHistory customer gateway get address history API
*/
func (a *Client) CustomerGatewayGetAddressHistory(params *CustomerGatewayGetAddressHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayGetAddressHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayGetAddressHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_GetAddressHistory",
		Method:             "GET",
		PathPattern:        "/v1/addresses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayGetAddressHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayGetAddressHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayGetAddressHistoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayGetApplication customer gateway get application API
*/
func (a *Client) CustomerGatewayGetApplication(params *CustomerGatewayGetApplicationParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayGetApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayGetApplicationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_GetApplication",
		Method:             "GET",
		PathPattern:        "/v1/applications/{app_type}/{app_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayGetApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayGetApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayGetApplicationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayGetNotifications customer gateway get notifications API
*/
func (a *Client) CustomerGatewayGetNotifications(params *CustomerGatewayGetNotificationsParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayGetNotificationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayGetNotificationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_GetNotifications",
		Method:             "GET",
		PathPattern:        "/v1/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayGetNotificationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayGetNotificationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayGetNotificationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayGetProfile customer gateway get profile API
*/
func (a *Client) CustomerGatewayGetProfile(params *CustomerGatewayGetProfileParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayGetProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayGetProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_GetProfile",
		Method:             "GET",
		PathPattern:        "/v1/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayGetProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayGetProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayGetProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayGetSummary customer gateway get summary API
*/
func (a *Client) CustomerGatewayGetSummary(params *CustomerGatewayGetSummaryParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayGetSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayGetSummaryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_GetSummary",
		Method:             "GET",
		PathPattern:        "/v1/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayGetSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayGetSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayGetSummaryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayListApplications customer gateway list applications API
*/
func (a *Client) CustomerGatewayListApplications(params *CustomerGatewayListApplicationsParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayListApplicationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayListApplicationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_ListApplications",
		Method:             "GET",
		PathPattern:        "/v1/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayListApplicationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayListApplicationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayListApplicationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayListCompanies customer gateway list companies API
*/
func (a *Client) CustomerGatewayListCompanies(params *CustomerGatewayListCompaniesParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayListCompaniesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayListCompaniesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_ListCompanies",
		Method:             "GET",
		PathPattern:        "/v1/companies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayListCompaniesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayListCompaniesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayListCompaniesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayListDocuments customer gateway list documents API
*/
func (a *Client) CustomerGatewayListDocuments(params *CustomerGatewayListDocumentsParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayListDocumentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayListDocumentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_ListDocuments",
		Method:             "GET",
		PathPattern:        "/v1/documents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayListDocumentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayListDocumentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayListDocumentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayListEmployment customer gateway list employment API
*/
func (a *Client) CustomerGatewayListEmployment(params *CustomerGatewayListEmploymentParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayListEmploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayListEmploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_ListEmployment",
		Method:             "GET",
		PathPattern:        "/v1/employment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayListEmploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayListEmploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayListEmploymentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayListNotificationPreferences customer gateway list notification preferences API
*/
func (a *Client) CustomerGatewayListNotificationPreferences(params *CustomerGatewayListNotificationPreferencesParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayListNotificationPreferencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayListNotificationPreferencesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_ListNotificationPreferences",
		Method:             "GET",
		PathPattern:        "/v1/settings/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayListNotificationPreferencesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayListNotificationPreferencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayListNotificationPreferencesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayListOwnerships customer gateway list ownerships API
*/
func (a *Client) CustomerGatewayListOwnerships(params *CustomerGatewayListOwnershipsParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayListOwnershipsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayListOwnershipsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_ListOwnerships",
		Method:             "GET",
		PathPattern:        "/v1/ownerships",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayListOwnershipsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayListOwnershipsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayListOwnershipsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayRemoveEmployment customer gateway remove employment API
*/
func (a *Client) CustomerGatewayRemoveEmployment(params *CustomerGatewayRemoveEmploymentParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayRemoveEmploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayRemoveEmploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_RemoveEmployment",
		Method:             "DELETE",
		PathPattern:        "/v1/employment/{employment_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayRemoveEmploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayRemoveEmploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayRemoveEmploymentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayRemoveOwnership customer gateway remove ownership API
*/
func (a *Client) CustomerGatewayRemoveOwnership(params *CustomerGatewayRemoveOwnershipParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayRemoveOwnershipOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayRemoveOwnershipParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_RemoveOwnership",
		Method:             "DELETE",
		PathPattern:        "/v1/ownerships/{ownership_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayRemoveOwnershipReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayRemoveOwnershipOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayRemoveOwnershipDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewaySetAddressHistory customer gateway set address history API
*/
func (a *Client) CustomerGatewaySetAddressHistory(params *CustomerGatewaySetAddressHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewaySetAddressHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewaySetAddressHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_SetAddressHistory",
		Method:             "POST",
		PathPattern:        "/v1/addresses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewaySetAddressHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewaySetAddressHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewaySetAddressHistoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewaySetNotificationPreferences merges the notification preferences which are currently available with the preferences provided in set notification preferences request
*/
func (a *Client) CustomerGatewaySetNotificationPreferences(params *CustomerGatewaySetNotificationPreferencesParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewaySetNotificationPreferencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewaySetNotificationPreferencesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_SetNotificationPreferences",
		Method:             "POST",
		PathPattern:        "/v1/settings/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewaySetNotificationPreferencesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewaySetNotificationPreferencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewaySetNotificationPreferencesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayUpdateEmployment customer gateway update employment API
*/
func (a *Client) CustomerGatewayUpdateEmployment(params *CustomerGatewayUpdateEmploymentParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayUpdateEmploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayUpdateEmploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_UpdateEmployment",
		Method:             "PUT",
		PathPattern:        "/v1/employment/{employment_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayUpdateEmploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayUpdateEmploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayUpdateEmploymentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CustomerGatewayUpdateOwnership customer gateway update ownership API
*/
func (a *Client) CustomerGatewayUpdateOwnership(params *CustomerGatewayUpdateOwnershipParams, authInfo runtime.ClientAuthInfoWriter) (*CustomerGatewayUpdateOwnershipOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerGatewayUpdateOwnershipParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomerGateway_UpdateOwnership",
		Method:             "PUT",
		PathPattern:        "/v1/ownerships/{ownership_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerGatewayUpdateOwnershipReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerGatewayUpdateOwnershipOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerGatewayUpdateOwnershipDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
