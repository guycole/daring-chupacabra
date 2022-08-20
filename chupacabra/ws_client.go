// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

const maxPubSubChannels = 10

type pubSubChannelEnum int

const (
	pubSubChannel0 pubSubChannelEnum = iota
	pubSubChannel1
	pubSubChannel2
	pubSubChannel3
	pubSubChannel4
	pubSubChannel5
	pubSubChannel6
	pubSubChannel7
	pubSubChannel8
	pubSubChannel9
	pubSubChannelUnknown
)

// must match order for pubSubChannelEnum
var pubSubChannelNames = [...]string{
	"pubSubChannel0",
	"pubSubChannel1",
	"pubSubChannel2",
	"pubSubChannel3",
	"pubSubChannel4",
	"pubSubChannel5",
	"pubSubChannel6",
	"pubSubChannel7",
	"pubSubChannel8",
	"pubSubChannel9",
	"pubSubChannelUnknown",
}

func (channelName pubSubChannelEnum) String() string {
	return pubSubChannelNames[channelName]
}

type wsClientType struct {
	pubSubChannelName string
	connection        *websocket.Conn
	send              chan []byte
}

type wsClientArrayType [maxPubSubChannels]*wsClientType

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var wsClientArray wsClientArrayType

func newWsClient(ww http.ResponseWriter, rr *http.Request) (*wsClientType, error) {
	log.Println("---x---x---x---")

	var freshChannel int = maxPubSubChannels
	for ndx := 0; ndx < maxPubSubChannels; ndx++ {
		if wsClientArray[ndx] == nil {
			log.Printf("assignment %d\n", ndx)
			freshChannel = ndx
			break
		}
	}

	if freshChannel >= maxPubSubChannels {
		return nil, errors.New("pubsub channels full")
	}

	connection, err := upgrader.Upgrade(ww, rr, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	freshChannelName := pubSubChannelNames[freshChannel]
	log.Println(freshChannelName)

	result := &wsClientType{connection: connection, pubSubChannelName: freshChannelName, send: make(chan []byte, 256)}
	wsClientArray[freshChannel] = result

	backEndChannelName := os.Getenv("BE_CHANNEL")
	pt := newRegisterPayload(freshChannelName)
	pt.publishPayload(backEndChannelName, newRedisClient())

	return result, nil
}

func serveWs(ww http.ResponseWriter, rr *http.Request) {
	log.Println("servWs entry")

	client, err := newWsClient(ww, rr)
	if err != nil {
		log.Panic(err)
	}

	log.Println(client)
}
