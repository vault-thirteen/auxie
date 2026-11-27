package mcci

import "github.com/vault-thirteen/auxie/errors"

const (
	Err_DimensionsMismatch = "dimensions mismatch"
	Err_AreaIsOutsideImage = "area is outside the image"
)

// DrawOverArea replaces the area of a destination image with pixels taken from
// the area of a source image. Destination and source areas must be within
// their images and must have same dimensions.
func (img *Image[T]) DrawOverArea(srcImage *Image[T], srcArea *Area[Size], dstArea *Area[Size]) (err error) {
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

// DrawOverImage replaces the area of a destination image with pixels taken from
// a source image. Destination area must be within destination image and must
// have same dimensions as source image.
func (img *Image[T]) DrawOverImage(srcImage *Image[T], dstArea *Area[Size]) (err error) {
	srcArea := NewAreaFromDimensions(srcImage.Settings.Width, srcImage.Settings.Height)
	return img.DrawOverArea(srcImage, srcArea, dstArea)
}
