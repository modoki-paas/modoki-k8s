package store

import (
	"context"
	"time"

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

type tokensStore struct {
	db *dbContext
}

func newTokensStore(db *dbContext) *tokensStore {
	return &tokensStore{db: db}
}

func (s *tokensStore) AddToken(t *Token) (ret *Token, err error) {
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
				ret = nil
			}
		}
	}()

	res, err := dbx.db.ExecContext(
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
		return nil, xerrors.Errorf("failed to register new token: %w", err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, xerrors.Errorf("failed to register new token: %w", err)
	}

	return store.Token().GetToken(int(id))
}

func (s *tokensStore) GetToken(id int) (*Token, error) {
	var ts Token
	err := s.db.db.
		QueryRowxContext(context.Background(), "SELECT * FROM tokens WHERE seq=?", id).
		StructScan(&ts)

	if err != nil {
		return nil, xerrors.Errorf("failed to scan: %w", err)
	}

	return &ts, nil
}
func (s *tokensStore) GetFromToken(token string) (*Token, error) {
	var ts Token
	err := s.db.db.
		QueryRowxContext(context.Background(), "SELECT * FROM tokens WHERE token=?", token).
		StructScan(&ts)

	if err != nil {
		return nil, xerrors.Errorf("failed to scan: %w", err)
	}

	return &ts, nil
}
