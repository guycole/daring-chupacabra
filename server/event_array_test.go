package main

import (
	"testing"
)

func TestEventArrayInitialization(t *testing.T) {
	eventArray := initializeEventArray()

	for ii := 0; ii < maxEventNodeHeader; ii++ {
		if eventArray[ii] == nil {
			t.Errorf("TestEventArray(%d) failure", ii)
		}

		if eventArray[ii].Population != 0 {
			t.Errorf("TestEventArray(%d) failure", ii)
		}

		if eventArray[ii].Next != nil {
			t.Errorf("TestEventArray(%d) failure", ii)
		}
	}
}

func TestEventArrayOperations(t *testing.T) {
	const turn = 5

	eventArray := initializeEventArray()

	tests := []struct {
		candidate string
		turn      int
	}{
		{"81837d8a-2925-4b52-ab4f-31177a6b2f83", turn},
		{"4d0c6caa-5ad4-4505-b3d2-e951f5c838fc", turn + 100},
	}

	for _, ndx := range tests {
		eventArray.insert(ndx.candidate, ndx.turn)
	}

	if eventArray[turn].Population != 2 {
		t.Errorf("TestEventArrayOperations failure")
	}

	temp1, err1 := eventArray.selectNext(turn)
	if err1 != nil {
		t.Errorf("TestEventArrayOperations failure")
	}
	if temp1.ItemID != tests[1].candidate {
		t.Errorf("TestEventArrayOperations failure")
	}
	if eventArray[turn].Population != 1 {
		t.Errorf("TestEventArrayOperations failure")
	}

	temp2, err2 := eventArray.selectNext(turn)
	if err2 != nil {
		t.Errorf("TestEventArrayOperations failure")
	}
	if temp2.ItemID != tests[0].candidate {
		t.Errorf("TestEventArrayOperations failure")
	}
	if eventArray[turn].Population != 0 {
		t.Errorf("TestEventArrayOperations failure")
	}

	_, err3 := eventArray.selectNext(turn)
	if err3 == nil {
		t.Errorf("TestEventArrayOperations failure")
	}
	if eventArray[turn].Population != 0 {
		t.Errorf("TestEventArrayOperations failure")
	}
}
