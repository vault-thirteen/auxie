package i

import (
	"syscall"
	"unsafe"

	b "github.com/vault-thirteen/auxie/os/win32/base"
	c "github.com/vault-thirteen/auxie/os/win32/common"
)

var IID_IUnknown = b.IID{
	Data1: 0x00000000,
	Data2: 0x0000,
	Data3: 0x0000,
	Data4: [8]byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46},
}

type IUnknown struct {
	*c.VTable
}

type LPUNKNOWN = *IUnknown

func (x *IUnknown) QueryInterface(riid *b.IID, ppvObject *unsafe.Pointer) (hr b.HRESULT, err error) {
	r0, _, _ := syscall.SyscallN(x.VTable[0], uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(riid)), uintptr(unsafe.Pointer(ppvObject)))
	hr = b.HRESULT(r0)
	if hr.FAILED() {
		err = syscall.Errno(hr)
	}
	return
}

func (x *IUnknown) AddRef() b.ULONG {
	r0, _, _ := syscall.SyscallN(x.VTable[1], uintptr(unsafe.Pointer(x)))
	return b.ULONG(r0)
}

func (x *IUnknown) Release() b.ULONG {
	r0, _, _ := syscall.SyscallN(x.VTable[2], uintptr(unsafe.Pointer(x)))
	return b.ULONG(r0)
}
