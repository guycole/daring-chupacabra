// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"log"
)

type messageType struct {
	payload string
}

func newMessage(payload string) (*messageType, error) {
	if len(payload) < 1 {
		return nil, errors.New("bad payload")
	}

	result := messageType{payload: payload}

	//	start := time.Now()

	return &result, nil
}

func encodeMessage(message messageType) {
	payload, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	log.Println(payload)
}
