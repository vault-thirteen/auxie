package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type RECT struct {
	Left   b.LONG
	Top    b.LONG
	Right  b.LONG
	Bottom b.LONG
}
type PRECT = *RECT
type NPRECT = *RECT
type LPRECT = *RECT
type LPCRECT = *RECT
