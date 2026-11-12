package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type CREATESTRUCTW struct {
	LpCreateParams b.LPVOID
	HInstance      b.HINSTANCE
	HMenu          b.HMENU
	HwndParent     b.HWND
	Cy             int32
	Cx             int32
	Y              int32
	X              int32
	Style          b.LONG
	LpszName       b.LPCWSTR
	LpszClass      b.LPCWSTR
	DwExStyle      b.DWORD
}

type LPCREATESTRUCTW = *CREATESTRUCTW
type LPCREATESTRUCT = LPCREATESTRUCTW
