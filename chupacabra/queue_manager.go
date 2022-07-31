// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"errors"
	"log"
)

type queueType struct {
	name string
}

const maxQueue = 16

type queueArrayType [maxQueue]*queueType

func newQueue(name string) (*queueType, error) {
	if len(name) < 1 {
		return nil, errors.New("bad name")
	}

	result := queueType{name: name}
	return &result, nil
}

func userAdd(name string, queueArray queueArrayType) (int, error) {
	// duplicate name test
	for ndx := 0; ndx < maxQueue; ndx++ {
		if (queueArray[ndx] != nil) && (queueArray[ndx].name == name) {
			return -1, errors.New("duplicate name")
		}
	}

	// queue creation
	candidate, err := newQueue(name)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	// queue insertion
	for ndx := 0; ndx < maxQueue; ndx++ {
		if queueArray[ndx] == nil {
			queueArray[ndx] = candidate
			return ndx, nil
		}
	}

	return -1, errors.New("queueArray full")
}

func userDelete(name string, queueArray queueArrayType) int {
	for ndx := 0; ndx < maxQueue; ndx++ {
		if queueArray[ndx] == nil {
			queueArray[ndx] = nil
			return ndx
		}
	}

	return -1
}
