package roles

import "github.com/modoki-paas/modoki-k8s/pkg/rbac/permissions"

type Role struct {
	Name        string                    `json:"name" yaml:"name"`
	Type        permissions.NamespaceType `json:"type" yaml:"type"`
	Permissions []*permissions.Permission `json:"permissions" yaml:"permissions"`
}

func NewRole(name string, nsType permissions.NamespaceType, perms ...*permissions.Permission) *Role {
	r := &Role{
		Name:        name,
		Type:        nsType,
		Permissions: perms,
	}

	for i := range perms {
		if !perms[i].Namespaced(nsType) {
			panic("roles only can contain namespaced permissions")
		}
	}

	return r
}

type SystemRole struct {
	Name        string                    `json:"name" yaml:"name"`
	Permissions []*permissions.Permission `json:"permissions" yaml:"permissions"`
}

func NewSystemRole(name string, perms ...*permissions.Permission) *SystemRole {
	r := &SystemRole{
		Name:        name,
		Permissions: perms,
	}

	return r
}
