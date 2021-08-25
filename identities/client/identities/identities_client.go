// Code generated by go-swagger; DO NOT EDIT.

package identities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new identities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for identities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	IdentitiesCreateAccount(params *IdentitiesCreateAccountParams) (*IdentitiesCreateAccountOK, error)

	IdentitiesDoResetPassword(params *IdentitiesDoResetPasswordParams) (*IdentitiesDoResetPasswordOK, error)

	IdentitiesGetAccount(params *IdentitiesGetAccountParams, authInfo runtime.ClientAuthInfoWriter) (*IdentitiesGetAccountOK, error)

	IdentitiesLogin(params *IdentitiesLoginParams) (*IdentitiesLoginOK, error)

	IdentitiesQueryAccounts(params *IdentitiesQueryAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*IdentitiesQueryAccountsOK, error)

	IdentitiesSendResetPassword(params *IdentitiesSendResetPasswordParams) (*IdentitiesSendResetPasswordOK, error)

	IdentitiesTokenCheck(params *IdentitiesTokenCheckParams) (*IdentitiesTokenCheckOK, error)

	IdentitiesTokenRenew(params *IdentitiesTokenRenewParams) (*IdentitiesTokenRenewOK, error)

	IdentitiesUpdateAccountDetails(params *IdentitiesUpdateAccountDetailsParams, authInfo runtime.ClientAuthInfoWriter) (*IdentitiesUpdateAccountDetailsOK, error)

	IdentitiesUpdateAccountPassword(params *IdentitiesUpdateAccountPasswordParams, authInfo runtime.ClientAuthInfoWriter) (*IdentitiesUpdateAccountPasswordOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  IdentitiesCreateAccount identities create account API
*/
func (a *Client) IdentitiesCreateAccount(params *IdentitiesCreateAccountParams) (*IdentitiesCreateAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentitiesCreateAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Identities_CreateAccount",
		Method:             "POST",
		PathPattern:        "/v1/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentitiesCreateAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentitiesCreateAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentitiesCreateAccountDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IdentitiesDoResetPassword identities do reset password API
*/
func (a *Client) IdentitiesDoResetPassword(params *IdentitiesDoResetPasswordParams) (*IdentitiesDoResetPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentitiesDoResetPasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Identities_DoResetPassword",
		Method:             "POST",
		PathPattern:        "/v1/users/reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentitiesDoResetPasswordReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentitiesDoResetPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentitiesDoResetPasswordDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IdentitiesGetAccount identities get account API
*/
func (a *Client) IdentitiesGetAccount(params *IdentitiesGetAccountParams, authInfo runtime.ClientAuthInfoWriter) (*IdentitiesGetAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentitiesGetAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Identities_GetAccount",
		Method:             "GET",
		PathPattern:        "/v1/accounts/{account_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentitiesGetAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentitiesGetAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentitiesGetAccountDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IdentitiesLogin identities login API
*/
func (a *Client) IdentitiesLogin(params *IdentitiesLoginParams) (*IdentitiesLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentitiesLoginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Identities_Login",
		Method:             "POST",
		PathPattern:        "/v1/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentitiesLoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentitiesLoginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentitiesLoginDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IdentitiesQueryAccounts identities query accounts API
*/
func (a *Client) IdentitiesQueryAccounts(params *IdentitiesQueryAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*IdentitiesQueryAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentitiesQueryAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Identities_QueryAccounts",
		Method:             "GET",
		PathPattern:        "/v1/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentitiesQueryAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentitiesQueryAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentitiesQueryAccountsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IdentitiesSendResetPassword identities send reset password API
*/
func (a *Client) IdentitiesSendResetPassword(params *IdentitiesSendResetPasswordParams) (*IdentitiesSendResetPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentitiesSendResetPasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Identities_SendResetPassword",
		Method:             "POST",
		PathPattern:        "/v1/users/send-reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentitiesSendResetPasswordReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentitiesSendResetPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentitiesSendResetPasswordDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IdentitiesTokenCheck identities token check API
*/
func (a *Client) IdentitiesTokenCheck(params *IdentitiesTokenCheckParams) (*IdentitiesTokenCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentitiesTokenCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Identities_TokenCheck",
		Method:             "POST",
		PathPattern:        "/v1/token/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentitiesTokenCheckReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentitiesTokenCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentitiesTokenCheckDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IdentitiesTokenRenew identities token renew API
*/
func (a *Client) IdentitiesTokenRenew(params *IdentitiesTokenRenewParams) (*IdentitiesTokenRenewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentitiesTokenRenewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Identities_TokenRenew",
		Method:             "POST",
		PathPattern:        "/v1/token/renew",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentitiesTokenRenewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentitiesTokenRenewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentitiesTokenRenewDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IdentitiesUpdateAccountDetails identities update account details API
*/
func (a *Client) IdentitiesUpdateAccountDetails(params *IdentitiesUpdateAccountDetailsParams, authInfo runtime.ClientAuthInfoWriter) (*IdentitiesUpdateAccountDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentitiesUpdateAccountDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Identities_UpdateAccountDetails",
		Method:             "POST",
		PathPattern:        "/v1/accounts/{account_id}/update_details",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentitiesUpdateAccountDetailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentitiesUpdateAccountDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentitiesUpdateAccountDetailsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IdentitiesUpdateAccountPassword identities update account password API
*/
func (a *Client) IdentitiesUpdateAccountPassword(params *IdentitiesUpdateAccountPasswordParams, authInfo runtime.ClientAuthInfoWriter) (*IdentitiesUpdateAccountPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentitiesUpdateAccountPasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Identities_UpdateAccountPassword",
		Method:             "POST",
		PathPattern:        "/v1/accounts/{account_id}/change_password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentitiesUpdateAccountPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentitiesUpdateAccountPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentitiesUpdateAccountPasswordDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
