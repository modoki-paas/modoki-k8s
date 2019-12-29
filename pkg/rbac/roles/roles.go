package roles

import (
	"sync"

	"github.com/modoki-paas/modoki-k8s/pkg/rbac/permissions"
)

var (
	rolesLock sync.RWMutex
	roles     map[string]*Role

	systemRolesLock sync.RWMutex
	systemRoles     map[string]*SystemRole
)

func addRole(name string, role *Role) {
	rolesLock.Lock()
	defer rolesLock.Unlock()

	if roles == nil {
		roles = map[string]*Role{}
	}

	roles[name] = role
}

func FindRole(name string) *Role {
	rolesLock.RLock()
	defer rolesLock.RUnlock()

	r, ok := roles[name]

	if ok {
		return r
	}

	return nil
}

func addSystemRole(name string, role *SystemRole) {
	systemRolesLock.Lock()
	defer systemRolesLock.Unlock()

	if systemRoles == nil {
		systemRoles = map[string]*SystemRole{}
	}

	systemRoles[name] = role
}

func FindSystemRole(name string) *SystemRole {
	systemRolesLock.RLock()
	defer systemRolesLock.RUnlock()

	r, ok := systemRoles[name]
	if ok {
		return r
	}

	return nil
}

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

	addRole(name, r)

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

	addSystemRole(name, r)

	return r
}
