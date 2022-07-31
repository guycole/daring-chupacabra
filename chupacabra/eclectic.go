// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
)

type eclecticType struct {
	turnCounter int
	queueArray  queueArrayType
}

func stuntBox(state *eclecticType) {

}

func preTurnCycle(state *eclecticType) {
	log.Println("pre turn cycle")
}

func postTurnCycle(state *eclecticType) {
	log.Println("post turn cycle")
}

func turnCycle(state *eclecticType) {
	state.turnCounter++
	log.Println("turn counter:", state.turnCounter)

	preTurnCycle(state)

	for ndx := 0; ndx < maxQueue; ndx++ {
		if state.queueArray[ndx] == nil {
			log.Println("skipping nil queue:", ndx)
		} else {
			log.Println("working queue:", ndx)
		}
	}

	postTurnCycle(state)
}

// runs forever, loops through queues
func eclectic() {
	state := eclecticType{}

	log.Println("eclectic entry")

	stuntBox(&state)

	for ndx := 0; ndx < 10; ndx++ {
		turnCycle(&state)
	}

	log.Println("eclectic exit")
}
