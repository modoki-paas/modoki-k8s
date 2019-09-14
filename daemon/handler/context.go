package handler

import (
	"github.com/modoki-paas/modoki-k8s/daemon/store"
)

// ServerContext contains accessor used by handlers
type ServerContext struct {
	store *store.Store
}