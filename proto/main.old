package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const banner = "prom-target 0.0"

var (
	guyRandom = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "guy_random",
			Help: "guy random value",
		},
	)
)

func getRoot(ww http.ResponseWriter, rr *http.Request) {
	log.Println("get root")
	ww.Header().Set("Content-Type", "text/html")
	ww.WriteHeader(http.StatusOK)
	fmt.Fprintf(ww, "Welcome to PromTarget")
}

func updateMetrics() {
	go func() {
		prometheus.MustRegister(guyRandom)

		for {
			log.Println("update metrics")
			guyRandom.Set(rand.NormFloat64())
			time.Sleep(3 * time.Second)
		}
	}()
}

func main() {
	log.Println(banner)

	rand.New(rand.NewSource(time.Now().UnixNano()))

	updateMetrics()

	router := mux.NewRouter()
	router.HandleFunc("/", getRoot).Methods("GET")
	router.Path("/metrics").Handler(promhttp.Handler())

	log.Fatal(http.ListenAndServe(":80", router))
}
