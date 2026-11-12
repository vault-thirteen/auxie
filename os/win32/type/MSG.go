package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type MSG struct {
	hwnd    b.HWND
	message b.UINT
	wParam  b.WPARAM
	lParam  b.LPARAM
	time    b.DWORD
	pt      POINT
	//lPrivate DWORD
}

type PMSG = *MSG
type NPMSG = *MSG
type LPMSG = *MSG
