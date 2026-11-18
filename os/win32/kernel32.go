package win32

//go:generate mkwinsyscall -output kernel32_functions.go kernel32.go

// Functions that set last error:
/*
 */

// Functions that do not set last error:
/*
//sys GetLastError()(lastErrorCode b.DWORD) = kernel32.GetLastError
*/

// Non-standard logics. Be very careful using these functions.
/*
//sys SetLastError(dwErrCode b.DWORD) = kernel32.SetLastError
*/
