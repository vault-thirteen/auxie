package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
	"golang.org/x/sys/windows"
)

type WNDCLASSEXW struct {
	Size       b.UINT      // cbSize
	Style      b.UINT      // style
	WndProc    WNDPROC     // lpfnWndProc
	ClsExtra   int32       // cbClsExtra
	WndExtra   int32       // cbWndExtra
	Instance   b.HINSTANCE // hInstance
	Icon       b.HICON     // hIcon
	Cursor     b.HCURSOR   // hCursor
	Background b.HBRUSH    // hbrBackground
	MenuName   b.LPCWSTR   // lpszMenuName
	ClassName  b.LPCWSTR   // lpszClassName
	IconSm     b.HICON     // hIconSm
}

type WNDCLASSEXW_Golang struct {
	Size       b.UINT      // cbSize
	Style      b.UINT      // style
	WndProc    uintptr     // lpfnWndProc wrapped over 'windows.NewCallback'
	ClsExtra   int32       // cbClsExtra
	WndExtra   int32       // cbWndExtra
	Instance   b.HINSTANCE // hInstance
	Icon       b.HICON     // hIcon
	Cursor     b.HCURSOR   // hCursor
	Background b.HBRUSH    // hbrBackground
	MenuName   b.LPCWSTR   // lpszMenuName
	ClassName  b.LPCWSTR   // lpszClassName
	IconSm     b.HICON     // hIconSm
}

func WNDCLASSEXW_ForGolang(wc WNDCLASSEXW) (wcg WNDCLASSEXW_Golang) {
	return WNDCLASSEXW_Golang{
		Size:       wc.Size,
		Style:      wc.Style,
		WndProc:    windows.NewCallback(wc.WndProc),
		ClsExtra:   wc.ClsExtra,
		WndExtra:   wc.WndExtra,
		Instance:   wc.Instance,
		Icon:       wc.Icon,
		Cursor:     wc.Cursor,
		Background: wc.Background,
		MenuName:   wc.MenuName,
		ClassName:  wc.ClassName,
		IconSm:     wc.IconSm,
	}
}

type PWNDCLASSEXW = *WNDCLASSEXW
type NPWNDCLASSEXW = *WNDCLASSEXW
type LPWNDCLASSEXW = *WNDCLASSEXW

type WNDCLASSEX = WNDCLASSEXW
type PWNDCLASSEX = PWNDCLASSEXW
type NPWNDCLASSEX = NPWNDCLASSEXW
type LPWNDCLASSEX = LPWNDCLASSEXW
