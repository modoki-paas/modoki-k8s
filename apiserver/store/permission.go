package store

import (
	"database/sql/driver"
	"encoding/json"

	"golang.org/x/xerrors"
)

type Permission struct {
}

func (j *Permission) Scan(src interface{}) error {
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

func (j Permission) Value() (driver.Value, error) {
	return json.Marshal(j)
}
