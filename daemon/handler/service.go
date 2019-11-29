package handler

import (
	"context"

	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/daemon/store"
	"golang.org/x/xerrors"
)

type ServiceServer struct {
	Context *ServerContext
}

var _ api.ServiceServer = &ServiceServer{}
var _ Authorizer = &ServiceServer{}

func (s *ServiceServer) Create(ctx context.Context, req *api.ServiceCreateRequest) (*api.ServiceCreateResponse, error) {
	tx, err := s.Context.DB.Begin(ctx, nil)

	if err != nil {
		return nil, xerrors.Errorf("failed to begin transaction: %w", err)
	}

	svc := &store.Service{
		Owner: int(req.Spec.Owner),
		Name:  req.Spec.Name,
		Spec:  req.Spec,
	}

	svc, err = tx.Service().AddService(svc)

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
