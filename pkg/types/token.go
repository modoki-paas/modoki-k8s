package types

import (
	"time"
)

type Token struct {
	SeqID     int       `db:"seq"`
	Token     string    `db:"token"`
	Owner     int       `db:"owner"`
	Author    int       `db:"author"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}