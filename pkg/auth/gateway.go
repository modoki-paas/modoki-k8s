package auth

import (
	"context"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	api "github.com/modoki-paas/modoki-k8s/api"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type IsPrivateService interface {
	IsPrivate(method string) bool
}

type GatewayAuthorizerInterceptor struct {
	tokenClient   api.TokenClient
	userOrgClient api.UserOrgClient
}

// getTokenFromHeader gets a token from header
func (ai *GatewayAuthorizerInterceptor) getTokenFromHeader(md metadata.MD) string {
	headers := md.Get("Authorization")

	if len(headers) == 0 {
		return ""
	}

	key := strings.TrimPrefix(headers[0], "Bearer ")

	return key
}

func (ai *GatewayAuthorizerInterceptor) getTargetID(md metadata.MD) (id string, ok bool) {
	arr := md.Get(TargetIDHeader)

	if len(arr) == 0 {
		return "", false
	}

	return arr[0], true
}

func (ai *GatewayAuthorizerInterceptor) wrapContext(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, xerrors.Errorf("unauthorized: metadata missing")
	}

	tk := ai.getTokenFromHeader(md)

	if tk == "" {
		return nil, status.Error(codes.Unauthenticated, "authorization header missing")
	}

	vt, err := ai.tokenClient.ValidateToken(ctx, &api.ValidateTokenRequest{
		Token: tk,
	})

	if err != nil {
		if stat, ok := status.FromError(err); ok {
			if stat.Code() == codes.NotFound {
				return nil, status.Error(codes.Unauthenticated, "invalid token")
			}

			return nil, stat.Err()
		}

		return nil, xerrors.Errorf("failed to validate token: %w", err)
	}

	performer, err := ai.userOrgClient.UserFindByID(ctx, &api.UserFindByIDRequest{
		UserId: vt.UserId,
	})

	if err != nil {
		if stat, ok := status.FromError(err); ok {
			return nil, stat.Err()
		}

		return nil, xerrors.Errorf("failed to validate token: %w", err)
	}

	targetID, ok := ai.getTargetID(md)

	if !ok {
		targetID = performer.User.UserId
	}

	rb, err := ai.userOrgClient.GetRoleBinding(ctx, &api.GetRoleBindingRequest{
		UserId:   performer.GetUser().UserId,
		TargetId: targetID,
	})

	if err != nil {
		if stat, ok := status.FromError(err); !ok {
			return nil, xerrors.Errorf("failed to get role binding: %w", err)
		} else if stat.Code() != codes.NotFound {
			return nil, stat.Err()
		}
		// continue if role binding is not found
	}

	ctx = AddUserIDContext(ctx, performer.User.UserId)

	ctx = AddTargetIDContext(ctx, targetID)

	roles := RoleBindings(map[string]string{
		"*": performer.User.SystemRoleName,
	})

	if rb != nil && rb.Role != "" {
		roles[targetID] = rb.Role
	}

	ctx = AddRolesContext(ctx, roles)

	perms := getPermissions(roles, targetID)
	ctx = AddPermissionsContext(ctx, perms)

	return ctx, nil
}

// UnaryGatewayServerInterceptor handles authentication for each call
func UnaryGatewayServerInterceptor(tokenClient api.TokenClient, userOrgClient api.UserOrgClient) grpc.UnaryServerInterceptor {
	ai := &GatewayAuthorizerInterceptor{
		tokenClient:   tokenClient,
		userOrgClient: userOrgClient,
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if ips, ok := info.Server.(IsPrivateService); ok && !ips.IsPrivate(info.FullMethod) {
			return handler(ctx, req)
		}

		ctx, err := ai.wrapContext(ctx)

		if err != nil {
			if stat, ok := status.FromError(err); ok {
				return nil, stat.Err()
			}

			return nil, status.Error(codes.PermissionDenied, err.Error())
		}

		return handler(ctx, req)
	}
}

// StreamGatewayServerInterceptor handles authentication for each call
func StreamGatewayServerInterceptor(tokenClient api.TokenClient, userOrgClient api.UserOrgClient) grpc.StreamServerInterceptor {
	ai := &GatewayAuthorizerInterceptor{
		tokenClient:   tokenClient,
		userOrgClient: userOrgClient,
	}

	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if ips, ok := srv.(IsPrivateService); ok && !ips.IsPrivate(info.FullMethod) {
			return handler(srv, stream)
		}

		newStream := grpc_middleware.WrapServerStream(stream)

		ctx, err := ai.wrapContext(stream.Context())

		if err != nil {
			if stat, ok := status.FromError(err); ok {
				return stat.Err()
			}

			return status.Error(codes.PermissionDenied, err.Error())
		}

		newStream.WrappedContext = ctx

		return handler(srv, newStream)
	}
}
