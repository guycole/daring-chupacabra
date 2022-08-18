// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
	"time"
)

type eventManagerType struct {
	shutDownFlag bool // true, graceful exit
	turnCounter  int  // current game turn
}

func eventManager() {
	log.Println("eventManager entry")

	var emt eventManagerType

	emt.shutDownFlag = false
	emt.turnCounter = 0

	for !emt.shutDownFlag {
		emt.turnCounter++

		log.Println("event manager loop")
		time.Sleep(time.Second * 10)
	}
}
