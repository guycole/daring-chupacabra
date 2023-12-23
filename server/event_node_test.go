// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestEventNodeHeader(t *testing.T) {
	tests := []struct {
		candidate string
	}{
		{"81837d8a-2925-4b52-ab4f-31177a6b2f83"},
		{"4d0c6caa-5ad4-4505-b3d2-e951f5c838fc"},
	}

	eventNodeHeader := EventNodeHeaderType{Population: 0, Next: nil}

	for _, ndx := range tests {
		eventNodeHeader.insertNode(ndx.candidate)
	}

	if eventNodeHeader.Population != 2 {
		t.Errorf("TestEventNodeOperations failure")
	}

	temp1, err1 := eventNodeHeader.selectNextNode()
	if err1 != nil {
		t.Errorf("TestEventNodeOperations failure")
	}
	if temp1.ItemID != tests[1].candidate {
		t.Errorf("TestEventNodeOperations failure")
	}
	if eventNodeHeader.Population != 1 {
		t.Errorf("TestEventNodeOperations failure")
	}

	temp2, err2 := eventNodeHeader.selectNextNode()
	if err2 != nil {
		t.Errorf("TestEventNodeOperations failure")
	}
	if temp2.ItemID != tests[0].candidate {
		t.Errorf("TestEventNodeOperations failure")
	}
	if eventNodeHeader.Population != 0 {
		t.Errorf("TestEventNodeOperations failure")
	}

	_, err3 := eventNodeHeader.selectNextNode()
	if err3 == nil {
		t.Errorf("TestEventNodeOperations failure")
	}
	if eventNodeHeader.Population != 0 {
		t.Errorf("TestEventNodeOperations failure")
	}
}
