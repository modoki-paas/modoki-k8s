package config

import (
	"context"
	"time"

	"github.com/modoki-paas/modoki-k8s/pkg/configloader"
	"golang.org/x/xerrors"
)

type Config struct {
	Address       string   `yaml:"address" json:"address" config:"modoki-address"`
	AppSecretName string   `json:"app_secret_name" yaml:"app_secret_name" config:"modoki-app-secret-name"`
	APIKeys       []string `yaml:"api_keys" json:"api_keys" config:"modoki-api-key"`
}

var (
	defaultConfig = &Config{
		Address:       ":443",
		AppSecretName: "modoki-apps-cert-secret",
	}
)

func ReadConfig() (*Config, error) {
	cfg := *defaultConfig

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := configloader.ReadConfig(ctx, "yamler", &cfg); err != nil {
		return nil, xerrors.Errorf("failed to load config: %w", err)
	}

	return &cfg, nil
}
