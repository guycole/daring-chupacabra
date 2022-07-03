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
	next           *messageType //single linked list sorted by turn
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

	current := queue.next
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

	if queue.next == nil {
		// new message root for empty list
		queue.next = message
	} else if message.turn < queue.next.turn {
		// new message root for earlier turn
		message.next = queue.next
		queue.next = message
	} else {
		// insert message into sorted list
		current := queue.next
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
