package b

import "golang.org/x/sys/windows"

type GUID = windows.GUID
type LPGUID = *GUID
type LPCGUID = *GUID

type IID = GUID
type LPIID = *IID

type CLSID = GUID
type LPCLSID = *CLSID
type FMTID = GUID
type LPFMTID = *FMTID

type REFIID = *GUID
type REFCLSID = *GUID
type REFFMTID = *GUID
