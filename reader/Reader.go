package reader

// This package provides is a convenient Reader for specific purposes.

import (
	"io"
)

// ASCII symbols.
const (
	CR = '\r'
	LF = '\n'
)

type Reader struct {
	reader io.Reader
}

func NewReader(
	reader io.Reader,
) *Reader {
	return &Reader{
		reader: reader,
	}
}

func (r *Reader) Read(dst []byte) (n int, err error) {
	return r.reader.Read(dst)
}

func (r *Reader) GetInternalReader() io.Reader {
	return r.reader
}
