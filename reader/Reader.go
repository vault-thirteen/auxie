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

func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

func (r *Reader) Read(dst []byte) (n int, err error) {
	return r.r.Read(dst)
}

func (r *Reader) GetInternalReader() io.Reader {
	return r.r
}
