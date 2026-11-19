package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type MSG struct {
	Hwnd    b.HWND
	Message b.UINT
	WParam  b.WPARAM
	LParam  b.LPARAM
	Time    b.DWORD
	Pt      POINT
	//lPrivate DWORD
}

type PMSG = *MSG
type NPMSG = *MSG
type LPMSG = *MSG
