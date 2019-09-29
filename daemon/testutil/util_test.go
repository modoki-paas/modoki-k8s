// build +test_db

package testutil

import (
	"testing"
)

func TestNewSQLConn(t *testing.T) {
	t.Run("sql conn test", func(t *testing.T) {
		conn := NewSQLConn(t)

		conn.Close()
	})
}
