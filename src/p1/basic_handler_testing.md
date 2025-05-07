# Basic Handler Testing

Testing HTTP handlers is an essential part of building reliable REST APIs. Go's standard library provides the `httptest` package, which makes testing HTTP handlers straightforward without requiring a running server.

## The httptest Package

The `httptest` package offers tools to simulate HTTP requests and record responses for verification:

```go
import (
	"net/http"
	"net/http/httptest"
	"testing"
)
```

## Testing a Simple Handler

Let's test our basic "Hello world" handler:

```go
{{#include ../../goexamples/example-helloworld/main_test.go}}
```

From within the `goexamples/example-helloworld` run:

```bash
go test -v ./...
```

output:

```bash
=== RUN   TestHandler
--- PASS: TestHandler (0.00s)
PASS
ok      goexamples/helloworld   0.003s
```

## Best Practices

When writing handler tests:

1. **Test status codes**: Verify your handler returns the expected HTTP status codes
2. **Test response headers**: Check for expected headers like `Content-Type`
3. **Test response body**: Validate the content of the response
4. **Test error cases**: Ensure handlers respond appropriately to invalid inputs
5. **Keep tests focused**: Each test should verify one specific aspect of the handler
