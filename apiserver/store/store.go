package store

import (
	"context"
	"database/sql"

	"golang.org/x/xerrors"

	"github.com/jmoiron/sqlx"
)

type dbContext struct {
	db sqlx.ExtContext

	opt *sql.TxOptions
	txRoot bool
}

// Beginnable represents db connection that can begin transaction
type Beginnable interface {
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}

func (d *dbContext) Begin(ctx context.Context, opt *sql.TxOptions) (*dbContext, error) {
	switch v := d.db.(type) {
	case Beginnable:
		tx, err := v.BeginTxx(ctx, opt)

		if err != nil {
			return nil, xerrors.Errorf("beginning transaction error: ", err)
		}

		return &dbContext{
			db:  tx,
			opt: opt,
			txRoot: true,
		}, nil
	default:
		// already begun

		if opt != nil && d.opt != opt {
			return nil, xerrors.Errorf("tx option does not match: %v <=> %v", d.opt, opt)
		}
		return &dbContext {
			db: d.db,
			opt: opt,
			txRoot: false,
		}, nil
	}
}

func (d *dbContext) Commit() error {
	if !d.txRoot {
		return nil
	}

	v, ok := d.db.(*sqlx.Tx)

	if !ok {
		return xerrors.New("tx is not begun")
	}

	return v.Commit()
}

func (d *dbContext) Rollback() error {
	if !d.txRoot {
		return nil
	}

	v, ok := d.db.(*sqlx.Tx)

	if !ok {
		return xerrors.New("tx is not begun")
	}

	return v.Rollback()
}

type DB struct {
	db *dbContext
}

func newDB(db *dbContext) *DB {
	return &DB{
		db: db,
	}
}

func NewDB(db sqlx.ExtContext) *DB {
	return newDB(&dbContext{db: db})
}

func (d *DB) Begin(ctx context.Context, opt *sql.TxOptions) (*DB, error) {
	c, err := d.db.Begin(ctx, opt)

	if err != nil {
		return nil, err
	}

	return newDB(c), nil
}

func (d *DB) Commit() error {
	return d.db.Commit()
}

func (d *DB) Rollback() error {
	return d.db.Rollback()
}

func (d *DB) User() *userStore {
	return newUserStore(d.db)
}

func (d *DB) Service() *serviceStore {
	return newServiceStore(d.db)
}

func (d *DB) UserGroupRelation() *userGroupRelationsStore {
	return newUserGroupRelationsStore(d.db)
}

func (d *DB) Token() *tokensStore {
	return newTokensStore(d.db)
}
