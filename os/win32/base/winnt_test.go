package b

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_hresult_overflow(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(hresult_overflow(0), HRESULT(0))
	aTest.MustBeEqual(D2DERR_RECREATE_TARGET, HRESULT(-2003238900))
	aTest.MustBeEqual(CONTEXT_E_FIRST, HRESULT(0x8004E000-0x100000000))
	aTest.MustBeEqual(CONTEXT_E_LAST, HRESULT(0x8004E02F-0x100000000))
}
