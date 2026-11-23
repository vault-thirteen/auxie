package mcci

func ConvertColor8bitTo16bit(c8 uint16) (c16 uint16) {
	return (c8 << 8) | c8
}
