package users

import (
	"context"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/modoki-paas/modoki-k8s/pkg/types"
	"golang.org/x/xerrors"
)

type UserStore struct {
	db sqlx.ExtContext
}

func NewUserStore(db sqlx.ExtContext) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) AddUser(id, name string, userType types.UserTypeEnum, systemRole string) (seqID int, err error) {
	u := &types.User{
		ID:         id,
		UserType:   userType,
		Name:       name,
		SystemRole: systemRole,
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
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return 0, ErrUserIDDuplicates
		}

		return 0, xerrors.Errorf("faield to add user: %w", err)
	}

	id64, err := res.LastInsertId()

	if err != nil {
		return 0, xerrors.Errorf("faield to retrieve last inserted id: %w", err)
	}

	return int(id64), nil
}

func (s *UserStore) GetUser(id int) (*types.User, error) {
	var u types.User

	if err := s.db.QueryRowxContext(context.Background(), "SElECT * FROM users WHERE seq = ?", id).StructScan(&u); err != nil {
		return nil, xerrors.Errorf("faield to retrieve user info: %w", err)
	}

	return &u, nil
}
