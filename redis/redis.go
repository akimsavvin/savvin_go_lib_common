package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewClient(ctx context.Context, clientURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(clientURL)

	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	const (
		retries  = 5
		waitTime = 5 * time.Second
	)

	ctx, cancel := context.WithTimeout(ctx, waitTime)
	defer cancel()

	pingSucceed := true

	for i := 0; i < retries; i++ {
		err = pingClient(ctx, client)

		if err != nil {
			pingSucceed = false
		} else {
			break
		}
	}

	if !pingSucceed {
		return nil, err
	}

	return client, nil
}

func pingClient(ctx context.Context, client *redis.Client) error {
	return client.Ping(ctx).Err()
}
