package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	// Set up routes and testserver
	mux := http.NewServeMux()
	mux.HandleFunc("GET /json/{word}", handler4)

	server := httptest.NewServer(mux)
	defer server.Close()

	// Create httpexpect instance
	e := httpexpect.Default(t, server.URL)

	var target map[string]string

	e.GET("/json/hello").
		Expect().
		Status(http.StatusOK).
		JSON().Object().Decode(&target)

	assert.Equal(t, "hello", target["word"])
}
