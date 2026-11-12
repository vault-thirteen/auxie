package b

type POINTER_64_INT = uint64

type INT8 = int8
type INT16 = int16
type INT32 = int32
type INT64 = int64

type UINT8 = uint8
type UINT16 = uint16
type UINT32 = uint32
type UINT64 = uint64

type LONG32 = int32
type PLONG32 = *LONG32

type ULONG32 = uint32
type PULONG32 = *ULONG32
type DWORD32 = uint32
type PDWORD32 = *DWORD32

type INT_PTR = uintptr
type PINT_PTR = *INT_PTR
type UINT_PTR = uintptr
type PUINT_PTR = *UINT_PTR

type LONG_PTR = uintptr
type PLONG_PTR = *LONG_PTR
type ULONG_PTR = uintptr
type PULONG_PTR = *ULONG_PTR

type HANDLE64 = uint64
type PHANDLE64 = *HANDLE64

// #define ADDRESS_TAG_BIT 0x40000000000UI64
const AddressTagBit = 0x40000000000

type SHANDLE_PTR = int64
type HANDLE_PTR = uint64
type UHALF_PTR = uint32
type PUHALF_PTR = *UHALF_PTR
type HALF_PTR = int32
type PHALF_PTR = *HALF_PTR

type SIZE_T = ULONG_PTR
type PSIZE_T = *SIZE_T
type SSIZE_T = LONG_PTR
type PSSIZE_T = *SSIZE_T

type DWORD_PTR = ULONG_PTR
type PDWORD_PTR = *ULONG_PTR

type LONG64 = int64
type PLONG64 = *LONG64

type ULONG64 = uint64
type PULONG64 = *ULONG64
type DWORD64 = uint64
type PDWORD64 = *DWORD64

type KAFFINITY = ULONG_PTR
type PKAFFINITY = *KAFFINITY
