package store

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"time"

	"golang.org/x/xerrors"
)

type TokenPermission struct {
}

func (j *TokenPermission) Scan(src interface{}) error {
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

func (j TokenPermission) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type Token struct {
	ID              int              `db:"id"`
	Token           string           `db:"token"`
	Organization    int              `db:"organization"`
	Author          int              `db:"author"`
	TokenPermission *TokenPermission `db:"permission"`
	CreatedAt       time.Time        `db:"created_at"`
	UpdatedAt       time.Time        `db:"updated_at"`
}

type tokensStore struct {
	db *dbContext
}

func newTokensStore(db *dbContext) *tokensStore {
	return &tokensStore{db: db}
}

func (s *tokensStore) NewToken(t *Token) (*Token, error) {
	res, err := s.db.db.ExecContext(
		context.Background(),
		"INSERT INTO tokens (token, organization, author, permission) VALUES (?, ?, ?, ?)",
		t.Token,
		t.Organization,
		t.Author,
		t.TokenPermission,
	)

	if err != nil {
		return nil, xerrors.Errorf("failed to register new token: %v", err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, xerrors.Errorf("failed to register new token: %v", err)
	}

	t.ID = int(id)

	return t, nil
}

func (s *tokensStore) GetFromToken(token string) (*Token, error) {
	var ts Token
	err := s.db.db.
		QueryRowxContext(context.Background(), "SELECT * FROM tokens WHERE token=?", token).
		StructScan(&ts)

	if err != nil {
		return nil, xerrors.Errorf("failed to scan: %v", err)
	}

	return &ts, nil
}
