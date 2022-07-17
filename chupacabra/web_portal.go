// Copyright 2022 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// IndexHandler bogus
func indexHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<h1>Daring Chupacabra</h1>"))
}

// AboutHandler bogus
func aboutHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<h1>This is the about page</h1>"))
}

func webPortal() {
	log.Println("webPortal entry")

	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/about", aboutHandler)

	// starting up the server
	server := &http.Server{
		// Addr:           configuration.Address,
		Addr:           "0.0.0.0:8088",
		Handler:        router,
		ReadTimeout:    time.Duration(configuration.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(configuration.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())

	log.Println("webPortal exit")
}
