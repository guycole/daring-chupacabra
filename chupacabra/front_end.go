// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func frontEnd2() {
	log.Println("frontEnd entry")

	flag.Parse()

	//rdb := newRedisClient()

	//backEndChannelName := os.Getenv("BE_CHANNEL")
	//frontEndChannelName := os.Getenv("FE_CHANNEL")

	//addNewSubscriber("channel1")
	//addNewSubscriber("channel2")
	//addNewSubscriber("channel3")

	/*
		pt := newPayload("newId", unknownPayload, frontEndChannelName)
		if err != nil {
			log.Panic("tire ripper")
		}

		pt = pt.newRegisterPayload("channel1")
		pt.publishPayload(backEndChannelName, rdb)

		pt = pt.newRegisterPayload("channel2")
		pt.publishPayload(backEndChannelName, rdb)

		pt = pt.newRegisterPayload("channel3")
		pt.publishPayload(backEndChannelName, rdb)
	*/

	//http.HandleFunc("/", echo)
	//http.ListenAndServe(":8080", nil)

	for {
		log.Println("Infinite Loop 2")
		time.Sleep(time.Second * 10)
	}
}

func serveHome(ww http.ResponseWriter, rr *http.Request) {
	log.Println(rr.URL)

	if rr.URL.Path != "/" {
		http.Error(ww, "not found", http.StatusNotFound)
		return
	}

	if rr.Method != http.MethodGet {
		http.Error(ww, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(ww, rr, "home.html")
}

var addr = flag.String("addr", ":8080", "http service address")

func frontEnd() {
	log.Println("frontEnd entry")

	flag.Parse()
	log.Println(addr)

	//hub := newHub()
	//go hub.run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(ww http.ResponseWriter, rr *http.Request) {
		serveWs(ww, rr)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
