package b

import (
	"golang.org/x/sys/windows"
)

type HWND = windows.Handle
type HHOOK = windows.Handle
type HEVENT = windows.Handle
type HGDIOBJ = windows.Handle
type HACCEL = windows.Handle
type HBITMAP = windows.Handle
type HBRUSH = windows.Handle
type HCOLORSPACE = windows.Handle
type HDC = windows.Handle
type HGLRC = windows.Handle
type HDESK = windows.Handle
type HENHMETAFILE = windows.Handle
type HFONT = windows.Handle
type HICON = windows.Handle
type HMENU = windows.Handle
type HPALETTE = windows.Handle
type HPEN = windows.Handle
type HWINEVENTHOOK = windows.Handle
type HMONITOR = windows.Handle
type HUMPD = windows.Handle

type HCURSOR = HICON
type COLORREF = DWORD
type LPCOLORREF = *DWORD
