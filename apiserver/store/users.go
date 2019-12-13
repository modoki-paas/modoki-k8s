package store

import (
	"context"
	"database/sql/driver"
	"time"

	sqlxselect "github.com/cs3238-tsuzu/sqlx-selector/v2"

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
	SeqID     int          `db:"seq"`
	ID        string       `db:"id"`
	UserType  UserTypeEnum `db:"type"`
	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}

type userStore struct {
	db *dbContext
}

func newUserStore(db *dbContext) *userStore {
	return &userStore{db: db}
}

func (s *userStore) AddUser(id, name, string, userType UserTypeEnum) (u *User, err error) {
	u = &User{
		UserType: userType,
		Name:     name,
	}

	dbx, err := s.db.Begin(context.Background(), nil)
	store := newDB(dbx)

	if err != nil {
		return nil, xerrors.Errorf("faield to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			dbx.Rollback()
		} else {
			if err := dbx.Commit(); err != nil {
				err = xerrors.Errorf("failed to commit transaction: %w", err)
				u = nil
			}
		}
	}()

	res, err := dbx.db.ExecContext(
		context.Background(),
		`INSERT INTO users (
			type,
			id,
			name,
		) VALUES (?, ?, ?)`, u.UserType, u.ID, u.Name)

	if err != nil {
		return nil, xerrors.Errorf("faield to add user: %w", err)
	}

	id64, err := res.LastInsertId()

	if err != nil {
		return nil, xerrors.Errorf("faield to retrieve last inserted id: %w", err)
	}

	return store.User().GetUser(int(id64))
}

func (s *userStore) GetUser(id int) (*User, error) {
	var u User

	if err := s.db.db.QueryRowxContext(context.Background(), "SElECT * FROM users WHERE id = ?", id).StructScan(&u); err != nil {
		return nil, xerrors.Errorf("faield to retrieve user info: %w", err)
	}

	return &u, nil
}

func (s *userStore) GetUserFromToken(token string) (*User, *Token, error) {
	// userToken represents target user or organization(NOT AUTHOR OF TOKEN) and token
	type userToken struct {
		User  *User  `db:"users"`
		Token *Token `db:"tokens"`
	}

	var ut userToken
	err := s.db.db.QueryRowxContext(
		context.Background(),
		"SELECT "+
			sqlxselect.New(&ut).
				SelectStruct("users.*").
				SelectStruct("tokens.*").
				String()+
			" FROM users INNER JOIN tokens ON tokens.organization = users.id WHERE tokens.token = ?",
		token,
	).StructScan(&ut)

	if err != nil {
		return nil, nil, xerrors.Errorf("failed to get token and user from db: %w", err)
	}

	return ut.User, ut.Token, nil
}
