package cache

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/akimsavvin/savvin_go_lib_common/logger/sl"
	"github.com/redis/go-redis/v9"
)

type Manager struct {
	log sl.Log
	c   *redis.Client
}

var ErrKeyNotFound = errors.New("value is not cached")

func NewManager(log sl.Log, redisClient *redis.Client) *Manager {
	log.Debug("creating cache manager", sl.Pkg("cache"), sl.Op("NewManager"))

	return &Manager{
		log: log.With(sl.Pkg("cache"), sl.Mdl("Manager")),
		c:   redisClient,
	}
}

func (m *Manager) Get(ctx context.Context, key string) (string, error) {
	const op = "Get"
	log := m.log.With(sl.Op(op))

	cmd := m.c.Get(ctx, key)

	err := cmd.Err()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			log.Error(
				"value is not cached",
				sl.Err(err),
				slog.String("key", key),
			)

			return "", ErrKeyNotFound
		}

		log.Error(
			"could not get cached value",
			sl.Err(err),
			slog.String("key", key),
		)

		return "", err
	}

	val := cmd.Val()

	log.Info(
		"got cached value",
		slog.String("key", key),
		slog.String("value", val),
	)

	return val, nil
}

func (m *Manager) Set(ctx context.Context, key string, val any, exp time.Duration) error {
	const op = "Set"
	log := m.log.With(sl.Op(op))

	cmd := m.c.Set(ctx, key, val, exp)

	err := cmd.Err()
	if err != nil {
		log.Error(
			"could not cache value",
			sl.Err(err),
			slog.String("key", key),
			slog.Any("value", val),
			slog.Duration("expiration", exp),
		)

		return err
	}

	log.Info(
		"cached value",
		slog.String("key", key),
		slog.Any("value", val),
		slog.Duration("expiration", exp),
	)

	return nil
}
