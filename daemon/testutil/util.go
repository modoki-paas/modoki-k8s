package testutil

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// NewSQLMock returns sqlmock with sql*x*.DB
func NewSQLMock(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	dbx := sqlx.NewDb(db, "mysql")
	return dbx, mock
}

// NewSQLConn returns connection with sql*x*.DB
func NewSQLConn(t *testing.T) *sqlx.DB {
	dbx, err := sqlx.Connect("mysql", "root:password@tcp(127.0.0.1)/testdb?parseTime=true&multiStatements=true")

	if err != nil {
		t.Fatalf("failed to connect to external db for test: %v", err)
	}

	if _, err := dbx.Exec("DROP DATABASE IF EXISTS testdb"); err != nil {
		t.Fatalf("failed to delete database: %v", err)
	}
	if _, err := dbx.Exec("CREATE DATABASE testdb"); err != nil {
		t.Fatalf("failed to delete database: %v", err)
	}
	dbx.Close()

	dbx, err = sqlx.Connect("mysql", "root:password@tcp(127.0.0.1)/testdb?parseTime=true&multiStatements=true")

	if err != nil {
		t.Fatalf("failed to connect to external db for test: %v", err)
	}

	createTable(t, dbx)

	return dbx
}

func createTable(t *testing.T, dbx *sqlx.DB) {
	sqls, err := filepath.Glob("../schema/*.sql")

	if err != nil {
		t.Fatalf("failed to open sql files: %v", err)
	}
	for i := range sqls {
		t.Logf("running %s", sqls[i])
		b, err := ioutil.ReadFile(sqls[i])

		if err != nil {
			t.Fatalf("failed to open %s: %v", sqls[i], err)
		}

		_, err = dbx.Exec(string(b))

		if err != nil {
			t.Fatalf("execute sql in %s: %v", sqls[i], err)
		}
	}
}
