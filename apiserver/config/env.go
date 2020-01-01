package config

import (
	"fmt"
	"os"
	"strings"
)

type EnvConfig struct {
	DB      string
	APIKeys []string
	Domain  string
}

func ReadEnv() *EnvConfig {
	apiKeys := strings.Split(os.Getenv("MODOKI_API_KEY"), ",")
	domain := os.Getenv("MODOKI_APP_DOMAIN")
	db := os.Getenv("MODOKI_DB")

	if db == "" {
		db = fmt.Sprintf(
			"mysql://%s:%s@%s:%s/%s?parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)
	}

	return &EnvConfig{
		APIKeys: apiKeys,
		Domain:  domain,
		DB:      db,
	}
}
