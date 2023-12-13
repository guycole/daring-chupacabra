// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
)

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

	newClientMap()

	flag.Parse()
	log.Println(addr)

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(ww http.ResponseWriter, rr *http.Request) {
		serveWebSocket(ww, rr)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
