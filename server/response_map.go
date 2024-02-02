// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

type ResponseMapType map[string]*ResponseNodeType

func initializeResponseMap() *ResponseMapType {
	responseMap := make(ResponseMapType)
	return &responseMap
}

func (responseMap *ResponseMapType) insertNode(candidate *ResponseNodeType, clientId string) {
	candidate.Next = nil

	temp, ok := (*responseMap)[clientId] // Dereference the pointer to access the underlying map
	if ok {
		// traffic already waiting for this client
		candidate.Next = temp
		(*responseMap)[clientId] = candidate // Dereference the pointer to access the underlying map
	} else {
		// no traffic waiting for this client
		(*responseMap)[clientId] = candidate // Dereference the pointer to access the underlying map
	}

	//(*responseMap)[candidate.ClientID] = candidate
}

func (responseMap *ResponseMapType) getResponse(clientId string) []*ResponseNodeType {
	return nil
}
