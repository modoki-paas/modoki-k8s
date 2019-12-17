package testutil

import (
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// NewSQLMock returns sqlmock with sql*x*.DB
func NewSQLMock(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	dbx := sqlx.NewDb(db, "mysql")
	return dbx, mock
}

type sqlxConn struct {
	dbName string
	*sqlx.DB
}

func (c *sqlxConn) Close() error {
	c.DB.Exec("DROP DATABASE " + c.dbName)

	return c.DB.Close()
}

// NewSQLConn returns connection with sql*x*.DB
func NewSQLConn(t *testing.T) *sqlxConn {
	t.Helper()
	dbx, err := sqlx.Connect("mysql", "root:password@tcp(127.0.0.1)/?parseTime=true&multiStatements=true")

	if err != nil {
		t.Fatalf("failed to connect to external db for test: %v", err)
	}

	var lts = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]byte, 10)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = lts[rand.Intn(len(lts))]
	}
	dbName := string(b)

	if _, err := dbx.Exec("CREATE DATABASE " + dbName); err != nil {
		t.Fatalf("failed to create database: %v", err)
	}
	if _, err := dbx.Exec("USE " + dbName); err != nil {
		t.Fatalf("failed to select database: %v", err)
	}

	createTable(t, dbx)

	return &sqlxConn{
		dbName: dbName,
		DB:     dbx,
	}
}

func createTable(t *testing.T, dbx *sqlx.DB) {
	t.Helper()
	sqls, err := filepath.Glob("../../../schema/*.sql")

	if err != nil {
		t.Fatalf("failed to open sql files: %v", err)
	}
	t.Log(len(sqls))
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
