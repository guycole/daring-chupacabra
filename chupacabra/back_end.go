// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"log"

	redis "github.com/go-redis/redis/v8"
)

const maxPayloadArguments = 3

type argumentArrayType [maxPayloadArguments]string

type PayloadType struct {
	ArgumentSize int
	Arguments    argumentArrayType
	PayloadId    string
	PayloadType  string
	ReplyChannel string
}

func backEnd() {
	log.Println("backEnd entry")

	// TODO get these arguments from secrets
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-master.chupacabra.svc.cluster.local:6379",
		Password: "bigSekret",
		DB:       0, // use default DB
	})

	channelName := "back_end"

	topic := rdb.Subscribe(context.Background(), channelName)

	for {
		// blocking read
		message, err := topic.ReceiveMessage(context.Background())
		if err != nil {
			log.Println(err)
			log.Println("backEnd skipping bad receive message")
			continue
		}

		log.Println(message)

		var pt PayloadType
		err = json.Unmarshal([]byte(message.Payload), &pt)
		if err != nil {
			log.Println(err)
			log.Println("backEnd skipping bad unmarshal")
			continue
		}

		log.Println(pt)
	}
}
