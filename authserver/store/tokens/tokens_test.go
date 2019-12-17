// +build use_external_db

package tokens

import (
	"testing"
	"time"

	"github.com/modoki-paas/modoki-k8s/internal/testutil"
	"github.com/modoki-paas/modoki-k8s/pkg/types"
)

func TestGetFromToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tk := &types.Token{
			Token:  "token",
			Owner:  1,
			Author: 10,
		}

		db := testutil.NewSQLConn(t)
		defer db.Close()

		store := NewTokenStore(db)
		_, err := store.AddToken(tk)

		if err != nil {
			t.Fatalf("failed to register new token: %v", err)
		}

		ret, err := store.GetFromToken("token")

		if err != nil {
			t.Fatalf("failed to retrieve token: %v", err)
		}

		if ret.SeqID == 0 {
			t.Error("id should be not zero")
		}
		if tk.Token != ret.Token {
			t.Errorf("invalid token: want %v got %v", tk.Token, ret.Token)
		}
		if tk.Owner != ret.Owner {
			t.Errorf("invalid owner: want %v got %v", tk.Owner, ret.Owner)
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
