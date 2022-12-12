package main

import (
	"context"
	"log"
	"os"
	"session-redis-migration/internal/ttl"
	config "session-redis-migration/pkg/config/ttl"
	"session-redis-migration/pkg/redis"
)

func main() {
	var ctx = context.Background()
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("could not load configuration: %s", err)
	}

	err = ttl.Run(ctx, redis.New(cfg), cfg)
	if err != nil {
		os.Exit(1)
	}
}
