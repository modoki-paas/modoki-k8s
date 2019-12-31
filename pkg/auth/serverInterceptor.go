package auth

import (
	"context"
	"encoding/json"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/modoki-paas/modoki-k8s/pkg/rbac/roles"
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

func (ai *AuthorizerInterceptor) getRoles(md metadata.MD) RoleBindings {
	arr := md.Get(RolesHeader)

	if len(arr) == 0 {
		return nil
	}

	var rb RoleBindings
	if err := json.Unmarshal([]byte(arr[0]), &rb); err != nil {
		return nil
	}

	return rb
}

func (ai *AuthorizerInterceptor) getPermissions(rb RoleBindings, targetID string) (permMap map[string]struct{}) {
	if rb == nil {
		return nil
	}

	perms := []string{}

	systemRoleName, ok := rb["*"]

	if ok {
		systemRole := roles.FindSystemRole(systemRoleName)

		if systemRole != nil {
			for i := range systemRole.Permissions {
				perms = append(perms, systemRole.Permissions[i].Name)
			}
		}
	}

	roleName, ok := rb[targetID]

	if ok {
		role := roles.FindRole(roleName)

		if role != nil {
			for i := range role.Permissions {
				perms = append(perms, role.Permissions[i].Name)
			}
		}
	}

	permMap = map[string]struct{}{}

	for i := range perms {
		permMap[perms[i]] = struct{}{}
	}

	return permMap
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

	ctx = addUserIDContext(ctx, id)

	targetID, ok := ai.getTargetID(md)

	if !ok {
		targetID = id
	}

	ctx = addTargetIDContext(ctx, targetID)

	roles := ai.getRoles(md)
	ctx = addRolesContext(ctx, roles)

	perms := ai.getPermissions(roles, targetID)
	ctx = addPermissionsContext(ctx, perms)

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
