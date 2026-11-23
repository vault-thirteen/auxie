package mcci

const (
	Err_UnsupportedChannelPixelFormat = "unsupported channel pixel format"
)

type ChannelPixelFormat byte

const (
	PixelFormat_Unknown = ChannelPixelFormat(0)

	// PixelFormat_Integer_1B
	// Each channel pixel is stored as an 8-bit integer number.
	// Each separate colour channel value uses 1 byte of data.
	PixelFormat_Integer_1B = ChannelPixelFormat(1)

	// PixelFormat_Integer_2B
	// Each channel pixel is stored as an 16-bit integer number.
	// Each separate colour channel value uses 2 bytes of data.
	PixelFormat_Integer_2B = ChannelPixelFormat(2)

	// PixelFormat_Float_4B
	// Each channel pixel is stored as an IEEE 754 floating-point number.
	// Each separate colour channel value uses 4 bytes of data.
	PixelFormat_Float_4B = ChannelPixelFormat(3)
)

func (cpf ChannelPixelFormat) IsValid() bool {
	switch cpf {
	case PixelFormat_Integer_1B,
		PixelFormat_Integer_2B,
		PixelFormat_Float_4B:
		return true

	default:
		return false
	}
}
