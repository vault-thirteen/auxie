package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type D2D1_SIZE_U = D2D_SIZE_U

func D2D1_SizeU(width b.UINT32, height b.UINT32) D2D1_SIZE_U {
	return D2D1_SIZE_U{Width: width, Height: height}
}
