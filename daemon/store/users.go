package store

import (
	"context"
	"time"

	"golang.org/x/xerrors"
)

type User struct {
	ID        int       `db:"id"`
	UserType  string    `db:"type"`
	Name      string    `db:"name"`
	Password  []byte    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// UserToken represents target user or organization(NOT AUTHOR OF TOKEN) and token
type UserToken struct {
	ID        int       `db:"id"`
	UserType  string    `db:"type"`
	Name      string    `db:"name"`
	Password  []byte    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Token           string           `db:"token"`
	Organization    int              `db:"organization"`
	Author          int              `db:"author"`
	TokenPermission *TokenPermission `db:"permission"`
}

type userStore struct {
	db *dbContext
}

func newUserStore(db *dbContext) *userStore {
	return &userStore{db: db}
}

func (s *userStore) GetUserFromToken(token string) (*UserToken, error) {
	var userToken UserToken
	err := s.db.db.QueryRowxContext(
		context.Background(),
		"SELECT users.id, users.type, users.name, users.password, users.created_at, users.updated_at, tokens.token, tokens.organization, tokens.author, tokens.permission FROM users INNER JOIN tokens ON tokens.organization = user.id WHERE tokens.token = ?",
		token,
	).StructScan(&userToken)

	if err != nil {
		return nil, xerrors.Errorf("failed to get token and user from db: %v", err)
	}

	return &userToken, nil
}
