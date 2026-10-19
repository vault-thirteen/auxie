package ver

import (
	"testing"

	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/tester"
)

const IgnoreGolangDevelVersionBug = true

func Test_ShowIntroText(t *testing.T) {
	aTest := tester.New(t)

	// This test will not pass locally while Go language considers local
	// project having the '(devel)' version. This is how Go language works.
	v, err := New(IgnoreGolangDevelVersionBug)
	aTest.MustBeNoError(err)

	v.ShowIntroText("")
}

func Test_GoVersion(t *testing.T) {
	aTest := tester.New(t)

	// This test will not pass locally while Go language considers local
	// project having the '(devel)' version. This is how Go language works.
	v, err := New(IgnoreGolangDevelVersionBug)
	aTest.MustBeNoError(err)

	cio.PrintMessage(v.GoVersion())
}
