package mcci

import (
	"image"
	"image/color"
	"math"
	"strings"

	c "github.com/vault-thirteen/auxie/colour"
	"github.com/vault-thirteen/auxie/errors"
)

const (
	Err_DimensionsError        = "dimensions error"
	Err_UnsupportedImageFormat = "unsupported image format"
	Err_InvalidPixelFormat     = "invalid pixel format"

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

	channelData := make([]ChannelDatum[T], 0, 4)

	cdRed := ChannelDatum[T]{
		Channel: &img.Channels[img.ChannelIndex.Red],
		Canvas:  make([]Row[T], 0, img.Settings.Height),
	}
	cdGreen := ChannelDatum[T]{
		Channel: &img.Channels[img.ChannelIndex.Green],
		Canvas:  make([]Row[T], 0, img.Settings.Height),
	}
	cdBlue := ChannelDatum[T]{
		Channel: &img.Channels[img.ChannelIndex.Blue],
		Canvas:  make([]Row[T], 0, img.Settings.Height),
	}
	cdAlpha := ChannelDatum[T]{
		Channel: &img.Channels[img.ChannelIndex.Alpha],
		Canvas:  make([]Row[T], 0, img.Settings.Height),
	}

	yLast := int(img.Settings.Height - 1)
	xLast := int(img.Settings.Width - 1)
	var rowRed, rowGreen, rowBlue, rowAlpha Row[T]
	var abstractPixel color.Color
	var ok bool
	var pixel_NRGBA color.NRGBA
	var pixel_RGBA color.RGBA
	var r_straight, g_straight, b_straight float32
	var r_premultiplied, g_premultiplied, b_premultiplied float32
	var alpha, alpha_normalised float32

	// Read rows from bottom to top.
	for y := yLast; y >= 0; y-- {
		rowRed = make(Row[T], 0, img.Settings.Width)
		rowGreen = make(Row[T], 0, img.Settings.Width)
		rowBlue = make(Row[T], 0, img.Settings.Width)
		rowAlpha = make(Row[T], 0, img.Settings.Width)

		for x := 0; x <= xLast; x++ {
			abstractPixel = src.At(x, y)

			// Part I. Get colour values (straight and pre-multiplied).
			{
				switch srcColourModel {
				case c.ColourModelType_RGBA: // Pre-multiplied, 8-bit.
					{
						pixel_RGBA, ok = abstractPixel.(color.RGBA)
						if !ok {
							return errors.ErrorFromText(Err_InvalidPixelFormat)
						}

						r_premultiplied = float32(pixel_RGBA.R)
						g_premultiplied = float32(pixel_RGBA.G)
						b_premultiplied = float32(pixel_RGBA.B)

						alpha = float32(pixel_RGBA.A)
						alpha_normalised = alpha / float32(math.MaxUint8) // 8-bit.

						if alpha_normalised != 0 {
							r_straight = r_premultiplied / alpha_normalised
							g_straight = g_premultiplied / alpha_normalised
							b_straight = b_premultiplied / alpha_normalised
						} else {
							r_straight = 0
							g_straight = 0
							b_straight = 0
						}
					}

				case c.ColourModelType_NRGBA: // Straight, 8-bit.
					{
						pixel_NRGBA, ok = abstractPixel.(color.NRGBA)
						if !ok {
							return errors.ErrorFromText(Err_InvalidPixelFormat)
						}

						r_straight = float32(pixel_NRGBA.R)
						g_straight = float32(pixel_NRGBA.G)
						b_straight = float32(pixel_NRGBA.B)

						alpha = float32(pixel_NRGBA.A)
						alpha_normalised = alpha / float32(math.MaxUint8) // 8-bit.

						r_premultiplied = r_straight * alpha_normalised
						g_premultiplied = g_straight * alpha_normalised
						b_premultiplied = b_straight * alpha_normalised
					}

				default: // Pre-multiplied, 16-bit.
					{
						r_premultiplied, g_premultiplied, b_premultiplied, alpha = GetRgbaAsFloat32(abstractPixel)
						alpha_normalised = alpha / float32(math.MaxUint16) // 16-bit.

						if alpha_normalised != 0 {
							r_straight = r_premultiplied / alpha_normalised
							g_straight = g_premultiplied / alpha_normalised
							b_straight = b_premultiplied / alpha_normalised
						} else {
							r_straight = 0
							g_straight = 0
							b_straight = 0
						}
					}
				}
			}

			// Part II. Store pixels into selected format.
			{
				switch img.Settings.ChannelPixelFormat {
				case PixelFormat_Float_4B:
					{
						// Normalise values before storing.
						var divider float32
						switch srcColourModel {
						case c.ColourModelType_RGBA: // 8-bit.
							divider = math.MaxUint8
						case c.ColourModelType_NRGBA: // 8-bit.
							divider = math.MaxUint8
						default: // 16-bit.
							divider = math.MaxUint16
						}

						switch img.Settings.AlphaChannelMode {
						case AlphaChannelMode_Straight:
							{
								rowRed = append(rowRed, T(r_straight/divider))
								rowGreen = append(rowGreen, T(g_straight/divider))
								rowBlue = append(rowBlue, T(b_straight/divider))
								rowAlpha = append(rowAlpha, T(alpha_normalised))
							}

						case AlphaChannelMode_Premultiplied:
							{
								rowRed = append(rowRed, T(r_premultiplied/divider))
								rowGreen = append(rowGreen, T(g_premultiplied/divider))
								rowBlue = append(rowBlue, T(b_premultiplied/divider))
								rowAlpha = append(rowAlpha, T(alpha_normalised))
							}
						}
					}

				case PixelFormat_Integer_1B:
					{
						// Store values directly.
						switch img.Settings.AlphaChannelMode {
						case AlphaChannelMode_Straight:
							{
								switch srcColourModel {
								case c.ColourModelType_RGBA, // 8-bit.
									c.ColourModelType_NRGBA: // 8-bit.
									{
										rowRed = append(rowRed, T(r_straight))
										rowGreen = append(rowGreen, T(g_straight))
										rowBlue = append(rowBlue, T(b_straight))
										rowAlpha = append(rowAlpha, T(alpha))
									}
								default: // 16-bit.
									{
										rowRed = append(rowRed, T(ConvertColor16bitTo8bit(uint16(r_straight))))
										rowGreen = append(rowGreen, T(ConvertColor16bitTo8bit(uint16(g_straight))))
										rowBlue = append(rowBlue, T(ConvertColor16bitTo8bit(uint16(b_straight))))
										rowAlpha = append(rowAlpha, T(ConvertColor16bitTo8bit(uint16(alpha))))
									}
								}
							}

						case AlphaChannelMode_Premultiplied:
							{
								switch srcColourModel {
								case c.ColourModelType_RGBA, // 8-bit.
									c.ColourModelType_NRGBA: // 8-bit.
									{
										rowRed = append(rowRed, T(r_premultiplied))
										rowGreen = append(rowGreen, T(g_premultiplied))
										rowBlue = append(rowBlue, T(b_premultiplied))
										rowAlpha = append(rowAlpha, T(alpha))
									}
								default: // 16-bit.
									{
										rowRed = append(rowRed, T(ConvertColor16bitTo8bit(uint16(r_premultiplied))))
										rowGreen = append(rowGreen, T(ConvertColor16bitTo8bit(uint16(g_premultiplied))))
										rowBlue = append(rowBlue, T(ConvertColor16bitTo8bit(uint16(b_premultiplied))))
										rowAlpha = append(rowAlpha, T(ConvertColor16bitTo8bit(uint16(alpha))))
									}
								}
							}
						}
					}

				case PixelFormat_Integer_2B:
					{
						// Store 16-bit values.
						switch img.Settings.AlphaChannelMode {
						case AlphaChannelMode_Straight:
							{
								switch srcColourModel {
								case c.ColourModelType_RGBA, // 8-bit.
									c.ColourModelType_NRGBA: // 8-bit.
									{
										rowRed = append(rowRed, T(ConvertColor8bitTo16bit(uint8(r_straight))))
										rowGreen = append(rowGreen, T(ConvertColor8bitTo16bit(uint8(g_straight))))
										rowBlue = append(rowBlue, T(ConvertColor8bitTo16bit(uint8(b_straight))))
										rowAlpha = append(rowAlpha, T(ConvertColor8bitTo16bit(uint8(alpha))))
									}
								default: // 16-bit.
									{
										rowRed = append(rowRed, T(r_straight))
										rowGreen = append(rowGreen, T(g_straight))
										rowBlue = append(rowBlue, T(b_straight))
										rowAlpha = append(rowAlpha, T(alpha))
									}
								}
							}

						case AlphaChannelMode_Premultiplied:
							{
								switch srcColourModel {
								case c.ColourModelType_RGBA, // 8-bit.
									c.ColourModelType_NRGBA: // 8-bit.
									{
										rowRed = append(rowRed, T(ConvertColor8bitTo16bit(uint8(r_premultiplied))))
										rowGreen = append(rowGreen, T(ConvertColor8bitTo16bit(uint8(g_premultiplied))))
										rowBlue = append(rowBlue, T(ConvertColor8bitTo16bit(uint8(b_premultiplied))))
										rowAlpha = append(rowAlpha, T(ConvertColor8bitTo16bit(uint8(alpha))))
									}
								default: // 16-bit.
									{
										rowRed = append(rowRed, T(r_premultiplied))
										rowGreen = append(rowGreen, T(g_premultiplied))
										rowBlue = append(rowBlue, T(b_premultiplied))
										rowAlpha = append(rowAlpha, T(alpha))
									}
								}
							}
						}
					}
				}
			}
		}

		cdRed.Canvas = append(cdRed.Canvas, rowRed)
		cdGreen.Canvas = append(cdGreen.Canvas, rowGreen)
		cdBlue.Canvas = append(cdBlue.Canvas, rowBlue)
		cdAlpha.Canvas = append(cdAlpha.Canvas, rowAlpha)
	}

	channelData = append(channelData, cdRed, cdGreen, cdBlue, cdAlpha)
	img.ChannelData = channelData

	return
}
