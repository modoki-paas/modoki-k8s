package handler

import (
	"context"

	modoki "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/internal/k8s"
	"github.com/modoki-paas/modoki-k8s/pkg/kustomizer"
	"github.com/modoki-paas/modoki-k8s/yamler/config"
	"golang.org/x/xerrors"
)

type Handler struct {
	Config    *config.Config
	K8sClient *k8s.Client
}

var _ modoki.GeneratorServer

func (h *Handler) Operate(ctx context.Context, req *modoki.OperateRequest) (*modoki.OperateResponse, error) {
	ws, err := kustomizer.NewWorkspace()

	if err != nil {
		return nil, xerrors.Errorf("failed to initializer workspace for yamler: %w", err)
	}

	defer ws.Close()

	y, err := ws.LoadConfig()

	if err != nil {
		return nil, xerrors.Errorf("failed to load config in workspace: %w", err)
	}

	ks := []struct {
		name string
		yk   yamlerKustomizer
	}{
		{"setup namespace", setupNamespace},
		{"setup name", setupName},
		{"setup labels", setupLabels},
		{"setup ingress", setupIngress},
		{"setup pod", setupPod},
	}

	for i := range ks {
		y, err = ks[i].yk(ctx, h.Config, y, req)

		if err != nil {
			return nil, xerrors.Errorf("failed to update kustomize yaml(step: %s)", err, ks[i].name)
		}
	}

	if err := ws.SaveConfig(y); err != nil {
		return nil, xerrors.Errorf("failed to save config in workspace: %w", err)
	}

	res, err := ws.Build(ctx)

	if err != nil {
		return nil, xerrors.Errorf("failed to build yaml config in workspace(message: %s): %w", res, err)
	}

	stat := req.Status

	if stat != nil {
		stat.State = "Updating"
	}

	switch req.Kind {
	case modoki.OperateKind_Apply:
		return &modoki.OperateResponse{
			Yaml:   &modoki.YAML{Config: res},
			Status: stat,
		}, nil

	case modoki.OperateKind_Delete:
		return &modoki.OperateResponse{
			Yaml:   &modoki.YAML{Config: res},
			Status: stat,
		}, nil

	default:
		return nil, xerrors.Errorf("unknown operate kind: " + req.Kind.String())
	}
}

func (h *Handler) Metrics(ctx context.Context, req *modoki.MetricsRequest) (*modoki.MetricsResponse, error) {
	stat, err := h.K8sClient.Status(
		ctx,
		req.K8SConfig.Namespace,
		"modoki-app-deploy-"+req.Status.Id,
		"server",
	)

	if err != nil {
		return nil, xerrors.Errorf("failed to get status for service(%s): %w", req.Status.Id, err)
	}

	stat.Id = req.Status.Id
	stat.Domain = req.Status.Domain
	stat.Spec = req.Status.Spec
	stat.CreatedAt = req.Status.CreatedAt
	stat.UpdatedAt = req.Status.UpdatedAt

	return &modoki.MetricsResponse{
		Status: stat,
	}, nil
}
