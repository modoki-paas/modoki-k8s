package auth

import (
	"context"

	"github.com/modoki-paas/modoki-k8s/pkg/rbac/permissions"
	"golang.org/x/xerrors"
)

type contextKey string

const (
	UserIDContext      contextKey = "user-id-context"
	TargetIDContext    contextKey = "target-id-context"
	RolesContext       contextKey = "roles-context"
	PermissionsContext contextKey = "role-context"

	UserIDHeader   = "X-Modoki-Executor-User-ID"
	TargetIDHeader = "X-Modoki-Target-ID"
	RolesHeader    = "X-Modoki-Roles"
)

func addUserIDContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, UserIDContext, id)
}

func addTargetIDContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, TargetIDContext, id)
}

func addRolesContext(ctx context.Context, rb RoleBindings) context.Context {
	return context.WithValue(ctx, RolesContext, rb)
}

func addPermissionsContext(ctx context.Context, perms map[string]struct{}) context.Context {
	return context.WithValue(ctx, PermissionsContext, perms)
}

func getUserIDContext(ctx context.Context) string {
	return ctx.Value(UserIDContext).(string)
}

func getTargetIDContext(ctx context.Context) string {
	return ctx.Value(TargetIDContext).(string)
}

func getRolesContext(ctx context.Context) RoleBindings {
	return ctx.Value(RolesContext).(RoleBindings)
}

func getPermissionsContext(ctx context.Context) (perms map[string]struct{}) {
	return ctx.Value(PermissionsContext).(map[string]struct{})
}

func IsAuthorized(ctx context.Context, required ...*permissions.Permission) error {
	permsMap := getPermissionsContext(ctx)

	for i := range required {
		if _, ok := permsMap[required[i].Name]; !ok {
			return xerrors.Errorf("lacking permission: %s", required[i].Name)
		}
	}

	return nil
}
