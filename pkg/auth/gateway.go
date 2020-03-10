package auth

import (
	"context"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	api "github.com/modoki-paas/modoki-k8s/api"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrUnauthenticated returns an error when authentication failed
	ErrUnauthenticated = xerrors.Errorf("unauthenticated")
)

type GatewayAuthorizerInterceptor struct {
	tokenClient   api.TokenClient
	userOrgClient api.UserOrgClient
}

// AuthenticatedMetadata represents data retrieved from the access token
type AuthenticatedMetadata struct {
	UserID string
	Roles  RoleBindings

	TargetID             string
	PermissionsForTarget map[string]struct{}
}

// GetAuthenticatedMetadata returns all metadata to authorize users
func (ai *GatewayAuthorizerInterceptor) GetAuthenticatedMetadata(ctx context.Context, tk, targetID string) (*AuthenticatedMetadata, error) {
	vt, err := ai.tokenClient.ValidateToken(ctx, &api.ValidateTokenRequest{
		Token: tk,
	})

	if err != nil {
		if stat, ok := status.FromError(err); ok {
			switch stat.Code() {
			case codes.NotFound:
				return nil, ErrUnauthenticated
			}
		}

		return nil, xerrors.Errorf("failed to validate token: %w", err)
	}

	performer, err := ai.userOrgClient.UserFindByID(ctx, &api.UserFindByIDRequest{
		UserId: vt.UserId,
	})

	if err != nil {
		return nil, xerrors.Errorf("failed to validate token: %w", err)
	}

	if targetID == "" {
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
		// continue even if role binding is not found
	}

	roles := RoleBindings(map[string]string{
		"*": performer.User.SystemRoleName,
	})

	if rb != nil && rb.Role != "" {
		roles[targetID] = rb.Role
	}

	perms := getPermissions(roles, targetID)

	return &AuthenticatedMetadata{
		UserID: performer.User.UserId,
		Roles:  roles,

		TargetID:             targetID,
		PermissionsForTarget: perms,
	}, nil
}

func (ai *GatewayAuthorizerInterceptor) wrapContext(ctx context.Context) (context.Context, error) {
	token := GetToken(ctx)
	targetID := GetTargetID(ctx)

	md, err := ai.GetAuthenticatedMetadata(ctx, token, targetID)

	if xerrors.Is(err, ErrUnauthenticated) {
		return nil, status.Error(codes.Unauthenticated, "token is missing or invalid")
	}

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "unauthorized")
	}

	ctx = AddUserIDContext(ctx, md.UserID)
	ctx = AddTargetIDContext(ctx, targetID)
	ctx = AddRolesContext(ctx, md.Roles)
	ctx = AddPermissionsContext(ctx, md.PermissionsForTarget)

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
