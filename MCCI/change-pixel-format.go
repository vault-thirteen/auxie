package mcci

import (
	"math"

	"github.com/vault-thirteen/auxie/errors"
)

func (img *Image[T]) ChangePixelFormatT[T2 IPixel](channelPixelFormat ChannelPixelFormat) (copy *Image[T2], err error) {
	switch img.Settings.ChannelPixelFormat { // Source format.
	case PixelFormat_Integer_1B: // 8-bit, 0..255.
		{
			switch channelPixelFormat { // Destination format.
			case PixelFormat_Integer_1B: // 8-bit, 0..255.
				return nil, errors.ErrorFromText(Err_NoChangeIsNeeded)

			case PixelFormat_Integer_2B: // 16-bit, 0..65535.
				return img.changePixelFormat_I1B_to_I2B[T2]()

			case PixelFormat_Float_4B: // 32-bit, 0..1.
				return img.changePixelFormat_I1B_to_F4B[T2]()

			default:
				return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
			}
		}

	case PixelFormat_Integer_2B: // 16-bit, 0..65535.
		{
			switch channelPixelFormat { // Destination format.
			case PixelFormat_Integer_1B: // 8-bit, 0..255.
				return img.changePixelFormat_I2B_to_I1B[T2]()

			case PixelFormat_Integer_2B: // 16-bit, 0..65535.
				return nil, errors.ErrorFromText(Err_NoChangeIsNeeded)

			case PixelFormat_Float_4B: // 32-bit, 0..1.
				return img.changePixelFormat_I2B_to_F4B[T2]()

			default:
				return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
			}
		}

	case PixelFormat_Float_4B: // 32-bit, 0..1.
		{
			switch channelPixelFormat { // Destination format.
			case PixelFormat_Integer_1B: // 8-bit, 0..255.
				return img.changePixelFormat_F4B_to_I1B[T2]()

			case PixelFormat_Integer_2B: // 16-bit, 0..65535.
				return img.changePixelFormat_F4B_to_I2B[T2]()

			case PixelFormat_Float_4B: // 32-bit, 0..1.
				return nil, errors.ErrorFromText(Err_NoChangeIsNeeded)

			default:
				return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
			}
		}

	default:
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}
}
func (img *Image[T]) changePixelFormat_I1B_to_I2B[T2 IPixel]() (copy *Image[T2], err error) {
	if img.Settings.ChannelPixelFormat != PixelFormat_Integer_1B {
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}

	copy, err = img.MakeCopyT[T2]()
	if err != nil {
		return nil, err
	}

	srcAlphaChannel := img.ChannelData[img.ChannelIndex.Alpha]
	srcRedChannel := img.ChannelData[img.ChannelIndex.Red]
	srcGreenChannel := img.ChannelData[img.ChannelIndex.Green]
	srcBlueChannel := img.ChannelData[img.ChannelIndex.Blue]

	dstAlphaChannel := copy.ChannelData[copy.ChannelIndex.Alpha]
	dstRedChannel := copy.ChannelData[copy.ChannelIndex.Red]
	dstGreenChannel := copy.ChannelData[copy.ChannelIndex.Green]
	dstBlueChannel := copy.ChannelData[copy.ChannelIndex.Blue]

	for y, alpha_row := range srcAlphaChannel.Canvas {
		for x, _ := range alpha_row {
			dstAlphaChannel.Canvas[y][x] = T2(ConvertColour_8bitTo16bit(uint8(srcAlphaChannel.Canvas[y][x])))
			dstRedChannel.Canvas[y][x] = T2(ConvertColour_8bitTo16bit(uint8(srcRedChannel.Canvas[y][x])))
			dstGreenChannel.Canvas[y][x] = T2(ConvertColour_8bitTo16bit(uint8(srcGreenChannel.Canvas[y][x])))
			dstBlueChannel.Canvas[y][x] = T2(ConvertColour_8bitTo16bit(uint8(srcBlueChannel.Canvas[y][x])))
		}
	}

	img.Settings.ChannelPixelFormat = PixelFormat_Integer_2B

	return copy, nil
}
func (img *Image[T]) changePixelFormat_I1B_to_F4B[T2 IPixel]() (copy *Image[T2], err error) {
	if img.Settings.ChannelPixelFormat != PixelFormat_Integer_1B {
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}

	divider := float32(math.MaxUint8) // 8-bit.

	return img.changePixelFormat_Integer_to_F4B[T2](divider)
}
func (img *Image[T]) changePixelFormat_I2B_to_I1B[T2 IPixel]() (copy *Image[T2], err error) {
	if img.Settings.ChannelPixelFormat != PixelFormat_Integer_2B {
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}

	copy, err = img.MakeCopyT[T2]()
	if err != nil {
		return nil, err
	}

	srcAlphaChannel := img.ChannelData[img.ChannelIndex.Alpha]
	srcRedChannel := img.ChannelData[img.ChannelIndex.Red]
	srcGreenChannel := img.ChannelData[img.ChannelIndex.Green]
	srcBlueChannel := img.ChannelData[img.ChannelIndex.Blue]

	dstAlphaChannel := copy.ChannelData[copy.ChannelIndex.Alpha]
	dstRedChannel := copy.ChannelData[copy.ChannelIndex.Red]
	dstGreenChannel := copy.ChannelData[copy.ChannelIndex.Green]
	dstBlueChannel := copy.ChannelData[copy.ChannelIndex.Blue]

	for y, alpha_row := range srcAlphaChannel.Canvas {
		for x, _ := range alpha_row {
			dstAlphaChannel.Canvas[y][x] = T2(ConvertColour_16bitTo8bit(uint16(srcAlphaChannel.Canvas[y][x])))
			dstRedChannel.Canvas[y][x] = T2(ConvertColour_16bitTo8bit(uint16(srcRedChannel.Canvas[y][x])))
			dstGreenChannel.Canvas[y][x] = T2(ConvertColour_16bitTo8bit(uint16(srcGreenChannel.Canvas[y][x])))
			dstBlueChannel.Canvas[y][x] = T2(ConvertColour_16bitTo8bit(uint16(srcBlueChannel.Canvas[y][x])))
		}
	}

	img.Settings.ChannelPixelFormat = PixelFormat_Integer_1B

	return copy, nil
}
func (img *Image[T]) changePixelFormat_I2B_to_F4B[T2 IPixel]() (copy *Image[T2], err error) {
	if img.Settings.ChannelPixelFormat != PixelFormat_Integer_2B {
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}

	divider := float32(math.MaxUint16) // 16-bit.

	return img.changePixelFormat_Integer_to_F4B[T2](divider)
}
func (img *Image[T]) changePixelFormat_Integer_to_F4B[T2 IPixel](divider float32) (copy *Image[T2], err error) {
	copy, err = img.MakeCopyT[T2]()
	if err != nil {
		return nil, err
	}

	srcAlphaChannel := img.ChannelData[img.ChannelIndex.Alpha]
	srcRedChannel := img.ChannelData[img.ChannelIndex.Red]
	srcGreenChannel := img.ChannelData[img.ChannelIndex.Green]
	srcBlueChannel := img.ChannelData[img.ChannelIndex.Blue]

	dstAlphaChannel := copy.ChannelData[copy.ChannelIndex.Alpha]
	dstRedChannel := copy.ChannelData[copy.ChannelIndex.Red]
	dstGreenChannel := copy.ChannelData[copy.ChannelIndex.Green]
	dstBlueChannel := copy.ChannelData[copy.ChannelIndex.Blue]

	for y, alpha_row := range srcAlphaChannel.Canvas {
		for x, _ := range alpha_row {
			dstAlphaChannel.Canvas[y][x] = T2(float32(srcAlphaChannel.Canvas[y][x]) / divider)
			dstRedChannel.Canvas[y][x] = T2(float32(srcRedChannel.Canvas[y][x]) / divider)
			dstGreenChannel.Canvas[y][x] = T2(float32(srcGreenChannel.Canvas[y][x]) / divider)
			dstBlueChannel.Canvas[y][x] = T2(float32(srcBlueChannel.Canvas[y][x]) / divider)
		}
	}

	img.Settings.ChannelPixelFormat = PixelFormat_Float_4B

	return copy, nil
}
func (img *Image[T]) changePixelFormat_F4B_to_I1B[T2 IPixel]() (copy *Image[T2], err error) {
	if img.Settings.ChannelPixelFormat != PixelFormat_Float_4B {
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}

	multiplier := float32(math.MaxUint8) // 8-bit.

	copy, err = img.changePixelFormat_F4B_to_Integer[T2](multiplier)
	if err != nil {
		return nil, err
	}

	img.Settings.ChannelPixelFormat = PixelFormat_Integer_1B

	return copy, nil
}
func (img *Image[T]) changePixelFormat_F4B_to_I2B[T2 IPixel]() (copy *Image[T2], err error) {
	if img.Settings.ChannelPixelFormat != PixelFormat_Float_4B {
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}

	multiplier := float32(math.MaxUint16) // 16-bit.

	copy, err = img.changePixelFormat_F4B_to_Integer[T2](multiplier)
	if err != nil {
		return nil, err
	}

	img.Settings.ChannelPixelFormat = PixelFormat_Integer_2B

	return copy, nil
}
func (img *Image[T]) changePixelFormat_F4B_to_Integer[T2 IPixel](multiplier float32) (copy *Image[T2], err error) {
	copy, err = img.MakeCopyT[T2]()
	if err != nil {
		return nil, err
	}

	srcAlphaChannel := img.ChannelData[img.ChannelIndex.Alpha]
	srcRedChannel := img.ChannelData[img.ChannelIndex.Red]
	srcGreenChannel := img.ChannelData[img.ChannelIndex.Green]
	srcBlueChannel := img.ChannelData[img.ChannelIndex.Blue]

	dstAlphaChannel := copy.ChannelData[copy.ChannelIndex.Alpha]
	dstRedChannel := copy.ChannelData[copy.ChannelIndex.Red]
	dstGreenChannel := copy.ChannelData[copy.ChannelIndex.Green]
	dstBlueChannel := copy.ChannelData[copy.ChannelIndex.Blue]

	for y, alpha_row := range srcAlphaChannel.Canvas {
		for x, _ := range alpha_row {
			dstAlphaChannel.Canvas[y][x] = T2(ConvertColour_NormalisedToInteger(float32(srcAlphaChannel.Canvas[y][x]), multiplier))
			dstRedChannel.Canvas[y][x] = T2(ConvertColour_NormalisedToInteger(float32(srcRedChannel.Canvas[y][x]), multiplier))
			dstGreenChannel.Canvas[y][x] = T2(ConvertColour_NormalisedToInteger(float32(srcGreenChannel.Canvas[y][x]), multiplier))
			dstBlueChannel.Canvas[y][x] = T2(ConvertColour_NormalisedToInteger(float32(srcBlueChannel.Canvas[y][x]), multiplier))
		}
	}

	return copy, nil
}
