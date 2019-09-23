package store

import (
	"context"
	"database/sql/driver"
	"time"

	"golang.org/x/crypto/bcrypt"

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
	ID        int          `db:"id"`
	UserType  UserTypeEnum `db:"type"`
	Name      string       `db:"name"`
	Password  []byte       `db:"password"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}

func (u *User) ComparePassword(passwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(passwd))

	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

type userStore struct {
	db *dbContext
}

func newUserStore(db *dbContext) *userStore {
	return &userStore{db: db}
}

func (s *userStore) AddUser(name, password string, userType UserTypeEnum) (u *User, err error) {
	passwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, xerrors.Errorf("failed to hash password: %v", err)
	}
	u = &User{
		UserType: userType,
		Name:     name,
		Password: passwd,
	}

	dbx, err := s.db.Begin(context.Background(), nil)

	if err != nil {
		return nil, xerrors.Errorf("faield to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			dbx.Rollback()
		} else {
			if err := dbx.Commit(); err != nil {
				err = xerrors.Errorf("failed to commit transaction: %v", err)
				u = nil
			}
		}
	}()

	res, err := s.db.db.ExecContext(context.Background(), "INSERT INTO users (type, name, password) VALUES (?, ?, ?)", u.UserType, u.Name, u.Password)

	if err != nil {
		return nil, xerrors.Errorf("faield to begin transaction: %v", err)
	}

	id64, err := res.LastInsertId()

	if err != nil {
		return nil, xerrors.Errorf("faield to retrieve last inserted id: %v", err)
	}

	if err := s.db.db.QueryRowxContext(context.Background(), "SElECT * FROM users WHERE id = ?", int(id64)).Scan(&u); err != nil {
		return nil, xerrors.Errorf("faield to retrieve user info: %v", err)
	}

	return u, nil
}

func (s *userStore) GetUserFromToken(token string) (*User, *Token, error) {
	// userToken represents target user or organization(NOT AUTHOR OF TOKEN) and token
	type userToken struct {
		Users  *User
		Tokens *Token
	}

	var ut userToken
	err := s.db.db.QueryRowxContext(
		context.Background(),
		"SELECT * FROM users INNER JOIN tokens ON tokens.organization = user.id WHERE tokens.token = ?",
		token,
	).StructScan(&ut)

	if err != nil {
		return nil, nil, xerrors.Errorf("failed to get token and user from db: %v", err)
	}

	return ut.Users, ut.Tokens, nil
}
