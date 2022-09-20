// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"context"
	"log"
	"os"

	redis "github.com/go-redis/redis/v8"
)

func handler(emt *eventManagerType, pt *PayloadType, rdb *redis.Client) {
	response := pt.newErrorPayload()

	switch pt.PayloadType {
	case unknownPayload:
		//should never get these
		log.Panic("unknown payload noted")
	case okPayload:
		//should never get these
		log.Panic("ok payload noted")
	case registerPayload:
		log.Println("register noted")
		emt.subscriberAdd(pt.ReplyChannel)
		response = pt.newSubscribePayload()
	case unregisterPayload:
		// always succeeds
		log.Println("unregister noted")
		emt.subscriberDelete(pt.ReplyChannel)
		response = pt.newOkPayload()
	default:
		log.Panic("unsupported payload type")
	}

	response.publishPayload(response.ReplyChannel, rdb)
}

func backEnd() {
	log.Println("backEnd entry")

	var emt eventManagerType

	emt.turnCounter = 0
	emt.shutDownFlag = false

	for ndx := 0; ndx < maxSubscribers; ndx++ {
		log.Printf("sub %d\n", ndx)
		emt.subscribers[ndx].active = false
	}

	rdb := newRedisClient()

	go eventManager(&emt)

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

		handler(&emt, pt, rdb)
	}
}
