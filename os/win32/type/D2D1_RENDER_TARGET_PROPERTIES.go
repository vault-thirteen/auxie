package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
	e "github.com/vault-thirteen/auxie/os/win32/enum"
)

type D2D1_RENDER_TARGET_PROPERTIES struct {
	Type        e.D2D1_RENDER_TARGET_TYPE
	PixelFormat D2D1_PIXEL_FORMAT
	DpiX        b.FLOAT
	DpiY        b.FLOAT
	Usage       e.D2D1_RENDER_TARGET_USAGE
	MinLevel    e.D2D1_FEATURE_LEVEL
}

func RenderTargetProperties() D2D1_RENDER_TARGET_PROPERTIES {
	return D2D1_RENDER_TARGET_PROPERTIES{
		Type: e.D2D1_RENDER_TARGET_TYPE_DEFAULT,
		PixelFormat: D2D1_PIXEL_FORMAT{
			Format:    e.DXGI_FORMAT_UNKNOWN,
			AlphaMode: e.D2D1_ALPHA_MODE_UNKNOWN,
		},
		DpiX:     0.0, // 0.0 means default system DPI
		DpiY:     0.0,
		Usage:    e.D2D1_RENDER_TARGET_USAGE_NONE,
		MinLevel: e.D2D1_FEATURE_LEVEL_DEFAULT,
	}
}
