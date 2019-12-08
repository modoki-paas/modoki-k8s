package handler

import (
	"context"

	modoki "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/pkg/kustomizer"
	"github.com/modoki-paas/modoki-k8s/yamler/config"
	"golang.org/x/xerrors"
)

type Handler struct {
	Config *config.Config
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

	switch req.Kind {
	case modoki.OperateKind_Apply:
		return &modoki.OperateResponse{
			ApplyYaml:  &modoki.YAML{Config: res},
			DeleteYaml: &modoki.YAML{Config: ""},
		}, nil

	case modoki.OperateKind_Delete:
		return &modoki.OperateResponse{
			ApplyYaml:  &modoki.YAML{Config: ""},
			DeleteYaml: &modoki.YAML{Config: res},
		}, nil

	default:
		return nil, xerrors.Errorf("unknown operate kind: " + req.Kind.String())
	}
}

func (h *Handler) Metrics(context.Context, *modoki.MetricsRequest) (*modoki.MetricsResponse, error) {
	panic("not implemented")
}
