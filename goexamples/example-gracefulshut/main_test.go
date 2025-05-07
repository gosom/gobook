package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerEndpoints(t *testing.T) {
	// Start the test server
	ts := setupTestServer(t)
	defer ts.Close()

	// Test the root endpoint
	resp, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.Status)
	}

	// Read and check the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "Hello world" {
		t.Errorf("Expected 'Hello world'; got %q", string(body))
	}
}

func setupTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	// Set up your router with all handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	// Create and return the test server
	return httptest.NewServer(mux)
}
