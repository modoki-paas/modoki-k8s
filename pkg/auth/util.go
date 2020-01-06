package auth

import (
	"encoding/json"

	"github.com/modoki-paas/modoki-k8s/pkg/rbac/roles"
	"google.golang.org/grpc/metadata"
)

func getRoles(md metadata.MD) RoleBindings {
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

func getPermissions(rb RoleBindings, targetID string) (permMap map[string]struct{}) {
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
