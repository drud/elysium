package elysium

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

var session *AuthSession
var sessionFilePath string

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	session = NewAuthSession(os.Getenv("DRUD_TERMINUS_TOKEN"))
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("could not determine home directory")
	}

	// Try to read from a saved session, if we can.
	sessionFilePath = filepath.Join(home, ".ddev", "pantheonapi")
	err = session.Read(sessionFilePath)

	if err != nil {
		// If we can't load a session, try to auth directly.
		err = session.Auth()
		if err != nil {
			log.Fatalf("Could not authenticate: %v", err)
		}

		session.Write(sessionFilePath)
	}
}

func TestAuthSession(t *testing.T) {
	assert := assert.New(t)

	req := Request{
		Auth: session,
	}

	SiteList := &SiteList{}
	err := req.Do("GET", SiteList)
	assert.NoError(err)
}
