package store

import (
	"context"
	"time"

	"golang.org/x/xerrors"
)

type TokenPermission struct {
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
