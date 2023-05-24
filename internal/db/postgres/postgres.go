package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nkolosov/whip-round/internal/db"
)

var (
	ErrParseConfig = errors.New("failed to parse config: %v")
	ErrCreatePool  = errors.New("failed to create connection pool: %v")
)

func NewPostgresConnection(cfg *db.Config) (*pgxpool.Pool, error) {
	if cfg == nil {
		return nil, db.ErrConfigIsNil
	}

	poolConfig, err := pgxpool.ParseConfig(cfg.String())
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParseConfig, err)
	}

	poolConfig.MaxConns = int32(cfg.MaxOpenConns)
	poolConfig.MinConns = int32(cfg.MaxIdleConns)

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreatePool, err)
	}

	return pool, nil
}
