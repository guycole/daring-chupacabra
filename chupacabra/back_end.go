// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"context"
	"time"

	"log"

	redis "github.com/go-redis/redis/v8"
)

type RequestType struct {
	RequestId    string
	RequestType  string
	ReplyChannel string
}

type ResponseType struct {
	RequestId   string
	RequestType string
	Response    string
}

//rdb *redis.Client

func backEnd() {
	log.Println("back end")

	// TODO get these arguments from secrets
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-master.chupacabra.svc.cluster.local:6379",
		Password: "bigSekret",
		DB:       0, // use default DB
	})

	channelName := "admin_in"

	log.Println("prepare to Subscribe")
	topic := rdb.Subscribe(context.Background(), channelName)
	log.Println("back from Subscribe")
	log.Println(topic)

	for {
		// blocking read
		message, err := topic.ReceiveMessage(context.Background())
		if err != nil {
			log.Println(err)
			log.Println("requestFromWorker skipping bad receive message")
			continue
		}

		log.Println(message)
	}

	/*
		var rt RequestType
		err = json.Unmarshal([]byte(msg.Payload), &rt)
		if err != nil {
				log.Println(err)
				log.Println("requestFromManager skipping bad unmarshal")
				continue
		}

		requestQueue.enqueue(&rt)
	*/
	//}

	/*
		rt := RequestType{Name: "name", RequestId: "reqid", ArgumentSize: 11}

		payload, err := json.Marshal(rt)
		if err != nil {
			log.Println(err)
		}

		log.Println(payload)

		err = rdb.Publish(context.Background(), channel, payload).Err()
		if err != nil {
			log.Fatal(err)
		}
	*/

	for true {
		log.Println("Infinite Loop 2")
		time.Sleep(time.Second * 10)
	}
}

//
