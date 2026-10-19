package bom

const (
	EncodingUTF8       = Encoding(1)  // UTF-8 Encoding.
	EncodingUTF16BE    = Encoding(2)  // UTF-16 (BE, Big Endian) Encoding.
	EncodingUTF16LE    = Encoding(3)  // UTF-16 (LE, Little Endian) Encoding.
	EncodingUTF32BE    = Encoding(4)  // UTF-32 (BE, Big Endian) Encoding.
	EncodingUTF32LE    = Encoding(5)  // UTF-32 (LE, Little Endian) Encoding.
	EncodingUTF7       = Encoding(6)  // UTF-7 Encoding.
	EncodingUTF1       = Encoding(8)  // UTF-1 Encoding.
	EncodingUTF_EBCDIC = Encoding(9)  // UTF-EBCDIC Encoding.
	EncodingSCSU       = Encoding(10) // SCSU Encoding.
	EncodingBOCU1      = Encoding(11) // BOCU-1 Encoding.
	EncodingGB18030    = Encoding(12) // GB18030 Encoding.
)

// Encoding is an encoding type.
// Usually it is a text encoding using Unicode symbols.
// Unicode on Wikipedia:
// https://en.wikipedia.org/wiki/Unicode
type Encoding byte

// possibleEncodings is a list of encodings known to this library.
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
