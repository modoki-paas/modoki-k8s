package handler

import (
	"github.com/modoki-paas/modoki-k8s/authserver/config"
)

// ServerContext contains accessor used by handlers
type ServerContext struct {
	Config    *config.Config
	EnvConfig *config.EnvConfig
}
