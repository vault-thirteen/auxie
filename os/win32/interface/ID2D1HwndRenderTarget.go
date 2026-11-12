package i

import (
	"math"
	"syscall"
	"unsafe"

	b "github.com/vault-thirteen/auxie/os/win32/base"
	e "github.com/vault-thirteen/auxie/os/win32/enum"
	t "github.com/vault-thirteen/auxie/os/win32/type"
)

type ID2D1HwndRenderTarget struct {
	Vtbl *ID2D1HwndRenderTargetVtbl
}

// ID2D1HwndRenderTargetVtbl is a VTable for ID2D1HwndRenderTarget.
// All the fields were tested with C++ VTable tester.
type ID2D1HwndRenderTargetVtbl struct {
	// IUnknown methods.
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// ID2D1Resource methods.
	GetFactory uintptr

	// ID2D1RenderTarget methods.
	CreateBitmap                 uintptr
	CreateBitmapFromWicBitmap    uintptr
	CreateSharedBitmap           uintptr
	CreateBitmapBrush            uintptr
	CreateSolidColorBrush        uintptr
	CreateGradientStopCollection uintptr
	CreateLinearGradientBrush    uintptr
	CreateRadialGradientBrush    uintptr
	CreateCompatibleRenderTarget uintptr
	CreateLayer                  uintptr
	CreateMesh                   uintptr
	DrawLine                     uintptr
	DrawRectangle                uintptr
	FillRectangle                uintptr
	DrawRoundedRectangle         uintptr
	FillRoundedRectangle         uintptr
	DrawEllipse                  uintptr
	FillEllipse                  uintptr
	DrawGeometry                 uintptr
	FillGeometry                 uintptr
	FillMesh                     uintptr
	FillOpacityMask              uintptr
	DrawBitmap                   uintptr
	DrawTextW                    uintptr
	DrawTextLayout               uintptr
	DrawGlyphRun                 uintptr
	SetTransform                 uintptr
	GetTransform                 uintptr
	SetAntialiasMode             uintptr
	GetAntialiasMode             uintptr
	SetTextAntialiasMode         uintptr
	GetTextAntialiasMode         uintptr
	SetTextRenderingParams       uintptr
	GetTextRenderingParams       uintptr
	SetTags                      uintptr
	GetTags                      uintptr
	PushLayer                    uintptr
	PopLayer                     uintptr
	Flush                        uintptr
	SaveDrawingState             uintptr
	RestoreDrawingState          uintptr
	PushAxisAlignedClip          uintptr
	PopAxisAlignedClip           uintptr
	Clear                        uintptr
	BeginDraw                    uintptr
	EndDraw                      uintptr
	GetPixelFormat               uintptr
	SetDpi                       uintptr
	GetDpi                       uintptr
	GetSize                      uintptr
	GetPixelSize                 uintptr
	GetMaximumBitmapSize         uintptr
	IsSupported                  uintptr

	// ID2D1HwndRenderTarget methods.
	CheckWindowState uintptr
	Resize           uintptr
	GetHwnd          uintptr
}

func (wrt *ID2D1HwndRenderTarget) Release() b.ULONG {
	r0, _, _ := syscall.SyscallN(wrt.Vtbl.Release, uintptr(unsafe.Pointer(wrt)))
	return b.ULONG(r0)
}
func (wrt *ID2D1HwndRenderTarget) CreateBitmap(
	size t.D2D1_SIZE_U,
	srcData []byte,
	pitch b.UINT32,
	bitmapProperties *t.D2D1_BITMAP_PROPERTIES,
	bitmap **ID2D1Bitmap,
) (hr b.HRESULT, err error) {

	r0, _, _ := syscall.SyscallN(
		wrt.Vtbl.CreateBitmap,
		uintptr(unsafe.Pointer(wrt)),
		uintptr(size.ToUintptr()),
		uintptr(unsafe.Pointer(unsafe.SliceData(srcData))),
		uintptr(pitch),
		uintptr(unsafe.Pointer(bitmapProperties)),
		uintptr(unsafe.Pointer(bitmap)),
	)
	hr = b.HRESULT(r0)
	if hr != b.S_OK {
		err = syscall.Errno(hr)
	}
	return
}
func (wrt *ID2D1HwndRenderTarget) DrawBitmap(
	bitmap *ID2D1Bitmap,
	destinationRectangle *t.D2D1_RECT_F,
	opacity b.FLOAT,
	interpolationMode e.D2D1_BITMAP_INTERPOLATION_MODE,
	sourceRectangle *t.D2D1_RECT_F,
) {

	_, _, _ = syscall.SyscallN(
		wrt.Vtbl.DrawBitmap,
		uintptr(unsafe.Pointer(wrt)),
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(destinationRectangle)),
		uintptr(math.Float32bits(opacity)),
		uintptr(interpolationMode),
		uintptr(unsafe.Pointer(sourceRectangle)),
	)
}
func (wrt *ID2D1HwndRenderTarget) SetTransform(transform *t.D2D1_MATRIX_3X2_F) {
	_, _, _ = syscall.SyscallN(wrt.Vtbl.SetTransform, uintptr(unsafe.Pointer(wrt)), uintptr(unsafe.Pointer(transform)))
}
func (wrt *ID2D1HwndRenderTarget) Clear(clearColor *t.D2D1_COLOR_F) {
	_, _, _ = syscall.SyscallN(wrt.Vtbl.Clear, uintptr(unsafe.Pointer(wrt)), uintptr(unsafe.Pointer(clearColor)))
}
func (wrt *ID2D1HwndRenderTarget) BeginDraw() {
	_, _, _ = syscall.SyscallN(wrt.Vtbl.BeginDraw, uintptr(unsafe.Pointer(wrt)))
}
func (wrt *ID2D1HwndRenderTarget) EndDraw() (hr b.HRESULT, err error) {
	// EndDraw returns an HRESULT code.
	// tagSize1, tagSize2 can be zero pointers if you don't trace errors to specific shapes
	//
	// If hr is D2DERR_RECREATE_TARGET, then caller must destroy and re-create
	// the render target.
	r0, _, _ := syscall.SyscallN(wrt.Vtbl.EndDraw, uintptr(unsafe.Pointer(wrt)),
		0, // D2D1_TAG* tag1
		0, // D2D1_TAG* tag2
	)
	hr = b.HRESULT(r0)
	if hr != b.S_OK {
		err = syscall.Errno(hr)
	}
	return
}
func (wrt *ID2D1HwndRenderTarget) GetMaximumBitmapSize() (bs b.UINT32) {
	r0, _, _ := syscall.SyscallN(wrt.Vtbl.GetMaximumBitmapSize, uintptr(unsafe.Pointer(wrt)))
	return b.UINT32(r0)
}
func (wrt *ID2D1HwndRenderTarget) CheckWindowState() (ws e.D2D1_WINDOW_STATE) {
	r0, _, _ := syscall.SyscallN(wrt.Vtbl.CheckWindowState, uintptr(unsafe.Pointer(wrt)))
	return e.D2D1_WINDOW_STATE(r0)
}
func (wrt *ID2D1HwndRenderTarget) Resize(size *t.D2D1_SIZE_U) (hr b.HRESULT, err error) {
	r0, _, _ := syscall.SyscallN(wrt.Vtbl.Resize, uintptr(unsafe.Pointer(wrt)), uintptr(unsafe.Pointer(size)))
	hr = b.HRESULT(r0)
	if hr != b.S_OK {
		err = syscall.Errno(hr)
	}
	return
}
