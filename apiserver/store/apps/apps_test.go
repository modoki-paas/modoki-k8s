// build +use_external_db

package apps

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/modoki-paas/modoki-k8s/internal/testutil"
	"github.com/modoki-paas/modoki-k8s/pkg/types"
)

func TestAddApp(t *testing.T) {
	t.Run("success_normal", func(t *testing.T) {
		s := &types.App{
			Owner: 10,
			Name:  "app-name",
			Spec: &types.AppSpec{
				Image: "image-name",
			},
		}

		db := testutil.NewSQLConn(t)
		defer db.Close()

		store := NewAppStore(db)
		seq, err := store.AddApp(s)

		if err != nil {
			t.Fatalf("failed to add app: %v", err)
		}

		ret, err := store.GetApp(seq)

		if err != nil {
			t.Fatalf("failed to get app: %v", err)
		}

		if ret.SeqID <= 0 {
			t.Errorf("id should be >0, but got %v", ret.ID)
		}
		if ret.ID == "" {
			t.Errorf("id should not be empty, but got empty id")
		}
		if ret.Owner != s.Owner {
			t.Errorf("invalid owner: want %v got %v", s.Owner, ret.Owner)
		}
		if ret.Name != s.Name {
			t.Errorf("invalid name: want %v got %v", s.Name, ret.Name)
		}
		if !cmp.Equal(ret.Spec, s.Spec) {
			t.Errorf("invalid spec: want %v got %v", s.Spec, ret.Spec)
		}

		if ret.CreatedAt == (time.Time{}) {
			t.Errorf("invalid created_at: %v", ret.CreatedAt)
		}
		if ret.UpdatedAt == (time.Time{}) {
			t.Errorf("invalid updated_at: %v", ret.UpdatedAt)
		}
	})
}
