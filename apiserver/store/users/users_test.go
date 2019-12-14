// +build use_external_db

package users

import (
	"testing"
	"time"

	"github.com/modoki-paas/modoki-k8s/apiserver/testutil"
)

func TestAddUser(t *testing.T) {
	t.Run("success_normal", func(t *testing.T) {
		db := testutil.NewSQLConn(t)
		defer db.Close()

		store := NewUserStore(db)

		seq, err := store.AddUser("test-id", "test-name", UserNormal, UserRoleAdmin)

		if err != nil {
			t.Fatalf("failed to add user: %v", err)
		}

		u, err := store.GetUser(seq)

		if err != nil {
			t.Fatalf("failed to get user: %v", err)
		}

		if u.SeqID <= 0 {
			t.Errorf("id should be >0, but got %v", u.SeqID)
		}
		if u.ID != "test-id" {
			t.Errorf("id should be test-id, but got %v", u.ID)
		}
		if u.UserType != UserNormal {
			t.Errorf("type should be UserNormal, but got %v", u.UserType)
		}
		if u.SystemRole != UserRoleAdmin {
			t.Errorf("role should be UserRoleAdmin(admin), but got %v", u.SystemRole)
		}
		if u.Name != "test-name" {
			t.Errorf("name should be %v, but got %v", "test-name", u.Name)
		}
		if u.CreatedAt == (time.Time{}) {
			t.Error("created_at is not set")
		}
		if u.UpdatedAt == (time.Time{}) {
			t.Error("updated_at is not set")
		}
	})

	t.Run("success_organization", func(t *testing.T) {
		db := testutil.NewSQLConn(t)
		defer db.Close()

		store := NewUserStore(db)

		seq, err := store.AddUser("test-id", "test-name", UserOrganization, UserRoleAdmin)

		if err != nil {
			t.Fatalf("failed to add user: %v", err)
		}

		u, err := store.GetUser(seq)

		if err != nil {
			t.Fatalf("failed to get user: %v", err)
		}

		if u.SeqID <= 0 {
			t.Errorf("id should be >0, but got %v", u.SeqID)
		}
		if u.ID != "test-id" {
			t.Errorf("id should be test-id, but got %v", u.ID)
		}
		if u.UserType != UserOrganization {
			t.Errorf("type should be UserOrganization, but got %v", u.UserType)
		}
		if u.SystemRole != UserRoleAdmin {
			t.Errorf("role should be UserRoleAdmin(admin), but got %v", u.SystemRole)
		}
		if u.Name != "test-name" {
			t.Errorf("name should be %v, but got %v", "test-name", u.Name)
		}
		if u.CreatedAt == (time.Time{}) {
			t.Error("created_at is not set")
		}
		if u.UpdatedAt == (time.Time{}) {
			t.Error("updated_at is not set")
		}
	})
}
