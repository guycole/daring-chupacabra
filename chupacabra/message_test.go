// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestNewMessage(t *testing.T) {
	payload := "testaroo"

	msg, err := newMessage(payload)
	if err != nil {
		t.Error("new message failure")
	}

	if msg.payload != payload {
		t.Error("payload failure")
	}
}

func TestNewMessageBadPayload(t *testing.T) {
	msg, err := newMessage("")
	if err == nil {
		t.Error("new message should fail")
	}

	if msg != nil {
		t.Error("message should be nil")
	}
}
