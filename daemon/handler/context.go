package handler

import (
	"github.com/modoki-paas/modoki-k8s/daemon/config"
	"github.com/modoki-paas/modoki-k8s/daemon/store"
)

// ServerContext contains accessor used by handlers
type ServerContext struct {
	DB     *store.DB
	Config *config.Config
}
