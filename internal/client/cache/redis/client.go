package redis

import (
	"context"
	"github.com/biryanim/leadgenmarket_tz/internal/client/cache"
	"github.com/biryanim/leadgenmarket_tz/internal/config"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

type handler func(cxt context.Context, conn redis.Conn) error

type client struct {
	pool   *redis.Pool
	config config.RedisConfig
}

func NewClient(pool *redis.Pool, config config.RedisConfig) cache.RedisClient {
	return &client{
		pool:   pool,
		config: config,
	}
}

func (c *client) Get(ctx context.Context, key string) (interface{}, error) {
	var value interface{}
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		var errEx error
		value, errEx = conn.Do("GET", key)
		if errEx != nil {
			return errEx
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return value, nil
}

func (c *client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("SET", redis.Args{key}.Add(value)...)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *client) execute(ctx context.Context, handler handler) error {
	conn, err := c.getConnect(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Printf("error closing connection: %v", err)
		}
	}()

	err = handler(ctx, conn)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) getConnect(ctx context.Context) (redis.Conn, error) {
	getConnTimeoutCtx, cancel := context.WithTimeout(ctx, c.config.ConnectionTimeout())
	defer cancel()

	conn, err := c.pool.GetContext(getConnTimeoutCtx)
	if err != nil {
		log.Printf("redis connection error: %v", err)

		_ = conn.Close()
		return nil, err
	}

	return conn, nil
}

func (c *client) Close() error {
	return c.pool.Close()
}
