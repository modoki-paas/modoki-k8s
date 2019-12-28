package roles

import "github.com/modoki-paas/modoki-k8s/pkg/rbac/permissions"

var (
	SystemAdmin = NewSystemRole(
		"system_admin",
		permissions.UserList, permissions.UserDelete, permissions.UserGet, permissions.UserUpdate,
		permissions.TokenIssue, permissions.TokenDelete,
		permissions.OrgCreate, permissions.OrgList, permissions.OrgListAll, permissions.OrgDelete, permissions.OrgUpdate, permissions.OrgUpdate,
		permissions.AppCreate, permissions.AppUpdate, permissions.AppDelete, permissions.AppList,
	)

	SystemDeveloper = NewSystemRole(
		"system_developer",
		permissions.UserList,
		permissions.OrgCreate, permissions.OrgList,
	)
)
