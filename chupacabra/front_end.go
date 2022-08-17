// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"time"

	redis "github.com/go-redis/redis/v8"
)

func frontEnd() {
	log.Println("frontEnd entry")

	redisAddress := os.Getenv("REDIS_ADDRESS")
	log.Println(redisAddress)

	redisPassword := os.Getenv("REDIS_PASSWORD")
	log.Println(redisPassword)

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0, // use default DB
	})

	backEndChannelName := os.Getenv("BE_CHANNEL")
	//frontEndChannelName := os.Getenv("FE_CHANNEL")

	pt := newRegisterPayload("channel1")
	publishPayload(pt, backEndChannelName, rdb)

	pt = newRegisterPayload("channel2")
	publishPayload(pt, backEndChannelName, rdb)

	pt = newRegisterPayload("channel3")
	publishPayload(pt, backEndChannelName, rdb)

	//http.HandleFunc("/", echo)
	//http.ListenAndServe(":8080", nil)

	for {
		log.Println("Infinite Loop 2")
		time.Sleep(time.Second * 10)
	}
}

/*
func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "home.html")
}

var addr = flag.String("addr", ":8080", "http service address")

func frontEnd2() {
	log.Println("frontEnd entry")

	flag.Parse()
	log.Println(addr)

	hub := newHub()
	go hub.run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
*/
