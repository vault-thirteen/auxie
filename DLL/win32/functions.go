package dll

import (
	"syscall"
	"unsafe"

	b "github.com/vault-thirteen/auxie/os/win32/base"
)

func errnoErr(e syscall.Errno) error {
	return e
}

// GetLastError
// _Post_equals_last_error_ DWORD GetLastError();
// https://learn.microsoft.com/en-us/windows/win32/api/errhandlingapi/nf-errhandlingapi-getlasterror
func (dc *DllController) GetLastError() (ret b.DWORD) {
	r0, _, _ := syscall.SyscallN(dc.DllFunc.GetLastError)
	ret = b.DWORD(r0)
	return
}

// SetLastError
// void SetLastError( [in] DWORD dwErrCode );
// https://learn.microsoft.com/en-us/windows/win32/api/errhandlingapi/nf-errhandlingapi-setlasterror
func (dc *DllController) SetLastError(error b.DWORD) {
	_, _, _ = syscall.SyscallN(dc.DllFunc.SetLastError, uintptr(error))
}

// GetStdHandle
// WINBASEAPI HANDLE WINAPI GetStdHandle(_In_ DWORD nStdHandle);
func (dc *DllController) GetStdHandle(handle b.DWORD) (ret b.HANDLE, err error) {
	r0, _, e1 := syscall.SyscallN(dc.DllFunc.GetStdHandle, uintptr(handle))

	// 1. Compare the raw value.
	// Go language can destroy raw pointers.
	if r0 == b.INVALID_HANDLE_VALUE {
		err = errnoErr(e1)
	}

	// 2. Return the pointer directly to avoid garbage collection.
	return b.HANDLE(unsafe.Add(nil, r0)), nil
}

// GetConsoleMode
// WINBASEAPI BOOL WINAPI GetConsoleMode(_In_ HANDLE hConsoleHandle, _Out_ LPDWORD lpMode);
func (dc *DllController) GetConsoleMode(consoleHandle b.HANDLE, lpMode *b.DWORD) (ret b.BOOL, err error) {
	r0, _, e1 := syscall.SyscallN(dc.DllFunc.GetConsoleMode, uintptr(consoleHandle), uintptr(unsafe.Pointer(lpMode)))
	ret = b.BOOL(r0)
	if ret == 0 {
		err = errnoErr(e1)
	}
	return
}

// SetConsoleMode
// WINBASEAPI BOOL WINAPI SetConsoleMode(_In_ HANDLE hConsoleHandle, _In_ DWORD dwMode);
func (dc *DllController) SetConsoleMode(consoleHandle b.HANDLE, dwMode b.DWORD) (ret b.BOOL, err error) {
	r0, _, e1 := syscall.SyscallN(dc.DllFunc.SetConsoleMode, uintptr(consoleHandle), uintptr(dwMode))
	ret = b.BOOL(r0)
	if ret == 0 {
		err = errnoErr(e1)
	}
	return
}
