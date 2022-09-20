// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"reflect"

	redis "github.com/go-redis/redis/v8"
)

func newRedisClient() *redis.Client {
	redisAddress := os.Getenv("REDIS_ADDRESS")
	log.Println(redisAddress)

	redisPassword := os.Getenv("REDIS_PASSWORD")
	log.Println(redisPassword)

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0, // use default DB
	})

	log.Println(reflect.TypeOf(rdb))

	return rdb
}
