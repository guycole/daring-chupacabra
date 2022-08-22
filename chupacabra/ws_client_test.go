// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestClientMap(t *testing.T) {
	newClientMap()

	// assign all channels
	for ndx := 0; ndx < maxPubSubChannels; ndx++ {
		key, err := selectFreeClient()
		if err != nil {
			t.Errorf("unexpected selectFreeClient error:%s", err)
		}

		temp := wsClientMap[key]
		temp.active = true
	}

	// ensure all channels assigned
	for key, value := range wsClientMap {
		if !value.active {
			t.Errorf("unexpected available channel:%s", key)
		}
	}

	// should reject assignment
	key, err := selectFreeClient()
	if err == nil {
		t.Errorf("expected channel full error:%s", key)
	}
}
