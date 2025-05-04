# net/http package

This package contains implementations for HTTP servers (and clients).
Understanding the basic structs and functions this package offers is crucial for 
getting a solid foundation of how Web Programming can be made in Go.

We are going to explore the package by using our simple Hello world example

On a high level the hello world program does:

- starts an HTTP Server: The HTTP Server is a program that "listens/binds" to a network port is a capable of serving HTTP requests to that port. 
- It registers a function to be called when a specific HTTP request is made.

The HTTP server starts using:

[http.ListenAndServe](https://pkg.go.dev/net/http#ListenAndServe)

The [Documentation](https://pkg.go.dev/net/http#hdr-Servers) explains it a clean way:

>ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux

However there are some terms that the reader may not be familiar with.

## Exercise

> Add a route /info that returns information about Golang's runtime

Solution:

```go
{{#include ../../goexamples/example-helloinfo/main.go}}
```
