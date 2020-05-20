package roles

import "github.com/modoki-paas/modoki-k8s/pkg/rbac/permissions"

var (
	SystemAuth = NewSystemRole(
		"auth",
		permissions.UserCreate, permissions.UserGetAll, permissions.UserDelete, permissions.UserGet, permissions.UserUpdate,
		permissions.OrgCreate, permissions.OrgList, permissions.OrgListAll, permissions.OrgDelete, permissions.OrgUpdate, permissions.OrgMemberManagement, permissions.OrgMemberList,
		permissions.UserOrgGetRoleBinding,
		permissions.TokenIssue, permissions.TokenValidate,
	)

	SystemAdmin = NewSystemRole(
		"system_admin",
		permissions.UserGetAll, permissions.UserDelete, permissions.UserGet, permissions.UserUpdate,
		permissions.TokenIssue, permissions.TokenDelete,
		permissions.OrgCreate, permissions.OrgList, permissions.OrgListAll, permissions.OrgDelete, permissions.OrgUpdate, permissions.OrgUpdate,
		permissions.AppCreate, permissions.AppUpdate, permissions.AppDelete, permissions.AppList, permissions.AppStatus,
	)

	SystemDeveloper = NewSystemRole(
		"system_developer",
		permissions.UserGetAll,
		permissions.OrgCreate, permissions.OrgList,
	)
)
