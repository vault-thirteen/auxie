package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
	"golang.org/x/sys/windows"
)

type OBJECTID struct { // size is 20
	Lineage    windows.GUID
	Uniquifier b.DWORD
}
