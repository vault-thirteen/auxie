package mcci

import (
	"image/color"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	Err_DimensionsMismatch = "dimensions mismatch"
	Err_AreaIsOutsideImage = "area is outside the image"
)

// ReplaceWithArea replaces the area of a destination image with pixels taken
// from the area of a source image. Destination and source areas must be within
// their images and must have same dimensions.
func (img *Image[T]) ReplaceWithArea(srcImage *Image[T], srcArea *Area[Size], dstArea *Area[Size]) (err error) {
	if srcArea.IsMalformed() || dstArea.IsMalformed() {
		return errors.ErrorFromText(Err_AreaIsMalformed)
	}
	if !srcArea.HasSameDimension(dstArea) {
		return errors.ErrorFromText(Err_DimensionsMismatch)
	}
	if !img.Area().ContainsArea(dstArea) {
		return errors.ErrorFromText(Err_AreaIsOutsideImage)
	}
	if !srcImage.Area().ContainsArea(srcArea) {
		return errors.ErrorFromText(Err_AreaIsOutsideImage)
	}

	dstChannelRed := img.ChannelData[img.ChannelIndex.Red]
	dstChannelGreen := img.ChannelData[img.ChannelIndex.Green]
	dstChannelBlue := img.ChannelData[img.ChannelIndex.Blue]
	dstChannelAlpha := img.ChannelData[img.ChannelIndex.Alpha]

	srcChannelRed := srcImage.ChannelData[srcImage.ChannelIndex.Red]
	srcChannelGreen := srcImage.ChannelData[srcImage.ChannelIndex.Green]
	srcChannelBlue := srcImage.ChannelData[srcImage.ChannelIndex.Blue]
	srcChannelAlpha := srcImage.ChannelData[srcImage.ChannelIndex.Alpha]

	dstXFirst := dstArea.BottomLeft.X
	dstXLast := dstArea.TopRight.X
	dstYFirst := dstArea.BottomLeft.Y
	dstYLast := dstArea.TopRight.Y

	srcXFirst := srcArea.BottomLeft.X
	srcXLast := srcArea.TopRight.X
	srcYFirst := srcArea.BottomLeft.Y
	srcYLast := srcArea.TopRight.Y

	// Go language does not support iteration on two variables as in C++ !
	// What a shame ! We will be inventing the wheel again and again ...

	// Y loop.
	srcY := srcYFirst
	dstY := dstYFirst
	for {
		if (srcY > srcYLast) || (dstY > dstYLast) {
			break
		}

		// X loop.
		srcX := srcXFirst
		dstX := dstXFirst
		for {
			if (srcX > srcXLast) || (dstX > dstXLast) {
				break
			}

			// Copy pixels.
			dstChannelRed.Canvas[dstY][dstX] = srcChannelRed.Canvas[srcY][srcX]
			dstChannelGreen.Canvas[dstY][dstX] = srcChannelGreen.Canvas[srcY][srcX]
			dstChannelBlue.Canvas[dstY][dstX] = srcChannelBlue.Canvas[srcY][srcX]
			dstChannelAlpha.Canvas[dstY][dstX] = srcChannelAlpha.Canvas[srcY][srcX]

			// Next.
			srcX++
			dstX++
		}

		// Next.
		srcY++
		dstY++
	}

	return nil
}

// ReplaceWithImage replaces the area of a destination image with pixels taken
// from a source image. Destination area must be within destination image and
// must have same dimensions as source image.
func (img *Image[T]) ReplaceWithImage(srcImage *Image[T], dstArea *Area[Size]) (err error) {
	srcArea := NewAreaFromDimensions(srcImage.Settings.Width, srcImage.Settings.Height)
	return img.ReplaceWithArea(srcImage, srcArea, dstArea)
}

// ReplaceWithImageAt replaces the area of a destination image selected by its
// bottom left corner with pixels taken from a source image. Destination area
// must be within destination image and must have enough space to fit the
// source image.
func (img *Image[T]) ReplaceWithImageAt(srcImage *Image[T], at Point[Size]) (err error) {
	srcArea := NewAreaFromDimensions(srcImage.Settings.Width, srcImage.Settings.Height)
	vec := NewVector[Size](srcImage.Settings.Width-1, srcImage.Settings.Height-1)
	dstArea := NewAreaFromPointAndVector[Size](at, vec)
	return img.ReplaceWithArea(srcImage, srcArea, dstArea)
}

// FillWithColour fill selected area with selected colour.
func (img *Image[T]) FillWithColour(c color.Color, dstArea *Area[Size]) (err error) {
	if dstArea.IsMalformed() {
		return errors.ErrorFromText(Err_AreaIsMalformed)
	}
	if !img.Area().ContainsArea(dstArea) {
		return errors.ErrorFromText(Err_AreaIsOutsideImage)
	}

	// Prepare the pixel.
	var p Pixel[T]
	{
		var pd *PixelData
		pd, err = GetPixelData_Abstract(c)
		if err != nil {
			return err
		}

		switch img.Settings.ChannelPixelFormat {
		case PixelFormat_Integer_1B:
			p = MakePixel_I1B_16bit[T](pd, img.Settings.AlphaChannelMode)

		case PixelFormat_Integer_2B:
			p = MakePixel_I2B_16bit[T](pd, img.Settings.AlphaChannelMode)

		case PixelFormat_Float_4B:
			p = MakePixel_F4B_16bit[T](pd, img.Settings.AlphaChannelMode)

		default:
			return errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
		}
	}

	// Draw pixels.
	dstChannelRed := img.ChannelData[img.ChannelIndex.Red]
	dstChannelGreen := img.ChannelData[img.ChannelIndex.Green]
	dstChannelBlue := img.ChannelData[img.ChannelIndex.Blue]
	dstChannelAlpha := img.ChannelData[img.ChannelIndex.Alpha]

	dstXFirst := dstArea.BottomLeft.X
	dstXLast := dstArea.TopRight.X
	dstYFirst := dstArea.BottomLeft.Y
	dstYLast := dstArea.TopRight.Y

	for dstY := dstYFirst; dstY <= dstYLast; dstY++ {
		for dstX := dstXFirst; dstX <= dstXLast; dstX++ {
			dstChannelRed.Canvas[dstY][dstX] = p.R
			dstChannelGreen.Canvas[dstY][dstX] = p.G
			dstChannelBlue.Canvas[dstY][dstX] = p.B
			dstChannelAlpha.Canvas[dstY][dstX] = p.A
		}
	}

	return nil
}
