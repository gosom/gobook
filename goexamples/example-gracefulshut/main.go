package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := run()

	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}

	log.Println("Server stopped gracefully")
}

func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	errc := make(chan error, 1)

	// start the server in a goroutine
	go func() {
		defer close(errc)
		// when shutdown is called, the server will immediatelly return and stop accepting new connections
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errc <- err

			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	// we give the server 5 seconds to finish the in progress requests
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Disable keep-alives
	server.SetKeepAlivesEnabled(false)

	// we should wait for the server to finish
	err := server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return <-errc
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
