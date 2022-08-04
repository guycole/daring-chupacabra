// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	name := "testaroo"

	queue, err := newQueue(name)
	if err != nil {
		t.Error("new queue failure")
	}

	if queue.name != name {
		t.Error("queue failure")
	}
}

func TestNewQueueBadName(t *testing.T) {
	queue, err := newQueue("")
	if err == nil {
		t.Error("new queue should fail")
	}

	if queue != nil {
		t.Error("message should be nil")
	}
}

func TestQueueMessageAddDelete(t *testing.T) {
	queue, err := newQueue("queue1")
	if err != nil {
		t.Error("new queue failure")
	}

	msg1, err := newMessage(1, "payload1")
	if err != nil {
		t.Error("new message1 failure")
	}

	msg2, err := newMessage(1, "payload2")
	if err != nil {
		t.Error("new message2 failure")
	}

	msg3, err := newMessage(5, "payload3")
	if err != nil {
		t.Error("new message3 failure")
	}

	msg4, err := newMessage(3, "payload4")
	if err != nil {
		t.Error("new message4 failure")
	}

	// add messages
	messageAdd(msg1, queue)
	messageAdd(msg2, queue)
	messageAdd(msg3, queue)
	messageAdd(msg4, queue)

	queueDump(queue)

	// ensure proper message list
	if queue.messageCounter != 4 {
		t.Error("bad message counter")
	}

	if queue.messageList.turn != 1 {
		t.Error("bad list root")
	}

	if queue.messageList.next.turn != 1 {
		t.Error("bad list 2")
	}

	if queue.messageList.next.next.turn != 3 {
		t.Error("bad list 3")
	}

	if queue.messageList.next.next.next.turn != 5 {
		t.Error("bad list 4")
	}

	// consume message list
	temp := messageRead(queue)
	temp = messageRead(queue)
	temp = messageRead(queue)
	temp = messageRead(queue)

	if queue.messageCounter != 0 {
		t.Error("bad message counter")
	}

	if temp.payload != "payload3" {
		t.Error("bad payload3")
	}

	queueDump(queue)

	// attempt to read empty list
	temp = messageRead(queue)

	if queue.messageCounter != 0 {
		t.Error("bad message counter")
	}

	if temp != nil {
		t.Error("bad nil payload")
	}
}
