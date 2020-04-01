package apps

import (
	"context"
	"time"

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

func (ss *AppStore) AddApp(s *types.App) (seqID int, id string, err error) {
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
			return 0, "", ErrAppNameDuplicates
		}

		return 0, "", xerrors.Errorf("failed to add app to db: %w", err)
	}

	id64, err := res.LastInsertId()

	if err != nil {
		return 0, "", xerrors.Errorf("failed to add app to db: %w", err)
	}

	return int(id64), s.ID, nil
}

func (ss *AppStore) UpdateApp(seq int, s *types.AppSpec) error {
	_, err := ss.db.ExecContext(
		context.Background(),
		`UPDATE apps SET spec=? WHERE seq=?`,
		s, seq,
	)

	if err != nil {
		return xerrors.Errorf("failed to update app in db: %w", err)
	}

	return nil
}

func (ss *AppStore) GetUpdatedTime(seq int) (time.Time, error) {
	var t time.Time

	err := ss.db.QueryRowxContext(
		context.Background(),
		`SELECT updated_at FROM apps WHERE seq=?`,
		seq,
	).Scan(&t)

	if err != nil {
		return time.Time{}, xerrors.Errorf("failed to update app in db: %w", err)
	}

	return time.Time{}, nil
}

func (ss *AppStore) GetApp(seq int) (*types.App, error) {
	var app types.App
	err := ss.db.
		QueryRowxContext(context.Background(), "SELECT * FROM apps WHERE seq=? FOR UPDATE", seq).
		StructScan(&app)

	if err != nil {
		return nil, xerrors.Errorf("failed to retrieve app: %w", err)
	}

	return &app, nil
}

func (ss *AppStore) FindAppByID(id string) (*types.App, error) {
	var app types.App
	err := ss.db.
		QueryRowxContext(context.Background(), "SELECT * FROM apps WHERE id=? FOR UPDATE", id).
		StructScan(&app)

	if err != nil {
		return nil, xerrors.Errorf("failed to retrieve app: %w", err)
	}

	return &app, nil
}
