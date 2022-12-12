package ttl

import (
	"context"
	"fmt"
	"log"
	config "session-redis-migration/pkg/config/ttl"
	"session-redis-migration/pkg/redis"
	"time"
)

func Run(ctx context.Context, redisClient redis.Client, cfg *config.Config) error {
	// get keys conforming to the provided pattern
	keys, err := GetKeys(ctx, redisClient, cfg.KeyPattern)
	if err != nil {
		return err
	}

	// set configured TTL
	for _, key := range keys {
		log.Printf("setting TTL (%s) for %s", cfg.TargetTTL.String(), key)
		err = SetTTL(ctx, redisClient, key, cfg.TargetTTL)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetKeys(ctx context.Context, redisClient redis.Client, pattern string) ([]string, error) {
	res, err := redisClient.Keys(ctx, pattern).Result()
	if err != nil {
		return []string{}, fmt.Errorf("could not get keys for %s: %s", pattern, err)
	}
	return res, nil
}

func SetTTL(ctx context.Context, redisClient redis.Client, key string, ttl time.Duration) error {
	err := redisClient.Expire(ctx, key, ttl).Err()
	if err != nil {
		return fmt.Errorf("could not set TTL: %s", err)
	}

	return nil
}
