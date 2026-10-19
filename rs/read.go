package rs

import (
	bt "github.com/vault-thirteen/auxie/BasicTypes"
)

// Read is a wrapper of the child's method.
func (rs *ReaderSeeker) Read(dst []byte) (n int, err error) {
	return rs.ir.Read(dst)
}

// ReadLineEndingWithCRLF is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadLineEndingWithCRLF() (line []byte, err error) {
	return rs.ir.ReadLineEndingWithCRLF()
}

// ReadBytes is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadBytes(size int) (bytes []byte, err error) {
	return rs.ir.ReadBytes(size)
}

// ReadByte is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadByte() (ub byte, err error) {
	return rs.ir.ReadByte()
}

// ReadSByte is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadSByte() (sb int8, err error) {
	return rs.ir.ReadSByte()
}

// Read2Bytes is a wrapper of the child's method.
func (rs *ReaderSeeker) Read2Bytes() (bytes []byte, err error) {
	return rs.ir.Read2Bytes()
}

// Read4Bytes is a wrapper of the child's method.
func (rs *ReaderSeeker) Read4Bytes() (bytes []byte, err error) {
	return rs.ir.Read4Bytes()
}

// Read8Bytes is a wrapper of the child's method.
func (rs *ReaderSeeker) Read8Bytes() (bytes []byte, err error) {
	return rs.ir.Read8Bytes()
}

// ReadWord_BE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadWord_BE() (w bt.Word, err error) {
	return rs.ir.ReadWord_BE()
}

// ReadWord_LE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadWord_LE() (w bt.Word, err error) {
	return rs.ir.ReadWord_LE()
}

// ReadDWord_BE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadDWord_BE() (dw bt.DWord, err error) {
	return rs.ir.ReadDWord_BE()
}

// ReadDWord_LE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadDWord_LE() (dw bt.DWord, err error) {
	return rs.ir.ReadDWord_LE()
}

// ReadUShort_BE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadUShort_BE() (us uint16, err error) {
	return rs.ir.ReadUShort_BE()
}

// ReadUShort_LE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadUShort_LE() (us uint16, err error) {
	return rs.ir.ReadUShort_LE()
}

// ReadULong_BE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadULong_BE() (ul uint32, err error) {
	return rs.ir.ReadULong_BE()
}

// ReadULong_LE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadULong_LE() (ul uint32, err error) {
	return rs.ir.ReadULong_LE()
}

// ReadSShort_BE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadSShort_BE() (ss int16, err error) {
	return rs.ir.ReadSShort_BE()
}

// ReadSShort_LE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadSShort_LE() (ss int16, err error) {
	return rs.ir.ReadSShort_LE()
}

// ReadSLong_BE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadSLong_BE() (sl int32, err error) {
	return rs.ir.ReadSLong_BE()
}

// ReadSLong_LE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadSLong_LE() (sl int32, err error) {
	return rs.ir.ReadSLong_LE()
}

// ReadFloat_BE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadFloat_BE() (f float32, err error) {
	return rs.ir.ReadFloat_BE()
}

// ReadFloat_LE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadFloat_LE() (f float32, err error) {
	return rs.ir.ReadFloat_LE()
}

// ReadDouble_BE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadDouble_BE() (d float64, err error) {
	return rs.ir.ReadDouble_BE()
}

// ReadDouble_LE is a wrapper of the child's method.
func (rs *ReaderSeeker) ReadDouble_LE() (d float64, err error) {
	return rs.ir.ReadDouble_LE()
}
