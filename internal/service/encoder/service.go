package encoder

import (
	"github.com/biryanim/leadgenmarket_tz/internal/client/cache"
	"github.com/biryanim/leadgenmarket_tz/internal/service"
)

type serv struct {
	redisClient cache.RedisClient
}

func NewService(redisClient cache.RedisClient) service.EncodeService {
	return &serv{
		redisClient: redisClient,
	}
}
