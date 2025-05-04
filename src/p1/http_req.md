# http.Request

[http.Request](https://pkg.go.dev/net/http#Request) is a fundamental component of Go's web programming toolkit. It encapsulates all the information received from an incoming HTTP request and allows you to process and read the data.

**What you need to know**
When your handler function is called, it receives an http.Request pointer that contains everything about the request:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // r contains all the request information
}
```

**Key Features**
The request objects gives you access to:

- **Request Method**: via `r.Method`
- **URL and Path**: via `r.URL`
- **Headers**: via `r.Header.Get("header-name")`
- **Query Params**: via `r.URL.Query().Get("param")`
- **Form Data**: via `r.ParseForm()`
- **Request Body**: via `io.ReadAll(r.Body))`
- **Path Variables**: via `r.PathValue("variable-name")`


Understanding the `http.Request` type is crucial because it's the gateway to all client information. When building a web application you will constantly work with the request object.

I prefer not to explain all the possible methods and fields. We will focus to the ones that
solve the immediate problem. As we build a more complex application, we will discover and use
more and more of the `http.Request` functionality.


