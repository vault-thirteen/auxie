package mcci

import (
	"github.com/vault-thirteen/auxie/errors"
	e "github.com/vault-thirteen/auxie/os/win32/enum"
)

func (img *Image[T]) ExportToDirect2D(pixelFormat e.DXGI_FORMAT, alphaMode e.D2D1_ALPHA_MODE) (out []byte, err error) {
	// Formats supported by Direct2D bitmaps, i.e. the ID2D1Bitmap type, are listed at the following link:
	// https: //learn.microsoft.com/en-gb/windows/win32/direct2d/supported-pixel-formats-and-alpha-modes#specifying-a-pixel-format-for-an-id2d1bitmap
	// From all those formats only the DXGI_FORMAT_B8G8R8A8_UNORM seems to be acceptable by Go language.
	switch pixelFormat {
	case e.DXGI_FORMAT_B8G8R8A8_UNORM:
		{
			switch alphaMode {
			case e.D2D1_ALPHA_MODE_PREMULTIPLIED:
				return img.exportToDirect2D_BGRA_Premultiplied()

			default:
				return nil, errors.ErrorFromText(Err_UnsupportedAlphaChannelMode)
			}
		}

	default:
		return nil, errors.ErrorFromText(Err_UnsupportedPixelFormat)
	}
}
func (img *Image[T]) exportToDirect2D_BGRA_Premultiplied() (out []byte, err error) {
	// To export for Direct2D we must:
	// - make an exact copy of the object to save original data;
	// - change alpha channel mode (if needed);
	// - change pixel format (if needed);
	// - reverse rows order;
	// - swap colour channels for BGRA;
	// - convert the 2D canvas pixel array into a 1D array of bytes.

	var cópy *Image[T]
	if img.Settings.AlphaChannelMode != AlphaChannelMode_Premultiplied {
		// Copy the image if alpha channel mode is straight.
		// Premultiplied mode destroys information irreversibly !
		cópy, err = img.MakeCopy()
		if err != nil {
			return nil, err
		}

		err = cópy.ChangeAlphaMode(AlphaChannelMode_Premultiplied)
		if err != nil {
			return nil, err
		}
	} else {
		// Do not copy image if alpha channel mode is already premultiplied.
		cópy = img
	}

	var cópy2 *Image[uint8]
	if cópy.Settings.ChannelPixelFormat != PixelFormat_Integer_1B {
		// Create a copy with different pixel format.
		cópy2, err = img.ChangePixelFormatT[uint8](PixelFormat_Integer_1B)
		if err != nil {
			return nil, err
		}
	} else {
		// Pixel format is already 'PixelFormat_Integer_1B' => T is uint8.
		// Use type casting and avoid copying.
		var probe T
		switch any(probe).(type) {
		case uint8:
			cópy2 = any(cópy).(*Image[uint8])
		default:
			return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
		}
	}

	// Collect bytes.
	alphaChannel := cópy2.ChannelData[cópy2.ChannelIndex.Alpha]
	redChannel := cópy2.ChannelData[cópy2.ChannelIndex.Red]
	greenChannel := cópy2.ChannelData[cópy2.ChannelIndex.Green]
	blueChannel := cópy2.ChannelData[cópy2.ChannelIndex.Blue]

	width := cópy2.Settings.Width
	height := cópy2.Settings.Height

	out = make([]byte, 0, width*height*4*1) // Four 8-bit channels per pixel.
	var yFirst = int(height - 1)            // Y: Read rows in reverse order.
	var xLast = int(width - 1)              // X: Left to Right.
	for y := yFirst; y >= 0; y-- {
		for x := 0; x <= xLast; x++ {
			// Swap channels to get BGRA sequence.
			out = append(out, blueChannel.Canvas[y][x])
			out = append(out, greenChannel.Canvas[y][x])
			out = append(out, redChannel.Canvas[y][x])
			out = append(out, alphaChannel.Canvas[y][x])
		}
	}

	return out, nil
}
