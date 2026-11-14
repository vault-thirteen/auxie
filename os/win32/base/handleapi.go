package b

/* #define INVALID_HANDLE_VALUE ((HANDLE)(LONG_PTR)-1) */
const INVALID_HANDLE_VALUE = ^uintptr(0) // Emulate overflow with bitwise NOT operator (^).
