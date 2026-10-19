package reader

import (
	"io"

	bt "github.com/vault-thirteen/auxie/BasicTypes"
)

type IReader interface {
	io.Reader
	GetInternalReader() io.Reader
	ReadLineEndingWithCRLF() (line []byte, err error)
	ReadBytes(size int) (bytes []byte, err error)
	ReadByte() (ub byte, err error)
	ReadSByte() (sb int8, err error)
	Read2Bytes() (bytes []byte, err error)
	Read4Bytes() (bytes []byte, err error)
	Read8Bytes() (bytes []byte, err error)
	ReadWord_BE() (w bt.Word, err error)
	ReadWord_LE() (w bt.Word, err error)
	ReadDWord_BE() (dw bt.DWord, err error)
	ReadDWord_LE() (dw bt.DWord, err error)
	ReadUShort_BE() (us uint16, err error)
	ReadUShort_LE() (us uint16, err error)
	ReadULong_BE() (ul uint32, err error)
	ReadULong_LE() (ul uint32, err error)
	ReadSShort_BE() (ss int16, err error)
	ReadSShort_LE() (ss int16, err error)
	ReadSLong_BE() (sl int32, err error)
	ReadSLong_LE() (sl int32, err error)
	ReadFloat_BE() (f float32, err error)
	ReadFloat_LE() (f float32, err error)
	ReadDouble_BE() (d float64, err error)
	ReadDouble_LE() (d float64, err error)
}
