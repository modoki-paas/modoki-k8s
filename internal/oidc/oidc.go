package oidc

import (
	"context"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"golang.org/x/xerrors"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

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
