// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestEventArrayInitialization(t *testing.T) {
	eventArray := initializeEventArray()

	for ii := 0; ii < maxEventNodeHeader; ii++ {
		if eventArray[ii] == nil || eventArray[ii].Population != 0 || eventArray[ii].Next != nil {
			t.Errorf("TestEventArray(%d) failure", ii)
		}
	}
}

func TestEventArrayOperations(t *testing.T) {
	const turn = 5

	tests := []struct {
		action    EventActionEnum
		candidate string
		tokenType CatalogTokenEnum
		turn      int
	}{
		{nominalAction, "81837d8a-2925-4b52-ab4f-31177a6b2f83", obj1Token, turn},
		{moveAction, "4d0c6caa-5ad4-4505-b3d2-e951f5c838fc", obj2Token, turn + maxEventNodeHeader},
	}

	eventArray := initializeEventArray()

	for _, ndx := range tests {
		eventArray.insertNode(newEventNode(ndx.action, ndx.candidate, ndx.tokenType), ndx.turn)
	}

	if eventArray[turn].Population != 2 {
		t.Errorf("TestEventArrayOperations failure")
	}

	// event node operations thoroughly tested in event_node_test.go
	temp1, err1 := eventArray.selectNextNode(turn)
	if err1 != nil {
		t.Errorf("TestEventArrayOperations failure")
	}
	if temp1.ItemID != tests[1].candidate || temp1.TokenType != tests[1].tokenType {
		t.Errorf("TestEventArrayOperations failure")
	}

	if eventArray[turn].Population != 1 {
		t.Errorf("TestEventArrayOperations failure")
	}
}
