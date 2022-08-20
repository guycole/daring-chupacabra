// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

type subscriberType struct {
	active      bool
	channelName string
}

// maximum subscriber connections
const maxSubscribers = 5

//
type subscriberArrayType [maxSubscribers]subscriberType

func newSubscriber(channelName string) *subscriberType {
	result := subscriberType{active: true, channelName: channelName}
	return &result
}

func addNewSubscriber(channelName string) {
	st := newSubscriber(channelName)
	pt = pt.newRegisterPayload(channelName)
	//	pt.publishPayload(backEndChannelName, rdb)
}
