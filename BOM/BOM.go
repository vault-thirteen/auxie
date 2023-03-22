package bom

import (
	"errors"
	"fmt"
	"io"

	rs "github.com/vault-thirteen/auxie/ReaderSeeker"
	"github.com/vault-thirteen/auxie/reader"
	"github.com/vault-thirteen/errorz"
)

const (
	ErrUnknownEncoding = "unknown encoding: %v"
	ErrBOMIsNotFound   = "byte order mark is not found"
)

// GetEncoding tries to get all possible encodings for the stream.
// Please note that some encodings have similar BOMs, that is why an array is
// returned instead of a single value. The reader is reset after the reading.
//
// The 'shouldIncludeUnknownEncoding' flag switches the way of result telling.
// If the flag is true, the UnknownEncoding is always included into the result.
// This is a pedantic way of telling the result for those who want a scientific
// approach. If the flag is false, UnknownEncoding is not included into the
// result. This means that if no encoding is found, the result will be an empty
// array. This is a simplistic approach. Independently of the setting, the
// result array is always null (nil) on reader's error.
func GetEncoding(rs rs.ReaderSeeker, shouldIncludeUnknownEncoding bool) (result []Encoding, err error) {
	result = make([]Encoding, 0)

	var isEncoding bool
	for _, pe := range possibleEncodings {
		isEncoding, err = IsEncoding(rs, pe)
		if err != nil {
			return nil, err
		}

		if isEncoding {
			result = append(result, pe)
		}
	}

	if shouldIncludeUnknownEncoding {
		result = append(result, EncodingUnknown)
	}

	return result, nil
}

