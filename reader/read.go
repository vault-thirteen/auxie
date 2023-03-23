package reader

import (
	"encoding/binary"
	"fmt"
	"io"

	bt "github.com/vault-thirteen/auxie/BasicTypes"
)

const ErrUnexpectedDataSize = "unexpected data size: %v vs %v"

// ReadLineEndingWithCRLF reads a line ending exactly with the 'CR'+'LF'
// symbols sequence. The two symbols at the end of the line (CR+LF) are
// included into the returned result. On error, returns the last read sequence
// of bytes, even if it does not have a correct ending.
func (r *Reader) ReadLineEndingWithCRLF() (line []byte, err error) {
	var b byte
	var success bool

	// We must find the exact combination (sequence) of CR and LF, where LF is
	// right after the CR.

	// Read the first byte.
	b, err = r.ReadByte()
	if err != nil {
		return line, err
	}
	line = append(line, b)

	// Read next bytes.
	for !success {

		// Read a single byte.
		b, err = r.ReadByte()
		if err != nil {
			return line, err
		}
		line = append(line, b)

		// LF ?
		if b == LF {
			// Previous symbol must be CR to exit the loop.
			if line[len(line)-2] == CR {
				success = true
			}
		}
	}

	return line, nil
}

// ReadBytes reads an exact number of bytes.
func (r *Reader) ReadBytes(size int) (bytes []byte, err error) {
	bytes = make([]byte, size)
	var n int
	n, err = io.ReadFull(r, bytes)
	if err != nil {
		return bytes, err
	}
	if n != size {
		return bytes, fmt.Errorf(ErrUnexpectedDataSize, size, n)
	}

	return bytes, nil
}

// ReadByte reads one (unsigned) byte.
func (r *Reader) ReadByte() (ub byte, err error) {
	var bytes []byte
	bytes, err = r.ReadBytes(1)
	if err != nil {
		return ub, err
	}

	return bytes[0], nil
}

// ReadSByte reads one signed byte.
func (r *Reader) ReadSByte() (sb int8, err error) {
	var ub byte
	ub, err = r.ReadByte()
	if err != nil {
		return sb, err
	}

	return int8(ub), nil
}

// Read2Bytes reads two bytes.
func (r *Reader) Read2Bytes() (bytes []byte, err error) {
	return r.ReadBytes(2)
}

// Read4Bytes reads four bytes.
func (r *Reader) Read4Bytes() (bytes []byte, err error) {
	return r.ReadBytes(4)
}

// Read8Bytes reads eight bytes.
func (r *Reader) Read8Bytes() (bytes []byte, err error) {
	return r.ReadBytes(8)
}

// ReadWord_BE reads a word using the big endian technique.
func (r *Reader) ReadWord_BE() (w bt.Word, err error) {
	var bytes []byte
	bytes, err = r.Read2Bytes()
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint16(bytes), nil
}

// ReadWord_LE reads a word using the little endian technique.
func (r *Reader) ReadWord_LE() (w bt.Word, err error) {
	var bytes []byte
	bytes, err = r.Read2Bytes()
	if err != nil {
		return w, err
	}

	return binary.LittleEndian.Uint16(bytes), nil
}

// ReadDWord_BE reads a double word using the big endian technique.
func (r *Reader) ReadDWord_BE() (dw bt.DWord, err error) {
	var bytes []byte
	bytes, err = r.Read4Bytes()
	if err != nil {
		return dw, err
	}

	return binary.BigEndian.Uint32(bytes), nil
}

// ReadDWord_LE reads a double word using the little endian technique.
func (r *Reader) ReadDWord_LE() (dw bt.DWord, err error) {
	var bytes []byte
	bytes, err = r.Read4Bytes()
	if err != nil {
		return dw, err
	}

	return binary.LittleEndian.Uint32(bytes), nil
}

// ReadUShort_BE reads an unsigned short using the big endian technique.
func (r *Reader) ReadUShort_BE() (us uint16, err error) {
	return r.ReadWord_BE()
}

// ReadUShort_LE reads an unsigned short using the little endian technique.
func (r *Reader) ReadUShort_LE() (us uint16, err error) {
	return r.ReadWord_LE()
}

// ReadULong_BE reads an unsigned long using the big endian technique.
func (r *Reader) ReadULong_BE() (ul uint32, err error) {
	return r.ReadDWord_BE()
}

// ReadULong_LE reads an unsigned long using the little endian technique.
func (r *Reader) ReadULong_LE() (ul uint32, err error) {
	return r.ReadDWord_LE()
}

// ReadSShort_BE reads a signed short using the big endian technique.
func (r *Reader) ReadSShort_BE() (ss int16, err error) {
	err = binary.Read(r, binary.BigEndian, &ss)
	if err != nil {
		return ss, err
	}

	return ss, nil
}

// ReadSShort_LE reads a signed short using the little endian technique.
func (r *Reader) ReadSShort_LE() (ss int16, err error) {
	err = binary.Read(r, binary.LittleEndian, &ss)
	if err != nil {
		return ss, err
	}

	return ss, nil
}

// ReadSLong_BE reads a signed long using the big endian technique.
func (r *Reader) ReadSLong_BE() (sl int32, err error) {
	err = binary.Read(r, binary.BigEndian, &sl)
	if err != nil {
		return sl, err
	}

	return sl, nil
}

// ReadSLong_LE reads a signed long using the little endian technique.
func (r *Reader) ReadSLong_LE() (sl int32, err error) {
	err = binary.Read(r, binary.LittleEndian, &sl)
	if err != nil {
		return sl, err
	}

	return sl, nil
}

// ReadFloat_BE reads a Float using the big endian technique.
func (r *Reader) ReadFloat_BE() (f float32, err error) {
	err = binary.Read(r, binary.BigEndian, &f)
	if err != nil {
		return f, err
	}

	return f, nil
}

// ReadFloat_LE reads a Float using the little endian technique.
func (r *Reader) ReadFloat_LE() (f float32, err error) {
	err = binary.Read(r, binary.LittleEndian, &f)
	if err != nil {
		return f, err
	}

	return f, nil
}

// ReadDouble_BE reads a Double using the big endian technique.
func (r *Reader) ReadDouble_BE() (d float64, err error) {
	err = binary.Read(r, binary.BigEndian, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}

// ReadDouble_LE reads a Double using the little endian technique.
func (r *Reader) ReadDouble_LE() (d float64, err error) {
	err = binary.Read(r, binary.LittleEndian, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}
