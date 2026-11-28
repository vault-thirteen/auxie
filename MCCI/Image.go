package mcci

import (
	"image"
	"image/color"
	"os"
	"strings"

	c "github.com/vault-thirteen/auxie/colour"
	"github.com/vault-thirteen/auxie/errors"
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

// NewImageWithSizeAndColour creates a new image with specified dimensions and
// colour.
func NewImageWithSizeAndColour(width Size, height Size, colour color.Color, channelPixelFormat ChannelPixelFormat, alphaChannelMode AlphaChannelMode) (img any, err error) {
	if !channelPixelFormat.IsValid() {
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}
	if !alphaChannelMode.IsValid() {
		return nil, errors.ErrorFromText(Err_UnsupportedAlphaChannelMode)
	}

	settings := &ImageSettings{
		FileFormatRevision: FileFormatRevision_1,
		Width:              width,
		Height:             height,
		ChannelCount:       4,
		ChannelPixelFormat: channelPixelFormat,
		AlphaChannelMode:   alphaChannelMode,
	}

	switch channelPixelFormat {
	case PixelFormat_Float_4B:
		return newImageWithSizeAndColour[float32](colour, settings)

	case PixelFormat_Integer_2B:
		return newImageWithSizeAndColour[uint16](colour, settings)

	case PixelFormat_Integer_1B:
		return newImageWithSizeAndColour[uint8](colour, settings)

	default:
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}
}
func newImageWithSizeAndColour[T IPixel](colour color.Color, settings *ImageSettings) (img *Image[T], err error) {
	img = &Image[T]{
		Settings: settings,
	}

	img.initRgbaChannels()

	err = img.initRgbaChannelIndices()
	if err != nil {
		return nil, err
	}

	err = img.fillChannelDataWithSizeAndColour(colour)
	if err != nil {
		return nil, err
	}

	return img, nil
}
func (img *Image[T]) fillChannelDataWithSizeAndColour(colour color.Color) (err error) {
	var cd []ChannelDatum[T]
	var cdRed, cdGreen, cdBlue, cdAlpha ChannelDatum[T]
	cd, cdRed, cdGreen, cdBlue, cdAlpha = img.prepareChannelData()

	yLast := int(img.Settings.Height - 1)
	xLast := int(img.Settings.Width - 1)
	var rowRed, rowGreen, rowBlue, rowAlpha Row[T]

	// Prepare the pixel.
	var p Pixel[T]
	err = img.PreparePixel(colour, &p)
	if err != nil {
		return err
	}

	// Read rows from bottom to top.
	for y := yLast; y >= 0; y-- {
		rowRed = make(Row[T], 0, img.Settings.Width)
		rowGreen = make(Row[T], 0, img.Settings.Width)
		rowBlue = make(Row[T], 0, img.Settings.Width)
		rowAlpha = make(Row[T], 0, img.Settings.Width)

		for x := 0; x <= xLast; x++ {
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

// NewImageFromImage creates a new image from Go 'image.Image' type object.
// Both 'channelPixelFormat' and 'alphaChannelMode' arguments set the desired
// format into which the source image will be converted.
func NewImageFromImage(src image.Image, channelPixelFormat ChannelPixelFormat, alphaChannelMode AlphaChannelMode) (img any, err error) {
	if !channelPixelFormat.IsValid() {
		return nil, errors.ErrorFromText(Err_UnsupportedChannelPixelFormat)
	}
	if !alphaChannelMode.IsValid() {
		return nil, errors.ErrorFromText(Err_UnsupportedAlphaChannelMode)
	}

	width := src.Bounds().Dx()
	height := src.Bounds().Dy()

	if (width > int(SizeMax)) || (height > int(SizeMax)) {
		return nil, errors.ErrorFromText(Err_DimensionsError)
	}

	settings := &ImageSettings{
		FileFormatRevision: FileFormatRevision_1,
		Width:              Size(width),
		Height:             Size(height),
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
	var p Pixel[T]

	// Read rows from bottom to top.
	for y := yLast; y >= 0; y-- {
		rowRed = make(Row[T], 0, img.Settings.Width)
		rowGreen = make(Row[T], 0, img.Settings.Width)
		rowBlue = make(Row[T], 0, img.Settings.Width)
		rowAlpha = make(Row[T], 0, img.Settings.Width)

		for x := 0; x <= xLast; x++ {
			// Prepare the pixel.
			err = img.PreparePixel(src.At(x, y), &p)
			if err != nil {
				return err
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

// NewImageFromFile creates a new image from file.
// Both 'channelPixelFormat' and 'alphaChannelMode' arguments set the desired
// format into which the source image will be converted.
func NewImageFromFile(filePath string, channelPixelFormat ChannelPixelFormat, alphaChannelMode AlphaChannelMode) (img any, err error) {
	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		derr := f.Close()
		err = errors.Combine(err, derr)
	}()

	var goImg image.Image
	goImg, _, err = image.Decode(f)
	if err != nil {
		return nil, err
	}

	return NewImageFromImage(goImg, channelPixelFormat, alphaChannelMode)
}

func (img *Image[T]) Area() *Area[Size] {
	return NewAreaFromDimensions[Size](img.Settings.Width, img.Settings.Height)
}
