package httphelper

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/vault-thirteen/auxie/errors"
	"github.com/vault-thirteen/auxie/tester"
)

const (
	Err_Text = "error-x"
)

func Test_logErrorIfSet(t *testing.T) {
	aTest := tester.New(t)
	var err error
	var buf bytes.Buffer

	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	// Test.
	err = errors.ErrorFromText(Err_Text)
	logErrorIfSet(err)
	// Check.
	loggedMsg := strings.TrimSpace(buf.String())
	loggedMsgParts := strings.Split(loggedMsg, " ") // Error has no space.
	aTest.MustBeEqual(len(loggedMsgParts), 3)       // Three fields: date, time, message.
	aTest.MustBeEqual(loggedMsgParts[2], Err_Text)
}
