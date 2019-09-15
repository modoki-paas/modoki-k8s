package store

import (
	"database/sql/driver"
	"encoding/json"
	"time"

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
