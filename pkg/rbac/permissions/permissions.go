package permissions

type NamespaceType string

var (
	NamespaceUser NamespaceType = "user"
	NamespaceOrg  NamespaceType = "organization"
)

type Permission struct {
	Name       string          `json:"name" yaml:"name"`
	Namespaces []NamespaceType `json:"namespaces" yaml:"namespaces"`
}

func NewPermission(name string, namespaces ...NamespaceType) *Permission {
	return &Permission{
		Name:       name,
		Namespaces: namespaces,
	}
}

func (p *Permission) Namespaced(nsType NamespaceType) bool {
	for i := range p.Namespaces {
		if p.Namespaces[i] == nsType {
			return true
		}
	}

	return false
}

type Role struct {
	Permissions []*Permission `json:"permissions" yaml:"permissions"`
}
