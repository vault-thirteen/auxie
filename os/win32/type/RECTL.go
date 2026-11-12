package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type RECTL struct {
	Left   b.LONG
	Top    b.LONG
	Right  b.LONG
	Bottom b.LONG
}
type PRECTL = *RECTL
type LPRECTL = *RECTL
type LPCRECTL = *RECTL
