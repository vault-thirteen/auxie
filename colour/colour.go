package c

import "image/color"

func GetColourModelType(cm color.Model) (cmt ColourModelType) {
	switch cm {
	case color.RGBAModel:
		return ColourModelType_RGBA
	case color.RGBA64Model:
		return ColourModelType_RGBA64
	case color.NRGBAModel:
		return ColourModelType_NRGBA
	case color.NRGBA64Model:
		return ColourModelType_NRGBA64
	case color.AlphaModel:
		return ColourModelType_Alpha
	case color.Alpha16Model:
		return ColourModelType_Alpha16
	case color.GrayModel:
		return ColourModelType_Gray
	case color.Gray16Model:
		return ColourModelType_Gray16
	case color.YCbCrModel:
		return ColourModelType_YCbCr
	case color.NYCbCrAModel:
		return ColourModelType_NYCbCrA
	case color.CMYKModel:
		return ColourModelType_CMYK
	default:
		return ColourModelType_Unknown
	}
}
