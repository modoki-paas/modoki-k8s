package handler

import (
	"context"

	api "github.com/modoki-paas/modoki-k8s/api"
)

type ServiceServer struct {
	Context *ServerContext
}

func (s *ServiceServer) Create(ctx context.Context, req *api.ServiceCreateRequest) (*api.ServiceCreateResponse, error) {
	panic("not implemented")
}

var _ api.ServiceServer = &ServiceServer{}
