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
	svc, err := s.Context.DB.Service().AddService(req.Spec)

	if err != nil {
		return nil, xerrors.Errorf("failed to store service config in db :%w", err)
	}

	return &api.ServiceCreateResponse{
		Id:   int32(svc.ID),
		Spec: req.GetSpec(),
	}, nil
}

func (s *ServiceServer) Authorize(ctx context.Context, route string) error {
	user, _ := GetValuesFromContext(ctx)

	if user != nil {
		return nil
	}

	return xerrors.New("not authorized")
}
