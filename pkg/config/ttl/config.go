package config

import (
	"flag"
	"fmt"
	"log"
	"session-redis-migration/pkg/redis"
	"time"
)

type Config struct {
	RedisHost     string
	RedisPort     int
	RedisDB       int
	RedisPassword string
	TargetTTL     time.Duration
	KeyPattern    string
}

func New() (*Config, error) {
	redisHostFlag := flag.String("host", "localhost", "Redis host")
	redisPortFlag := flag.Int("port", 6379, "Redis port")
	redisDBFlag := flag.Int("db", 0, "Redis DB index")
	redisPasswordFlag := flag.String("password", "", "Redis password")
	targetTTLFlag := flag.String("ttl", "336h", "TTL for Redis keys")
	keyPatternFlag := flag.String("keys", "", "KEYS patter")

	flag.Parse()

	// parse and convert values
	targetTTL, err := time.ParseDuration(*targetTTLFlag)
	if err != nil {
		return nil, fmt.Errorf("could not parse ttl: %s", err)
	}

	// perform basic sanity checks
	if *keyPatternFlag == "" {
		log.Fatalf("-keys flag must not be empty")
	}

	return &Config{
		RedisHost:     *redisHostFlag,
		RedisPort:     *redisPortFlag,
		RedisDB:       *redisDBFlag,
		RedisPassword: *redisPasswordFlag,
		TargetTTL:     targetTTL,
		KeyPattern:    *keyPatternFlag,
	}, nil
}

func (c *Config) GetRedisConfig() redis.Config {
	return redis.Config{
		RedisHost:     c.RedisHost,
		RedisPort:     c.RedisPort,
		RedisDB:       c.RedisDB,
		RedisPassword: c.RedisPassword,
	}
}
