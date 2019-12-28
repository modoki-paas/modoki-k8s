package types

type RoleBinding struct {
	Seq         int    `db:"seq"`
	UserSeqID   int    `db:"user_seq"`
	TargetSeqID int    `db:"org_seq"`
	Role        string `db:"role"`
}
