// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

// banner splash message
const banner = "daring chupacabra 0.0"

func main() {
	log.Println(banner)

	rand.Seed(time.Now().UnixNano())

	run_mode := os.Getenv("RUN_MODE")
	log.Println("RunMode:", run_mode)

	switch run_mode {
	case "backend":
		backEnd()
	case "frontend":
		frontEnd()
	default:
		log.Printf("unknown run mode %s\n", run_mode)
	}
}
