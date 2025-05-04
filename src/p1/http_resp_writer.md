# http.ResponseWriter

[http.ResponseWriter](https://pkg.go.dev/net/http#ResponseWriter) is an interface that is used to construct the HTTP Response.

In the `helloworld` example we used its `Write` method.

```go
Write([]byte) (int, error)
```


> The `Write` method writes the data to the underlying TCP connection as part of an HTTP reply.

Each HTTP Response contains a Body and a Status Code. The write method writes the body. 

What about the Status code? 

When we run `curl -i http://localhost:8080` the response was something like:

```bash
HTTP/1.1 200 OK
Date: Sat, 03 May 2025 16:33:22 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

Hello world!
```

The HTTP response status code is `200`.

In order to use another status code you need to call the `WriteHeader` method **before** calling `Write` . 


```go
w.WriteHeader(201)
w.Write([]byte("this is other status code`)
```

The above will make an HTTP response with response code `201`

**For HTTP response codes go has [constants defined](https://pkg.go.dev/net/http#pkg-constants) and these the ones that should be used**

Example:

```go
w.WriteHeader(http.StatusCreated) // 201 
```

`WriteHeader` method is used to write all the required HTTP Headers.

Example:

```go
w.WriteHeader("Content-Type", "application/json")
```
