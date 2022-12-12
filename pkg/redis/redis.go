package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Keys(ctx context.Context, pattern string) *redis.StringSliceCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
}

type Config struct {
	RedisHost     string
	RedisPort     int
	RedisDB       int
	RedisPassword string
}

type RedisConfigurator interface {
	GetRedisConfig() Config
}

func New(cfg RedisConfigurator) *redis.Client {
	c := cfg.GetRedisConfig()
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.RedisHost, c.RedisPort),
		Password: c.RedisPassword,
		DB:       c.RedisDB,
	})
}
