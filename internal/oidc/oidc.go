package oidc

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/modoki-paas/modoki-k8s/internal/tokenutil"
	"golang.org/x/oauth2"
	"golang.org/x/xerrors"
)

// Authenticator is a helper to sign in with OpenID Connect
type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
}

// AuthResult represents results of OpenID Connect authentication flow
type AuthResult struct {
	Token   *oauth2.Token
	IDToken *oidc.IDToken
	Profile map[string]interface{}
}

// NewAuthenticator initializes an Authenticator
func NewAuthenticator(ctx context.Context, clientID, clientSecret, redirectURL, providerURL string, scopes []string) (*Authenticator, error) {
	provider, err := oidc.NewProvider(ctx, providerURL)
	if err != nil {
		return nil, xerrors.Errorf("failed to initialize oidc provider: %w", err)
	}

	conf := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       scopes,
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}

// Callback handles callback requests from IdP
func (author *Authenticator) Callback(ctx context.Context, expectedState, actualState, code string) (*AuthResult, error) {
	token, err := author.Config.Exchange(ctx, code)
	if err != nil {
		return nil, xerrors.Errorf("failed to exchange code: %w", err)
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, xerrors.Errorf("No id_token field in oauth2 token: %w", http.StatusInternalServerError)
	}

	oidcConfig := &oidc.Config{
		ClientID: author.Config.ClientID,
	}

	idToken, err := author.Provider.Verifier(oidcConfig).Verify(ctx, rawIDToken)

	if err != nil {
		return nil, xerrors.Errorf("Failed to verify ID Token: %w", err)
	}

	// Getting now the userInfo
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return nil, xerrors.Errorf("failed to get userinfo: %w", err)
	}

	return &AuthResult{
		Token:   token,
		IDToken: idToken,
		Profile: profile,
	}, nil
}

// Login generates state and redirect url to IdP
func (author *Authenticator) Login(ctx context.Context, salt string) (redirectURL, state string, err error) {
	state, err = tokenutil.GenerateRandomToken()

	if err != nil {
		return "", "", xerrors.Errorf("failed to generate a random token: %w", err)
	}

	hashedState, err := bcrypt.GenerateFromPassword([]byte(state), bcrypt.DefaultCost)

	if err != nil {
		return "", "", xerrors.Errorf("failed to hash the state string: %w", err)
	}

	redirectURL = author.Config.AuthCodeURL(string(hashedState))

	return
}
