package types

import (
	"time"
)

type Token struct {
	SeqID     int       `db:"seq"`
	ID        string    `db:"id"`
	Token     string    `db:"token"`
	Owner     string    `db:"owner"`
	Author    string    `db:"author"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
