package win32

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// GetLastError
// _Post_equals_last_error_ DWORD GetLastError();
// https://learn.microsoft.com/en-us/windows/win32/api/errhandlingapi/nf-errhandlingapi-getlasterror
func (dc *DllController) GetLastError() (err DWORD) {
	ret, _, callErr := syscall.SyscallN(dc.DllFunc.GetLastError)
	MustBeNoCallError(callErr)
	return DWORD(ret)
}

// SetLastError
// void SetLastError( [in] DWORD dwErrCode );
// https://learn.microsoft.com/en-us/windows/win32/api/errhandlingapi/nf-errhandlingapi-setlasterror
func (dc *DllController) SetLastError(err DWORD) {
	_, _, callErr := syscall.SyscallN(dc.DllFunc.SetLastError, uintptr(err))
	MustBeNoCallError(callErr)
}

// WINBASEAPI HANDLE WINAPI GetStdHandle(_In_ DWORD nStdHandle);
func (dc *DllController) GetStdHandle(handle DWORD) (h windows.Handle) {
	ret, _, callErr := syscall.SyscallN(dc.DllFunc.GetStdHandle, uintptr(handle))
	MustBeNoCallError(callErr)
	return windows.Handle(ret)
}

// WINBASEAPI BOOL WINAPI GetConsoleMode(_In_ HANDLE hConsoleHandle, _Out_ LPDWORD lpMode);
func (dc *DllController) GetConsoleMode(consoleHandle windows.Handle, lpMode *DWORD) bool {
	ret, _, callErr := syscall.SyscallN(dc.DllFunc.GetConsoleMode, uintptr(consoleHandle), uintptr(unsafe.Pointer(lpMode)))
	MustBeNoCallError(callErr)
	return ret != 0
}

// WINBASEAPI BOOL WINAPI SetConsoleMode(_In_ HANDLE hConsoleHandle, _In_ DWORD dwMode);
func (dc *DllController) SetConsoleMode(consoleHandle windows.Handle, dwMode DWORD) bool {
	ret, _, callErr := syscall.SyscallN(dc.DllFunc.SetConsoleMode, uintptr(consoleHandle), uintptr(dwMode))
	MustBeNoCallError(callErr)
	return ret != 0
}
