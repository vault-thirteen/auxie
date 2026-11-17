package t

import (
	"testing"
	"unsafe"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_INPUT(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(int(unsafe.Sizeof(MouseInputStruct{})), 40)
	aTest.MustBeEqual(int(unsafe.Sizeof(KeyboardInputStruct{})), 40)
	aTest.MustBeEqual(int(unsafe.Sizeof(HardwareInputStruct{})), 40)
}
