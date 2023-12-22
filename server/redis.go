package main

import (
	"context"

	redis "github.com/go-redis/redis/v8"
)

var redisCtx = context.Background()

func redisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
