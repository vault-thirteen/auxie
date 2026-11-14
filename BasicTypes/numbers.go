package bt

import (
	"math/big"

	b "github.com/vault-thirteen/auxie/os/win32/base"
)

// Microsoft's types.

// Word is the word type.
// It is a 16-bit unsigned integer.
// WORD type in the open specification by Microsoft:
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp/f8573df3-a44a-4a50-b070-ac4c3aa78e3c
type Word = b.WORD

// DWord is the double word type.
// It is a 32-bit unsigned integer.
// DWORD type in the open specification by Microsoft:
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp/262627d8-3418-4627-9218-4ffe110850b2
type DWord = b.DWORD

// Common sense types.

// Byte is a byte type.
type Byte = b.BYTE

// SByte means a signed byte type.
type SByte = b.INT8

// UByte means an unsigned byte type.
type UByte = b.CHAR

// SShort means a signed short type.
type SShort = b.SHORT

// UShort means an unsigned short type.
type UShort = b.USHORT

// SLong means a signed long type.
type SLong = b.LONG

// ULong means an unsigned long type.
type ULong = b.ULONG

// Float is a 32-bit floating point number.
type Float = b.FLOAT

// Double is a 64-bit floating point number.
type Double = b.DOUBLE

// Rational is a rational type.
type Rational = *big.Rat
