package pantheon

import (
	"fmt"
	"bytes"
	"io/ioutil"
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
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			assert.Error(err)
		}

		// Check that the request body is carrying our value.
		actual := string(body)
		assert.Equal("Bar", actual)
		// Write a response body with the value extracted from the request body.
		fmt.Fprint(w, actual)
	})

	// Construct a new request with custom header.
	resp, err := httpRequest("GET", "/foo/bar", []byte("Bar"), map[string]string{"Foo": expected})
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