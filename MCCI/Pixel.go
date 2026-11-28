package mcci

import (
	"math"
)

const (
	Err_UnsupportedPixelFormat = "unsupported pixel format"
)

type Pixel[T IPixel] struct {
	R T
	G T
	B T
	A T
}

// MakePixel_I1B_8bit makes a pixel using 8-bit integer values.
// Such pixel stores data within a [0..255] range.
// Source must have 8-bit colour information.
func MakePixel_I1B_8bit[T IPixel](pd *PixelData, alphaChannelMode AlphaChannelMode) (p Pixel[T]) {
	switch alphaChannelMode {
	case AlphaChannelMode_Straight:
		p = Pixel[T]{
			R: T(pd.R_straight),
			G: T(pd.G_straight),
			B: T(pd.B_straight),
			A: T(pd.Alpha),
		}

	case AlphaChannelMode_Premultiplied:
		p = Pixel[T]{
			R: T(pd.R_premultiplied),
			G: T(pd.G_premultiplied),
			B: T(pd.B_premultiplied),
			A: T(pd.Alpha),
		}
	}

	return p
}

// MakePixel_I1B_16bit makes a pixel using 8-bit integer values.
// Such pixel stores data within a [0..255] range.
// Source must have 16-bit colour information.
func MakePixel_I1B_16bit[T IPixel](pd *PixelData, alphaChannelMode AlphaChannelMode) (p Pixel[T]) {
	switch alphaChannelMode {
	case AlphaChannelMode_Straight:
		p = Pixel[T]{
			R: T(ConvertColour_16bitTo8bit(uint16(pd.R_straight))),
			G: T(ConvertColour_16bitTo8bit(uint16(pd.G_straight))),
			B: T(ConvertColour_16bitTo8bit(uint16(pd.B_straight))),
			A: T(ConvertColour_16bitTo8bit(uint16(pd.Alpha))),
		}

	case AlphaChannelMode_Premultiplied:
		p = Pixel[T]{
			R: T(ConvertColour_16bitTo8bit(uint16(pd.R_premultiplied))),
			G: T(ConvertColour_16bitTo8bit(uint16(pd.G_premultiplied))),
			B: T(ConvertColour_16bitTo8bit(uint16(pd.B_premultiplied))),
			A: T(ConvertColour_16bitTo8bit(uint16(pd.Alpha))),
		}
	}

	return p
}

// MakePixel_I2B_8bit makes a pixel using 16-bit integer values.
// Such pixel stores data within a [0..65535] range.
// Source must have 8-bit colour information.
func MakePixel_I2B_8bit[T IPixel](pd *PixelData, alphaChannelMode AlphaChannelMode) (p Pixel[T]) {
	switch alphaChannelMode {
	case AlphaChannelMode_Straight:
		p = Pixel[T]{
			R: T(ConvertColour_8bitTo16bit(uint8(pd.R_straight))),
			G: T(ConvertColour_8bitTo16bit(uint8(pd.G_straight))),
			B: T(ConvertColour_8bitTo16bit(uint8(pd.B_straight))),
			A: T(ConvertColour_8bitTo16bit(uint8(pd.Alpha))),
		}

	case AlphaChannelMode_Premultiplied:
		p = Pixel[T]{
			R: T(ConvertColour_8bitTo16bit(uint8(pd.R_premultiplied))),
			G: T(ConvertColour_8bitTo16bit(uint8(pd.G_premultiplied))),
			B: T(ConvertColour_8bitTo16bit(uint8(pd.B_premultiplied))),
			A: T(ConvertColour_8bitTo16bit(uint8(pd.Alpha))),
		}
	}

	return p
}

// MakePixel_I2B_16bit makes a pixel using 16-bit integer values.
// Such pixel stores data within a [0..65535] range.
// Source must have 16-bit colour information.
func MakePixel_I2B_16bit[T IPixel](pd *PixelData, alphaChannelMode AlphaChannelMode) (p Pixel[T]) {
	switch alphaChannelMode {
	case AlphaChannelMode_Straight:
		p = Pixel[T]{
			R: T(pd.R_straight),
			G: T(pd.G_straight),
			B: T(pd.B_straight),
			A: T(pd.Alpha),
		}

	case AlphaChannelMode_Premultiplied:
		p = Pixel[T]{
			R: T(pd.R_premultiplied),
			G: T(pd.G_premultiplied),
			B: T(pd.B_premultiplied),
			A: T(pd.Alpha),
		}
	}

	return p
}

// MakePixel_F4B_8bit makes a pixel using 32-bit floating point values.
// Such pixel stores data in a normalised form, i.e. within a [0..1] range.
// Source must have 8-bit colour information.
func MakePixel_F4B_8bit[T IPixel](pd *PixelData, alphaChannelMode AlphaChannelMode) (p Pixel[T]) {
	return makePixel_F4B[T](pd, float32(math.MaxUint8), alphaChannelMode)
}

// MakePixel_F4B_16bit makes a pixel using 32-bit floating point values.
// Such pixel stores data in a normalised form, i.e. within a [0..1] range.
// Source must have 16-bit colour information.
func MakePixel_F4B_16bit[T IPixel](pd *PixelData, alphaChannelMode AlphaChannelMode) (p Pixel[T]) {
	return makePixel_F4B[T](pd, float32(math.MaxUint16), alphaChannelMode)
}

// makePixel_F4B makes a pixel using 32-bit floating point values.
// Such pixel stores data in a normalised form, i.e. within a [0..1] range.
// Divider depends on the source colour information (8-bit or 16-bit).
func makePixel_F4B[T IPixel](pd *PixelData, divider float32, alphaChannelMode AlphaChannelMode) (p Pixel[T]) {
	switch alphaChannelMode {
	case AlphaChannelMode_Straight:
		p = Pixel[T]{
			R: T(pd.R_straight / divider),
			G: T(pd.G_straight / divider),
			B: T(pd.B_straight / divider),
			A: T(pd.Alpha_normalised),
		}

	case AlphaChannelMode_Premultiplied:
		p = Pixel[T]{
			R: T(pd.R_premultiplied / divider),
			G: T(pd.G_premultiplied / divider),
			B: T(pd.B_premultiplied / divider),
			A: T(pd.Alpha_normalised),
		}
	}

	return p
}
