package colorF

import (
	"github.com/vault-thirteen/auxie/os/win32/base"
	"github.com/vault-thirteen/auxie/os/win32/class/D2D1/ColorF/enum"
	t "github.com/vault-thirteen/auxie/os/win32/type"
)

const (
	sc_redShift   b.UINT32 = 16
	sc_greenShift b.UINT32 = 8
	sc_blueShift  b.UINT32 = 0
)

const (
	sc_redMask   b.UINT32 = 0xff << sc_redShift
	sc_greenMask b.UINT32 = 0xff << sc_greenShift
	sc_blueMask  b.UINT32 = 0xff << sc_blueShift
)

type ColorF struct {
	t.D2D1_COLOR_F
}

// Following functions are trying to emulate overloaded constructors of the
// 'ColorF' class and C++ default parameters functionality.

func NewColorF_1A(rgb b.UINT32, a b.FLOAT) (cf ColorF) {
	cf.Init(rgb, a)
	return cf
}
func NewColorF_1B(rgb b.UINT32) (cf ColorF) {
	cf.Init(rgb, 1.0)
	return cf
}
func NewColorF_2A(knownColor enum.Enum, a b.FLOAT) (cf ColorF) {
	cf.Init(b.UINT32(knownColor), a)
	return cf
}
func NewColorF_2B(knownColor enum.Enum) (cf ColorF) {
	cf.Init(b.UINT32(knownColor), 1.0)
	return cf
}
func NewColorF_3A(red, green, blue, alpha b.FLOAT) (cf ColorF) {
	return ColorF{
		R: red,
		G: green,
		B: blue,
		A: alpha,
	}
}
func NewColorF_3B(red, green, blue b.FLOAT) (cf ColorF) {
	return ColorF{
		R: red,
		G: green,
		B: blue,
		A: 1.0,
	}
}

func (cf *ColorF) Init(rgb b.UINT32, alpha b.FLOAT) {
	cf.R = b.FLOAT((rgb&sc_redMask)>>sc_redShift) / b.FLOAT(255.0)
	cf.G = b.FLOAT((rgb&sc_greenMask)>>sc_greenShift) / b.FLOAT(255.0)
	cf.B = b.FLOAT((rgb&sc_blueMask)>>sc_blueShift) / b.FLOAT(255.0)
	cf.A = alpha
}
