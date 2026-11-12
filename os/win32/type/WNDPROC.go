package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
	e "github.com/vault-thirteen/auxie/os/win32/enum"
)

// WNDPROC means Window Procedure.
// typedef LRESULT (CALLBACK* WNDPROC)(HWND, UINT, WPARAM, LPARAM);
type WNDPROC = func(hwnd b.HWND, msg e.WM, wParam b.WPARAM, lParam b.LPARAM) b.LRESULT
