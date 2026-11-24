package mcci

import (
	"image/color"
	"math"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	Err_InvalidPixelFormat = "invalid pixel format"
)

type PixelData struct {
	R_straight float32
	G_straight float32
	B_straight float32

	R_premultiplied float32
	G_premultiplied float32
	B_premultiplied float32

	Alpha            float32
	Alpha_normalised float32
}

// GetPixelData_RGBA reads pixel information directly from the raw pixel.
// Raw RGBA pixel stores pre-multiplied 8-bit colour data.
func GetPixelData_RGBA(abstractPixel color.Color) (pd *PixelData, err error) {
	var pixel_RGBA color.RGBA
	var ok bool
	pixel_RGBA, ok = abstractPixel.(color.RGBA)
	if !ok {
		return nil, errors.ErrorFromText(Err_InvalidPixelFormat)
	}

	pd = &PixelData{
		R_premultiplied: float32(pixel_RGBA.R),
		G_premultiplied: float32(pixel_RGBA.G),
		B_premultiplied: float32(pixel_RGBA.B),

		Alpha:            float32(pixel_RGBA.A),
		Alpha_normalised: float32(pixel_RGBA.A) / float32(math.MaxUint8), // 8-bit.
	}

	if pd.Alpha_normalised != 0 {
		pd.R_straight = pd.R_premultiplied / pd.Alpha_normalised
		pd.G_straight = pd.G_premultiplied / pd.Alpha_normalised
		pd.B_straight = pd.B_premultiplied / pd.Alpha_normalised
	} else {
		pd.R_straight = 0
		pd.G_straight = 0
		pd.B_straight = 0
	}

	return pd, nil
}

// GetPixelData_NRGBA reads pixel information directly from the raw pixel.
// Raw NRGBA pixel stores straight 8-bit colour data.
func GetPixelData_NRGBA(abstractPixel color.Color) (pd *PixelData, err error) {
	var pixel_NRGBA color.NRGBA
	var ok bool
	pixel_NRGBA, ok = abstractPixel.(color.NRGBA)
	if !ok {
		return nil, errors.ErrorFromText(Err_InvalidPixelFormat)
	}

	pd = &PixelData{
		R_straight: float32(pixel_NRGBA.R),
		G_straight: float32(pixel_NRGBA.G),
		B_straight: float32(pixel_NRGBA.B),

		Alpha:            float32(pixel_NRGBA.A),
		Alpha_normalised: float32(pixel_NRGBA.A) / float32(math.MaxUint8), // 8-bit.
	}

	pd.R_premultiplied = pd.R_straight * pd.Alpha_normalised
	pd.G_premultiplied = pd.G_straight * pd.Alpha_normalised
	pd.B_premultiplied = pd.B_straight * pd.Alpha_normalised

	return pd, nil
}

// GetPixelData_Abstract reads pixel information from a common interface.
// Common interface returns pixel with a pre-multiplied 16-bit colour data.
func GetPixelData_Abstract(abstractPixel color.Color) (pd *PixelData, err error) {
	r_temp, g_temp, b_temp, a_temp := abstractPixel.RGBA()

	pd = &PixelData{
		R_premultiplied: float32(r_temp),
		G_premultiplied: float32(g_temp),
		B_premultiplied: float32(b_temp),

		Alpha:            float32(a_temp),
		Alpha_normalised: float32(a_temp) / float32(math.MaxUint16), // 16-bit.
	}

	if pd.Alpha_normalised != 0 {
		pd.R_straight = pd.R_premultiplied / pd.Alpha_normalised
		pd.G_straight = pd.G_premultiplied / pd.Alpha_normalised
		pd.B_straight = pd.B_premultiplied / pd.Alpha_normalised
	} else {
		pd.R_straight = 0
		pd.G_straight = 0
		pd.B_straight = 0
	}

	return pd, nil
}
