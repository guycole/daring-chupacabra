// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
	"strings"
	"time"
)

type subscriberType struct {
	active      bool
	channelName string
}

// maximum subscriber connections
const maxSubscribers = 5

//
type subscriberArrayType [maxSubscribers]subscriberType

type eventManagerType struct {
	shutDownFlag bool // true, graceful exit
	subscribers  subscriberArrayType
	turnCounter  int // current game turn
}

func eventManager(emt *eventManagerType) {
	log.Println("eventManager entry")

	//var emt eventManagerType
	for ndx := 0; ndx < maxSubscribers; ndx++ {
		emt.subscribers[ndx].active = false
	}

	emt.shutDownFlag = false
	emt.turnCounter = 0

	rdb := newRedisClient()

	for !emt.shutDownFlag {
		emt.turnCounter++

		log.Printf("event manager loop w/subscribers\n")
		for ndx := 0; ndx < maxSubscribers; ndx++ {
			if emt.subscribers[ndx].active {
				log.Printf("subscriber %d active\n", ndx)
				result, _ := newPayload("woot", stubPayload, emt.subscribers[ndx].channelName)
				result.publishPayload(emt.subscribers[ndx].channelName, rdb)
			} else {
				log.Printf("subscriber %d inactive\n", ndx)
			}
		}

		time.Sleep(time.Second * 10)
	}
}

func (emt *eventManagerType) subscriberAdd(channelName string) {
	// test for duplicate channel name
	for ndx := 0; ndx < maxSubscribers; ndx++ {
		if emt.subscribers[ndx].active == true {
			if strings.Compare(emt.subscribers[ndx].channelName, channelName) == 0 {
				log.Printf("duplicate channel noted")
				//duplicate noted
			}
		}
	}

	for ndx := 0; ndx < maxSubscribers; ndx++ {
		if emt.subscribers[ndx].active == false {
			log.Printf("adding new subscriber %s\n", channelName)
			emt.subscribers[ndx].channelName = channelName
			emt.subscribers[ndx].active = true
			return
		}
	}

	log.Println("unable to add new subscriber")
}

func (emt *eventManagerType) subscriberDelete(channelName string) {
	for ndx := 0; ndx < maxSubscribers; ndx++ {
		if emt.subscribers[ndx].active == true {
			if strings.Compare(emt.subscribers[ndx].channelName, channelName) == 0 {
				emt.subscribers[ndx].active = false
			}
		}
	}
}
