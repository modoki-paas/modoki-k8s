package handler

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/store/apps"
	"github.com/modoki-paas/modoki-k8s/internal/dbutil"
	"github.com/modoki-paas/modoki-k8s/internal/grpcutil"
	"github.com/modoki-paas/modoki-k8s/internal/imageutil"
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
			domain = domain + s.Context.Config.Domain[1:]
		} else {
			domain = domain + "." + s.Context.Config.Domain
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

					log.Println(err)
					return status.Error(codes.Internal, "generator error")
				}

				log.Println(err)

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

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AppServer) Deploy(ctx context.Context, req *api.AppDeployRequest) (res *api.AppDeployResponse, err error) {
	if err := auth.IsAuthorized(ctx, permissions.AppUpdate); err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	err = dbutil.Transaction(ctx, s.Context.DB, func(tx *sqlx.Tx) error {
		store := apps.NewAppStore(tx)

		app, err := store.FindAppByID(req.Id)

		if err != nil {
			return status.Error(codes.Unknown, "unknown app")
		}

		if app.Owner != auth.GetTargetIDContext(ctx) {
			return status.Error(codes.Unknown, "unknown app")
		}

		ow, err := imageutil.ParseOverwrite(req.Spec.Image, true)

		if err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("image name format error: %v", err))
		}

		hash, err := imageutil.GetImageHash(req.Spec.Image)

		if err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("failed to get hash of images: %v", err))
		}

		req.Spec.Image = fmt.Sprintf("%s@sha256:%s", ow.Name, hash)

		err = store.UpdateApp(app.SeqID, (*types.AppSpec)(req.Spec))

		if err != nil {
			return xerrors.Errorf("failed to store app config in db: %w", err)
		}

		y := &api.YAML{}
		for i := range s.Context.Generators {
			res, err := s.Context.Generators[i].Client.Operate(
				ctx,
				&api.OperateRequest{
					Id:     app.ID,
					Domain: app.Name,
					Kind:   api.OperateKind_Apply,
					Spec:   req.Spec,
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

		res = &api.AppDeployResponse{
			Status: &api.AppStatus{
				Id:         app.ID,
				Domain:     app.Name,
				Spec:       req.Spec,
				State:      "deploying",
				StartedAt:  grpcutil.GRPCTimestamp(app.UpdatedAt), // TODO: Fix timestamp
				ExitCode:   0,
				CreatedAt:  grpcutil.GRPCTimestamp(app.CreatedAt),
				UpdatedAt:  grpcutil.GRPCTimestamp(time.Now()), // TODO: Fix timestamp
				Attributes: map[string]string{},
			},
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
