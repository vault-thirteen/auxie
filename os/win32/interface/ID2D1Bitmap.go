package i

import (
	"syscall"
	"unsafe"

	b "github.com/vault-thirteen/auxie/os/win32/base"
	t "github.com/vault-thirteen/auxie/os/win32/type"
)

type ID2D1Bitmap struct {
	Vtbl *ID2D1BitmapVtbl
}

// ID2D1BitmapVtbl is a VTable for ID2D1Bitmap.
// All the fields were tested with C++ VTable tester.
type ID2D1BitmapVtbl struct {
	// IUnknown methods
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// ID2D1Resource methods
	GetFactory uintptr

	// ID2D1Image methods (if compiling for newer Direct2D versions, otherwise flat)
	// ID2D1Bitmap methods
	GetSize              uintptr
	GetPixelSize         uintptr
	GetPixelFormat       uintptr
	GetDpi               uintptr
	CopyFromBitmap       uintptr
	CopyFromRenderTarget uintptr
	CopyFromMemory       uintptr
}

func (bm *ID2D1Bitmap) Release() b.ULONG {
	ret, _, _ := syscall.SyscallN(bm.Vtbl.Release, uintptr(unsafe.Pointer(bm)))
	return b.ULONG(ret)
}
func (bm *ID2D1Bitmap) GetSize() (sf t.D2D1_SIZE_F) {
	_, _, _ = syscall.SyscallN(bm.Vtbl.GetSize, uintptr(unsafe.Pointer(bm)), uintptr(unsafe.Pointer(&sf))) //!
	return sf
}
func (bm *ID2D1Bitmap) GetPixelSize() (su t.D2D1_SIZE_U) {
	_, _, _ = syscall.SyscallN(bm.Vtbl.GetPixelSize, uintptr(unsafe.Pointer(bm)), uintptr(unsafe.Pointer(&su))) //!
	return su
}
func (bm *ID2D1Bitmap) CopyFromMemory(dstRect *t.D2D1_RECT_U, srcData []byte, pitch b.UINT32) (hr b.HRESULT, err error) {
	r0, _, _ := syscall.SyscallN(bm.Vtbl.CopyFromMemory, uintptr(unsafe.Pointer(bm)), uintptr(unsafe.Pointer(dstRect)), uintptr(unsafe.Pointer(unsafe.SliceData(srcData))), uintptr(pitch))
	hr = b.HRESULT(r0)
	if hr != b.S_OK {
		err = syscall.Errno(hr)
	}
	return
}
