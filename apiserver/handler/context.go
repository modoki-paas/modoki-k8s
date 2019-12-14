package handler

import (
	"github.com/jmoiron/sqlx"
	"github.com/modoki-paas/modoki-k8s/apiserver/config"
)

// ServerContext contains accessor used by handlers
type ServerContext struct {
	DB        *sqlx.DB
	Config    *config.Config
	EnvConfig *config.EnvConfig
}
