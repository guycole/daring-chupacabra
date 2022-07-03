// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestNewMessage(t *testing.T) {
	turn := 1
	payload := "testaroo"

	msg, err := newMessage(turn, payload)
	if err != nil {
		t.Error("new message failure")
	}

	if msg.turn != turn {
		t.Error("turn failure")
	}

	if msg.payload != payload {
		t.Error("payload failure")
	}
}

func TestNewMessageBadTurn(t *testing.T) {
	msg, err := newMessage(0, "bogus")
	if err == nil {
		t.Error("new message should fail")
	}

	if msg != nil {
		t.Error("message should be nil")
	}
}

func TestNewMessageBadPayload(t *testing.T) {
	msg, err := newMessage(1, "")
	if err == nil {
		t.Error("new message should fail")
	}

	if msg != nil {
		t.Error("message should be nil")
	}
}
