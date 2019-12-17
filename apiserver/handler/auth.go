package handler

import (
	"context"
	"fmt"

	"strings"

	"github.com/modoki-paas/modoki-k8s/pkg/types"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const (
	TokenContext contextKey = "token-context"
	UserContext  contextKey = "user-context"
)

type Authorizer interface {
	Authorize(ctx context.Context, route string) error
}

type AuthorizerInterceptor struct {
	serverCtx *ServerContext
}

// validateTokenFromHeader validates tokens in env config
func (ai *AuthorizerInterceptor) validateTokenFromHeader(ctx context.Context) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}

	headers := md.Get("Authorization")

	if len(headers) == 0 {
		return false
	}

	key := strings.TrimPrefix(headers[0], "Bearer ")

	for _, k := range ai.serverCtx.EnvConfig.APIKeys {
		if k == key {
			return true
		}
	}

	return false
}

func (ai *AuthorizerInterceptor) addUserTokenContext(ctx context.Context) (context.Context, error) {
	var user *types.User
	var tk *types.Token

	ctx = context.WithValue(ctx, TokenContext, tk)
	ctx = context.WithValue(ctx, UserContext, user)

	return ctx, nil
}

// UnaryServerInterceptor handles authentication for each call
func UnaryServerInterceptor(serverCtx *ServerContext) grpc.UnaryServerInterceptor {
	ai := &AuthorizerInterceptor{
		serverCtx: serverCtx,
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !ai.validateTokenFromHeader(ctx) {
			return nil, status.Error(codes.PermissionDenied, "unauthorized")
		}

		ctx, err := ai.addUserTokenContext(ctx)

		if err != nil {
			return nil, xerrors.Errorf("failed to add user and token data to context: %w", err)
		}

		srv, ok := info.Server.(Authorizer)
		if !ok {
			return nil, fmt.Errorf("each service should implement an authorization")
		}

		if err := srv.Authorize(ctx, info.FullMethod); err != nil {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}

		return handler(ctx, req)
	}
}

// GetValuesFromContext returns user and token stored in context
func GetValuesFromContext(ctx context.Context) (user *types.User, token *types.Token) {
	u := ctx.Value(UserContext)

	if u == nil {
		user = nil
	} else {
		user = u.(*types.User)
	}

	tk := ctx.Value(TokenContext)

	if tk != nil {
		token = tk.(*types.Token)
	}

	return
}
