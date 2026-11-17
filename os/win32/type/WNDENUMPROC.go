package t

import b "github.com/vault-thirteen/auxie/os/win32/base"

// typedef BOOL (CALLBACK* WNDENUMPROC)(HWND, LPARAM);
type WNDENUMPROC = func(hwnd b.HWND, lParam b.LPARAM) b.BOOL
type WNDENUMPROC_ForGolang = func(hwnd b.HWND, lParam b.LPARAM) uintptr

func Convert_WNDENUMPROC_ToGolang(inFn WNDENUMPROC) (outFn WNDENUMPROC_ForGolang) {
	outFn = func(hwnd b.HWND, lParam b.LPARAM) uintptr {
		return uintptr(inFn(hwnd, lParam))
	}
	return outFn
}
