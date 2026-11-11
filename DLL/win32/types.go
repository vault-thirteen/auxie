package win32

import bt "github.com/vault-thirteen/auxie/BasicTypes"

// typedef unsigned long DWORD; // minwindef.h:156.
type DWORD bt.DWord

// typedef void *HANDLE; // WinBase.h:712.
type HANDLE uintptr //unsafe.Pointer
