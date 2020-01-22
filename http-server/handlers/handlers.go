package handlers

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
)

// Router registers routes and returns a instance of a router.
func Router(commit, buildTime, release string) *mux.Router {
	isReady := &atomic.Value{}
	isReady.Store(false)

	go func() {
		log.Printf("Readyz probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()

	r := mux.NewRouter()
	r.HandleFunc("/home", home(commit, buildTime, release)).Methods("GET")
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz(isReady))
	return r
}
