package cio

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

// Go language does not support terminal usage in tests, e.g. 'term.IsTerminal'
// returns false. Haha ! Great language (false) !

func Test_FormatMessage(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(FormatMessage("%s:%s", "a", "b"), "a:b")
}
