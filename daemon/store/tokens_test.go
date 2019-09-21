// +build test_db

package store

import (
	"testing"

	"github.com/modoki-paas/modoki-k8s/daemon/testutil"
)

func TestGetFromToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := testutil.NewSQLMock(t)

		mock.ExpectQuery("SELECT ")

		store := NewDB(db)

		store.Token().GetFromToken("my-token")
	})
}
