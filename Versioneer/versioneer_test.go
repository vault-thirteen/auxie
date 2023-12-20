package ver

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_ShowIntroText(t *testing.T) {
	aTest := tester.New(t)

	v, err := New()
	aTest.MustBeNoError(err)
	v.ShowIntroText("")
}
