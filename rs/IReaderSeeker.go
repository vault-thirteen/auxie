package rs

import (
	"io"

	iors "github.com/vault-thirteen/auxie/ReaderSeeker"
	"github.com/vault-thirteen/auxie/reader"
)

type IReaderSeeker interface {
	io.Reader
	io.Seeker
	reader.IReader
	GetInternalSeeker() io.Seeker
	GetInternalReaderSeeker() iors.ReaderSeeker
}
