package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	_, path, _ := getParams()

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/hello")
		// Send response to be tested
		// rw.Write([]byte(`GET Hello!`))
		fmt.Fprintf(rw, "GET Hello!")
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	body, _ := get(server.URL, path)

	assert.Equal(t, "GET Hello!", body)
}
