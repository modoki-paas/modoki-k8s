package apps

import (
	"bytes"
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
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

type AppStore struct {
	db sqlx.ExtContext
}

func NewAppStore(db sqlx.ExtContext) *AppStore {
	return &AppStore{db: db}
}

func (ss *AppStore) AddApp(s *App) (seqID int, err error) {
	s.ID = xid.New().String()

	res, err := ss.db.ExecContext(
		context.Background(),
		`INSERT INTO apps
			(id, owner, name, spec)
			VALUES (?, ?, ?, ?)`,
		s.ID, s.Owner, s.Name, s.Spec,
	)

	if err != nil {
		return 0, xerrors.Errorf("failed to add app to db: %w", err)
	}

	id64, err := res.LastInsertId()

	if err != nil {
		return 0, xerrors.Errorf("failed to add app to db: %w", err)
	}

	return int(id64), nil
}

func (ss *AppStore) GetApp(seq int) (*App, error) {
	var app App
	err := ss.db.
		QueryRowxContext(context.Background(), "SELECT * FROM apps WHERE seq=?", seq).
		StructScan(&app)

	if err != nil {
		return nil, xerrors.Errorf("failed to retrieve app: %w", err)
	}

	return &app, nil
}

func (ss *AppStore) FindAppByID(id string) (*App, error) {
	var app App
	err := ss.db.
		QueryRowxContext(context.Background(), "SELECT * FROM apps WHERE seq=?", id).
		StructScan(&app)

	if err != nil {
		return nil, xerrors.Errorf("failed to retrieve app: %w", err)
	}

	return &app, nil
}
