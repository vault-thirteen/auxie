package b

import (
	"syscall"
	"unsafe"
)

type ULONG = uint32
type PULONG = *ULONG
type USHORT = uint16
type PUSHORT = *USHORT
type UCHAR = uint8
type PUCHAR = *UCHAR

type DWORD = uint32
type BOOL = int32
type BYTE = uint8
type WORD = uint16
type FLOAT = float32
type PFLOAT = *FLOAT

/*
typedef BOOL near           *PBOOL;
typedef BOOL far            *LPBOOL;
typedef BYTE near           *PBYTE;
typedef BYTE far            *LPBYTE;
typedef int near            *PINT;
typedef int far             *LPINT;
typedef WORD near           *PWORD;
typedef WORD far            *LPWORD;
typedef long far            *LPLONG;
typedef DWORD near          *PDWORD;
typedef DWORD far           *LPDWORD;
typedef void far            *LPVOID;
typedef CONST void far      *LPCVOID;
*/
// Go language does not support near and far pointers.
type PBOOL = *BOOL
type LPBOOL = *BOOL
type PBYTE = *BYTE
type LPBYTE = *BYTE
type PINT = *int32
type LPINT = *int32
type PWORD = *WORD
type LPWORD = *WORD
type LPLONG = *int32
type PDWORD = *DWORD
type LPDWORD = *DWORD
type LPVOID = unsafe.Pointer
type LPCVOID = unsafe.Pointer

type INT = int32
type UINT = uint32
type PUINT = *UINT

type WPARAM = UINT_PTR
type LPARAM = LONG_PTR
type LRESULT = LONG_PTR

type SPHANDLE = *syscall.Handle
type LPHANDLE = *syscall.Handle

type HGLOBAL = syscall.Handle
type HLOCAL = syscall.Handle
type GLOBALHANDLE = syscall.Handle
type LOCALHANDLE = syscall.Handle

type ATOM = WORD

type HKEY = syscall.Handle
type PHKEY = *HKEY
type HMETAFILE = syscall.Handle
type HINSTANCE = syscall.Handle
type HMODULE = HINSTANCE
type HRGN = syscall.Handle
type HRSRC = syscall.Handle
type HSPRITE = syscall.Handle
type HLSURF = syscall.Handle
type HSTR = syscall.Handle
type HTASK = syscall.Handle
type HWINSTA = syscall.Handle
type HKL = syscall.Handle

type HFILE = int32

func LOWORD(l uintptr) WORD {
	return WORD(l & 0xffff)
}
func HIWORD(l uintptr) WORD {
	return WORD((l >> 16) & 0xffff)
}
