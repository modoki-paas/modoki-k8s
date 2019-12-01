package handler

import (
	"context"

	modoki "github.com/modoki-paas/modoki-k8s/api"
)

type Handler struct {
}

var _ modoki.GeneratorServer

func (h *Handler) Operate(ctx context.Context, req *modoki.OperateRequest) (*modoki.OperateResponse, error) {
	return &modoki.OperateResponse{
		Yaml: req.Yaml,
	}, nil
}
