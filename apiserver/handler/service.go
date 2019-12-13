package handler

import (
	"context"

	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/store"
	"golang.org/x/xerrors"
)

type AppServer struct {
	Context *ServerContext
}

var _ api.AppServer = &AppServer{}
var _ Authorizer = &AppServer{}

func (s *AppServer) Create(ctx context.Context, req *api.AppCreateRequest) (*api.AppCreateResponse, error) {
	tx, err := s.Context.DB.Begin(ctx, nil)

	if err != nil {
		return nil, xerrors.Errorf("failed to begin transaction: %w", err)
	}

	svc := &store.App{
		Owner: int(req.Spec.Owner),
		Name:  req.Spec.Name,
		Spec:  (*store.AppSpec)(req.Spec),
	}

	svc, err = tx.App().AddApp(svc)

	if err != nil {
		return nil, xerrors.Errorf("failed to store app config in db :%w", err)
	}

	return &api.AppCreateResponse{
		Id:   svc.ID,
		Spec: req.GetSpec(),
	}, nil
}

func (s *AppServer) Authorize(ctx context.Context, route string) error {
	user, _ := GetValuesFromContext(ctx)

	if user != nil {
		return nil
	}

	return xerrors.New("not authorized")
}
