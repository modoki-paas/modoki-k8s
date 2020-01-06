// build +use_external_db

package tokens

import (
	"testing"
	"time"

	"github.com/modoki-paas/modoki-k8s/internal/testutil"
	"github.com/modoki-paas/modoki-k8s/pkg/types"
)

func TestAddToken(t *testing.T) {
	t.Run("success_normal", func(t *testing.T) {
		tk := &types.Token{
			ID:     "first-id",
			Owner:  "owner-id",
			Author: "aurhor-id",
		}

		db := testutil.NewSQLConn(t)
		defer db.Close()

		store := NewTokenStore(db)
		seq, token, err := store.AddToken(tk)

		if err != nil {
			t.Fatalf("failed to add token: %v", err)
		}

		ret, err := store.GetToken(seq)

		if err != nil {
			t.Fatalf("failed to get token: %v", err)
		}

		if ret.SeqID <= 0 {
			t.Errorf("id should be >0, but got %v", ret.ID)
		}
		if ret.ID != "first-id" {
			t.Errorf("id should not be empty, but got empty id")
		}
		if ret.ID != tk.ID {
			t.Errorf("App.ID: %s, returned id: %s", ret.ID, tk.ID)
		}
		if ret.Owner != tk.Owner {
			t.Errorf("invalid owner: want %v got %v", tk.Owner, ret.Owner)
		}
		if ret.Author != tk.Author {
			t.Errorf("invalid author: want %v got %v", tk.Author, ret.Author)
		}
		if ret.Token != token {
			t.Errorf("invalid name: want %v got %v", token, ret.Token)
		}

		if ret.CreatedAt == (time.Time{}) {
			t.Errorf("invalid created_at: %v", ret.CreatedAt)
		}
		if ret.UpdatedAt == (time.Time{}) {
			t.Errorf("invalid updated_at: %v", ret.UpdatedAt)
		}
	})
}

func TestFindTokenByID(t *testing.T) {
	t.Run("success_normal", func(t *testing.T) {
		tk := &types.Token{
			ID:     "first-id",
			Owner:  "owner-id",
			Author: "aurhor-id",
		}
		db := testutil.NewSQLConn(t)
		defer db.Close()

		store := NewTokenStore(db)
		seq, token, err := store.AddToken(tk)

		if err != nil {
			t.Fatalf("failed to add token: %+v", err)
		}

		ret, err := store.FindTokenByID(tk.Owner, tk.ID)

		if err != nil {
			t.Fatalf("failed to find token: %+v", err)
		}

		if ret.SeqID != seq {
			t.Errorf("id should be >0, but got %v", ret.ID)
		}
		if ret.ID != "first-id" {
			t.Errorf("id should not be empty, but got empty id")
		}
		if ret.ID != tk.ID {
			t.Errorf("App.ID: %s, returned id: %s", ret.ID, tk.ID)
		}
		if ret.Owner != tk.Owner {
			t.Errorf("invalid owner: want %v got %v", tk.Owner, ret.Owner)
		}
		if ret.Author != tk.Author {
			t.Errorf("invalid author: want %v got %v", tk.Author, ret.Author)
		}
		if ret.Token != token {
			t.Errorf("invalid name: want %v got %v", token, ret.Token)
		}

		if ret.CreatedAt == (time.Time{}) {
			t.Errorf("invalid created_at: %v", ret.CreatedAt)
		}
		if ret.UpdatedAt == (time.Time{}) {
			t.Errorf("invalid updated_at: %v", ret.UpdatedAt)
		}
	})
}

func TestFindTokenByToken(t *testing.T) {
	t.Run("success_normal", func(t *testing.T) {
		tk := &types.Token{
			ID:     "first-id",
			Owner:  "owner-id",
			Author: "aurhor-id",
		}
		db := testutil.NewSQLConn(t)
		defer db.Close()

		store := NewTokenStore(db)
		seq, token, err := store.AddToken(tk)

		if err != nil {
			t.Fatalf("failed to add token: %+v", err)
		}

		ret, err := store.FindTokenByToken(token)

		if err != nil {
			t.Fatalf("failed to find token: %+v", err)
		}

		if ret.SeqID != seq {
			t.Errorf("id should be >0, but got %v", ret.ID)
		}
		if ret.ID != "first-id" {
			t.Errorf("id should not be empty, but got empty id")
		}
		if ret.ID != tk.ID {
			t.Errorf("App.ID: %s, returned id: %s", ret.ID, tk.ID)
		}
		if ret.Owner != tk.Owner {
			t.Errorf("invalid owner: want %v got %v", tk.Owner, ret.Owner)
		}
		if ret.Author != tk.Author {
			t.Errorf("invalid author: want %v got %v", tk.Author, ret.Author)
		}
		if ret.Token != token {
			t.Errorf("invalid name: want %v got %v", token, ret.Token)
		}

		if ret.CreatedAt == (time.Time{}) {
			t.Errorf("invalid created_at: %v", ret.CreatedAt)
		}
		if ret.UpdatedAt == (time.Time{}) {
			t.Errorf("invalid updated_at: %v", ret.UpdatedAt)
		}
	})
}
