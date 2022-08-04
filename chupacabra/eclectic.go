// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
)

type eclecticType struct {
	turnCounter int
	queueArray  queueArrayType
}

func stuntBox(state *eclecticType) {
	q1, err := newQueue("q1")
	q3, err := newQueue("q3")
	q5, err := newQueue("q5")

	state.queueArray[1] = q1
	state.queueArray[3] = q3
	state.queueArray[5] = q5

	msg1, err := newMessage(1, "payload1")
	msg2, err := newMessage(2, "payload2")
	msg3, err := newMessage(3, "payload3")
	msg4, err := newMessage(4, "payload4")
	msg5, err := newMessage(5, "payload5")
	msg6, err := newMessage(6, "payload6")
	msg7, err := newMessage(7, "payload7")
	msg8, err := newMessage(8, "payload8")
	msg9, err := newMessage(9, "payload9")

	messageAdd(msg1, q1)
	messageAdd(msg2, q3)
	messageAdd(msg3, q5)
	messageAdd(msg4, q1)
	messageAdd(msg5, q3)
	messageAdd(msg6, q5)
	messageAdd(msg7, q1)
	messageAdd(msg8, q3)
	messageAdd(msg9, q5)

	log.Println(err)
	log.Println(state)

	queueDump(q1)
	queueDump(q3)
	queueDump(q5)
}

func preTurnCycle(state *eclecticType) {
	log.Println("pre turn cycle")
}

func postTurnCycle(state *eclecticType) {
	log.Println("post turn cycle")
}

func turnCycleQueue(turn int, queue *queueType) {
	if queue == nil {
		// inactive queue
		return
	}

	for (queue.messageList != nil) && (queue.messageList.turn <= turn) {
		// consume all messages up to/including turn
		current := messageRead(queue)
		log.Println(current)
	}
}

func turnCycle(state *eclecticType) {
	state.turnCounter++
	log.Println("turn counter:", state.turnCounter)

	preTurnCycle(state)

	for ndx := 0; ndx < maxQueue; ndx++ {
		turnCycleQueue(state.turnCounter, state.queueArray[ndx])
	}

	postTurnCycle(state)
}

// runs forever, loops through queues
func eclectic() {
	state := eclecticType{}

	log.Println("eclectic entry")

	stuntBox(&state)

	for ndx := 0; ndx < 10; ndx++ {
		turnCycle(&state)
	}

	log.Println("eclectic exit")
}
