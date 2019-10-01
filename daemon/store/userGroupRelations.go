package store

import (
	"database/sql/driver"
	"encoding/json"

	"golang.org/x/xerrors"
)

type UserPermission struct {
}

func (j *UserPermission) Scan(src interface{}) error {
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

func (j UserPermission) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type UserGroupRelation struct {
	ID         int             `db:"id"`
	GroupID    int             `db:"group_id"`
	UserID     int             `db:"user_id"`
	Permission *UserPermission `db:"permission"`
}

type userGroupRelationsStore struct {
	db *dbContext
}

func newUserGroupRelationsStore(db *dbContext) *userGroupRelationsStore {
	return &userGroupRelationsStore{db: db}
}
