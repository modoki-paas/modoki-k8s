package handler

import (
	"github.com/jmoiron/sqlx"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/config"
)

type Plugin struct {
	Name    string
	Client  api.GeneratorClient
	Metrics bool
}

// ServerContext contains accessor used by handlers
type ServerContext struct {
	DB        *sqlx.DB
	Config    *config.Config
	EnvConfig *config.EnvConfig

	AppClient     api.AppClient
	UserOrgClient api.UserOrgClient
	Plugsins      []*Plugin
}
