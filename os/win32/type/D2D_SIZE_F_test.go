package t

import (
	"math"
	"testing"
	"unsafe"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_D2D_SIZE_F(t *testing.T) {
	aTest := tester.New(t)

	x := D2D_SIZE_F{Width: 11.11, Height: 22.22}
	aTest.MustBeEqual(int(unsafe.Sizeof(x)), 8)
}

func Test_D2D_SIZE_F_ToUint64(t *testing.T) {
	aTest := tester.New(t)

	var sf = D2D_SIZE_F{Width: 123.01, Height: 456.02}
	var sf64 = sf.ToUint64()
	wBits := uint64(math.Float32bits(sf.Width))
	hBits := uint64(math.Float32bits(sf.Height))
	var sf64expected uint64 = uint64(hBits<<32) | uint64(wBits)
	aTest.MustBeEqual(sf64, sf64expected)
}
func Test_D2D_SIZE_F_ToUintptr(t *testing.T) {
	aTest := tester.New(t)

	var sf = D2D_SIZE_F{Width: 123.01, Height: 456.02}
	var x = sf.ToUintptr()
	var sfExpected = *(*D2D_SIZE_F)(unsafe.Pointer(&x))
	aTest.MustBeEqual(sf, sfExpected)
	aTest.MustBeEqual(x, uintptr(sf.ToUint64()))
}

func Test_D2D_SIZE_F_FromUint64(t *testing.T) {
	aTest := tester.New(t)

	wBits := uint64(math.Float32bits(123.01))
	hBits := uint64(math.Float32bits(456.02))
	var x uint64 = uint64(hBits<<32) | uint64(wBits)
	var sf = D2D_SIZE_F_FromUint64(x)
	var sfExpected = D2D_SIZE_F{Width: 123.01, Height: 456.02}
	aTest.MustBeEqual(sf, sfExpected)
}
func Test_D2D_SIZE_F_FromUintptr(t *testing.T) {
	aTest := tester.New(t)

	var sf = D2D_SIZE_F{Width: 123.01, Height: 456.02}
	var x = uintptr(sf.ToUintptr())
	aTest.MustBeEqual(D2D_SIZE_F_FromUintptr(x), sf)
	aTest.MustBeEqual(D2D_SIZE_F_FromUint64(uint64(x)), sf)
}
