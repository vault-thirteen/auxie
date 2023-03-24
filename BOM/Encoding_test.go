package bom

import (
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_PossibleEncodings(t *testing.T) {
	tst := tester.New(t)
	tst.MustBeEqual(len(PossibleEncodings()), 11)
}
