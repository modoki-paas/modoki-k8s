package dbutil

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"golang.org/x/xerrors"
)

type TxInterface interface {
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}

func Transaction(ctx context.Context, db TxInterface, fn func(tx *sqlx.Tx) error) (err error) {
	tx, err := db.BeginTxx(ctx, nil)

	if err != nil {
		return xerrors.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if e := recover(); e != nil {
			if e, ok := e.(error); ok {
				err = xerrors.Errorf("recovered from panic: %w", e)
			} else {
				err = xerrors.Errorf("recovered from panic: %v", e)
			}
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()

		// Don't wrap for status.Error()
		return err
	}

	if err := tx.Commit(); err != nil {
		return xerrors.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
