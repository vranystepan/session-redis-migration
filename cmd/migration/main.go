package main

import (
	"context"
	"log"
	"os"
	"session-redis-migration/internal/migration"
	"session-redis-migration/pkg/config"
)

func main() {
	var ctx = context.Background()
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("could not load configuration: %s", err)
	}

	err = migration.Run(ctx, cfg)
	if err != nil {
		os.Exit(1)
	}
}
