# HTTP Request multiplexer

`DefaultServerMux`: Is a default implementation of the [ServeMux](https://pkg.go.dev/net/http#ServeMux). 

Quoting the docs:

> ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

Let provide an example of what the above means in the context of our hello world:

```go
http.HandleFunc("/", handler)
```

That line instructs the `DefaultServerMux` to "map" any HTTP request with to call
the `handler` function. 


A careful reader will notice that the `handler` function is just a function and it is not 
a struct implementing the `http.Handler` interface. 

`http.HandleFunc` is a convenience function that allows you to pass a function instead of a struct. 

```go
http.HandleFunc("/", handler)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
```

is equivalent to:

```go
hn := HelloHandler{}
http.Handle("/", hn)

type HelloHandler struct {
}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello world"))
}
```

Let's see some examples of different patterns that can be used.

```go
{{#include ../../goexamples/example-muxpatterns/main.go}}
```

Start the server:

```bash
go run ./example-muxpatterns
```

```bash
curl -i http://localhost:8080/hello
curl -XPOST -i http://localhost:8080/hello
curl -XGET -i http://localhost:8080/say-bye`
curl -XPOST -i http://localhost:8080/say-bye
curl -XGET -i http://localhost:8080/say/this-works
```

See in the [documentation](https://pkg.go.dev/net/http#hdr-Patterns-ServeMux) how patterns work

Also notice the response of `curl -XPOST -i http://localhost:8080/say-bye`:

```bash
HTTP/1.1 405 Method Not Allowed
Allow: GET, HEAD
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sun, 04 May 2025 05:44:41 GMT
Content-Length: 19

Method Not Allowed
```

We are not going to dive into more details at the moment.

>The important part is to understand that Go gives you the ability to register a method
that is executed based on the pattern. The pattern can contain a **Method** a **Path** and that
the path can contain a **wildcard** (see the handler3 example).

