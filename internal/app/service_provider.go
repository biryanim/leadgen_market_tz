package app

import (
	"context"
	"github.com/biryanim/leadgenmarket_tz/internal/api"
	"github.com/biryanim/leadgenmarket_tz/internal/client/cache"
	"github.com/biryanim/leadgenmarket_tz/internal/client/cache/redis"
	"github.com/biryanim/leadgenmarket_tz/internal/config"
	"github.com/biryanim/leadgenmarket_tz/internal/config/env"
	"github.com/biryanim/leadgenmarket_tz/internal/service"
	"github.com/biryanim/leadgenmarket_tz/internal/service/encoder"
	redigo "github.com/gomodule/redigo/redis"
	"log"
)

type serviceProvider struct {
	httpConfig  config.HTTPConfig
	redisConfig config.RedisConfig

	redisPool      *redigo.Pool
	redisClient    cache.RedisClient
	encoderService service.EncodeService

	apiImpl *api.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHttpConfig()
		if err != nil {
			log.Fatalf("failed to load http config: %v", err)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := env.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to load redis config: %v", err)
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) RedisPool() *redigo.Pool {
	if s.redisPool == nil {
		s.redisPool = &redigo.Pool{
			MaxIdle:     s.RedisConfig().MaxIdle(),
			IdleTimeout: s.RedisConfig().IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", s.RedisConfig().Address())
			},
		}
	}

	return s.redisPool
}

func (s *serviceProvider) RedisClient() cache.RedisClient {
	if s.redisClient == nil {
		s.redisClient = redis.NewClient(s.RedisPool(), s.RedisConfig())
	}

	return s.redisClient
}

func (s *serviceProvider) EncoderService(_ context.Context) service.EncodeService {
	if s.encoderService == nil {
		s.encoderService = encoder.NewService(s.RedisClient())
	}

	return s.encoderService
}

func (s *serviceProvider) ApiImpl(ctx context.Context) *api.Implementation {
	if s.apiImpl == nil {
		s.apiImpl = api.NewImplementation(s.EncoderService(ctx))
	}

	return s.apiImpl
}
