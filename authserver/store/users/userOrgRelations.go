package users

import (
	"context"

	"github.com/jmoiron/sqlx"
	"golang.org/x/xerrors"
)

type UserGroupRelation struct {
	Seq       int `db:"seq"`
	OrgSeqID  int `db:"org_seq"`
	UserSeqID int `db:"user_seq"`
}

type UserOrgRelationsStore struct {
	db sqlx.ExtContext
}

func NewUserOrgRelationsStore(db sqlx.ExtContext) *UserOrgRelationsStore {
	return &UserOrgRelationsStore{db: db}
}

func (s *UserOrgRelationsStore) RegisterUserToOrg(orgSeq, userSeq int) (int, error) {
	res, err := s.db.ExecContext(
		context.Background(),
		`INSERT INTO user_org_relations (
			org_seq,
			user_seq
		) VALUES (?, ?)`,
		orgSeq, userSeq,
	)

	id64, err := res.LastInsertId()

	if err != nil {
		return 0, xerrors.Errorf("faield to retrieve last inserted id for user_org_relations: %w", err)
	}

	return int(id64), nil

}
