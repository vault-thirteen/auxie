package i

import (
	"syscall"
	"unsafe"

	b "github.com/vault-thirteen/auxie/os/win32/base"
	"golang.org/x/sys/windows"
)

type IWICImagingFactory struct {
	Vtbl *IWICImagingFactoryVtbl
}

// IWICImagingFactoryVtbl is a VTable for IWICImagingFactory.
// All the fields were tested with C++ VTable tester.
type IWICImagingFactoryVtbl struct {
	// IUnknown methods
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IWICImagingFactory methods
	CreateDecoderFromFilename                uintptr
	CreateDecoderFromStream                  uintptr
	CreateDecoderFromFileHandle              uintptr
	CreateComponentInfo                      uintptr
	CreateDecoder                            uintptr
	CreateEncoder                            uintptr
	CreatePalette                            uintptr
	CreateFormatConverter                    uintptr
	CreateBitmapScaler                       uintptr
	CreateBitmapClipper                      uintptr
	CreateBitmapFlipRotator                  uintptr
	CreateStream                             uintptr
	CreateColorContext                       uintptr
	CreateColorTransformer                   uintptr
	CreateBitmap                             uintptr
	CreateBitmapFromSource                   uintptr
	CreateBitmapFromSourceRect               uintptr
	CreateBitmapFromMemory                   uintptr
	CreateBitmapFromHBITMAP                  uintptr
	CreateBitmapFromHICON                    uintptr
	CreateComponentEnumerator                uintptr
	CreateFastMetadataEncoderFromDecoder     uintptr
	CreateFastMetadataEncoderFromFrameDecode uintptr
	CreateQueryWriter                        uintptr
	CreateQueryWriterFromReader              uintptr
}

func (f *IWICImagingFactory) Release() b.ULONG {
	r0, _, _ := syscall.SyscallN(f.Vtbl.Release, uintptr(unsafe.Pointer(f)))
	return b.ULONG(r0)
}

var CLSID_WICImagingFactory = windows.GUID{
	Data1: 0xcacaf262,
	Data2: 0x9370,
	Data3: 0x4615,
	Data4: [8]byte{0xa1, 0x3b, 0x9f, 0x55, 0x39, 0xda, 0x4c, 0xa},
}

var CLSID_WICImagingFactory1 = windows.GUID{
	Data1: 0xcacaf262,
	Data2: 0x9370,
	Data3: 0x4615,
	Data4: [8]byte{0xa1, 0x3b, 0x9f, 0x55, 0x39, 0xda, 0x4c, 0xa},
}

var CLSID_WICImagingFactory2 = windows.GUID{
	Data1: 0x317d06e8,
	Data2: 0x5f24,
	Data3: 0x433d,
	Data4: [8]byte{0xbd, 0xf7, 0x79, 0xce, 0x68, 0xd8, 0xab, 0xc2},
}

// IID_IWICImagingFactory is an Interface Identifier for IWICImagingFactory.
// MIDL_INTERFACE("ec5ec8a9-c395-4314-9c77-54d7a935ff70")
var IID_IWICImagingFactory = windows.GUID{
	Data1: 0xec5ec8a9,
	Data2: 0xc395,
	Data3: 0x4314,
	Data4: [8]byte{0x9c, 0x77, 0x54, 0xd7, 0xa9, 0x35, 0xff, 0x70},
}
