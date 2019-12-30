package handler

import (
	"context"

	"github.com/jmoiron/sqlx"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/store/apps"
	"github.com/modoki-paas/modoki-k8s/internal/dbutil"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"github.com/modoki-paas/modoki-k8s/pkg/rbac/permissions"
	"github.com/modoki-paas/modoki-k8s/pkg/types"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppServer struct {
	Context *ServerContext
}

var _ api.AppServer = &AppServer{}

func (s *AppServer) Create(ctx context.Context, req *api.AppCreateRequest) (res *api.AppCreateResponse, err error) {
	if err := auth.IsAuthorized(ctx, permissions.AppCreate); err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	err = dbutil.Transaction(ctx, s.Context.DB, func(tx *sqlx.Tx) error {
		store := apps.NewAppStore(tx)

		app := &types.App{
			Owner: req.Spec.Owner,
			Name:  req.Spec.Name,
			Spec:  (*types.AppSpec)(req.Spec),
		}

		_, id, err := store.AddApp(app)

		if err != nil {
			return xerrors.Errorf("failed to store app config in db: %w", err)
		}

		res = &api.AppCreateResponse{
			Id:   id,
			Spec: req.GetSpec(),
		}

		return nil
	})

	return res, nil
}
