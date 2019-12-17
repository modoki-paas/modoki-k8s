package types

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	api "github.com/modoki-paas/modoki-k8s/api"
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
