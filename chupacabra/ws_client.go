// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"sync"
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
	active            bool
	connection        *websocket.Conn
	mootex            sync.Mutex
	pubSubChannelName string
	send              chan []byte
}

type wsClientArrayType [maxPubSubChannels]*wsClientType

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var wsClientArray wsClientArrayType

var wsClientMap map[string]*wsClientType

func newClientMap() {
	wsClientMap = make(map[string]*wsClientType)

	for ndx := 0; ndx < maxPubSubChannels; ndx++ {
		temp := wsClientType{pubSubChannelName: pubSubChannelNames[ndx], active: false}
		wsClientMap[pubSubChannelNames[ndx]] = &temp
	}
}

func selectFreeClient() (string, error) {
	for key, value := range wsClientMap {
		if !value.active {
			return key, nil
		}
	}

	return "", errors.New("pubsub channels full")
}

// read message from back end and write to web socket
func (ws *wsClientType) webSocketProxy() {
	log.Printf("webSocketProx:%s\n", ws.pubSubChannelName)

	rdb := newRedisClient()

	backEndChannelName := os.Getenv("BE_CHANNEL")
	pt := newRegisterPayload(ws.pubSubChannelName)
	pt.publishPayload(backEndChannelName, rdb)

	topic := rdb.Subscribe(context.Background(), ws.pubSubChannelName)

	for {
		// blocking read
		message, err := topic.ReceiveMessage(context.Background())
		if err != nil {
			log.Println(err)
			log.Println("backEnd skipping bad receive message")
			continue
		}

		log.Printf("fresh message noted:%s\n", ws.pubSubChannelName)
		log.Println(message)

		//time.Sleep(time.Second * 10)

		pt := decodePayload(message)
		log.Println(pt)

		w, err := ws.connection.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Println("error from next writer")
			log.Println(err)
		} else {
			w.Write([]byte("message from ws_client"))
		}

		if err := w.Close(); err != nil {
			log.Println("error from close")
		}
	}
}

func serveWebSocket(ww http.ResponseWriter, rr *http.Request) {
	log.Println("servWs entry")

	key, err := selectFreeClient()
	if err != nil {
		log.Println("selectFromClient error noted")
		return
	}

	result := wsClientMap[key]
	result.mootex.Lock()
	result.active = true
	result.mootex.Unlock()

	connection, err := upgrader.Upgrade(ww, rr, nil)
	if err != nil {
		return
	} else {
		result.connection = connection
	}

	go result.webSocketProxy()
}
