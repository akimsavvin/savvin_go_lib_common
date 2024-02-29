package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	const (
		retries  = 5
		waitTime = 5 * time.Second
	)

	ctx, cancel := context.WithTimeout(ctx, waitTime)
	defer cancel()

	pingSucceed := true

	for i := 0; i < retries; i++ {
		err = pool.Ping(ctx)

		if err != nil {
			pingSucceed = false
		} else {
			break
		}
	}

	if !pingSucceed {
		return nil, err
	}

	return pool, nil
}
