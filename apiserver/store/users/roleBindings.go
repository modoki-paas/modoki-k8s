package users

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/modoki-paas/modoki-k8s/pkg/rbac/roles"
	"golang.org/x/xerrors"
)

type RoleBindingsStore struct {
	db sqlx.ExtContext
}

func NewRoleBindingsStore(db sqlx.ExtContext) *RoleBindingsStore {
	return &RoleBindingsStore{db: db}
}

// RegisterRoleBinding registers or updated role binding
// MUST BE CALLED IN TRANSACTION
func (s *RoleBindingsStore) RegisterRoleBinding(userSeqID, targetSeqID int, roleName string) (int, error) {
	var seq int
	err := s.db.QueryRowxContext(
		context.Background(),
		"SELECT seq FROM role_bindings WHERE user_seq=? AND target_seq=? FOR UPDATE",
		userSeqID, targetSeqID,
	).Scan(&seq)

	if err != nil && err != sql.ErrNoRows {
		return 0, xerrors.Errorf("failed to retrieve existing role_binding: %w", err)
	}

	if err == sql.ErrNoRows {
		_, err := s.db.ExecContext(
			context.Background(),
			"UPDATE role_bindings SET role_name=? WHERE seq=?",
			roleName, seq,
		)

		if err != nil {
			return 0, xerrors.Errorf("failed to update existing role_binding: %w", err)
		}
	}

	res, err := s.db.ExecContext(
		context.Background(),
		"INSERT INTO role_bindings (user_seq, target_seq, role_name) VALUES (?, ?, ?)",
		userSeqID, targetSeqID, roleName,
	)

	if err != nil {
		return 0, xerrors.Errorf("failed to insert new role_binding: %w", err)
	}

	id64, err := res.LastInsertId()

	if err != nil {
		return 0, xerrors.Errorf("failed to get last insert id: %w", err)
	}

	return int(id64), nil
}

func (s *RoleBindingsStore) GetRoleBinding(userSeqID, targetSeqID int) (*roles.Role, error) {
	var roleName string
	err := s.db.QueryRowxContext(
		context.Background(),
		"SELECT role_name FROM role_bindings WHERE user_seq=? AND target_seq=?",
		userSeqID, targetSeqID,
	).Scan(&roleName)

	if err != nil {
		return nil, xerrors.Errorf("failed to get role binding from db: %w", err)
	}

	role := roles.FindRole(roleName)

	if role == nil {
		return nil, xerrors.Errorf("unknown role name: %s", roleName)
	}

	return role, nil
}
