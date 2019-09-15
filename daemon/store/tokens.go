package store

import "time"

type Token struct {
	ID        int       `db:"id"`
	UserType  string    `db:"type"`
	Name      string    `db:"name"`
	Password  []byte    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type tokensStore struct {
	db *dbContext
}

func newTokensStore(db *dbContext) *tokensStore {
	return &tokensStore{db: db}
}
