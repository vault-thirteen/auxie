package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type FILETIME struct {
	DwLowDateTime  b.DWORD
	DwHighDateTime b.DWORD
}

type PFILETIME = *FILETIME
type LPFILETIME = *FILETIME
