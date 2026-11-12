package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type POINTL struct {
	X b.LONG
	Y b.LONG
}
type PPOINTL = *POINTL
