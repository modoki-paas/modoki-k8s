package handler

import (
	"context"
	"strings"

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

		spec := &api.AppSpec{
			Image: "nginx",
		}
		domain := req.Domain

		if strings.HasPrefix(s.Context.Config.Domain, "*.") {
			domain = domain + s.Context.Config.Domain
		} else {
			domain = domain + s.Context.Config.Domain[2:]
		}

		app := &types.App{
			Owner: auth.GetTargetIDContext(ctx),
			Name:  domain,
			Spec:  (*types.AppSpec)(spec),
		}

		_, id, err := store.AddApp(app)

		if err != nil {
			return xerrors.Errorf("failed to store app config in db: %w", err)
		}

		y := &api.YAML{}
		for i := range s.Context.Generators {
			res, err := s.Context.Generators[i].Client.Operate(
				ctx,
				&api.OperateRequest{
					Id:     id,
					Domain: domain,
					Kind:   api.OperateKind_Apply,
					Spec:   spec,
					Yaml:   y,
					K8SConfig: &api.KubernetesConfig{
						Namespace: s.Context.Config.Namespace,
					},
				},
			)

			if err != nil {
				if stat, ok := status.FromError(err); ok {
					switch stat.Code() {
					case codes.PermissionDenied:
						return stat.Err()
					case codes.InvalidArgument:
						return stat.Err()
					}

					return status.Error(codes.Internal, "generator error")
				}

				return status.Error(codes.Internal, "generator failed due to unknown reason")
			}

			y = res.Yaml
		}

		if output, err := s.Context.K8s.Apply(ctx, strings.NewReader(y.Config)); err != nil {
			return xerrors.Errorf("failed to apply k8s config(message: %s): %w", output, err)
		}

		res = &api.AppCreateResponse{
			Id:     id,
			Domain: domain,
		}

		return nil
	})

	return res, nil
}

func (s *AppServer) Deploy(ctx context.Context, req *api.AppDeployRequest) (res *api.AppDeployResponse, err error) {
	if err := auth.IsAuthorized(ctx, permissions.AppUpdate); err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	err = dbutil.Transaction(ctx, s.Context.DB, func(tx *sqlx.Tx) error {
		/*store := apps.NewAppStore(tx)

		app := &types.App{
			Owner: auth.GetTargetIDContext(ctx),
			Name:  req.,
			Spec:  (*types.AppSpec)(req.Spec),
		}

		imageutil.GetImageHash(req.Spec.Image)

		_, id, err := store.AddApp(app)

		if err != nil {
			return xerrors.Errorf("failed to store app config in db: %w", err)
		}

		y := &api.YAML{}
		for i := range s.Context.Generators {
			res, err := s.Context.Generators[i].Client.Operate(
				ctx,
				&api.OperateRequest{
					Id:   id,
					Kind: api.OperateKind_Apply,
					Spec: req.Spec,
					Yaml: y,
					K8SConfig: &api.KubernetesConfig{
						Namespace: s.Context.Config.Namespace,
					},
				},
			)

			if err != nil {
				if stat, ok := status.FromError(err); ok {
					switch stat.Code() {
					case codes.PermissionDenied:
						return stat.Err()
					case codes.InvalidArgument:
						return stat.Err()
					}

					return status.Error(codes.Internal, "generator error")
				}

				return status.Error(codes.Internal, "generator failed due to unknown reason")
			}

			y = res.Yaml
		}

		res = &api.AppCreateResponse{
			Id:   id,
			Spec: req.GetSpec(),
		}
		*/
		return nil
	})

	return res, nil
}
