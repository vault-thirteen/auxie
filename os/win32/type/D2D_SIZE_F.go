package t

import (
	"math"
	"unsafe"

	b "github.com/vault-thirteen/auxie/os/win32/base"
	c "github.com/vault-thirteen/auxie/os/win32/common"
)

type D2D_SIZE_F c.Size[b.FLOAT]

func (sf D2D_SIZE_F) ToUint64() (out uint64) {
	wBits := uint64(math.Float32bits(sf.Width))
	hBits := uint64(math.Float32bits(sf.Height))
	return ((hBits) << 32) | wBits
}
func (sf D2D_SIZE_F) ToUintptr() (out uintptr) {
	return *(*uintptr)(unsafe.Pointer(&sf))
}

func D2D_SIZE_F_FromUint64(in uint64) (out D2D_SIZE_F) {
	wBits := uint32(in & 0xFFFFFFFF)
	hBits := uint32(in >> 32)
	return D2D1_SIZE_F{
		Width:  math.Float32frombits(wBits),
		Height: math.Float32frombits(hBits),
	}
}
func D2D_SIZE_F_FromUintptr(in uintptr) (out D2D_SIZE_F) {
	return *(*D2D_SIZE_F)(unsafe.Pointer(&in))
}
