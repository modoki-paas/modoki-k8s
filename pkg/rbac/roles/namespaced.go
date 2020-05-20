package roles

import "github.com/modoki-paas/modoki-k8s/pkg/rbac/permissions"

var (
	OrgOwner = NewRole(
		"org_owner",
		permissions.NamespaceOrg,
		permissions.AppCreate, permissions.AppDelete, permissions.AppUpdate, permissions.AppList, permissions.AppStatus,
		permissions.OrgUpdate, permissions.OrgDelete,
		permissions.TokenIssue, permissions.TokenDelete,
	)

	OrgAdmin = NewRole(
		"org_developer",
		permissions.NamespaceOrg,
		permissions.AppCreate, permissions.AppDelete, permissions.AppUpdate, permissions.AppList, permissions.AppStatus,
		permissions.TokenIssue, permissions.TokenDelete,
	)
)

var (
	Self = NewRole(
		"user_self",
		permissions.NamespaceUser,
		permissions.AppCreate, permissions.AppDelete, permissions.AppUpdate, permissions.AppList, permissions.AppStatus,
		permissions.TokenIssue, permissions.TokenDelete,
		permissions.UserDelete, permissions.UserGet, permissions.UserUpdate,
	)
)
