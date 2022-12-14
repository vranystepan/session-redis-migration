package migration

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	config "session-redis-migration/pkg/config/migration"
	"session-redis-migration/pkg/redis"
	"strings"
	"time"
)

func Run(ctx context.Context, redisClient redis.Client, cfg *config.Config) error {
	// get files
	files, err := filepath.Glob(cfg.FilePattern)
	if err != nil {
		return fmt.Errorf("could not get list of files: %s", err)
	}

	// go through files and put them to redis
	for _, f := range files {
		log.Printf("processing %s", f)
		err := processFile(ctx, redisClient, f, cfg.TargetTTL)
		if err != nil {
			return fmt.Errorf("could not process file %s: %s", f, err)
		}
	}

	return nil
}

func processFile(ctx context.Context, redisClient redis.Client, path string, ttl time.Duration) error {
	contents, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read file %s: %s", path, err)
	}

	// prepare key
	basename := filepath.Base(path)
	keyName := strings.Replace(basename, "sess_", "", -1)
	key := "PHPREDIS_SESSION:" + keyName

	// put session to redis
	err = redisClient.Set(ctx, key, string(contents), ttl).Err()
	if err != nil {
		return fmt.Errorf("could not put key %s to redis: %s", keyName, err)
	}

	return nil
}
