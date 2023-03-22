package reader

// This package provides a convenient Reader-Seeker for specific purposes.

import (
	rs "github.com/vault-thirteen/auxie/ReaderSeeker"
)

// ASCII symbols.
const (
	CR = '\r'
	LF = '\n'
)

type Reader struct {
	rs rs.ReaderSeeker
}

func NewReader(rs rs.ReaderSeeker) *Reader {
	return &Reader{
		rs: rs,
	}
}

func (r *Reader) Read(dst []byte) (n int, err error) {
	return r.rs.Read(dst)
}

func (r *Reader) Seek(offset int64, whence int) (int64, error) {
	return r.rs.Seek(offset, whence)
}

func (r *Reader) GetInternalReader() rs.ReaderSeeker {
	return r.rs
}
