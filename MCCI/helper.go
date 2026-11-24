package mcci

func ConvertColor8bitTo16bit(c8 uint8) (c16 uint16) {
	x := uint16(c8)
	return (x << 8) | x
}
func ConvertColor16bitTo8bit(c16 uint16) (c8 uint8) {
	return uint8(c16 >> 8)
}
