package mcci

import "image/color"

func ConvertColor8bitTo16bit(c8 uint8) (c16 uint16) {
	x := uint16(c8)
	return (x << 8) | x
}
func ConvertColor16bitTo8bit(c16 uint16) (c8 uint8) {
	return uint8(c16 >> 8)
}

func GetRgbaAsFloat32(c color.Color) (r, g, b, a float32) {
	r_temp, g_temp, b_temp, a_temp := c.RGBA()
	return float32(r_temp), float32(g_temp), float32(b_temp), float32(a_temp)
}
