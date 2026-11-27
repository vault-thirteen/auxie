package mcci

import (
	"math"

	"github.com/vault-thirteen/auxie/errors"
)

func (img *Image[T]) ChangeAlphaMode(alphaMode AlphaChannelMode) (err error) {
	if img.Settings.AlphaChannelMode == alphaMode {
		return errors.ErrorFromText(Err_NoChangeIsNeeded)
	}

	switch alphaMode {
	case AlphaChannelMode_Straight:
		{
			switch img.Settings.ChannelPixelFormat {
			case PixelFormat_Integer_1B: // 8-bit, 0..255.
				return img.changeAlphaModeToStraight_I1B()

			case PixelFormat_Integer_2B: // 16-bit, 0..65535.
				return img.changeAlphaModeToStraight_I2B()

			case PixelFormat_Float_4B: // 32-bit, 0..1.
				return img.changeAlphaModeToStraight_F4B()

			default:
				return errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
			}
		}

	case AlphaChannelMode_Premultiplied:
		{
			switch img.Settings.ChannelPixelFormat {
			case PixelFormat_Integer_1B: // 8-bit, 0..255.
				return img.changeAlphaModeToPremultiplied_I1B()

			case PixelFormat_Integer_2B: // 16-bit, 0..65535.
				return img.changeAlphaModeToPremultiplied_I2B()

			case PixelFormat_Float_4B: // 32-bit, 0..1.
				return img.changeAlphaModeToPremultiplied_F4B()

			default:
				return errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
			}
		}

	default:
		return errors.ErrorFromText(Err_UnsupportedAlphaChannelMode)
	}
}
func (img *Image[T]) changeAlphaModeToStraight_I1B() (err error) {
	if img.Settings.AlphaChannelMode != AlphaChannelMode_Premultiplied {
		return errors.ErrorFromText(Err_UnsupportedAlphaChannelMode)
	}

	alphaChannel := img.ChannelData[img.ChannelIndex.Alpha]
	redChannel := img.ChannelData[img.ChannelIndex.Red]
	greenChannel := img.ChannelData[img.ChannelIndex.Green]
	blueChannel := img.ChannelData[img.ChannelIndex.Blue]

	var pixelMaxValue = any(uint8(math.MaxUint8)).(T)

	var alpha_norm float32
	for y, alpha_row := range alphaChannel.Canvas {
		for x, alpha_value := range alpha_row {
			if alpha_value == 0 {
				// Important note ! Frankly speaking, such operation is not
				// possible when alpha channel value is zero, as we can not
				// "resurrect" information which is already lost. For
				// compatibility purposes we simply return junk. If you need
				// to use colours with extremely low alpha values, never use
				// premultiplied format, as it will destroy your source data !
				redChannel.Canvas[y][x] = pixelMaxValue
				greenChannel.Canvas[y][x] = pixelMaxValue
				blueChannel.Canvas[y][x] = pixelMaxValue
			} else {
				alpha_norm = float32(alpha_value) / float32(math.MaxUint8)
				redChannel.Canvas[y][x] = T(float32(redChannel.Canvas[y][x]) / alpha_norm)
				greenChannel.Canvas[y][x] = T(float32(greenChannel.Canvas[y][x]) / alpha_norm)
				blueChannel.Canvas[y][x] = T(float32(blueChannel.Canvas[y][x]) / alpha_norm)
			}
		}
	}

	img.Settings.AlphaChannelMode = AlphaChannelMode_Straight

	return nil
}
func (img *Image[T]) changeAlphaModeToStraight_I2B() (err error) {
	if img.Settings.AlphaChannelMode != AlphaChannelMode_Premultiplied {
		return errors.ErrorFromText(Err_UnsupportedAlphaChannelMode)
	}

	alphaChannel := img.ChannelData[img.ChannelIndex.Alpha]
	redChannel := img.ChannelData[img.ChannelIndex.Red]
	greenChannel := img.ChannelData[img.ChannelIndex.Green]
	blueChannel := img.ChannelData[img.ChannelIndex.Blue]

	var pixelMaxValue = any(uint16(math.MaxUint16)).(T)

	var alpha_norm float32
	for y, alpha_row := range alphaChannel.Canvas {
		for x, alpha_value := range alpha_row {
			if alpha_value == 0 {
				// Important note ! Frankly speaking, such operation is not
				// possible when alpha channel value is zero, as we can not
				// "resurrect" information which is already lost. For
				// compatibility purposes we simply return junk. If you need
				// to use colours with extremely low alpha values, never use
				// premultiplied format, as it will destroy your source data !
				redChannel.Canvas[y][x] = pixelMaxValue
				greenChannel.Canvas[y][x] = pixelMaxValue
				blueChannel.Canvas[y][x] = pixelMaxValue
			} else {
				alpha_norm = float32(alpha_value) / float32(math.MaxUint16)
				redChannel.Canvas[y][x] = T(float32(redChannel.Canvas[y][x]) / alpha_norm)
				greenChannel.Canvas[y][x] = T(float32(greenChannel.Canvas[y][x]) / alpha_norm)
				blueChannel.Canvas[y][x] = T(float32(blueChannel.Canvas[y][x]) / alpha_norm)
			}
		}
	}

	img.Settings.AlphaChannelMode = AlphaChannelMode_Straight

	return nil
}
func (img *Image[T]) changeAlphaModeToStraight_F4B() (err error) {
	if img.Settings.AlphaChannelMode != AlphaChannelMode_Premultiplied {
		return errors.ErrorFromText(Err_UnsupportedAlphaChannelMode)
	}

	alphaChannel := img.ChannelData[img.ChannelIndex.Alpha]
	redChannel := img.ChannelData[img.ChannelIndex.Red]
	greenChannel := img.ChannelData[img.ChannelIndex.Green]
	blueChannel := img.ChannelData[img.ChannelIndex.Blue]

	var alpha_norm float32
	for y, alpha_row := range alphaChannel.Canvas {
		for x, alpha_value := range alpha_row {
			if alpha_value == 0 {
				// Important note ! Frankly speaking, such operation is not
				// possible when alpha channel value is zero, as we can not
				// "resurrect" information which is already lost. For
				// compatibility purposes we simply return junk. If you need
				// to use colours with extremely low alpha values, never use
				// premultiplied format, as it will destroy your source data !
				redChannel.Canvas[y][x] = T(float32(1.0))
				greenChannel.Canvas[y][x] = T(float32(1.0))
				blueChannel.Canvas[y][x] = T(float32(1.0))
			} else {
				alpha_norm = float32(alpha_value) // Alpha is already normalised.
				redChannel.Canvas[y][x] = T(float32(redChannel.Canvas[y][x]) / alpha_norm)
				greenChannel.Canvas[y][x] = T(float32(greenChannel.Canvas[y][x]) / alpha_norm)
				blueChannel.Canvas[y][x] = T(float32(blueChannel.Canvas[y][x]) / alpha_norm)
			}
		}
	}

	img.Settings.AlphaChannelMode = AlphaChannelMode_Straight

	return nil
}
func (img *Image[T]) changeAlphaModeToPremultiplied_I1B() (err error) {
	if img.Settings.AlphaChannelMode != AlphaChannelMode_Straight {
		return errors.ErrorFromText(Err_UnsupportedAlphaChannelMode)
	}

	alphaChannel := img.ChannelData[img.ChannelIndex.Alpha]
	redChannel := img.ChannelData[img.ChannelIndex.Red]
	greenChannel := img.ChannelData[img.ChannelIndex.Green]
	blueChannel := img.ChannelData[img.ChannelIndex.Blue]

	var alpha_norm float32
	for y, alpha_row := range alphaChannel.Canvas {
		for x, alpha_value := range alpha_row {
			// Important note ! Never use premultiplied format with low
			// alpha channel values, as it will destroy your source data !
			alpha_norm = float32(alpha_value) / float32(math.MaxUint8)
			redChannel.Canvas[y][x] = T(float32(redChannel.Canvas[y][x]) * alpha_norm)
			greenChannel.Canvas[y][x] = T(float32(greenChannel.Canvas[y][x]) * alpha_norm)
			blueChannel.Canvas[y][x] = T(float32(blueChannel.Canvas[y][x]) * alpha_norm)
		}
	}

	img.Settings.AlphaChannelMode = AlphaChannelMode_Premultiplied

	return nil
}
func (img *Image[T]) changeAlphaModeToPremultiplied_I2B() (err error) {
	if img.Settings.AlphaChannelMode != AlphaChannelMode_Straight {
		return errors.ErrorFromText(Err_UnsupportedAlphaChannelMode)
	}

	alphaChannel := img.ChannelData[img.ChannelIndex.Alpha]
	redChannel := img.ChannelData[img.ChannelIndex.Red]
	greenChannel := img.ChannelData[img.ChannelIndex.Green]
	blueChannel := img.ChannelData[img.ChannelIndex.Blue]

	var alpha_norm float32
	for y, alpha_row := range alphaChannel.Canvas {
		for x, alpha_value := range alpha_row {
			// Important note ! Never use premultiplied format with low
			// alpha channel values, as it will destroy your source data !
			alpha_norm = float32(alpha_value) / float32(math.MaxUint16)
			redChannel.Canvas[y][x] = T(float32(redChannel.Canvas[y][x]) * alpha_norm)
			greenChannel.Canvas[y][x] = T(float32(greenChannel.Canvas[y][x]) * alpha_norm)
			blueChannel.Canvas[y][x] = T(float32(blueChannel.Canvas[y][x]) * alpha_norm)
		}
	}

	img.Settings.AlphaChannelMode = AlphaChannelMode_Premultiplied

	return nil
}
func (img *Image[T]) changeAlphaModeToPremultiplied_F4B() (err error) {
	if img.Settings.AlphaChannelMode != AlphaChannelMode_Straight {
		return errors.ErrorFromText(Err_UnsupportedAlphaChannelMode)
	}

	alphaChannel := img.ChannelData[img.ChannelIndex.Alpha]
	redChannel := img.ChannelData[img.ChannelIndex.Red]
	greenChannel := img.ChannelData[img.ChannelIndex.Green]
	blueChannel := img.ChannelData[img.ChannelIndex.Blue]

	var alpha_norm float32
	for y, alpha_row := range alphaChannel.Canvas {
		for x, alpha_value := range alpha_row {
			// Important note ! Never use premultiplied format with low
			// alpha channel values, as it will destroy your source data !
			alpha_norm = float32(alpha_value) // Alpha is already normalised.
			redChannel.Canvas[y][x] = T(float32(redChannel.Canvas[y][x]) * alpha_norm)
			greenChannel.Canvas[y][x] = T(float32(greenChannel.Canvas[y][x]) * alpha_norm)
			blueChannel.Canvas[y][x] = T(float32(blueChannel.Canvas[y][x]) * alpha_norm)
		}
	}

	img.Settings.AlphaChannelMode = AlphaChannelMode_Premultiplied

	return nil
}
