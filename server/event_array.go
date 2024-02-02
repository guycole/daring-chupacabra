// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
	"sync"
)

const maxEventNodeHeader = 100

type EventArrayType [maxEventNodeHeader]*EventNodeHeaderType

var eventArrayLock = sync.RWMutex{}

func initializeEventArray() *EventArrayType {
	eventArray := new(EventArrayType)

	for ii := 0; ii < maxEventNodeHeader; ii++ {
		eventArray[ii] = &EventNodeHeaderType{Population: 0, Next: nil}
	}

	return eventArray
}

func (eventArray *EventArrayType) dumper() {
	for ii := 0; ii < maxEventNodeHeader; ii++ {
		if eventArray[ii].Population > 0 {
			for temp := eventArray[ii].Next; temp != nil; temp = temp.Next {
				log.Println("ndx:", ii, "itemID:", temp.ItemID)
			}
		}
	}
}

func (eventArray *EventArrayType) insertNode(candidate *EventNodeType, turn int) {
	eventArrayLock.Lock()
	defer eventArrayLock.Unlock()

	eventArray[turn%maxEventNodeHeader].insertNode(candidate)
}

func (eventArray *EventArrayType) selectNextNode(turn int) (*EventNodeType, error) {
	eventArrayLock.RLock()
	defer eventArrayLock.RUnlock()

	result, err := eventArray[turn%maxEventNodeHeader].selectNextNode()
	return result, err
}
