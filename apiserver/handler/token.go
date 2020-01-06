package handler

import (
	"context"

	modoki "github.com/modoki-paas/modoki-k8s/api"
)

type TokenServer struct {
}

func (s *TokenServer) IssueToken(_ context.Context, _ *modoki.IssueTokenRequest) (*modoki.IssueTokenResponse, error) {
	panic("not implemented")
}

func (s *TokenServer) ValidateToken(_ context.Context, _ *modoki.ValidateTokenRequest) (*modoki.ValidateTokenResponse, error) {
	panic("not implemented")
}
