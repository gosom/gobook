package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// Create our handler
	handler := http.HandlerFunc(helloHandler)

	// Wrap it with the middleware
	wrappedHandler := loggingMiddleware(handler)

	// Register the wrapped handler
	http.Handle("/", wrappedHandler)

	// Start the server
	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Code executed before the handler
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Call the wrapped handler
		next.ServeHTTP(w, r)

		// Code executed after the handler
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
