package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nkolosov/whip-round/internal/db"
	"sync"
)

var (
	ErrParseConfig = errors.New("failed to parse config: %v")
	ErrCreatePool  = errors.New("failed to create connection pool: %v")
)

func NewPostgresConnection(ctx context.Context, cfg *db.Config) (*pgxpool.Pool, error) {
	if cfg == nil {
		return nil, db.ErrConfigIsNil
	}

	poolConfig, err := pgxpool.ParseConfig(cfg.String())
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParseConfig, err)
	}

	poolConfig.MaxConns = int32(cfg.MaxOpenConns)
	poolConfig.MinConns = int32(cfg.MaxIdleConns)

	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreatePool, err)
	}

	err = checkConnections(ctx, pool, int(poolConfig.MaxConns))
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func checkConnections(ctx context.Context, pool *pgxpool.Pool, numConns int) error {
	var wg sync.WaitGroup
	errChan := make(chan error, numConns)

	for i := 0; i < numConns; i++ {
		wg.Add(1)
		go func(errChan chan<- error, count int) {
			defer wg.Done()

			_, err := pool.Exec(ctx, "SELECT 1")
			if err != nil {
				errChan <- err
			}
		}(errChan, i)
	}

	wg.Wait()

	select {
	case err := <-errChan:
		return err
	default:
		close(errChan)
		return nil
	}
}
