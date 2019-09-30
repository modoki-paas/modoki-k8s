package handler

import (
	"context"

	api "github.com/modoki-paas/modoki-k8s/api"
	"golang.org/x/xerrors"
)

type ServiceServer struct {
	Context *ServerContext
}

var _ api.ServiceServer = &ServiceServer{}
var _ Authorizer = &ServiceServer{}

func (s *ServiceServer) Create(ctx context.Context, req *api.ServiceCreateRequest) (*api.ServiceCreateResponse, error) {
	panic("not implemented")
}

func (s *ServiceServer) Authorize(ctx context.Context, route string) error {
	user, _ := GetValuesFromContext(ctx)

	if user != nil {
		return nil
	}

	return xerrors.New("not authorized")
}
