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

func newPayload(id string, payType string, reply string) (*PayloadType, error) {
	result := PayloadType{PayloadId: id, PayloadType: payType, ReplyChannel: reply}
	return &result, nil
}

func decodePayload(message *redis.Message) *PayloadType {
	var pt PayloadType

	err := json.Unmarshal([]byte(message.Payload), &pt)
	if err != nil {
		log.Println(err)
		log.Println("skipping bad unmarshal")
	}

	return &pt
}

func publishPayload(pt *PayloadType, channelName string, rdb *redis.Client) {
	payload, err := json.Marshal(pt)
	if err != nil {
		log.Println(err)
	}

	err = rdb.Publish(context.Background(), channelName, payload).Err()
	if err != nil {
		log.Fatal(err)
	}
}
