package win32

//go:generate mkwinsyscall -output ole32_functions.go ole32.go

// Functions that set last error:
/*
 */

// Functions that do not set last error:
/*
//sys CoCreateInstance(rclsid *windows.GUID, pUnkOuter *i.IUnknown, dwClsContext b.DWORD, riid *windows.GUID, ppv *unsafe.Pointer) (ret b.HRESULT) = ole32.CoCreateInstance
*/
