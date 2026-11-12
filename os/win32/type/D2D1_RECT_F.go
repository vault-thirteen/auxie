package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type D2D1_RECT_F = D2D_RECT_F

func D2D1_RectF(left, top, right, bottom b.FLOAT) D2D1_RECT_F {
	return D2D1_RECT_F{Left: left, Top: top, Right: right, Bottom: bottom}
}
