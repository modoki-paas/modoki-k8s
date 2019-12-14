package handler

import (
	"context"

	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/store/apps"
	"golang.org/x/xerrors"
)

type AppServer struct {
	Context *ServerContext
}

var _ api.AppServer = &AppServer{}
var _ Authorizer = &AppServer{}

func (s *AppServer) Create(ctx context.Context, req *api.AppCreateRequest) (res *api.AppCreateResponse, err error) {
	tx, err := s.Context.DB.BeginTxx(ctx, nil)

	if err != nil {
		return nil, xerrors.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			if err := tx.Commit(); err != nil {
				err = xerrors.Errorf("failed to commit transaction: %w", err)
				res = nil
			}
		}
	}()

	store := apps.NewAppStore(tx)

	svc := &apps.App{
		Owner: int(req.Spec.Owner),
		Name:  req.Spec.Name,
		Spec:  (*apps.AppSpec)(req.Spec),
	}

	seq, err := store.AddApp(svc)

	if err != nil {
		return nil, xerrors.Errorf("failed to store app config in db: %w", err)
	}

	app, err := store.GetApp(seq)

	if err != nil {
		return nil, xerrors.Errorf("failed to get app config in db: %w", err)
	}

	return &api.AppCreateResponse{
		Id:   app.ID,
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
