package b

import (
	"syscall"
	"unsafe"
)

type PVOID = unsafe.Pointer
type PVOID64 = unsafe.Pointer

type CHAR = uint8
type SHORT = int16
type LONG = int32

// INT type is already declared in 'minwindef' file.
//type INT = int32

type WCHAR = uint16
type PWCHAR = *WCHAR
type LPWCH = *WCHAR
type PWCH = *WCHAR
type LPCWCH = *WCHAR
type PCWCH = *WCHAR

// Pointers to null-terminated strings.

type NWPSTR = *WCHAR
type LPWSTR = *WCHAR
type PWSTR = *WCHAR
type PZPWSTR = *PWSTR
type PCZPWSTR = *PWSTR
type LPUWSTR = *WCHAR
type PUWSTR = *WCHAR
type LPCWSTR = *WCHAR
type PCWSTR = *WCHAR
type PZPCWSTR = *PCWSTR
type PCZPCWSTR = *PCWSTR
type LPCUWSTR = *WCHAR
type PCUWSTR = *WCHAR
type PZZWSTR = *WCHAR
type PCZZWSTR = *WCHAR
type PUZZWSTR = *WCHAR
type PCUZZWSTR = *WCHAR
type PNZWCH = *WCHAR
type PCNZWCH = *WCHAR
type PUNZWCH = *WCHAR
type PCUNZWCH = *WCHAR
type LPCWCHAR = *WCHAR
type PCWCHAR = *WCHAR
type LPCUWCHAR = *WCHAR
type PCUWCHAR = *WCHAR

type UCSCHAR = uint32
type PUCSCHAR = *UCSCHAR
type PCUCSCHAR = *UCSCHAR
type PUCSSTR = *UCSCHAR
type PUUCSSTR = *UCSCHAR
type PCUCSSTR = *UCSCHAR
type PCUUCSSTR = *UCSCHAR
type PUUCSCHAR = *UCSCHAR
type PCUUCSCHAR = *UCSCHAR

type PCHAR = *CHAR
type LPCH = *CHAR
type PCH = *CHAR
type LPCCH = *CHAR
type PCCH = *CHAR
type NPSTR = *CHAR
type LPSTR = *CHAR
type PSTR = *CHAR
type LPCSTR = *CHAR
type PCSTR = *CHAR
type PZZSTR = *CHAR
type PCZZSTR = *CHAR
type PNZCH = *CHAR
type PCNZCH = *CHAR

type PZPSTR = *PSTR
type PCZPSTR = *PSTR
type PZPCSTR = *PCSTR
type PCZPCSTR = *PCSTR

type TCHAR = WCHAR
type PTCHAR = *WCHAR
type TBYTE = WCHAR
type PTBYTE = *WCHAR

type LPTCH = LPWCH
type PTCH = LPWCH
type LPCTCH = LPCWCH
type PCTCH = LPCWCH
type PTSTR = LPWSTR
type LPTSTR = LPWSTR
type PCTSTR = LPCWSTR
type LPCTSTR = LPCWSTR
type PUTSTR = LPUWSTR
type LPUTSTR = LPUWSTR
type PCUTSTR = LPCUWSTR
type LPCUTSTR = LPCUWSTR
type LP = LPWSTR
type PZZTSTR = PZZWSTR
type PCZZTSTR = PCZZWSTR
type PUZZTSTR = PUZZWSTR
type PCUZZTSTR = PCUZZWSTR
type PZPTSTR = PZPWSTR
type PNZTCH = PNZWCH
type PCNZTCH = PCNZWCH
type PUNZTCH = PUNZWCH
type PCUNZTCH = PCUNZWCH

type PSHORT = *SHORT
type PLONG = *LONG

type HANDLE = unsafe.Pointer
type PHANDLE = *HANDLE

type FCHAR = BYTE
type FSHORT = WORD
type FLONG = DWORD

/*
HRESULT is a numeric code of result used in Microsoft Windows API.

Definition in 'winnt.h' file:

	typedef _Return_type_success_(return >= 0) long HRESULT;

Helper macros in 'intsafe.h' file:

	#define SUCCEEDED(hr)   (((HRESULT)(hr)) >= 0)
	#define FAILED(hr)      (((HRESULT)(hr)) < 0)

Notes

	It is critically important that Microsoft uses 'HRESULT' as a signed
	integer. Go language's 'uintptr' type is an unsigned integer which is not
	directly compatible with Microsoft's API !

	It is crucially important first to cast type of values returned by
	Microsoft's API into a signed integer before any further processing !
*/
type HRESULT int32

func (hr HRESULT) SUCCEEDED() bool {
	return hr >= 0
}
func (hr HRESULT) FAILED() bool {
	return hr < 0
}
func (hr HRESULT) Error() string {
	return syscall.Errno(hr).Error()
}

// hresult_overflow is a function to simulate an ugly hack of integer overflow
// used by Microsoft API. Such C/C++ hack is not supported by Go language.
func hresult_overflow(val uint32) HRESULT {
	return HRESULT(val)
}

type CCHAR = uint8
type LCID = DWORD
type PLCID = PDWORD
type LANGID = WORD

type FLOAT128 struct {
	LowPart  int64
	HighPart int64
}
type PFLOAT128 = *FLOAT128

type LONGLONG = int64
type ULONGLONG = uint64

type PLONGLONG = *LONGLONG
type PULONGLONG = *ULONGLONG

type USN = LONGLONG

type RTL_REFERENCE_COUNT = LONG_PTR
type PRTL_REFERENCE_COUNT = *LONG_PTR
type RTL_REFERENCE_COUNT32 = LONG
type PRTL_REFERENCE_COUNT32 = *LONG

type LUID struct {
	LowPart  DWORD
	HighPart LONG
}
type PLUID = *LUID

type DWORDLONG = ULONGLONG
type PDWORDLONG = *DWORDLONG

type BOOLEAN = BYTE
type PBOOLEAN = *BOOLEAN

type KSPIN_LOCK = ULONG_PTR
type PKSPIN_LOCK = *KSPIN_LOCK

const (
	APPLICATION_ERROR_MASK       = 0x20000000
	ERROR_SEVERITY_SUCCESS       = 0x00000000
	ERROR_SEVERITY_INFORMATIONAL = 0x40000000
	ERROR_SEVERITY_WARNING       = 0x80000000
	ERROR_SEVERITY_ERROR         = 0xC0000000
)
