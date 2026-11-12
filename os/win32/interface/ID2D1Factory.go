package i

import (
	"syscall"
	"unsafe"

	b "github.com/vault-thirteen/auxie/os/win32/base"
	t "github.com/vault-thirteen/auxie/os/win32/type"
)

type ID2D1Factory struct {
	Vtbl *ID2D1FactoryVtbl
}

// ID2D1FactoryVtbl is a VTable for ID2D1Factory.
// All the fields were tested with C++ VTable tester.
type ID2D1FactoryVtbl struct {
	// IUnknown methods
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// ID2D1Factory methods
	ReloadSystemMetrics            uintptr
	GetDesktopDpi                  uintptr
	CreateRectangleGeometry        uintptr
	CreateRoundedRectangleGeometry uintptr
	CreateEllipseGeometry          uintptr
	CreateGeometryGroup            uintptr
	CreateTransformedGeometry      uintptr
	CreatePathGeometry             uintptr
	CreateStrokeStyle              uintptr
	CreateDrawingStateBlock        uintptr
	CreateWicBitmapRenderTarget    uintptr
	CreateHwndRenderTarget         uintptr
	CreateDxgiSurfaceRenderTarget  uintptr
	CreateDCRenderTarget           uintptr
}

func (f *ID2D1Factory) Release() b.ULONG {
	r0, _, _ := syscall.SyscallN(f.Vtbl.Release, uintptr(unsafe.Pointer(f)))
	return b.ULONG(r0)
}
func (f *ID2D1Factory) CreateHwndRenderTarget(
	renderTargetProperties *t.D2D1_RENDER_TARGET_PROPERTIES,
	hwndRenderTargetProperties *t.D2D1_HWND_RENDER_TARGET_PROPERTIES,
	hwndRenderTarget **ID2D1HwndRenderTarget,
) (hr b.HRESULT, err error) {

	r0, _, _ := syscall.SyscallN(
		f.Vtbl.CreateHwndRenderTarget,
		uintptr(unsafe.Pointer(f)),
		uintptr(unsafe.Pointer(renderTargetProperties)),
		uintptr(unsafe.Pointer(hwndRenderTargetProperties)),
		uintptr(unsafe.Pointer(hwndRenderTarget)),
	)

	hr = b.HRESULT(r0)
	if hr != b.S_OK {
		err = syscall.Errno(hr)
	}
	return
}
