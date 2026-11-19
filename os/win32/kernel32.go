package win32

//go:generate mkwinsyscall -output kernel32_functions.go kernel32.go

// Functions that set last error:
/*
//sys CreateEventW(lpEventAttributes *t.SECURITY_ATTRIBUTES, bManualReset b.BOOL, bInitialState b.BOOL, lpName *b.WCHAR)(ret b.HANDLE, err error)[failretval==nil] = kernel32.CreateEventW
//sys SetEvent(hEvent b.HANDLE)(ret b.BOOL, err error)[failretval==0] = kernel32.SetEvent
*/

// Functions that do not set last error:
/*
//sys GetLastError()(lastErrorCode b.DWORD) = kernel32.GetLastError
*/

// Non-standard logics. Be very careful using these functions.
/*
//sys SetLastError(dwErrCode b.DWORD) = kernel32.SetLastError
*/
