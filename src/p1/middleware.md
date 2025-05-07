# Understanding Middleware in Go REST APIs

## What is Middleware?

Middleware in the context of Go web applications acts as an intermediary layer between the incoming HTTP request and your application's handlers. 

It intercepts HTTP requests before they reach your final handler, allowing you to perform operations on the request or response. 

Essentially, middleware is an implementation of the http.Handler interface that wraps another handler, executing code before and/or after the handler is called.


## Basic Example

```go
{{#include ../../goexamples/example-middleware/main.go}}
```

In this example, loggingMiddleware is a function that takes an http.Handler and returns a new http.Handler that wraps the original one. 

It logs the start time of the request, calls the original handler, and then logs how long the request took to process.

## Chaining Middleware

Middleware chaining is simple but powerful. It's like passing a request through a series of checkpoints before reaching its destination:

```go
func main() {
    // Our core handler function
    handler := http.HandlerFunc(helloHandler)
    
    // Apply middleware in sequence
    handler = loggingMiddleware(handler)
    
    handler = authMiddleware(handler)
    
    handler = timeoutMiddleware(handler)
    
    http.Handle("/", handler)
    
    http.ListenAndServe(":8080", nil)
}
```

In this example, the middleware execution happens in the opposite order of how they're applied:

1. When a request comes in, it first goes through timeoutMiddleware
2. Then it passes through authMiddleware
3. Next it goes through loggingMiddleware
4. Finally, it reaches the helloHandler

This happens because each middleware wraps the next one in the chain, creating nested layers. 
>Understanding this execution order is crucial when designing your middleware chain, especially when certain middleware depends on the processing done by others.


## Common Use Cases

Some common use cases include:

1. **Logging**: Recording request details, response times, and errors for monitoring and debugging.
2. **Authentication and Authorization**: Verifying user identity and checking if they have permission to access certain resources.
3. **Request Validation**: Ensuring requests contain required fields or meet specific criteria before they reach your handlers.
4. **CORS (Cross-Origin Resource Sharing)**: Managing which domains can access your API.
5. **Rate Limiting**: Preventing abuse by limiting the number of requests from a single client.
6. **Response Compression**: Automatically compressing HTTP responses to reduce bandwidth.
7. **Error Handling**: Providing consistent error responses across your API.
8. **Content Type Negotiation**: Processing requests based on their content type and formatting responses accordingly.
9. **Request Parsing**: Automatically parsing JSON or form data into structured types.
10. **Caching**: Storing response data to improve performance for frequent requests.

