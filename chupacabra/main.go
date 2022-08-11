// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

// banner splash message
const banner = "daring chupacabra 0.0"

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

	run_mode := os.Getenv("RUN_MODE")
	log.Println("RunMode:", run_mode)

	rand.Seed(time.Now().UnixNano())

	// TODO get these arguments from secrets
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-master:6379",
		Password: "bigSekret",
		DB:       0, // use default DB
	})

	log.Println(rdb)

	if strings.Compare(run_mode, "backend") == 0 {
		log.Println("backend mode")
	}

	fmt.Println(strings.Compare("GeeksforGeeks",
		"GeeksforGeeks"))

	for true {
		log.Println("Infinite Loop 2")
		time.Sleep(time.Second)
	}

	//    topic := rdb.Subscribe(context.Background(), channelName)

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
