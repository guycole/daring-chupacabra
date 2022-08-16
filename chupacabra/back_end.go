// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"context"
	"log"
	"os"

	redis "github.com/go-redis/redis/v8"
)

func handler(pt *PayloadType, replyChannel string, rdb *redis.Client) {
	response, err := newPayload(pt.PayloadId, "OK", replyChannel)
	if err != nil {
		log.Println("new payload failure")
	}

	log.Println(response)

	publishPayload(response, replyChannel, rdb)
}

func backEnd() {
	log.Println("backEnd entry")

	redisAddress := os.Getenv("REDIS_ADDRESS")
	log.Println(redisAddress)

	redisPassword := os.Getenv("REDIS_PASSWORD")
	log.Println(redisPassword)

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0, // use default DB
	})

	channelName := os.Getenv("BE_CHANNEL")
	log.Println(channelName)

	topic := rdb.Subscribe(context.Background(), channelName)

	for {
		// blocking read
		message, err := topic.ReceiveMessage(context.Background())
		if err != nil {
			log.Println(err)
			log.Println("backEnd skipping bad receive message")
			continue
		}

		log.Println("fresh message noted")
		log.Println(message)

		//time.Sleep(time.Second * 10)

		pt := decodePayload(message)
		log.Println(pt)

		handler(pt, pt.ReplyChannel, rdb)
	}
}
