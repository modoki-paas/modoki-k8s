package handler

import (
	"github.com/modoki-paas/modoki-k8s/apiserver/config"
	"github.com/modoki-paas/modoki-k8s/apiserver/store"
)

// ServerContext contains accessor used by handlers
type ServerContext struct {
	DB     *store.DB
	Config *config.Config
}
