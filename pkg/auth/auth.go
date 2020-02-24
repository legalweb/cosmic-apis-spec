package auth

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TokenProvider ...
type TokenProvider struct {
	Token string
}

// NewTokenProvider ...
func NewTokenProvider(t string) *TokenProvider {
	return &TokenProvider{
		Token: t,
	}
}

// AuthenticateRequest ...
func (t *TokenProvider) AuthenticateRequest(req runtime.ClientRequest, _ strfmt.Registry) error {
	return req.SetHeaderParam("Authorization", "Bearer "+t.Token)
}
