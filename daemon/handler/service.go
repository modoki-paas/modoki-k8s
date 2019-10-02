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
	specOpt := req.Spec.Options

	opt := map[string]string{}
	for k, v := range specOpt {
		opt[k] = v.String()
	}

	serviceConfig := &store.Service{
		Name:  req.Spec.Name,
		Owner: int(req.Spec.Owner),
		Config: &store.ServiceConfig{
			Image:   req.Spec.Image,
			Command: req.Spec.Command,
			Args:    req.Spec.Args,
			Options: opt,
		},
	}

	svc, err := s.Context.DB.Service().AddService(serviceConfig)

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
