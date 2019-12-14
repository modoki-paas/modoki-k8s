package tokens

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/xerrors"
)

type Token struct {
	SeqID     int       `db:"seq"`
	Token     string    `db:"token"`
	Owner     int       `db:"owner"`
	Author    int       `db:"author"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TokenStore struct {
	db sqlx.ExtContext
}

func NewTokenStore(db sqlx.ExtContext) *TokenStore {
	return &TokenStore{db: db}
}

func (s *TokenStore) AddToken(t *Token) (seqID int, err error) {
	res, err := s.db.ExecContext(
		context.Background(),
		`INSERT INTO tokens (
			token,
			owner,
			author
		) VALUES (?, ?, ?)`,
		t.Token,
		t.Owner,
		t.Author,
	)

	if err != nil {
		return 0, xerrors.Errorf("failed to register new token: %w", err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, xerrors.Errorf("failed to register new token: %w", err)
	}

	return int(id), nil
}

func (s *TokenStore) GetToken(id int) (*Token, error) {
	var ts Token
	err := s.db.
		QueryRowxContext(context.Background(), "SELECT * FROM tokens WHERE seq=?", id).
		StructScan(&ts)

	if err != nil {
		return nil, xerrors.Errorf("failed to scan: %w", err)
	}

	return &ts, nil
}
func (s *TokenStore) GetFromToken(token string) (*Token, error) {
	var ts Token
	err := s.db.
		QueryRowxContext(context.Background(), "SELECT * FROM tokens WHERE token=?", token).
		StructScan(&ts)

	if err != nil {
		return nil, xerrors.Errorf("failed to scan: %w", err)
	}

	return &ts, nil
}
