package store

type UserGroupRelation struct {
	Seq            int         `db:"seq"`
	GroupSeqID     int         `db:"group_seq"`
	UserSeqID      int         `db:"user_seq"`
	UserPermission *Permission `db:"permission"`
}

type userGroupRelationsStore struct {
	db *dbContext
}

func newUserGroupRelationsStore(db *dbContext) *userGroupRelationsStore {
	return &userGroupRelationsStore{db: db}
}