// IsEncoding checks the beginning of the stream and compares it's BOM with the
// BOM of the specified encoding. Simply, it tries to check whether the stream
// has a BOM of the specified encoding or not. The reader is reset after the
// reading.
func IsEncoding(rs rs.ReaderSeeker, enc Encoding) (isEncoding bool, err error) {
	var initialOffset int64
	initialOffset, err = rs.Seek(0, io.SeekCurrent)
	if err != nil {
		return false, err
	}

	defer func() {
		_, derr := rs.Seek(initialOffset, io.SeekStart)
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	rdr := reader.NewReader(rs)

	var getErr error
	switch enc {
	case EncodingUTF8:
		_, getErr = GetBOMForEncodingUTF8(rdr)
	case EncodingUTF16BE:
		_, getErr = GetBOMForEncodingUTF16BE(rdr)
	case EncodingUTF16LE:
		_, getErr = GetBOMForEncodingUTF16LE(rdr)
	case EncodingUTF32BE:
		_, getErr = GetBOMForEncodingUTF32BE(rdr)
	case EncodingUTF32LE:
		_, getErr = GetBOMForEncodingUTF32LE(rdr)
	case EncodingUTF7:
		_, getErr = GetBOMForEncodingUTF7(rdr)
	case EncodingUTF1:
		_, getErr = GetBOMForEncodingUTF1(rdr)
	case EncodingUTF_EBCDIC:
		_, getErr = GetBOMForEncodingUTF_EBCDIC(rdr)
	case EncodingSCSU:
		_, getErr = GetBOMForEncodingSCSU(rdr)
	case EncodingBOCU1:
		_, getErr = GetBOMForEncodingBOCU1(rdr)
	case EncodingGB18030:
		_, getErr = GetBOMForEncodingGB18030(rdr)
	default:
		return false, fmt.Errorf(ErrUnknownEncoding, enc)
	}
	if getErr != nil {
		return false, nil
	}

	return true, nil
}

// SkipBOMPrefix tries to skip BOM prefix from the data.
// It reads the BOM and returns the reader.
func SkipBOMPrefix(rs rs.ReaderSeeker, enc Encoding) (newRS rs.ReaderSeeker, err error) {
	rdr := reader.NewReader(rs)

	switch enc {
	case EncodingUTF8:
		_, err = GetBOMForEncodingUTF8(rdr)
	case EncodingUTF16BE:
		_, err = GetBOMForEncodingUTF16BE(rdr)
	case EncodingUTF16LE:
		_, err = GetBOMForEncodingUTF16LE(rdr)
	case EncodingUTF32BE:
		_, err = GetBOMForEncodingUTF32BE(rdr)
	case EncodingUTF32LE:
		_, err = GetBOMForEncodingUTF32LE(rdr)
	case EncodingUTF7:
		_, err = GetBOMForEncodingUTF7(rdr)
	case EncodingUTF1:
		_, err = GetBOMForEncodingUTF1(rdr)
	case EncodingUTF_EBCDIC:
		_, err = GetBOMForEncodingUTF_EBCDIC(rdr)
	case EncodingSCSU:
		_, err = GetBOMForEncodingSCSU(rdr)
	case EncodingBOCU1:
		_, err = GetBOMForEncodingBOCU1(rdr)
	case EncodingGB18030:
		_, err = GetBOMForEncodingGB18030(rdr)
	default:
		return nil, fmt.Errorf(ErrUnknownEncoding, enc)
	}
	if err != nil {
		return nil, err
	}

	return rdr.GetInternalReader(), nil
}

// GetBOMForEncodingUTF8 tries to read the BOM of UTF-8 encoding.
// If the BOM is found, error is nil; otherwise it is not nil.
// The prefix which was read from the stream is always returned as the first
// returned value. The reader is not reset after the reading.
func GetBOMForEncodingUTF8(r *reader.Reader) (prefix []byte, err error) {
	prefix, err = r.ReadBytes(3)
	if err != nil {
		return prefix, err
	}

	if (prefix[0] == 0xEF) && (prefix[1] == 0xBB) && (prefix[2] == 0xBF) {
		return prefix, nil
	}

	return prefix, errors.New(ErrBOMIsNotFound)
}

// GetBOMForEncodingUTF16BE tries to read the BOM of UTF-16 [BE] encoding.
// If the BOM is found, error is nil; otherwise it is not nil.
// The prefix which was read from the stream is always returned as the first
// returned value. The reader is not reset after the reading.
func GetBOMForEncodingUTF16BE(r *reader.Reader) (prefix []byte, err error) {
	prefix, err = r.ReadBytes(2)
	if err != nil {
		return prefix, err
	}

	if (prefix[0] == 0xFE) && (prefix[1] == 0xFF) {
		return prefix, nil
	}

	return prefix, errors.New(ErrBOMIsNotFound)
}

// GetBOMForEncodingUTF16LE tries to read the BOM of UTF-16 [LE] encoding.
// If the BOM is found, error is nil; otherwise it is not nil.
// The prefix which was read from the stream is always returned as the first
// returned value. The reader is not reset after the reading.
func GetBOMForEncodingUTF16LE(r *reader.Reader) (prefix []byte, err error) {
	prefix, err = r.ReadBytes(2)
	if err != nil {
		return prefix, err
	}

	if (prefix[0] == 0xFF) && (prefix[1] == 0xFE) {
		return prefix, nil
	}

	return prefix, errors.New(ErrBOMIsNotFound)
}

// GetBOMForEncodingUTF32BE tries to read the BOM of UTF-32 [BE] encoding.
// If the BOM is found, error is nil; otherwise it is not nil.
// The prefix which was read from the stream is always returned as the first
// returned value. The reader is not reset after the reading.
func GetBOMForEncodingUTF32BE(r *reader.Reader) (prefix []byte, err error) {
	prefix, err = r.ReadBytes(4)
	if err != nil {
		return prefix, err
	}

	if (prefix[0] == 0x00) && (prefix[1] == 0x00) &&
		(prefix[2] == 0xFE) && (prefix[3] == 0xFF) {
		return prefix, nil
	}

	return prefix, errors.New(ErrBOMIsNotFound)
}

// GetBOMForEncodingUTF32LE tries to read the BOM of UTF-32 [LE] encoding.
// If the BOM is found, error is nil; otherwise it is not nil.
// The prefix which was read from the stream is always returned as the first
// returned value. The reader is not reset after the reading.
func GetBOMForEncodingUTF32LE(r *reader.Reader) (prefix []byte, err error) {
	prefix, err = r.ReadBytes(4)
	if err != nil {
		return prefix, err
	}

	if (prefix[0] == 0xFF) && (prefix[1] == 0xFE) &&
		(prefix[2] == 0x00) && (prefix[3] == 0x00) {
		return prefix, nil
	}

	return prefix, errors.New(ErrBOMIsNotFound)
}

// GetBOMForEncodingUTF7 tries to read the BOM of UTF-7 encoding.
// If the BOM is found, error is nil; otherwise it is not nil.
// The prefix which was read from the stream is always returned as the first
// returned value. The reader is not reset after the reading.
func GetBOMForEncodingUTF7(r *reader.Reader) (prefix []byte, err error) {
	prefix, err = r.ReadBytes(3)
	if err != nil {
		return prefix, err
	}

	if (prefix[0] == 0x2B) && (prefix[1] == 0x2F) && (prefix[2] == 0x76) {
		return prefix, nil
	}

	return prefix, errors.New(ErrBOMIsNotFound)
}

// GetBOMForEncodingUTF1 tries to read the BOM of UTF-1 encoding.
// If the BOM is found, error is nil; otherwise it is not nil.
// The prefix which was read from the stream is always returned as the first
// returned value. The reader is not reset after the reading.
func GetBOMForEncodingUTF1(r *reader.Reader) (prefix []byte, err error) {
	prefix, err = r.ReadBytes(3)
	if err != nil {
		return prefix, err
	}

	if (prefix[0] == 0xF7) && (prefix[1] == 0x64) && (prefix[2] == 0x4C) {
		return prefix, nil
	}

	return prefix, errors.New(ErrBOMIsNotFound)
}

// GetBOMForEncodingUTF_EBCDIC tries to read the BOM of UTF-EBCDIC encoding.
// If the BOM is found, error is nil; otherwise it is not nil.
// The prefix which was read from the stream is always returned as the first
// returned value. The reader is not reset after the reading.
func GetBOMForEncodingUTF_EBCDIC(r *reader.Reader) (prefix []byte, err error) {
	prefix, err = r.ReadBytes(4)
	if err != nil {
		return prefix, err
	}

	if (prefix[0] == 0xDD) && (prefix[1] == 0x73) &&
		(prefix[2] == 0x66) && (prefix[3] == 0x73) {
		return prefix, nil
	}

	return prefix, errors.New(ErrBOMIsNotFound)
}

// GetBOMForEncodingSCSU tries to read the BOM of SCSU encoding.
// If the BOM is found, error is nil; otherwise it is not nil.
// The prefix which was read from the stream is always returned as the first
// returned value. The reader is not reset after the reading.
func GetBOMForEncodingSCSU(r *reader.Reader) (prefix []byte, err error) {
	prefix, err = r.ReadBytes(3)
	if err != nil {
		return prefix, err
	}

	if (prefix[0] == 0x0E) && (prefix[1] == 0xFE) && (prefix[2] == 0xFF) {
		return prefix, nil
	}

	return prefix, errors.New(ErrBOMIsNotFound)
}

// GetBOMForEncodingBOCU1 tries to read the BOM of BOCU-1 encoding.
// If the BOM is found, error is nil; otherwise it is not nil.
// The prefix which was read from the stream is always returned as the first
// returned value. The reader is not reset after the reading.
func GetBOMForEncodingBOCU1(r *reader.Reader) (prefix []byte, err error) {
	prefix, err = r.ReadBytes(3)
	if err != nil {
		return prefix, err
	}

	if (prefix[0] == 0xFB) && (prefix[1] == 0xEE) && (prefix[2] == 0x28) {
		return prefix, nil
	}

	return prefix, errors.New(ErrBOMIsNotFound)
}

// GetBOMForEncodingGB18030 tries to read the BOM of GB18030 encoding.
// If the BOM is found, error is nil; otherwise it is not nil.
// The prefix which was read from the stream is always returned as the first
// returned value. The reader is not reset after the reading.
func GetBOMForEncodingGB18030(r *reader.Reader) (prefix []byte, err error) {
	prefix, err = r.ReadBytes(4)
	if err != nil {
		return prefix, err
	}

	if (prefix[0] == 0x84) && (prefix[1] == 0x31) &&
		(prefix[2] == 0x95) && (prefix[3] == 0x33) {
		return prefix, nil
	}

	return prefix, errors.New(ErrBOMIsNotFound)
}
