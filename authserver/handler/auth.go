package handler

import (
	"context"
	"time"

	modoki "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/internal/oidc"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"github.com/modoki-paas/modoki-k8s/pkg/rbac/roles"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	Context *ServerContext
}

func (s *AuthServer) SignIn(ctx context.Context, in *modoki.SignInRequest) (*modoki.SignInResponse, error) {
	oidcConfig := s.Context.Config.OIDC
	auth, err := oidc.NewAuthenticator(ctx, oidcConfig.ClientID, oidcConfig.ClientSecret, oidcConfig.RedirectURL, oidcConfig.ProviderURL, oidcConfig.Scopes)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to initialize OpenID Connect authenticator")
	}

	redirectURI, state, err := auth.Login(ctx)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate redirect url for OpenID Connect")
	}

	return &modoki.SignInResponse{
		RedirectUri: redirectURI,
		State:       state,
	}, nil
}

func (s *AuthServer) SignOut(_ context.Context, _ *modoki.SignOutRequest) (*modoki.SignOutResponse, error) {
	panic("not implemented")
}

func (s *AuthServer) Callback(ctx context.Context, req *modoki.CallbackRequest) (*modoki.CallbackResponse, error) {
	oidcConfig := s.Context.Config.OIDC
	author, err := oidc.NewAuthenticator(ctx, oidcConfig.ClientID, oidcConfig.ClientSecret, oidcConfig.RedirectURL, oidcConfig.RedirectURL, oidcConfig.Scopes)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to initialize OpenID Connect authenticator")
	}

	res, err := author.Callback(ctx, req.ExpectedState, req.State, req.Code)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to handle callback requests")
	}

	{
		ctx := auth.OverwritePerfomerContext(ctx, "system:authserver", res.IDToken.Subject, roles.SystemAuth)

		name := res.IDToken.Subject
		if n, ok := res.Profile["name"]; ok {
			if n, ok := n.(string); ok {
				name = n
			}
		}

		_, err := s.Context.UserOrgClient.UserAdd(
			ctx, &modoki.UserAddRequest{
				User: &modoki.User{
					UserId:         res.IDToken.Subject,
					Name:           name,
					SystemRoleName: roles.SystemDeveloper.Name, // TODO: support group claim
				},
			},
		)

		if err != nil {
			if stat, ok := status.FromError(err); !ok {
				return nil, err
			} else if stat.Code() != codes.AlreadyExists {
				return nil, err
			}
		}

		resp, err := s.Context.TokenClient.IssueToken(ctx, &modoki.IssueTokenRequest{
			Id: "oidc:" + time.Now().Format(time.RFC3339Nano),
		})

		if err != nil {
			return nil, err
		}

		return &modoki.CallbackResponse{
			Token: resp.Token,
		}, nil
	}
}

func (s *AuthServer) IsPrivate(method string) bool {
	switch method {
	case "/modoki.Auth/SignIn", "/modoki.Auth/Callback":
		return false
	default:
		return true
	}
}
