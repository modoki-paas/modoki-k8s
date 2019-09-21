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

		store := NewDB(db)

		tk := &Token{
			Token:           "token",
			Organization:    1,
			Author:          10,
			TokenPermission: nil,
		}

		_, err := store.Token().NewToken(tk)

		if err != nil {
			t.Errorf("failed to register new token: %v", err)
		}

		ret, err := store.Token().GetFromToken("token")

		if err != nil {
			t.Errorf("failed to retrieve token: %v", err)
		}

		if tk.ID != ret.ID {
			t.Error("invalid id: want %v got %v", tk.ID, ret.ID)
		}
		if tk.Token != ret.Token {
			t.Error("invalid token: want %v got %v", tk.Token, ret.Token)
		}
		if tk.Organization != ret.Organization {
			t.Error("invalid org: want %v got %v", tk.Organization, ret.Organization)
		}
		if tk.Author != ret.Author {
			t.Error("invalid author: want %v got %v", tk.Author, ret.Author)
		}
		if ret.CreatedAt == (time.Time{}) {
			t.Error("invalid created_at", tk.CreatedAt)
		}
		if ret.UpdatedAt == (time.Time{}) {
			t.Error("invalid updated_at", tk.UpdatedAt)
		}
	})
}
