package config

import (
	"os"
	"strings"
)

type EnvConfig struct {
	APIKeys []string
}

func ReadEnv() (*EnvConfig, error) {
	apiKeys := strings.Split(os.Getenv("MODOKI_APIKEY"), ",")

	return &EnvConfig{
		APIKeys: apiKeys,
	}, nil
}
