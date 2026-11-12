package t

import (
	"testing"
	"unsafe"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_D2D_SIZE_U(t *testing.T) {
	aTest := tester.New(t)

	x := D2D_SIZE_U{Width: 1000000001, Height: 2000000002}
	aTest.MustBeEqual(int(unsafe.Sizeof(x)), 8)
}

func Test_D2D_SIZE_U_ToUint64(t *testing.T) {
	aTest := tester.New(t)

	var su = D2D_SIZE_U{Width: 123, Height: 456}
	var su64 = su.ToUint64()
	var su64expected uint64 = uint64(456<<32) | uint64(123)
	aTest.MustBeEqual(su64, su64expected)
}
func Test_D2D_SIZE_U_ToUintptr(t *testing.T) {
	aTest := tester.New(t)

	var su = D2D_SIZE_U{Width: 123, Height: 456}
	var x = su.ToUintptr()
	var suExpected = *(*D2D_SIZE_U)(unsafe.Pointer(&x))
	aTest.MustBeEqual(su, suExpected)
	aTest.MustBeEqual(x, uintptr(su.ToUint64()))
}

func Test_D2D_SIZE_U_FromUint64(t *testing.T) {
	aTest := tester.New(t)

	var x uint64 = uint64(456<<32) | uint64(123)
	var su = D2D_SIZE_U_FromUint64(x)
	var suExpected = D2D_SIZE_U{Width: 123, Height: 456}
	aTest.MustBeEqual(su, suExpected)
}
func Test_D2D_SIZE_U_FromUintptr(t *testing.T) {
	aTest := tester.New(t)

	var su = D2D_SIZE_U{Width: 123, Height: 456}
	var x = uintptr(su.ToUintptr())
	aTest.MustBeEqual(D2D_SIZE_U_FromUintptr(x), su)
	aTest.MustBeEqual(D2D_SIZE_U_FromUint64(uint64(x)), su)
}
