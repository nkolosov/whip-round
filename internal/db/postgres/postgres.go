package postgres

import (
	"database/sql"

	"github.com/nkolosov/whip-round/internal/db"
)

func NewPostgresConnection(cfg *db.Config) (*sql.DB, error) {
	if cfg == nil {
		return nil, db.ErrConfigIsNil
	}

	store, err := sql.Open("postgres", cfg.String())
	if err != nil {
		return nil, err
	}

	err = store.Ping()
	if err != nil {
		return nil, err
	}

	store.SetMaxIdleConns(cfg.MaxIdleConns)
	store.SetMaxOpenConns(cfg.MaxOpenConns)

	return store, nil
}
