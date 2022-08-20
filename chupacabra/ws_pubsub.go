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

/*
func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *wsClientType),
		unregister: make(chan *wsClientType),
		clients:    make(map[*wsClientType]bool),
	}
}
*/

/*
func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
*/
