package mcci

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"strings"

	c "github.com/vault-thirteen/auxie/colour"
	"github.com/vault-thirteen/auxie/errors"
	e "github.com/vault-thirteen/auxie/os/win32/enum"
)

const (
	Err_DimensionsError = "dimensions error"

	Errf_ChannelIsNotFound = "channel is not found, name=%s, waveLength=%v"
)

type Image[T IPixel] struct {
	Settings     *ImageSettings
	Channels     []Channel
	ChannelIndex ChannelIndex
	ChannelData  []ChannelDatum[T]
}

func NewImageFromImage(src image.Image, channelPixelFormat ChannelPixelFormat, alphaChannelMode AlphaChannelMode) (img any, err error) {
	if !channelPixelFormat.IsValid() {
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}
	if !alphaChannelMode.IsValid() {
		return nil, errors.ErrorFromText(Err_UnsupportedAlphaChannelMode)
	}

	width := src.Bounds().Dx()
	height := src.Bounds().Dy()

	if (width > math.MaxUint16) || (height > math.MaxUint16) {
		return nil, errors.ErrorFromText(Err_DimensionsError)
	}

	settings := &ImageSettings{
		FileFormatRevision: FileFormatRevision_1,
		Width:              uint16(width),
		Height:             uint16(height),
		ChannelCount:       4,
		ChannelPixelFormat: channelPixelFormat,
		AlphaChannelMode:   alphaChannelMode,
	}

	switch channelPixelFormat {
	case PixelFormat_Float_4B:
		return newImageFromImage[float32](src, settings)

	case PixelFormat_Integer_2B:
		return newImageFromImage[uint16](src, settings)

	case PixelFormat_Integer_1B:
		return newImageFromImage[uint8](src, settings)

	default:
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}
}
func newImageFromImage[T IPixel](src image.Image, settings *ImageSettings) (img *Image[T], err error) {
	img = &Image[T]{
		Settings: settings,
	}

	img.initRgbaChannels()

	err = img.initRgbaChannelIndices()
	if err != nil {
		return nil, err
	}

	err = img.fillChannelData(src)
	if err != nil {
		return nil, err
	}

	return img, nil
}
func (img *Image[T]) initRgbaChannels() {
	channels := []Channel{
		{
			Name:          ChannelName_Red,
			WaveLength_Nm: ChannelWaveLengthNm_Red,
		},
		{
			Name:          ChannelName_Green,
			WaveLength_Nm: ChannelWaveLengthNm_Green,
		},
		{
			Name:          ChannelName_Blue,
			WaveLength_Nm: ChannelWaveLengthNm_Blue,
		},
		{
			Name:          ChannelName_Alpha,
			WaveLength_Nm: WaveLengthNm_None,
		},
	}

	img.Channels = channels
}
func (img *Image[T]) initRgbaChannelIndices() (err error) {
	channelIndex := ChannelIndex{}

	channelIndex.Red, err = img.findChannelIndex(ChannelName_Red, ChannelWaveLengthNm_Red)
	if err != nil {
		return err
	}

	channelIndex.Green, err = img.findChannelIndex(ChannelName_Green, ChannelWaveLengthNm_Green)
	if err != nil {
		return err
	}

	channelIndex.Blue, err = img.findChannelIndex(ChannelName_Blue, ChannelWaveLengthNm_Blue)
	if err != nil {
		return err
	}

	channelIndex.Alpha, err = img.findChannelIndex(ChannelName_Alpha, WaveLengthNm_None)
	if err != nil {
		return err
	}

	img.ChannelIndex = channelIndex

	return nil
}
func (img *Image[T]) findChannelIndex(name string, waveLengthNm float32) (idx int, err error) {
	nameLC := strings.ToLower(name)

	for i, ch := range img.Channels {
		if strings.ToLower(ch.Name) == nameLC {
			return i, nil
		}
		if (waveLengthNm-WaveLengthNm_Tolerance <= ch.WaveLength_Nm) &&
			(ch.WaveLength_Nm <= waveLengthNm+WaveLengthNm_Tolerance) {
			return i, nil
		}
	}

	return ChannelIndexUnknown, errors.ErrorFromFormattedText(Errf_ChannelIsNotFound, name, waveLengthNm)
}
func (img *Image[T]) fillChannelData(src image.Image) (err error) {
	srcColourModel := c.GetColourModelType(src.ColorModel())
	switch srcColourModel {
	case c.ColourModelType_RGBA: // Pre-multiplied, 8-bit.
		return img.fillChannelData_RGBA(src)

	case c.ColourModelType_NRGBA: // Straight, 8-bit.
		return img.fillChannelData_NRGBA(src)

	default: // Pre-multiplied, 16-bit.
		return img.fillChannelData_Common(src)
	}
}
func (img *Image[T]) fillChannelData_RGBA(src image.Image) (err error) {
	var cd []ChannelDatum[T]
	var cdRed, cdGreen, cdBlue, cdAlpha ChannelDatum[T]
	cd, cdRed, cdGreen, cdBlue, cdAlpha = img.prepareChannelData()

	yLast := int(img.Settings.Height - 1)
	xLast := int(img.Settings.Width - 1)
	var rowRed, rowGreen, rowBlue, rowAlpha Row[T]
	var abstractPixel color.Color
	var pd *PixelData
	var p Pixel[T]

	// Read rows from bottom to top.
	for y := yLast; y >= 0; y-- {
		rowRed = make(Row[T], 0, img.Settings.Width)
		rowGreen = make(Row[T], 0, img.Settings.Width)
		rowBlue = make(Row[T], 0, img.Settings.Width)
		rowAlpha = make(Row[T], 0, img.Settings.Width)

		for x := 0; x <= xLast; x++ {
			abstractPixel = src.At(x, y)

			pd, err = GetPixelData_RGBA(abstractPixel)
			if err != nil {
				return err
			}

			switch img.Settings.ChannelPixelFormat {
			case PixelFormat_Integer_1B:
				p = MakePixel_I1B_8bit[T](pd, img.Settings.AlphaChannelMode)

			case PixelFormat_Integer_2B:
				p = MakePixel_I2B_8bit[T](pd, img.Settings.AlphaChannelMode)

			case PixelFormat_Float_4B:
				p = MakePixel_F4B_8bit[T](pd, img.Settings.AlphaChannelMode)
			}

			rowRed = append(rowRed, p.R)
			rowGreen = append(rowGreen, p.G)
			rowBlue = append(rowBlue, p.B)
			rowAlpha = append(rowAlpha, p.A)
		}

		cdRed.Canvas = append(cdRed.Canvas, rowRed)
		cdGreen.Canvas = append(cdGreen.Canvas, rowGreen)
		cdBlue.Canvas = append(cdBlue.Canvas, rowBlue)
		cdAlpha.Canvas = append(cdAlpha.Canvas, rowAlpha)
	}

	cd = append(cd, cdRed, cdGreen, cdBlue, cdAlpha)
	img.ChannelData = cd

	return
}
func (img *Image[T]) fillChannelData_NRGBA(src image.Image) (err error) {
	var cd []ChannelDatum[T]
	var cdRed, cdGreen, cdBlue, cdAlpha ChannelDatum[T]
	cd, cdRed, cdGreen, cdBlue, cdAlpha = img.prepareChannelData()

	yLast := int(img.Settings.Height - 1)
	xLast := int(img.Settings.Width - 1)
	var rowRed, rowGreen, rowBlue, rowAlpha Row[T]
	var abstractPixel color.Color
	var pd *PixelData
	var p Pixel[T]

	// Read rows from bottom to top.
	for y := yLast; y >= 0; y-- {
		rowRed = make(Row[T], 0, img.Settings.Width)
		rowGreen = make(Row[T], 0, img.Settings.Width)
		rowBlue = make(Row[T], 0, img.Settings.Width)
		rowAlpha = make(Row[T], 0, img.Settings.Width)

		for x := 0; x <= xLast; x++ {
			abstractPixel = src.At(x, y)

			pd, err = GetPixelData_NRGBA(abstractPixel)
			if err != nil {
				return err
			}

			switch img.Settings.ChannelPixelFormat {
			case PixelFormat_Integer_1B:
				p = MakePixel_I1B_8bit[T](pd, img.Settings.AlphaChannelMode)

			case PixelFormat_Integer_2B:
				p = MakePixel_I2B_8bit[T](pd, img.Settings.AlphaChannelMode)

			case PixelFormat_Float_4B:
				p = MakePixel_F4B_8bit[T](pd, img.Settings.AlphaChannelMode)
			}

			rowRed = append(rowRed, p.R)
			rowGreen = append(rowGreen, p.G)
			rowBlue = append(rowBlue, p.B)
			rowAlpha = append(rowAlpha, p.A)
		}

		cdRed.Canvas = append(cdRed.Canvas, rowRed)
		cdGreen.Canvas = append(cdGreen.Canvas, rowGreen)
		cdBlue.Canvas = append(cdBlue.Canvas, rowBlue)
		cdAlpha.Canvas = append(cdAlpha.Canvas, rowAlpha)
	}

	cd = append(cd, cdRed, cdGreen, cdBlue, cdAlpha)
	img.ChannelData = cd

	return
}
func (img *Image[T]) fillChannelData_Common(src image.Image) (err error) {
	var cd []ChannelDatum[T]
	var cdRed, cdGreen, cdBlue, cdAlpha ChannelDatum[T]
	cd, cdRed, cdGreen, cdBlue, cdAlpha = img.prepareChannelData()

	yLast := int(img.Settings.Height - 1)
	xLast := int(img.Settings.Width - 1)
	var rowRed, rowGreen, rowBlue, rowAlpha Row[T]
	var abstractPixel color.Color
	var pd *PixelData
	var p Pixel[T]

	// Read rows from bottom to top.
	for y := yLast; y >= 0; y-- {
		rowRed = make(Row[T], 0, img.Settings.Width)
		rowGreen = make(Row[T], 0, img.Settings.Width)
		rowBlue = make(Row[T], 0, img.Settings.Width)
		rowAlpha = make(Row[T], 0, img.Settings.Width)

		for x := 0; x <= xLast; x++ {
			abstractPixel = src.At(x, y)

			pd, err = GetPixelData_Abstract(abstractPixel)
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
			}

			rowRed = append(rowRed, p.R)
			rowGreen = append(rowGreen, p.G)
			rowBlue = append(rowBlue, p.B)
			rowAlpha = append(rowAlpha, p.A)
		}

		cdRed.Canvas = append(cdRed.Canvas, rowRed)
		cdGreen.Canvas = append(cdGreen.Canvas, rowGreen)
		cdBlue.Canvas = append(cdBlue.Canvas, rowBlue)
		cdAlpha.Canvas = append(cdAlpha.Canvas, rowAlpha)
	}

	cd = append(cd, cdRed, cdGreen, cdBlue, cdAlpha)
	img.ChannelData = cd

	return
}
func (img *Image[T]) prepareChannelData() (cd []ChannelDatum[T], cdRed, cdGreen, cdBlue, cdAlpha ChannelDatum[T]) {
	cd = make([]ChannelDatum[T], 0, 4)

	cdRed = ChannelDatum[T]{
		Channel: &img.Channels[img.ChannelIndex.Red],
		Canvas:  make([]Row[T], 0, img.Settings.Height),
	}
	cdGreen = ChannelDatum[T]{
		Channel: &img.Channels[img.ChannelIndex.Green],
		Canvas:  make([]Row[T], 0, img.Settings.Height),
	}
	cdBlue = ChannelDatum[T]{
		Channel: &img.Channels[img.ChannelIndex.Blue],
		Canvas:  make([]Row[T], 0, img.Settings.Height),
	}
	cdAlpha = ChannelDatum[T]{
		Channel: &img.Channels[img.ChannelIndex.Alpha],
		Canvas:  make([]Row[T], 0, img.Settings.Height),
	}

	return cd, cdRed, cdGreen, cdBlue, cdAlpha
}

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

	return nil
}

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

	return copy, nil
}
func (img *Image[T]) changePixelFormat_F4B_to_I1B[T2 IPixel]() (copy *Image[T2], err error) {
	if img.Settings.ChannelPixelFormat != PixelFormat_Float_4B {
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}

	multiplier := float32(math.MaxUint8) // 8-bit.

	return img.changePixelFormat_F4B_to_Integer[T2](multiplier)
}
func (img *Image[T]) changePixelFormat_F4B_to_I2B[T2 IPixel]() (copy *Image[T2], err error) {
	if img.Settings.ChannelPixelFormat != PixelFormat_Float_4B {
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}

	multiplier := float32(math.MaxUint16) // 16-bit.

	return img.changePixelFormat_F4B_to_Integer[T2](multiplier)
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
				return nil, fmt.Errorf(Err_UnsupportedAlphaChannelMode)
			}
		}

	default:
		return nil, fmt.Errorf(Err_UnsupportedPixelFormat)
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
	cópy, err = img.MakeCopy()
	if err != nil {
		return nil, err
	}

	if cópy.Settings.AlphaChannelMode != AlphaChannelMode_Premultiplied {
		err = cópy.ChangeAlphaMode(AlphaChannelMode_Premultiplied)
		if err != nil {
			return nil, err
		}
	}

	var cópy2 *Image[uint8]
	if cópy.Settings.ChannelPixelFormat != PixelFormat_Integer_1B {
		cópy2, err = img.ChangePixelFormatT[uint8](PixelFormat_Integer_1B)
		if err != nil {
			return nil, err
		}
	} else {
		// Pixel format is PixelFormat_Integer_1B => T is uint8.
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
