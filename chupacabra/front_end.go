// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
	"time"
)

func frontEnd() {
	log.Println("front end")

	for true {
		log.Println("Infinite Loop 2")
		time.Sleep(time.Second)
	}
}
