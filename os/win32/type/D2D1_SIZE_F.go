package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type D2D1_SIZE_F = D2D_SIZE_F

func D2D1_SizeF(width b.FLOAT, height b.FLOAT) D2D1_SIZE_F {
	return D2D1_SIZE_F{Width: width, Height: height}
}
