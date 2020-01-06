package auth

import (
	"context"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthorizerInterceptor struct {
	tokens []string
}

// validateTokenFromHeader validates tokens in env config
func (ai *AuthorizerInterceptor) validateTokenFromHeader(md metadata.MD) bool {
	headers := md.Get("Authorization")

	if len(headers) == 0 {
		return false
	}

	key := strings.TrimPrefix(headers[0], "Bearer ")

	for _, k := range ai.tokens {
		if k == key {
			return true
		}
	}

	return false
}

func (ai *AuthorizerInterceptor) getUserID(md metadata.MD) (id string, ok bool) {
	arr := md.Get(UserIDHeader)

	if len(arr) == 0 {
		return "", false
	}

	return arr[0], true
}

func (ai *AuthorizerInterceptor) getTargetID(md metadata.MD) (id string, ok bool) {
	arr := md.Get(TargetIDHeader)

	if len(arr) == 0 {
		return "", false
	}

	return arr[0], true
}

func (ai *AuthorizerInterceptor) wrapContext(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, xerrors.Errorf("unauthorized: metadata missing")
	}

	if !ai.validateTokenFromHeader(md) {
		return nil, xerrors.Errorf("unauthorized")
	}

	id, ok := ai.getUserID(md)

	if !ok {
		return nil, xerrors.Errorf("unknown user")
	}

	ctx = AddUserIDContext(ctx, id)

	targetID, ok := ai.getTargetID(md)

	if !ok {
		targetID = id
	}

	ctx = AddTargetIDContext(ctx, targetID)

	roles := getRoles(md)
	ctx = AddRolesContext(ctx, roles)

	perms := getPermissions(roles, targetID)
	ctx = AddPermissionsContext(ctx, perms)

	return ctx, nil
}

// UnaryServerInterceptor handles authentication for each call
func UnaryServerInterceptor(tokens []string) grpc.UnaryServerInterceptor {
	ai := &AuthorizerInterceptor{
		tokens: tokens,
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx, err := ai.wrapContext(ctx)

		if err != nil {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}

		return handler(ctx, req)
	}
}

// StreamServerInterceptor handles authentication for each call
func StreamServerInterceptor(tokens []string) grpc.StreamServerInterceptor {
	ai := &AuthorizerInterceptor{
		tokens: tokens,
	}

	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		newStream := grpc_middleware.WrapServerStream(stream)

		ctx, err := ai.wrapContext(stream.Context())

		if err != nil {
			return status.Error(codes.PermissionDenied, err.Error())
		}

		newStream.WrappedContext = ctx

		return handler(srv, newStream)
	}
}
