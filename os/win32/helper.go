package win32

import (
	"errors"
	"runtime"

	"golang.org/x/sys/windows"

	"github.com/vault-thirteen/auxie/cio"
	ae "github.com/vault-thirteen/auxie/errors"
)

const (
	Msg_GetStdHandleReturnedInvalidHandle = "GetStdHandle returned invalid handle"
	Msg_GetConsoleModeFailed              = "GetConsoleMode failed"
	Msg_SetConsoleModeFailed              = "SetConsoleMode failed"

	Msgf_SyscallFailed = "%s syscall failed: %v"
)

const (
	INVALID_HANDLE_VALUE = windows.InvalidHandle
)

// MustBeNoCallError checks for a call error.
// If call error is set, it finds a name of the function which started this
// function and panics about that function's failure.
func MustBeNoCallError(callErr windows.Errno) {
	if !errors.Is(callErr, windows.ERROR_SUCCESS) {
		caller, _, _, _ := runtime.Caller(1) // Skip 1 function up in stack: mustBeNoCallError -> caller.
		fn := runtime.FuncForPC(caller)
		panic(cio.FormatMessage(Msgf_SyscallFailed, fn.Name(), callErr.Error()))
	}
}

func (dc *DllController) EnableConsoleColours() (err error) {
	var hConsole windows.Handle
	hConsole = dc.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if hConsole == INVALID_HANDLE_VALUE {
		return ae.ErrorFromText(Msg_GetStdHandleReturnedInvalidHandle)
	}

	var dwMode DWORD
	if !dc.GetConsoleMode(hConsole, &dwMode) {
		return ae.ErrorFromText(Msg_GetConsoleModeFailed)
	}

	dwMode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING

	if !dc.SetConsoleMode(hConsole, dwMode) {
		return ae.ErrorFromText(Msg_SetConsoleModeFailed)
	}

	return nil
}
