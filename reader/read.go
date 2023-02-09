// read.go.

package reader

import (
	"errors"
	"io"
)

const ErrSizeError = "Size Error"

// ReadLineEndingWithCRLF reads a Line ending exactly with the 'CR'+'LF'
// Symbols Sequence. The two Symbols at the End of the Line (CR+LF) are
// included into the returned Result.
func (r *Reader) ReadLineEndingWithCRLF() ([]byte, error) {

	var b [1]byte
	var err error
	var line []byte
	var success bool

	// We must find the exact Combination (Sequence) of CR and LF, where LF is
	// right after the CR.

	// Read the first Byte.
	_, err = r.reader.Read(b[:])
	if err != nil {
		return []byte{}, err
	}
	line = append(line, b[0])

	// Read next Bytes.
	for !success {

		// Read a single Byte.
		_, err = r.reader.Read(b[:])
		if err != nil {
			return []byte{}, err
		}
		line = append(line, b[0])

		// LF?
		if b[0] == '\n' {
			// Previous Symbol must be CR to exit the Loop.
			if line[len(line)-2] == '\r' {
				success = true
			}
		}
	}

	return line, nil
}

// ReadBytes tries to read an exact Number of Bytes.
func (r *Reader) ReadBytes(
	size int,
) ([]byte, error) {

	var buffer []byte
	var err error
	var bytesReadCount int

	buffer = make([]byte, size)
	bytesReadCount, err = io.ReadAtLeast(r.reader, buffer, size)
	if err != nil {
		return []byte{}, err
	}
	if bytesReadCount != size {
		err = errors.New(ErrSizeError)
		return []byte{}, err
	}

	return buffer, nil
}
