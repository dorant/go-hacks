package main

// https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dorant/go-http/handlers"
	"github.com/dorant/go-http/version"
)

func main() {
	log.Print("Starting the service..")
	log.Printf("Commit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	router := handlers.Router(version.Commit, version.BuildTime, version.Release)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	log.Print("Service is ready to serve")

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}
	log.Print("Service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")
}
