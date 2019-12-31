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

type Plugin struct {
	Name       string    `yaml:"name" json:"name"`
	Endpoint   *Endpoint `yaml:"endpoint" json:"endpoint"`
	MetricsAPI bool      `yaml:"metrics_api" json:"metrics_api"`
}

type Endpoints struct {
	Generator *Endpoint `yaml:"generator" json:"generator"`
	App       *Endpoint `yaml:"app" json:"app"`
	UserOrg   *Endpoint `yaml:"user_org" json:"user_org"`

	Plugins []Plugin `yaml:"plugins" json:"plugins"`
}

type Config struct {
	DB        string    `yaml:"db" json:"db"`
	Namespace string    `yaml:"namespace" json:"namespace"`
	Address   string    `yaml:"address" json:"address"`
	Endpoints Endpoints `yaml:"endpoints" json:"endpoints"`
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
		&cfg.Endpoints.Generator,
		&cfg.Endpoints.UserOrg,
	}

	for i := range cfg.Endpoints.Plugins {
		targetEndpoints = append(targetEndpoints, &cfg.Endpoints.Plugins[i].Endpoint)
	}

	for _, e := range targetEndpoints {
		if e := *e; e == nil {
			e.Endpoint = cfg.Address
			e.Insecure = true
		}
	}
}
