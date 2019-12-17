package users

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"golang.org/x/xerrors"
)

type User struct {
	SeqID      int            `db:"seq"`
	ID         string         `db:"id"`
	UserType   UserTypeEnum   `db:"type"`
	Name       string         `db:"name"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  time.Time      `db:"updated_at"`
	SystemRole UserSystemRole `db:"system_role"`
}

type UserStore struct {
	db sqlx.ExtContext
}

func NewUserStore(db sqlx.ExtContext) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) AddUser(id, name string, userType UserTypeEnum, role UserSystemRole) (seqID int, err error) {
	u := &User{
		ID:         id,
		UserType:   userType,
		Name:       name,
		SystemRole: UserSystemRole(role),
	}

	res, err := s.db.ExecContext(
		context.Background(),
		`INSERT INTO users (
			type,
			id,
			name,
			system_role
		) VALUES (?, ?, ?, ?)`, u.UserType, u.ID, u.Name, u.SystemRole)

	if err != nil {
		return 0, xerrors.Errorf("faield to add user: %w", err)
	}

	id64, err := res.LastInsertId()

	if err != nil {
		return 0, xerrors.Errorf("faield to retrieve last inserted id: %w", err)
	}

	return int(id64), nil
}

func (s *UserStore) GetUser(id int) (*User, error) {
	var u User

	if err := s.db.QueryRowxContext(context.Background(), "SElECT * FROM users WHERE seq = ?", id).StructScan(&u); err != nil {
		return nil, xerrors.Errorf("faield to retrieve user info: %w", err)
	}

	return &u, nil
}
