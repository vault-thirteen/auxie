package c

import (
	"testing"
	"unsafe"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_Size(t *testing.T) {
	aTest := tester.New(t)

	x := Size[uint8]{Width: 101, Height: 202}
	aTest.MustBeEqual(int(unsafe.Sizeof(x)), 2)

	y := Size[uint16]{Width: 1001, Height: 2002}
	aTest.MustBeEqual(int(unsafe.Sizeof(y)), 4)

	z := Size[uint32]{Width: 1000000001, Height: 2000000002}
	aTest.MustBeEqual(int(unsafe.Sizeof(z)), 8)

	b := Size[float32]{Width: 1.1, Height: 2.2}
	aTest.MustBeEqual(int(unsafe.Sizeof(b)), 8)
}
