package bom

const (
	EncodingUnknown    = Encoding(0)
	EncodingUTF8       = Encoding(1)
	EncodingUTF16BE    = Encoding(2)
	EncodingUTF16LE    = Encoding(3)
	EncodingUTF32BE    = Encoding(4)
	EncodingUTF32LE    = Encoding(5)
	EncodingUTF7       = Encoding(6)
	EncodingUTF1       = Encoding(8)
	EncodingUTF_EBCDIC = Encoding(9)
	EncodingSCSU       = Encoding(10)
	EncodingBOCU1      = Encoding(11)
	EncodingGB18030    = Encoding(12)
)

type Encoding byte

var possibleEncodings = []Encoding{
	EncodingUTF8,
	EncodingUTF16BE,
	EncodingUTF16LE,
	EncodingUTF32BE,
	EncodingUTF32LE,
	EncodingUTF7,
	EncodingUTF1,
	EncodingUTF_EBCDIC,
	EncodingSCSU,
	EncodingBOCU1,
	EncodingGB18030,
}

// PossibleEncodings returns a list of possible encodings except the unknown
// encoding.
func PossibleEncodings() []Encoding {
	return possibleEncodings
}

// IsKnown tells whether the encoding is known or not.
func (e Encoding) IsKnown() bool {
	return e != EncodingUnknown
}

// IsUnknown tells whether the encoding is unknown or not.
func (e Encoding) IsUnknown() bool {
	return e == EncodingUnknown
}
