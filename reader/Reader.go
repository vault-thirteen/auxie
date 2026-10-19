package reader

import (
	"io"
)

// ASCII symbols.
const (
	CR = '\r'
	LF = '\n'
)

type Reader struct {
	r io.Reader
}

func New(stream io.Reader) *Reader {
	return &Reader{r: stream}
}

func (r *Reader) GetInternalReader() io.Reader {
	return r.r
}
