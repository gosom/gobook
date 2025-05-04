# HTTP Server

The [http.Server](https://pkg.go.dev/net/http#Server) struct provides a configurable to create
an HTTP-server compared to the simple `http.ListenAndServe` function we have used.

## Basic HTTP Server Configuration

`http.Server` allows you to explicitly define server parameters:

```go
package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
```

## Key Configuration Options

### Timeouts

Proper timeouts are critical for production REST APIs

- **ReadTimeout**: Maximum duration for reading the entire request, including body
- **WriteTimeout**: Maximum duration before timing out writes of the response
- **IdleTimeout**: Maximum amount of time to wait for the next request when keep-alives are enabled

Without proper timeouts, your server may be vulnerable to slow client attacks or resource exhaustion.


### TLS Configuration

For secure HTTPS connections, you can configure TLS directly:

```go
server := &http.Server{
	Addr:    "127.0.0.1:8443",
	Handler: mux,
	TLSConfig: &tls.Config{
		MinVersion: tls.VersionTLS12,
	},
}

err := server.ListenAndServeTLS("server.crt", "server.key")
if err != nil {
	panic(err)
}
```


### Graceful Shutdown

One of the most important features of using http.Server directly is the ability to implement graceful shutdown, which is essential for production REST APIs:


```go
{{#include ../../goexamples/example-gracefulshut/main.go}}
```

### When to use http.Server and best practices

> Always use the http.Server. An exception might be non production test apps

- **Always set timeouts**
- **Implement graceful shutdown**
- **Use TLS on production**
