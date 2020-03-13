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

type Plugin struct {
	Name       string `yaml:"name" json:"name"`
	MetricsAPI bool   `yaml:"metrics_api" json:"metrics_api"`
	*Endpoint
}

type Endpoints struct {
	Generator *Endpoint `yaml:"generator" json:"generator"`
	App       *Endpoint `yaml:"app" json:"app"`
	UserOrg   *Endpoint `yaml:"user_org" json:"user_org"`

	Plugins []Plugin `yaml:"plugins" json:"plugins"`
}

type Config struct {
	DB        string    `yaml:"db" json:"db" config:"modoki-db"`
	Domain    string    `yaml:"domain" json:"domain" config:"modoki-app-domain"`
	Namespace string    `yaml:"namespace" json:"namespace" config:"modoki-namespace"`
	Address   string    `yaml:"address" json:"address" config:"modoki-address"`
	Endpoints Endpoints `yaml:"endpoints" json:"endpoints" config:"-"`
	APIKeys   []string  `yaml:"api_keys" json:"api_keys" config:"modoki-app-key"` // TODO: Rename to modoki-api-keys
}

var (
	defaultConfig = &Config{
		Address:   ":443",
		Namespace: "modoki",

		Endpoints: Endpoints{
			Generator: &Endpoint{
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

	if err := configloader.ReadConfig(ctx, "apiserver", &cfg); err != nil {
		return nil, xerrors.Errorf("failed to load config: %w", err)
	}

	return &cfg, nil
}
