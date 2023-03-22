package bt

import "math/big"

// Microsoft's types.

// Word is the word type.
// It is a 16-bit unsigned integer.
// WORD type in the open specification by Microsoft:
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp/f8573df3-a44a-4a50-b070-ac4c3aa78e3c
type Word = uint16

// DWord is the double word type.
// It is a 32-bit unsigned integer.
// DWORD type in the open specification by Microsoft:
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp/262627d8-3418-4627-9218-4ffe110850b2
type DWord = uint32

// Common sense types.

// Byte is a byte type.
type Byte = byte

// SByte means a signed byte type.
type SByte = int8

// UByte means an unsigned byte type.
type UByte = uint8

// SShort means a signed short type.
type SShort = int16

// UShort means an unsigned short type.
type UShort = uint16

// SLong means a signed long type.
type SLong = int32

// ULong means an unsigned long type.
type ULong = uint32

// Float is a 32-bit floating point number.
type Float = float32

// Double is a 64-bit floating point number.
type Double = float64

// Rational is a rational type.
// Unfortunately Golang does not support unsigned rationals out-of-the-box, so
// we will not define signed or unsigned variants of the type here.
type Rational = *big.Rat
