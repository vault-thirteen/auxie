package reader

// This package provides is a convenient Reader for specific purposes.

import (
	"io"

	rs "github.com/vault-thirteen/auxie/ReaderSeeker"
)

// ASCII symbols.
const (
	CR = '\r'
	LF = '\n'
)

type Reader struct {
	reader rs.ReaderSeeker
}

func NewReader(
	reader rs.ReaderSeeker,
) *Reader {
	return &Reader{
		reader: reader,
	}
}

func (r *Reader) Read(dst []byte) (n int, err error) {
	return r.reader.Read(dst)
}

func (r *Reader) Seek(offset int64, whence int) (int64, error) {
	return r.reader.Seek(offset, whence)
}

func (r *Reader) GetInternalReader() io.Reader {
	return r.reader
}
