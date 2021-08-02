package redis

import "github.com/go-redis/redis"

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "@test#test", // no password set
		DB:       0,            // use default DB
	})
}
