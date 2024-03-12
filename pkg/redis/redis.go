package redis

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"practice/pkg/logger"
)

func Connect() *redis.Client {

	address := fmt.Sprintf("%s:%d", "127.0.0.1", 3307)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "127.0.0.1", 3307),
		Username: "",
		Password: "",
		DB:       0,
	})

	logger.Info(`🍟: Successfully connected to Redis at ` + address)

	return rdb
}
