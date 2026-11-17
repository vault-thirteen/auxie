package t

import (
	"unsafe"

	b "github.com/vault-thirteen/auxie/os/win32/base"
)

/*
typedef struct tagINPUT {
	DWORD   type;

	union
	{
		MOUSEINPUT      mi;
		KEYBDINPUT      ki;
		HARDWAREINPUT   hi;
	} DUMMYUNIONNAME;
} INPUT, *PINPUT, FAR* LPINPUT;
*/

// INPUT type can not be implemented in Go language directly while it is using
// C union feature which is not portable.

// LPINPUT is a pointer to INPUT.
type LPINPUT = unsafe.Pointer

const (
	INPUT_MOUSE    = 0
	INPUT_KEYBOARD = 1
	INPUT_HARDWARE = 2
)

type MouseInputStruct struct {
	Type uint32
	Mi   MOUSEINPUT
}
type MOUSEINPUT struct {
	Dx          b.LONG
	Dy          b.LONG
	MouseData   b.DWORD
	DwFlags     b.DWORD
	Time        b.DWORD
	DwExtraInfo b.ULONG_PTR
}

type KeyboardInputStruct struct {
	Type uint32
	Ki   KEYBDINPUT
	_    [8]byte
}

// KEYBDINPUT structure (16 bytes on 64-bit)
type KEYBDINPUT struct {
	WVk         b.WORD
	WScan       b.WORD
	DwFlags     b.DWORD
	Time        b.DWORD
	DwExtraInfo b.ULONG_PTR
}

type HardwareInputStruct struct {
	Type uint32
	Hi   HARDWAREINPUT
	_    [28]byte
}

// HARDWAREINPUT structure (8 bytes on 64-bit)
type HARDWAREINPUT struct {
	UMsg    b.DWORD
	WParamL b.WORD
	WParamH b.WORD
}
