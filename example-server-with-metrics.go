package main

import (
	"net/http"
	"time"

	"fmt"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter_requests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "example",
		Name:      "requests",
		Help:      "Number of requests",
	})

func main() {
	// Register request
	prometheus.MustRegister(counter_requests)

	// Add /metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter_requests.Add(1)
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/sleep10", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		counter_requests.Add(1)
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/sleep50", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		counter_requests.Add(1)
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/sleep100", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		counter_requests.Add(1)
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/sleep200", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		counter_requests.Add(1)
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/sleep500", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond)
		counter_requests.Add(1)
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	fmt.Println("Server started.")
	log.Fatal(http.ListenAndServe(":80", nil))
}
