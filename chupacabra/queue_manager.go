// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"errors"
	"log"
)

type queueType struct {
	name           string
	messageCounter int
	messageList    *messageType //single linked list sorted by turn
}

const maxQueue = 16

type queueArrayType [maxQueue]*queueType

func newQueue(name string) (*queueType, error) {
	if len(name) < 1 {
		return nil, errors.New("bad name")
	}

	result := queueType{name: name}
	return &result, nil
}

func queueAdd(name string, queueArray queueArrayType) (int, error) {
	// duplicate queue name test
	for ndx := 0; ndx < maxQueue; ndx++ {
		if (queueArray[ndx] != nil) && (queueArray[ndx].name == name) {
			return -1, errors.New("duplicate name")
		}
	}

	// queue creation
	candidate, err := newQueue(name)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	// queue insertion
	for ndx := 0; ndx < maxQueue; ndx++ {
		if queueArray[ndx] == nil {
			queueArray[ndx] = candidate
			return ndx, nil
		}
	}

	return -1, errors.New("queueArray full")
}

func queueDelete(name string, queueArray queueArrayType) int {
	for ndx := 0; ndx < maxQueue; ndx++ {
		if queueArray[ndx] == nil {
			queueArray[ndx] = nil
			return ndx
		}
	}

	return -1
}

func queueDump(queue *queueType) {
	log.Println("=-=-=-= Queue Dump =-=-=-=")
	log.Println(queue.name, queue.messageCounter)

	current := queue.messageList
	for {
		if current == nil {
			break
		}

		log.Println(current.turn, ":", current.payload)
		current = current.next
	}

	log.Println("=-=-=-= Queue Dump =-=-=-=")
}

func messageAdd(message *messageType, queue *queueType) {
	queue.messageCounter++

	if queue.messageList == nil {
		// new message root for empty list
		queue.messageList = message
	} else if message.turn < queue.messageList.turn {
		// new message root for earlier turn
		message.next = queue.messageList
		queue.messageList = message
	} else {
		// insert message into sorted list
		current := queue.messageList
		var last *messageType

		for (current != nil) && (current.turn < message.turn) {
			last = current
			current = current.next
		}

		if current == nil {
			// new tail
			last.next = message
		} else if last == nil {
			// new root
			current.next = message
		} else {
			// new middle
			last.next = message
			message.next = current
		}
	}
}

func messageRead(queue *queueType) *messageType {
	if queue.messageCounter < 1 {
		// empty message list
		return nil
	}

	temp := queue.messageList
	queue.messageList = queue.messageList.next

	queue.messageCounter--

	return temp
}
