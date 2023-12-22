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

func (eventArray *EventArrayType) insert(itemID string, turn int) {
	ndx := turn % maxEventNodeHeader
	eventArray[ndx].insert(itemID)
}

func (eventArray *EventArrayType) selectNext(turn int) (*EventNodeType, error) {
	ndx := turn % maxEventNodeHeader
	result, err := eventArray[ndx].selectNext()
	return result, err
}
