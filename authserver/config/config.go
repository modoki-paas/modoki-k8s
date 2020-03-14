package config

import (
	"context"
	"time"

	"github.com/modoki-paas/modoki-k8s/pkg/configloader"
	"golang.org/x/xerrors"
)

type Endpoint struct {
	Endpoint string `yaml:"endpoint" json:"endpoint"`
	Insecure bool   `yaml:"insecure" json:"insecure"`
}

type Endpoints struct {
	App     *Endpoint `yaml:"app" json:"app"`
	UserOrg *Endpoint `yaml:"user_org" json:"user_org"`
	Token   *Endpoint `yaml:"token" json:"token"`
}

type OpenIDConnect struct {
<<<<<<< HEAD
	ClientID     string   `yaml:"client_id" json:"client_id"`
	ClientSecret string   `yaml:"client_secret" json:"client_secret"`
	Scopes       []string `yaml:"scopes" json:"scopes"`
	RedirectURL  string   `yaml:"redirect_url" json:"redirect_url"`
	ProviderURL  string   `yaml:"provider_url" json:"provider_url"`
=======
	ClientID     string   `yaml:"client_id" json:"client_id" config:"oidc-client-id"`
	ClientSecret string   `yaml:"client_id" json:"client_id" config:"oidc-client-secret"`
	Scopes       []string `yaml:"scopes" json:"scopes" config:"oidc-scopes"`
	RedirectURL  string   `yaml:"redirect_url" json:"redirect_url" config:"oidc-redirect-url"`
	ProviderURL  string   `yaml:"provider_url" json:"provider_url" config:"oidc-provider-url"`
>>>>>>> origin/master
}

type Config struct {
	Address   string        `yaml:"address" json:"address" config:"modoki-address"`
	Endpoints Endpoints     `yaml:"endpoints" json:"endpoints" config:"-"`
	APIKeys   []string      `yaml:"api_keys" json:"api_keys" config:"modoki-api-key"` // TODO: Renamed to modoki-api-keys
	OIDC      OpenIDConnect `yaml:"oidc" json:"oidc"`
}

var (
	defaultConfig = &Config{
		Address: ":443",

		Endpoints: Endpoints{
			Token: &Endpoint{
				Endpoint: ":443",
				Insecure: true,
			},
			UserOrg: &Endpoint{
				Endpoint: ":443",
				Insecure: true,
			},
			App: &Endpoint{
				Endpoint: ":443",
				Insecure: true,
			},
		},
	}
)

func ReadConfig() (*Config, error) {
	cfg := *defaultConfig

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := configloader.ReadConfig(ctx, "authserver", &cfg); err != nil {
		return nil, xerrors.Errorf("failed to load config: %w", err)
	}

	return &cfg, nil
}
