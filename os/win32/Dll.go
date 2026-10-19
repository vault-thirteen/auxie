package win32

import "golang.org/x/sys/windows"

const (
	DllFile_Kernel32 = "kernel32.dll"
	DllFile_User32   = "user32.dll"
)

type Dll struct {
	handle      windows.Handle
	funcMapping []FuncMapping
}

func NewDll(fm []FuncMapping) Dll {
	return Dll{
		handle:      0,
		funcMapping: fm,
	}
}
