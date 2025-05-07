# Using testify & httpexpect

While Go's standard library provides excellent tools for testing HTTP handlers, third-party packages like [testify](https://github.com/stretchr/testify) and [httpexpect](https://github.com/gavv/httpexpect) can make your tests easier to write.


## Testify

The `testify` package provides enhanced assertion functions that can make the tests more readable and provide
better error messages.

### Installation

```bash
go get github.com/stretchr/testify
go get github.com/gavv/httpexpect/v2
```

### Basic Usage


```go
{{#include ../../goexamples/example-muxpatterns/main_test.go}}
```

