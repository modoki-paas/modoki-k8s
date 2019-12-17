package types

type UserGroupRelation struct {
	Seq       int `db:"seq"`
	OrgSeqID  int `db:"org_seq"`
	UserSeqID int `db:"user_seq"`
}
