// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestEventNodeHeader(t *testing.T) {
	tests := []struct {
		action    EventActionEnum
		candidate string
		tokenType CatalogTokenEnum
	}{
		{nominalAction, "81837d8a-2925-4b52-ab4f-31177a6b2f83", obj1Token},
		{moveAction, "4d0c6caa-5ad4-4505-b3d2-e951f5c838fc", obj2Token},
	}

	eventNodeHeader := EventNodeHeaderType{Population: 0, Next: nil}

	for _, ndx := range tests {
		eventNodeHeader.insertNode(newEventNode(ndx.action, ndx.candidate, ndx.tokenType))
	}

	if eventNodeHeader.Population != 2 {
		t.Errorf("TestEventNodeOperations failure")
	}

	temp1, err1 := eventNodeHeader.selectNextNode()
	if err1 != nil {
		t.Errorf("TestEventNodeOperations failure")
	}
	if temp1.Action != tests[1].action || temp1.ItemID != tests[1].candidate || temp1.TokenType != tests[1].tokenType {
		t.Errorf("TestEventNodeOperations failure")
	}

	if eventNodeHeader.Population != 1 {
		t.Errorf("TestEventNodeOperations failure")
	}

	temp2, err2 := eventNodeHeader.selectNextNode()
	if err2 != nil {
		t.Errorf("TestEventNodeOperations failure")
	}
	if temp2.Action != tests[0].action || temp2.ItemID != tests[0].candidate || temp2.TokenType != tests[0].tokenType {
		t.Errorf("TestEventNodeOperations failure")
	}

	if eventNodeHeader.Population != 0 {
		t.Errorf("TestEventNodeOperations failure")
	}

	if _, err3 := eventNodeHeader.selectNextNode(); err3 == nil {
		t.Errorf("TestEventNodeOperations failure")
	}

	if eventNodeHeader.Population != 0 {
		t.Errorf("TestEventNodeOperations failure")
	}
}
