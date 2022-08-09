// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"time"
)

// banner splash message
const banner = "daring chupacabra 0.0"

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func eclectic2() {
	done := make(chan bool)
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case <-ticker.C:
			fmt.Println("Hello !!")
		}
	}

	// wait for 10 seconds
	time.Sleep(10 * time.Second)
	done <- true
}

func main() {
	log.Println(banner)
	log.Println(configuration)

	//rand.Seed(time.Now().UnixNano())

	//eclectic()

	//webPortal()

	//	done := make(chan bool)
	//	ticker := time.NewTicker(1 * time.Second)

	//	go func() {
	//		for {
	//			select {
	//			case <-done:
	//				ticker.Stop()
	//				return
	//			case <-ticker.C:
	//				fmt.Println("Hello !!")
	//			}
	//		}
	//	}()

	// wait for 10 seconds
	//	time.Sleep(10 * time.Second)
	//	done <- true

	//go eclectic()
	//	log.Println("back")

	//go say("world")
	//say("hello")

	//	fmt.Println("Hello, Modules!")
	//http.Handle("/metrics", promhttp.Handler())
	//http.ListenAndServe(":4190", nil)
}
