package store

import "time"

type UserGroupRelation struct {
	ID        int       `db:"id"`
	UserType  string    `db:"type"`
	Name      string    `db:"name"`
	Password  []byte    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type userGroupRelationsStore struct {
	db *dbContext
}

func newUserGroupRelationsStore(db *dbContext) *userGroupRelationsStore {
	return &userGroupRelationsStore{db: db}
}

