package elysium

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidEnvironmentList(t *testing.T) {
	assert := assert.New(t)
	el := NewEnvironmentList("invalid-site-id")

	mux.HandleFunc(el.Path("GET", *session), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		assert.Contains(r.Header.Get("Authorization"), session.Session)

		http.Error(w, "Processing Failed", http.StatusInternalServerError)
	})

	err := session.Request("GET", el)
	assert.Error(err)
}

// TestEnvironmentList ensures EnvironmentLists can be retrieved as expected.
func TestEnvironmentList(t *testing.T) {
	assert := assert.New(t)
	el := NewEnvironmentList("some-site-id")
	mux.HandleFunc(el.Path("GET", *session), func(w http.ResponseWriter, r *http.Request) {
		// Ensure a HTTP GET request was made with the proper authorization headers.
		testMethod(t, r, "GET")
		assert.Contains(r.Header.Get("Authorization"), session.Session)

		// Send JSON response back.
		contents, err := ioutil.ReadFile("test/sites.json")
		assert.NoError(err)
		w.Write(contents)
	})

	session.Request("GET", el)

	// Ensure we got a valid response and were able to unmarshal it as expected.
	assert.Equal(len(el.Environments), 2)
}
