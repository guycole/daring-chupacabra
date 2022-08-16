// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"time"

	redis "github.com/go-redis/redis/v8"
)

func frontEnd() {
	log.Println("frontEnd entry")

	redisAddress := os.Getenv("REDIS_ADDRESS")
	log.Println(redisAddress)

	redisPassword := os.Getenv("REDIS_PASSWORD")
	log.Println(redisPassword)

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0, // use default DB
	})

	backEndChannelName := os.Getenv("BE_CHANNEL")
	frontEndChannelName := os.Getenv("FE_CHANNEL")

	pt, err := newPayload("id", "payType", frontEndChannelName)
	if err != nil {
		log.Panic(err)
	}

	publishPayload(pt, backEndChannelName, rdb)

	for {
		log.Println("Infinite Loop 2")
		time.Sleep(time.Second * 10)
	}
}
