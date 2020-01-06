package handler

import (
	"context"

	modoki "github.com/modoki-paas/modoki-k8s/api"
)

type TokenServer struct {
	Context *ServerContext
}

func (s *TokenServer) IssueToken(ctx context.Context, in *modoki.IssueTokenRequest) (*modoki.IssueTokenResponse, error) {
	return s.Context.TokenClient.IssueToken(ctx, in)
}

func (s *TokenServer) ValidateToken(ctx context.Context, in *modoki.ValidateTokenRequest) (*modoki.ValidateTokenResponse, error) {
	return s.Context.TokenClient.ValidateToken(ctx, in)
}
