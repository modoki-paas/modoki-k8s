package store

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"context"

	"golang.org/x/xerrors"
)

type ServiceConfig struct {
	Image   string            `json:"image"`
	Command []string          `json:"command"`
	Args    []string          `json:"args"`
	Options map[string]string `json:"options"`
}

func (j *ServiceConfig) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if err := json.Unmarshal(v, j); err != nil {
			return err
		}
	default:
		return xerrors.Errorf("failed to scan json for %v", v)
	}
	return nil
}

func (j ServiceConfig) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type Service struct {
	ID        int            `db:"id"`
	Owner     int            `db:"owner"`
	Name      string         `db:"name"`
	Config    *ServiceConfig `db:"config"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

type serviceStore struct {
	db *dbContext
}

func newServiceStore(db *dbContext) *serviceStore {
	return &serviceStore{db: db}
}

func (ss *serviceStore) AddService(s *Service) (ret *Service, err error) {
	dbx, err := ss.db.Begin(context.Background(), nil)
	store := newDB(dbx)

	if err != nil {
		return nil, xerrors.Errorf("faield to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			dbx.Rollback()
		} else {
			if err := dbx.Commit(); err != nil {
				err = xerrors.Errorf("failed to commit transaction: %v", err)
				ret = nil
			}
		}
	}()

	res, err := dbx.db.ExecContext(
		context.Background(),
		`INSERT INTO services
			(owner, name, config)
			VALUES (?, ?, ?)`,
		s.Owner, s.Name, s.Config,
	)

	if err != nil {
		return nil, xerrors.Errorf("failed to add service to db: %v", err)
	}

	id64, err := res.LastInsertId()

	if err != nil {
		return nil, xerrors.Errorf("failed to add service to db: %v", err)
	}

	return store.Service().GetService(int(id64))
}

func (ss *serviceStore) GetService(id int) (*Service, error) {
	var service Service
	err := ss.db.db.
		QueryRowxContext(context.Background(), "SELECT * FROM services WHERE id=?", id).
		StructScan(&service)

	if err != nil {
		return nil, xerrors.Errorf("failed to retrieve service: %v", err)
	}

	return &service, nil
}
