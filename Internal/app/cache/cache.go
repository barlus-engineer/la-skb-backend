package cache

import (
	"context"
	"fmt"
	"la-skb/config"
	"la-skb/pkg/logger"

	"github.com/redis/go-redis/v9"
)

var (
	cfg = config.LoadConfig()
	RDB *redis.Client
)

func InitializeCache() {
	options, err := redis.ParseURL(cfg.CacheURI)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse Redis URI: %v", err))
	}

	RDB = redis.NewClient(options)

	err = RDB.Ping(context.Background()).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	logger.Info("Connected to Redis cache successfully")
}

func GetCache() *redis.Client {
	if RDB == nil {
		panic("Redis client not initialized. Call InitializeCache() first.")
	}
	return RDB
}
