package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type M128A struct {
	Low  b.ULONGLONG
	High b.LONGLONG
}
type PM128A = *M128A
