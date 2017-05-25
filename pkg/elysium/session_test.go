package elysium

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	// session is a global AuthSession used by all tests.
	session *AuthSession

	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(1)
	//os.Exit(m.Run())
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	session = NewAuthSession("super-secret-terminus-token")
	host, _ := url.Parse(server.URL)
	APIHost = host.String()
	mux.HandleFunc("/"+session.Path("POST"), func(w http.ResponseWriter, r *http.Request) {
		expires := time.Now().UTC().Unix() + 100000
		fmt.Fprintf(w, `{"machine_token":"super-secret-terminus-token","email":"testuser@drud.com","client":"terminus","expires_at": %d,"session":"some-testsession","user_id":"some-testuser"}`, expires)
	})
	err := session.Auth()
	if err != nil {
		log.Fatalf("Could not auth: %v", err)
	}
}

func TestAuthSession(t *testing.T) {
	assert := assert.New(t)

	SiteList := &SiteList{}
	err := session.Request("GET", SiteList)
	assert.NoError(err)
	assert.NotEmpty(SiteList.Sites)

	site := SiteList.Sites[0]
	environmentList := NewEnvironmentList(site.ID)
	err = session.Request("GET", environmentList)
	assert.NoError(err)
	assert.NotEmpty(environmentList)

	env := environmentList.Environments["live"]
	bl := NewBackupList(site.ID, env.Name)
	err = session.Request("GET", bl)
	assert.NoError(err)

	if len(bl.Backups) > 0 {
		for i, backup := range bl.Backups {
			err = session.Request("POST", &backup)
			assert.NoError(err)
			bl.Backups[i] = backup
		}
	}
}
