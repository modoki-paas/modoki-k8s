package auth

import "encoding/json"

type RoleBindings map[string]string

func (r *RoleBindings) Marshal() string {
	b, _ := json.Marshal(r)

	return string(b)
}
