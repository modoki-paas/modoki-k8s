package types

import (
	"time"
)

type User struct {
	SeqID      int          `db:"seq"`
	ID         string       `db:"id"`
	UserType   UserTypeEnum `db:"type"`
	Name       string       `db:"name"`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  time.Time    `db:"updated_at"`
	SystemRole string       `db:"system_role"`
}
