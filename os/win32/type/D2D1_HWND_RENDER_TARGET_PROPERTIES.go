package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
	e "github.com/vault-thirteen/auxie/os/win32/enum"
)

type D2D1_HWND_RENDER_TARGET_PROPERTIES struct {
	Hwnd           b.HWND
	PixelSize      D2D1_SIZE_U
	PresentOptions e.D2D1_PRESENT_OPTIONS
}

func HwndRenderTargetProperties(hwnd b.HWND, pixelSize D2D1_SIZE_U, presentOptions e.D2D1_PRESENT_OPTIONS) D2D1_HWND_RENDER_TARGET_PROPERTIES {
	return D2D1_HWND_RENDER_TARGET_PROPERTIES{
		Hwnd:           hwnd,
		PixelSize:      pixelSize,
		PresentOptions: presentOptions,
	}
}
