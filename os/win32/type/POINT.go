package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type POINT struct {
	X b.LONG
	Y b.LONG
}
type PPOINT = *POINT
type NPPOINT = *POINT
type LPPOINT = *POINT
