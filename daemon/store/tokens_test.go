// +build test_db

package store

import (
	"testing"
	"time"

	"github.com/modoki-paas/modoki-k8s/daemon/testutil"
)

func TestGetFromToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db := testutil.NewSQLConn(t)
		defer db.Close()

		store := NewDB(db)

		tk := &Token{
			Token:           "token",
			Organization:    1,
			Author:          10,
			TokenPermission: nil,
		}

		_, err := store.Token().AddToken(tk)

		if err != nil {
			t.Fatalf("failed to register new token: %v", err)
		}

		ret, err := store.Token().GetFromToken("token")

		if err != nil {
			t.Fatalf("failed to retrieve token: %v", err)
		}

		if tk.ID != ret.ID {
			t.Errorf("invalid id: want %v got %v", tk.ID, ret.ID)
		}
		if tk.Token != ret.Token {
			t.Errorf("invalid token: want %v got %v", tk.Token, ret.Token)
		}
		if tk.Organization != ret.Organization {
			t.Errorf("invalid org: want %v got %v", tk.Organization, ret.Organization)
		}
		if tk.Author != ret.Author {
			t.Errorf("invalid author: want %v got %v", tk.Author, ret.Author)
		}
		if ret.CreatedAt == (time.Time{}) {
			t.Errorf("invalid created_at: %v", tk.CreatedAt)
		}
		if ret.UpdatedAt == (time.Time{}) {
			t.Errorf("invalid updated_at: %v", tk.UpdatedAt)
		}
	})
}
