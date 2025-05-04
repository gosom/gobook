# http.Handler

A Handler is an implementation of the [`http.Handler` interface]([https://pkg.go.dev/net/http#Handler)

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

This interface has only one method, `ServeHTTP`. Essentially, this is the code that is responsible to repond to an HTTP request.
