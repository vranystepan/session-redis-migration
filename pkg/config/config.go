package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	RedisHost     string
	RedisPort     int
	RedisDB       int
	RedisPassword string
	TargetTTL     time.Duration
	FilePattern   string
}

func New() (*Config, error) {
	redisHostEnv := os.Getenv("REDIS_HOST")
	redisPasswordEnv := os.Getenv("REDIS_PASSWORD")
	redisPortEnv := os.Getenv("REDIS_PORT")
	redisDBEnv := os.Getenv("REDIS_DB")
	targetTTLEnv := os.Getenv("TARGET_TTL")
	filePatternEnv := os.Getenv("FILE_PATTERN")

	// parse values and add some default values for the empty ones
	if redisHostEnv == "" {
		log.Println("defaulting host to localhost")
		redisHostEnv = "localhost"
	}

	redisPort, err := strconv.Atoi(redisPortEnv)
	if err != nil {
		return nil, fmt.Errorf("could not parse REDIS_PORT: %s", err)
	}
	if redisPort == 0 {
		log.Println("defaulting to port 6379")
		redisPort = 6379
	}

	redisDB, err := strconv.Atoi(redisDBEnv)
	if err != nil {
		return nil, fmt.Errorf("could not parse REDIS_DB: %s", err)
	}

	if targetTTLEnv == "" {
		log.Println("defaulting ttl to 14 days")
		targetTTLEnv = "336h"
	}
	targetTTL, err := time.ParseDuration(targetTTLEnv)
	if err != nil {
		return nil, fmt.Errorf("could not parse TARGET_TTL: %s", err)
	}

	return &Config{
		RedisHost:     redisHostEnv,
		RedisPort:     redisPort,
		RedisDB:       redisDB,
		RedisPassword: redisPasswordEnv,
		TargetTTL:     targetTTL,
		FilePattern:   filePatternEnv,
	}, nil
}
