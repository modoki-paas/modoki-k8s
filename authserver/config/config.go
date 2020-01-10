package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
	"gopkg.in/yaml.v2"
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
	ClientID     string   `yaml:"client_id" json:"client_id"`
	ClientSecret string   `yaml:"client_id" json:"client_id"`
	Scopes       []string `yaml:"scopes" json:"scopes"`
	RedirectURL  string   `yaml:"redirect_url" json:"redirect_url"`
	ProviderURL  string   `yaml:"provider_url" json:"provider_url"`
}

type Config struct {
	Address   string        `yaml:"address" json:"address"`
	Endpoints Endpoints     `yaml:"endpoints" json:"endpoints"`
	APIKeys   []string      `yaml:"api_keys" json:"api_keys"`
	OIDC      OpenIDConnect `yaml:"oidc" json:"oidc"`
}

func ReadConfig(name string) (*Config, error) {
	reader, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(name)
	var config Config
	switch ext {
	case ".json":
		if err := json.NewDecoder(reader).Decode(&config); err != nil {
			return nil, xerrors.Errorf("failed to parse config json: %w", err)
		}
	case ".yml", ".yaml":
		if err := yaml.NewDecoder(reader).Decode(&config); err != nil {
			return nil, xerrors.Errorf("failed to parse config yaml: %w", err)
		}
	default:
		return nil, xerrors.Errorf("unknown extension: %s", ext)
	}

	addDefaultValues(&config)

	return &config, nil
}

func addDefaultValues(cfg *Config) {
	if cfg.Address == "" {
		cfg.Address = ":443"
	}

	targetEndpoints := []**Endpoint{
		&cfg.Endpoints.App,
		&cfg.Endpoints.UserOrg,
		&cfg.Endpoints.Token,
	}

	for _, e := range targetEndpoints {
		if *e == nil {
			*e = &Endpoint{
				Endpoint: cfg.Address,
				Insecure: true,
			}
		}
	}

	envCfg := ReadEnv()

	cfg.APIKeys = append(cfg.APIKeys, envCfg.APIKeys...)

	if cfg.OIDC.ClientID == "" {
		cfg.OIDC.ClientID = envCfg.ClientID
	}
	if cfg.OIDC.ClientSecret == "" {
		cfg.OIDC.ClientSecret = envCfg.ClientSecret
	}

	cfg.OIDC.Scopes = append(cfg.OIDC.Scopes, envCfg.OpenIDConnect.Scopes...)

	if cfg.OIDC.RedirectURL == "" {
		cfg.OIDC.RedirectURL = envCfg.RedirectURL
	}
	if cfg.OIDC.ProviderURL == "" {
		cfg.OIDC.ProviderURL = envCfg.ProviderURL
	}
}
