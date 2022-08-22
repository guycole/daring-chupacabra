// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

type wsPubSubType struct {
	// Registered clients.
	clients map[*wsClientType]bool

	// Inbound messages from the clients.
	//broadcast chan []byte

	// Register requests from the clients.
	register chan *wsClientType

	// Unregister requests from clients.
	unregister chan *wsClientType
}
