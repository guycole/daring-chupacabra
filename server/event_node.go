// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"errors"
)

type EventActionEnum int

const (
	nothingAction EventActionEnum = iota
	createAction
	deleteAction
	houseKeepingAction
	moveAction
)

func (eae EventActionEnum) String() string {
	return [...]string{"nothing", "create", "delete", "housekeeping", "move"}[eae]
}

type EventNodeType struct {
	Action EventActionEnum
	ItemID string
	Next   *EventNodeType
}

type EventNodeHeaderType struct {
	Population int
	Next       *EventNodeType
}

func (eventNodeHeader *EventNodeHeaderType) insertNode(action EventActionEnum, itemID string) {
	candidate := EventNodeType{Action: action, ItemID: itemID, Next: nil}

	if eventNodeHeader.Population == 0 {
		eventNodeHeader.Next = &candidate
	} else {
		candidate.Next = eventNodeHeader.Next
		eventNodeHeader.Next = &candidate
	}

	eventNodeHeader.Population++
}

func (eventNodeHeader *EventNodeHeaderType) selectNextNode() (*EventNodeType, error) {
	var result *EventNodeType

	switch eventNodeHeader.Population {
	case 0:
		return nil, errors.New("empty event list")
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
