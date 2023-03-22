package bom

import (
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_PossibleEncodings(t *testing.T) {
	tst := tester.New(t)
	tst.MustBeEqual(len(PossibleEncodings()), 11)
}

func Test_IsKnown(t *testing.T) {
	tst := tester.New(t)
	tst.MustBeEqual(EncodingUTF8.IsKnown(), true)
	tst.MustBeEqual(EncodingUnknown.IsKnown(), false)
}

func Test_IsUnknown(t *testing.T) {
	tst := tester.New(t)
	tst.MustBeEqual(EncodingUnknown.IsUnknown(), true)
	tst.MustBeEqual(EncodingUTF8.IsUnknown(), false)
}
