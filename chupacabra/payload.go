// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"

	redis "github.com/go-redis/redis/v8"
)

type payloadEnum int

const (
	unknownPayload payloadEnum = iota
	errorPayload
	okPayload
	registerPayload
	stubPayload
	subscribePayload
	unregisterPayload
)

const maxPayloadArguments = 3

type argumentArrayType [maxPayloadArguments]string

type PayloadType struct {
	ArgumentSize int
	Arguments    argumentArrayType
	PayloadId    string
	PayloadType  payloadEnum
	ReplyChannel string
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

func newPayload(id string, payType payloadEnum, reply string) (*PayloadType, error) {
	result := PayloadType{PayloadId: id, PayloadType: payType, ReplyChannel: reply}
	return &result, nil
}

func (pt *PayloadType) newErrorPayload() *PayloadType {
	result, err := newPayload(pt.PayloadId, errorPayload, pt.ReplyChannel)
	if err != nil {
		log.Panic(err)
	}

	return result
}

func (pt *PayloadType) newOkPayload() *PayloadType {
	result, err := newPayload(pt.PayloadId, okPayload, pt.ReplyChannel)
	if err != nil {
		log.Panic(err)
	}

	return result
}

func (pt *PayloadType) newRegisterPayload(replyChannel string) *PayloadType {
	result, err := newPayload(uuid.NewString(), registerPayload, replyChannel)
	if err != nil {
		log.Panic(err)
	}

	return result
}

func (pt *PayloadType) newStubPayload() *PayloadType {
	result, err := newPayload(pt.PayloadId, stubPayload, pt.ReplyChannel)
	if err != nil {
		log.Panic(err)
	}

	return result
}

func (pt *PayloadType) newSubscribePayload() *PayloadType {
	result, err := newPayload(uuid.NewString(), subscribePayload, pt.ReplyChannel)
	if err != nil {
		log.Panic(err)
	}

	return result
}

func (pt *PayloadType) publishPayload(channelName string, rdb *redis.Client) {
	payload, err := json.Marshal(pt)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("publish to channel %s\n", channelName)
	err = rdb.Publish(context.Background(), channelName, payload).Err()
	if err != nil {
		log.Fatal(err)
	}
}
