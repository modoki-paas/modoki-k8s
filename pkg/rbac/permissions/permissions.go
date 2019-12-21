package permissions

type Permission struct {
	Name       string `json:"name" yaml:"name"`
	Namespaced bool   `json:"namespaced" yaml:"namespaced"`
}

// PermissionBinding represents binding b/w permission and user/org
type PermissionBinding struct {
	Permission string `json:"permission" yaml:"permission"`
	Target     string `json:"target" yaml:"target"`
}

func NewPermission(name string, namespaced bool) *Permission {
	return &Permission{
		Name:       name,
		Namespaced: namespaced,
	}
}

type Role struct {
	Permissions []*Permission `json:"permissions" yaml:"permissions"`
}
