package config

import (
	"os"
	"strings"
)

type EnvConfig struct {
	APIKeys []string

	OpenIDConnect
}

func ReadEnv() *EnvConfig {
	apiKeys := strings.Split(os.Getenv("MODOKI_API_KEY"), ",")

	if len(apiKeys) == 1 && apiKeys[0] == "" {
		apiKeys = nil
	}

	oidc := OpenIDConnect{
		ClientID:     os.Getenv("OIDC_CLIENT_ID"),
		ClientSecret: os.Getenv("OIDC_CLIENT_SECRET"),
		Scopes:       strings.Split(os.Getenv("OIDC_SCOPES"), ","),
		RedirectURL:  os.Getenv("OIDC_REDIRECT_URL"),
		ProviderURL:  os.Getenv("OIDC_PROVIDER_URL"),
	}

	if len(oidc.Scopes) == 1 && oidc.Scopes[0] == "" {
		oidc.Scopes = nil
	}

	return &EnvConfig{
		APIKeys:       apiKeys,
		OpenIDConnect: oidc,
	}
}
