package testutil

import (
	"testing"
)

func TestNewSQLConn(t *testing.T) {
	conn := NewSQLConn(t)

	conn.Close()
}
