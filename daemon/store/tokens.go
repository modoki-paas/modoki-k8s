package store

import "time"

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
