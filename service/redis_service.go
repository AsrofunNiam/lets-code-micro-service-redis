package service

import "github.com/redis/go-redis/v9"

type RedisService interface {
	ProcessCreateQueue(method, url, payload, key string)
	TakeRedisCacheConfig(redisClient *redis.Client, level uint) (string, error)
}
