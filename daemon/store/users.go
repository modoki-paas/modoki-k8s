package store

import (
	"context"
	"database/sql/driver"
	"time"

	"golang.org/x/xerrors"
)

// UserTypeEnum represents the type of users
type UserTypeEnum string

func (e *UserTypeEnum) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		*e = UserTypeEnum(v)
	case string:
		*e = UserTypeEnum(v)
	default:
		return xerrors.Errorf("failed to scan json for %v", v)
	}
	return nil
}

func (e UserTypeEnum) Value() (driver.Value, error) {
	return string(e), nil
}

var (
	// UserNormal means a individual user, not organization
	UserNormal UserTypeEnum = "user"
	// UserOrganization means an organization contains some users
	UserOrganization UserTypeEnum = "organization"
)

type User struct {
	ID        int       `db:"id"`
	UserType  string    `db:"type"`
	Name      string    `db:"name"`
	Password  []byte    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type userStore struct {
	db *dbContext
}

func newUserStore(db *dbContext) *userStore {
	return &userStore{db: db}
}

func (s *userStore) GetUserFromToken(token string) (*User, *Token, error) {
	// userToken represents target user or organization(NOT AUTHOR OF TOKEN) and token
	type userToken struct {
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
		TokenCreatedAt  time.Time        `db:"token_created_at"`
		TokenUpdatedAt  time.Time        `db:"token_updated_at"`
	}

	var ut userToken
	err := s.db.db.QueryRowxContext(
		context.Background(),
		"SELECT users.id, users.type, users.name, users.password, users.created_at, users.updated_at, tokens.token, tokens.organization, tokens.author, tokens.permission tokens.created_at AS token_created_at, tokens.updated_at AS token_updated_at FROM users INNER JOIN tokens ON tokens.organization = user.id WHERE tokens.token = ?",
		token,
	).StructScan(&ut)

	if err != nil {
		return nil, nil, xerrors.Errorf("failed to get token and user from db: %v", err)
	}

	u := &User{
		ID:        ut.ID,
		UserType:  ut.UserType,
		Name:      ut.Name,
		Password:  ut.Password,
		CreatedAt: ut.CreatedAt,
		UpdatedAt: ut.UpdatedAt,
	}

	t := &Token{
		Token:        ut.Token,
		Organization: ut.Organization,
		Author:       ut.Author,
		CreatedAt:    ut.TokenCreatedAt,
		UpdatedAt:    ut.TokenUpdatedAt,
	}

	return u, t, nil
}
