// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

type ResponseNodeType struct {
	Message   string
	ReceiptID string
	Next      *ResponseNodeType
}

func newResponseNode(message, receiptId string) *ResponseNodeType {
	return &ResponseNodeType{Message: message, ReceiptID: receiptId}
}

/*
func (eventNodeHeader *EventNodeHeaderType) insertNode(candidate *EventNodeType) {
	candidate.Next = nil

	if eventNodeHeader.Population == 0 {
		eventNodeHeader.Next = candidate
	} else {
		candidate.Next = eventNodeHeader.Next
		eventNodeHeader.Next = candidate
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
	result.Next = nil // prevent leak of next node

	return result, nil
}
*/
