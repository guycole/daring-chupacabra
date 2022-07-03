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

func TestQueueMessageAdd(t *testing.T) {
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

	messageAdd(msg1, queue)
	messageAdd(msg2, queue)
	messageAdd(msg3, queue)
	messageAdd(msg4, queue)

	queueDump(queue)

	if queue.messageCounter != 4 {
		t.Error("bad message counter")
	}

	if queue.next.turn != 1 {
		t.Error("bad list root")
	}

	if queue.next.next.turn != 1 {
		t.Error("bad list 2")
	}

	if queue.next.next.next.turn != 3 {
		t.Error("bad list 3")
	}

	if queue.next.next.next.next.turn != 5 {
		t.Error("bad list 4")
	}
}
