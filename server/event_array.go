// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
)

const maxEventNodeHeader = 100

type EventArrayType [maxEventNodeHeader]*EventNodeHeaderType

func initializeEventArray() *EventArrayType {
	eventArray := new(EventArrayType)

	for ii := 0; ii < maxEventNodeHeader; ii++ {
		candidate := EventNodeHeaderType{Population: 0, Next: nil}
		eventArray[ii] = &candidate
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

func (eventArray *EventArrayType) insertNode(itemID string, turn int) {
	ndx := turn % maxEventNodeHeader
	eventArray[ndx].insertNode(itemID)
}

func (eventArray *EventArrayType) selectNextNode(turn int) (*EventNodeType, error) {
	ndx := turn % maxEventNodeHeader
	result, err := eventArray[ndx].selectNextNode()
	return result, err
}
