package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
	"gopkg.in/yaml.v2"
)

type Plugin struct {
	Name              string `yaml:"name" json:"name"`
	GeneratorEndpoint string `yaml:"generator" json:"generator"`
	MetricsAPI        bool   `yaml:"metrics_api" json:"metrics_api"`
}

type Config struct {
	Plugins []Plugin `yaml:"plugins" json:"plugins"`

	Generator struct {
		Endpoint string `yaml:"endpoint" json:"endpoint"`
		Insecure bool   `yaml:"insecure" json:"insecure"`
	} `yaml:"generator" json:"generator"`

	DB      string `yaml:"db" json:"db"`
	Address string `yaml:"address" json:"address"`
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

	return &config, nil
}
