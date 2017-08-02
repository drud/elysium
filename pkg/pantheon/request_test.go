package pantheon

import (
	"fmt"
	"bytes"
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestRequestHttpRequest(t *testing.T) {
	assert := assert.New(t)
	expected := "Bar"

	// Define a route on the test server.
	mux.HandleFunc("/foo/bar", func(w http.ResponseWriter, r *http.Request) {
		// Check the request for default headers.
		assert.Equal("application/json", r.Header.Get("Content-Type"))
		assert.Equal("Terminus/1.3.1-dev (php_version=7.1.5&script=bin/terminus)", r.Header.Get("User-Agent"))
		// Check the request for custom headers.
		assert.Equal(expected, r.Header.Get("Foo"))

		// Write a test response body.
		fmt.Fprint(w, expected)
	})

	// Construct a new request with custom header.
	resp, err := httpRequest("GET", "/foo/bar", nil, map[string]string{"Foo": expected})
	if err != nil {
		assert.Error(err)
	}

	// Read the byte response.
	// Then convert it to a string.
	var actual string
	raw := bytes.NewBuffer(resp)
	actual = raw.String()
	// Check for equality between the response body and expected value.
	assert.Equal(expected, actual)
}
