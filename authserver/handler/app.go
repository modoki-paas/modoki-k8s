package handler

import (
	"context"

	modoki "github.com/modoki-paas/modoki-k8s/api"
)

type AppServer struct {
	Context *ServerContext
}

var _ modoki.AppServer = &AppServer{}

func (s *AppServer) Create(ctx context.Context, in *modoki.AppCreateRequest) (*modoki.AppCreateResponse, error) {
	return s.Context.AppClient.Create(ctx, in)
}

func (s *AppServer) Deploy(ctx context.Context, in *modoki.AppDeployRequest) (*modoki.AppDeployResponse, error) {
	return s.Context.AppClient.Deploy(ctx, in)
}
