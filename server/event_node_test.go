package main

import (
	"testing"
)

func TestEventNodeOperations(t *testing.T) {
	tests := []struct {
		candidate string
	}{
		{"81837d8a-2925-4b52-ab4f-31177a6b2f83"},
		{"4d0c6caa-5ad4-4505-b3d2-e951f5c838fc"},
	}

	eventNodeHeader := EventNodeHeaderType{Population: 0, Next: nil}

	for _, ndx := range tests {
		eventNodeHeader.insert(ndx.candidate)
	}

	if eventNodeHeader.Population != 2 {
		t.Errorf("TestEventNodeOperations failure")
	}

	temp1, err1 := eventNodeHeader.selectNext()
	if err1 != nil {
		t.Errorf("TestEventNodeOperations failure")
	}
	if temp1.ItemID != tests[1].candidate {
		t.Errorf("TestEventNodeOperations failure")
	}
	if eventNodeHeader.Population != 1 {
		t.Errorf("TestEventNodeOperations failure")
	}

	temp2, err2 := eventNodeHeader.selectNext()
	if err2 != nil {
		t.Errorf("TestEventNodeOperations failure")
	}
	if temp2.ItemID != tests[0].candidate {
		t.Errorf("TestEventNodeOperations failure")
	}
	if eventNodeHeader.Population != 0 {
		t.Errorf("TestEventNodeOperations failure")
	}

	_, err3 := eventNodeHeader.selectNext()
	if err3 == nil {
		t.Errorf("TestEventNodeOperations failure")
	}
	if eventNodeHeader.Population != 0 {
		t.Errorf("TestEventNodeOperations failure")
	}
}
