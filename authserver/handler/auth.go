package handler

import (
	"context"

	modoki "github.com/modoki-paas/modoki-k8s/api"
)

type AuthServer struct {
	Context *ServerContext
}

func (s *AuthServer) SignIn(ctx context.Context, in *modoki.SignInRequest) (*modoki.SignInResponse, error) {
	panic("not implemented")
}

func (s *AuthServer) SignOut(_ context.Context, _ *modoki.SignOutRequest) (*modoki.SignOutResponse, error) {
	panic("not implemented")
}

func (s *AuthServer) Callback(_ context.Context, _ *modoki.CallbackRequest) (*modoki.CallbackResponse, error) {
	panic("not implemented")
}

func (s *AuthServer) IsPrivate(method string) bool {
	switch method {
	case "modoki.Auth/SignIn", "modoki.Auth/Callback":
		return false
	default:
		return true
	}
}
