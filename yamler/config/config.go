package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Address       string   `yaml:"address" json:"address"`
	AppSecretName string   `json:"app_secret_name" yaml:"app_secret_name"`
	APIKeys       []string `yaml:"api_keys" json:"api_keys"`
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

	if cfg.AppSecretName == "" {
		cfg.AppSecretName = "modoki-apps-cert-secret"
	}

	envCfg := ReadEnv()

	cfg.APIKeys = append(cfg.APIKeys, envCfg.APIKeys...)
}
