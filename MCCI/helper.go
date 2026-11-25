package mcci

import (
	"math"
)

func ConvertColour_8bitTo16bit(c8 uint8) (c16 uint16) {
	x := uint16(c8)
	return (x << 8) | x
}
func ConvertColour_16bitTo8bit(c16 uint16) (c8 uint8) {
	return uint8(c16 >> 8)
}
func ConvertColour_NormalisedToInteger(in float32, multiplier float32) (out float32) {
	// Trim the value to avoid overflow.
	// HDR colours may have values greater than 1.0.
	if in > 1.0 {
		in = 1.0
	}

	return float32(math.Round(float64(in) * float64(multiplier)))
}
