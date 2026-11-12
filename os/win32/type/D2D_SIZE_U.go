package t

import (
	"unsafe"

	b "github.com/vault-thirteen/auxie/os/win32/base"
	c "github.com/vault-thirteen/auxie/os/win32/common"
)

type D2D_SIZE_U c.Size[b.UINT32]

func (su D2D_SIZE_U) ToUint64() (out uint64) {
	return (uint64(su.Height) << 32) | uint64(su.Width)
}
func (su D2D_SIZE_U) ToUintptr() (out uintptr) {
	return *(*uintptr)(unsafe.Pointer(&su))
}

func D2D_SIZE_U_FromUint64(in uint64) (out D2D_SIZE_U) {
	return D2D1_SIZE_U{
		Width:  uint32(in & 0xFFFFFFFF),
		Height: uint32(in >> 32),
	}
}
func D2D_SIZE_U_FromUintptr(in uintptr) (out D2D_SIZE_U) {
	return *(*D2D_SIZE_U)(unsafe.Pointer(&in))
}
