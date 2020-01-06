package tokens

import (
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/modoki-paas/modoki-k8s/internal/tokenutil"
	"github.com/modoki-paas/modoki-k8s/pkg/types"
	"golang.org/x/xerrors"
)

type TokenStore struct {
	db sqlx.ExtContext
}

func NewTokenStore(db sqlx.ExtContext) *TokenStore {
	return &TokenStore{db: db}
}

func (ss *TokenStore) AddToken(s *types.Token) (seq int, token string, err error) {
	token, err = tokenutil.GenerateRandomToken()

	if err != nil {
		return 0, "", xerrors.Errorf("failed to generate random token: %w", err)
	}
	s.Token = token

	res, err := ss.db.ExecContext(
		context.Background(),
		`INSERT INTO tokens
			(id, token, owner, author)
			VALUES (?, ?, ?, ?)`,
		s.ID, s.Token, s.Owner, s.Author,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return 0, "", ErrTokenIDDuplicates
		}

		return 0, "", xerrors.Errorf("failed to add Token to db: %w", err)
	}

	id64, err := res.LastInsertId()

	if err != nil {
		return 0, "", xerrors.Errorf("failed to add Token to db: %w", err)
	}

	return int(id64), token, nil
}

func (ss *TokenStore) GetToken(seq int) (*types.Token, error) {
	var token types.Token
	err := ss.db.
		QueryRowxContext(context.Background(), "SELECT * FROM tokens WHERE seq=? FOR UPDATE", seq).
		StructScan(&token)

	if err != nil {
		return nil, xerrors.Errorf("failed to retrieve Token: %w", err)
	}

	return &token, nil
}

func (ss *TokenStore) FindTokenByID(owner string, id string) (*types.Token, error) {
	var token types.Token
	err := ss.db.
		QueryRowxContext(context.Background(), "SELECT * FROM tokens WHERE owner=? AND id=? FOR UPDATE", owner, id).
		StructScan(&token)

	if err != nil {
		return nil, xerrors.Errorf("failed to retrieve Token: %w", err)
	}

	return &token, nil
}

func (ss *TokenStore) FindTokenByToken(token string) (*types.Token, error) {
	var res types.Token
	err := ss.db.
		QueryRowxContext(context.Background(), "SELECT * FROM tokens WHERE token=? FOR UPDATE", token).
		StructScan(&res)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUnknownToken
		}
		return nil, xerrors.Errorf("failed to retrieve Token: %w", err)
	}

	return &res, nil
}
