package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type SIZE struct {
	X b.LONG
	Y b.LONG
}
type PSIZE = *SIZE
type LPSIZE = *SIZE

type SIZEL = SIZE
type PSIZEL = *SIZE
type LPSIZEL = *SIZE
