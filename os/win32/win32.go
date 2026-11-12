package win32

import "unsafe"

//go:generate mkwinsyscall -output zsyscall_windows.go win32.go

// Functions that set last error:
/*
//sys CreateWindowExW(dwExStyle b.DWORD, lpClassName *uint16, lpWindowName *uint16, dwStyle b.DWORD, x int32, y int32, nWidth int32, nHeight int32, hWndParent b.HWND, hMenu b.HMENU, hInstance b.HINSTANCE, lpParam unsafe.Pointer) (ret b.HWND, err error)[failretval==0] = user32.CreateWindowExW
//sys DestroyWindow(hWnd b.HWND) (ret b.BOOL, err error)[failretval==0] = user32.DestroyWindow
//sys GetClientRect(hWnd b.HWND, lpRect *t.RECT)(ret b.BOOL, err error)[failretval==0] = user32.GetClientRect
//sys GetMessageW(lpMsg *t.MSG, hWnd b.HWND, wMsgFilterMin b.UINT, wMsgFilterMax b.UINT)(ret b.BOOL, err error)[failretval==-1] = user32.GetMessageW
//sys GetWindowLongPtrW(hWnd b.HWND, nIndex int32) (ret b.LONG_PTR, err error)[failretval==0] = user32.GetWindowLongPtrW
//sys LoadCursorW(hInstance b.HINSTANCE, lpCursorName *uint16) (ret b.HCURSOR, err error)[failretval==0] = user32.LoadCursorW
//sys RegisterClassExW(unnamedParam1 *t.WNDCLASSEXW_Golang) (ret b.ATOM, err error)[failretval==0] = user32.RegisterClassExW
//sys UnregisterClassW(lpClassName *uint16, hInstance b.HINSTANCE)(ret b.BOOL, err error)[failretval==0] = user32.UnregisterClassW
*/

// Functions that do not set last error:
/*
//sys CoCreateInstance(rclsid *windows.GUID, pUnkOuter *i.IUnknown, dwClsContext b.DWORD, riid *windows.GUID, ppv *unsafe.Pointer) (ret b.HRESULT) = ole32.CoCreateInstance
//sys D2D1CreateFactory(factoryType e.D2D1_FACTORY_TYPE, riid *b.IID, pFactoryOptions *t.D2D1_FACTORY_OPTIONS, ppIFactory **i.ID2D1Factory) (ret b.HRESULT) = d2d1.D2D1CreateFactory
//sys DefWindowProcW(hWnd b.HWND, Msg e.WM, wParam b.WPARAM, lParam b.LPARAM)(ret b.LRESULT) = user32.DefWindowProcW
//sys DispatchMessageW(lpMsg *t.MSG) (ret b.LRESULT) = user32.DispatchMessageW
//sys GetKeyState(nVirtKey int32)(ret b.SHORT) = user32.GetKeyState
//sys InvalidateRect(hWnd b.HWND, lpRect *t.RECT, bErase b.BOOL)(ret b.BOOL)[failretval==0] = user32.InvalidateRect
//sys PeekMessageW(lpMsg *t.MSG, hWnd b.HWND, wMsgFilterMin b.UINT, wMsgFilterMax b.UINT, wRemoveMsg b.UINT)(ret b.BOOL) = user32.PeekMessageW
//sys PostQuitMessage(nExitCode int32) = user32.PostQuitMessage
//sys ShowWindow(hWnd b.HWND, nCmdShow int32) (ret b.BOOL) = user32.ShowWindow
//sys TranslateMessage(lpMsg *t.MSG) (ret b.BOOL) = user32.TranslateMessage
//sys UpdateWindow(hWnd b.HWND)(ret b.BOOL) = user32.UpdateWindow
//sys ValidateRect(hWnd b.HWND, lpRect *t.RECT)(ret b.BOOL) = user32.ValidateRect
*/

// Non-standard logics. Be very careful using these functions.
/*
//sys SetLastError(dwErrCode b.DWORD) = kernel32.SetLastError
//sys SetWindowLongPtrW(hWnd b.HWND, nIndex int32, dwNewLong b.LONG_PTR)(ret b.LONG_PTR) = user32.SetWindowLongPtrW
*/

func StructArg[T any](v T) (x uintptr) {
	if unsafe.Sizeof(v) > unsafe.Sizeof(x) {
		panic("win32: by-value struct is too large to pass in a register")
	}
	*(*T)(unsafe.Pointer(&x)) = v
	return x
}
