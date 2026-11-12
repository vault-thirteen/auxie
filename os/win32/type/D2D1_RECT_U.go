package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type D2D1_RECT_U = D2D_RECT_U

func D2D1_RectU(left, top, right, bottom b.UINT32) D2D1_RECT_U {
	return D2D1_RECT_U{Left: left, Top: top, Right: right, Bottom: bottom}
}
