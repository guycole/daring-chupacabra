package main

import (
	"errors"
)

type EventNodeType struct {
	ItemID string
	Next   *EventNodeType
}

type EventNodeHeaderType struct {
	Population int
	Next       *EventNodeType
}

func (eventNodeHeader *EventNodeHeaderType) insert(itemID string) {
	candidate := EventNodeType{ItemID: itemID, Next: nil}

	if eventNodeHeader.Population == 0 {
		eventNodeHeader.Next = &candidate
	} else {
		candidate.Next = eventNodeHeader.Next
		eventNodeHeader.Next = &candidate
	}

	eventNodeHeader.Population++
}

func (eventNodeHeader *EventNodeHeaderType) selectNext() (*EventNodeType, error) {
	var result *EventNodeType

	switch eventNodeHeader.Population {
	case 0:
		return nil, errors.New("empty list")
	case 1:
		result = eventNodeHeader.Next
		eventNodeHeader.Next = nil
	default:
		result = eventNodeHeader.Next
		eventNodeHeader.Next = eventNodeHeader.Next.Next
	}

	eventNodeHeader.Population--
	result.Next = nil

	return result, nil
}
