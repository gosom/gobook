# Hello world

If you don't have Golang installed please 
Follow the instructions [here](https://go.dev/doc/install)

Now let's start our journey by just saying `Hello world`

```go
{{#include ../../goexamples/example-helloworld/main.go}}
```

### Run it

> All the book examples can be run if you clone the book repo.
> go to the goexamples directory and run them using
> `go run ./example-name`

start the HTTP server:

```bash
cd goexamples
go run ./example-helloworld
```

Try it:

> For testing the web server we are going to use curl


Open another terminal

```bash
curl -i http://localhost:8080
```

### What's happenning here

With only 17 lines of code Go:

1. Starts a webserver that is listening on port 8080
2. When an HTTP request is made it executes the `handler` function


### Next Steps

In the next section we are going to explain in more detail what is happening.
We are going to exampine the [net/http](https://pkg.go.dev/net/http) and in particular
we are going to focus on `http.Server, http.Request, http.ResponseWriter, http.ServeMux`.

