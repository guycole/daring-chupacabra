// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	redis "github.com/go-redis/redis/v8"
)

func frontEnd() {
	log.Println("frontEnd entry")

	// TODO get these arguments from secrets
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-master.chupacabra.svc.cluster.local:6379",
		Password: "bigSekret",
		DB:       0, // use default DB
	})

	var arguments argumentArrayType
	arguments[0] = "0"
	arguments[1] = "1"
	arguments[2] = "2"

	pt := PayloadType{ArgumentSize: 1, Arguments: arguments, PayloadId: "id", PayloadType: "test", ReplyChannel: "front_end"}

	payload, err := json.Marshal(pt)
	if err != nil {
		log.Println(err)
	}

	channelName := "front_end"
	err = rdb.Publish(context.Background(), channelName, payload).Err()
	if err != nil {
		log.Fatal(err)
	}

	for true {
		log.Println("Infinite Loop 2")
		time.Sleep(time.Second * 10)
	}
}
