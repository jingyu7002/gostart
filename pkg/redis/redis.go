package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis(addr, password string) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})
}
