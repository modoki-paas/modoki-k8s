// +build use_external_db

package users

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/modoki-paas/modoki-k8s/internal/dbutil"
	"github.com/modoki-paas/modoki-k8s/internal/testutil"
	"github.com/modoki-paas/modoki-k8s/pkg/rbac/roles"
)

func TestRegisterRoleBinding(t *testing.T) {
	t.Run("success_normal", func(t *testing.T) {
		db := testutil.NewSQLConn(t)
		defer db.Close()

		dbutil.Transaction(context.Background(), db, func(tx *sqlx.Tx) error {
			store := NewRoleBindingsStore(db)

			seq, err := store.RegisterRoleBinding(0, 10, roles.SystemAdmin.Name)

			if err != nil {
				t.Fatalf("failed to register role binding: %+v", err)
			}

			if seq <= 0 {
				t.Error("seq should be >0")
			}

			r, err := store.GetRoleBinding(0, 10)

			if err != nil {
				t.Fatalf("failed to get role binding: %+v", err)
			}

			if r.Name != roles.SystemAdmin.Name {
				t.Errorf("role names differ(actual: %s, expected: %s)", r.Name, roles.SystemAdmin.Name)
			}

			return nil
		})
	})
}
