// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

type ResponseMapType map[string]*ResponseNodeType

func initializeResponseMap() *ResponseMapType {
	responseMap := make(ResponseMapType)
	return &responseMap
}

func (responseMap *ResponseMapType) insertNode(candidate *ResponseNodeType) {
	candidate.Next = nil

	temp, ok := responseMap[candidate.ClientID]
	if ok {
		// traffic already waiting for this client
		candidate.Next = temp
		responseMap[candidate.ClientID] = candidate
	} else {
		// no traffic waiting for this client
		responseMap[candidate.ClientID] = candidate
	}

	(*responseMap)[candidate.ClientID] = candidate
}
