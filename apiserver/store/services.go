package store

import (
	"bytes"
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	api "github.com/modoki-paas/modoki-k8s/api"
	"golang.org/x/xerrors"
)

type ServiceSpec api.ServiceSpec

func (ss *ServiceSpec) Scan(src interface{}) error {
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

func (ss *ServiceSpec) Value() (driver.Value, error) {
	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	if err := json.NewEncoder(buf).Encode(ss); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type Service struct {
	ID        int          `db:"id"`
	Owner     int          `db:"owner"`
	Name      string       `db:"name"`
	Spec      *ServiceSpec `db:"spec"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
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
		s.Owner, s.Name, s.Spec,
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
