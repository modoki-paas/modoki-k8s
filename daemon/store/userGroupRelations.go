package store

type UserGroupRelation struct {
	ID             int         `db:"id"`
	GroupID        int         `db:"group_id"`
	UserID         int         `db:"user_id"`
	UserPermission *Permission `db:"permission"`
}

type userGroupRelationsStore struct {
	db *dbContext
}

func newUserGroupRelationsStore(db *dbContext) *userGroupRelationsStore {
	return &userGroupRelationsStore{db: db}
}
