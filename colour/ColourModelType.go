package c

type ColourModelType uint32

const (
	ColourModelType_Unknown ColourModelType = 0
	ColourModelType_RGBA    ColourModelType = 1
	ColourModelType_RGBA64  ColourModelType = 2
	ColourModelType_NRGBA   ColourModelType = 3
	ColourModelType_NRGBA64 ColourModelType = 4
	ColourModelType_Alpha   ColourModelType = 5
	ColourModelType_Alpha16 ColourModelType = 6
	ColourModelType_Gray    ColourModelType = 7
	ColourModelType_Gray16  ColourModelType = 8
	ColourModelType_YCbCr   ColourModelType = 9
	ColourModelType_NYCbCrA ColourModelType = 10
	ColourModelType_CMYK    ColourModelType = 11
)
