package handler

import (
	"context"

	modoki "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/pkg/kustomizer"
	"golang.org/x/xerrors"
)

type Handler struct {
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

	// Edit yaml

	if err := ws.SaveConfig(y); err != nil {
		return nil, xerrors.Errorf("failed to save config in workspace: %w", err)
	}

	res, err := ws.Build(ctx)

	if err != nil {
		return nil, xerrors.Errorf("failed to build yaml config in workspace: %w", err)
	}

	return &modoki.OperateResponse{
		Yaml: &modoki.YAML{Config: res},
	}, nil
}
