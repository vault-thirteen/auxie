package ver

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_ShowIntroText(t *testing.T) {
	aTest := tester.New(t)

	// This test will not pass locally while Go language considers local
	// project having the '(devel)' version. This is how Go language works.
	v, err := New()
	aTest.MustBeNoError(err)
	v.ShowIntroText("")
}
