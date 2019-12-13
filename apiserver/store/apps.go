package store

import (
	"bytes"
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/rs/xid"
	"golang.org/x/xerrors"
)

type AppSpec api.AppSpec

func (ss *AppSpec) Scan(src interface{}) error {
	var s []byte
	switch v := src.(type) {
	case []byte:
		s = v
	case string:
		s = []byte(v)
	default:
		return errors.New("failed to scan JsonObject")
	}

	if err := json.NewDecoder(bytes.NewReader(s)).Decode(ss); err != nil {
		return err
	}
	return nil
}

func (ss *AppSpec) Value() (driver.Value, error) {
	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	if err := json.NewEncoder(buf).Encode(ss); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type App struct {
	SeqID     int       `db:"seq"`
	ID        string    `db:"id"`
	Owner     int       `db:"owner"`
	Name      string    `db:"name"`
	Spec      *AppSpec  `db:"spec"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type appStore struct {
	db *dbContext
}

func newAppStore(db *dbContext) *appStore {
	return &appStore{db: db}
}

func (ss *appStore) AddApp(s *App) (ret *App, err error) {
	dbx, err := ss.db.Begin(context.Background(), nil)
	store := newDB(dbx)

	if err != nil {
		return nil, xerrors.Errorf("faield to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			dbx.Rollback()
		} else {
			if err := dbx.Commit(); err != nil {
				err = xerrors.Errorf("failed to commit transaction: %w", err)
				ret = nil
			}
		}
	}()

	s.ID = xid.New().String()

	res, err := dbx.db.ExecContext(
		context.Background(),
		`INSERT INTO apps
			(app_id, owner, name, spec)
			VALUES (?, ?, ?, ?)`,
		s.ID, s.Owner, s.Name, s.Spec,
	)

	if err != nil {
		return nil, xerrors.Errorf("failed to add app to db: %w", err)
	}

	id64, err := res.LastInsertId()

	if err != nil {
		return nil, xerrors.Errorf("failed to add app to db: %w", err)
	}

	return store.App().GetApp(int(id64))
}

func (ss *appStore) GetApp(seq int) (*App, error) {
	var app App
	err := ss.db.db.
		QueryRowxContext(context.Background(), "SELECT * FROM apps WHERE seq=?", seq).
		StructScan(&app)

	if err != nil {
		return nil, xerrors.Errorf("failed to retrieve app: %w", err)
	}

	return &app, nil
}

func (ss *appStore) FindAppByID(id string) (*App, error) {
	var app App
	err := ss.db.db.
		QueryRowxContext(context.Background(), "SELECT * FROM apps WHERE seq=?", id).
		StructScan(&app)

	if err != nil {
		return nil, xerrors.Errorf("failed to retrieve app: %w", err)
	}

	return &app, nil
}
