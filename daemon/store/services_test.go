// build +use_external_db

package store

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/modoki-paas/modoki-k8s/daemon/testutil"
)

func TestAddService(t *testing.T) {
	t.Run("success_normal", func(t *testing.T) {
		db := testutil.NewSQLConn(t)
		defer db.Close()

		store := NewDB(db)

		s := &Service{
			Owner: 10,
			Name:  "service-name",
			Config: &ServiceConfig{
				Image: "image-name",
			},
		}

		ret, err := store.Service().AddService(s)

		if err != nil {
			t.Fatalf("failed to add user: %v", err)
		}

		if ret.ID <= 0 {
			t.Errorf("id should be >0, but got %v", ret.ID)
		}
		if ret.Owner != s.Owner {
			t.Errorf("invalid owner: want %v got %v", s.Owner, ret.Owner)
		}
		if ret.Name != s.Name {
			t.Errorf("invalid name: want %v got %v", s.Name, ret.Name)
		}
		if !cmp.Equal(ret.Config, s.Config) {
			t.Errorf("invalid config: want %v got %v", s.Config, ret.Config)
		}

		if ret.CreatedAt == (time.Time{}) {
			t.Errorf("invalid created_at: %v", ret.CreatedAt)
		}
		if ret.UpdatedAt == (time.Time{}) {
			t.Errorf("invalid updated_at: %v", ret.UpdatedAt)
		}
	})
}
