package apps

import (
	"context"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/modoki-paas/modoki-k8s/pkg/types"
	"github.com/rs/xid"
	"golang.org/x/xerrors"
)

type AppStore struct {
	db sqlx.ExtContext
}

func NewAppStore(db sqlx.ExtContext) *AppStore {
	return &AppStore{db: db}
}

func (ss *AppStore) AddApp(s *types.App) (seqID int, err error) {
	s.ID = xid.New().String()

	res, err := ss.db.ExecContext(
		context.Background(),
		`INSERT INTO apps
			(id, owner, name, spec)
			VALUES (?, ?, ?, ?)`,
		s.ID, s.Owner, s.Name, s.Spec,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return 0, ErrAppNameDuplicates
		}

		return 0, xerrors.Errorf("failed to add app to db: %w", err)
	}

	id64, err := res.LastInsertId()

	if err != nil {
		return 0, xerrors.Errorf("failed to add app to db: %w", err)
	}

	return int(id64), nil
}

func (ss *AppStore) GetApp(seq int) (*types.App, error) {
	var app types.App
	err := ss.db.
		QueryRowxContext(context.Background(), "SELECT * FROM apps WHERE seq=?", seq).
		StructScan(&app)

	if err != nil {
		return nil, xerrors.Errorf("failed to retrieve app: %w", err)
	}

	return &app, nil
}

func (ss *AppStore) FindAppByID(id string) (*types.App, error) {
	var app types.App
	err := ss.db.
		QueryRowxContext(context.Background(), "SELECT * FROM apps WHERE id=?", id).
		StructScan(&app)

	if err != nil {
		return nil, xerrors.Errorf("failed to retrieve app: %w", err)
	}

	return &app, nil
}
