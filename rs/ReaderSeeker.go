package rs

import (
	"errors"
	"io"

	iors "github.com/vault-thirteen/auxie/ReaderSeeker"
	"github.com/vault-thirteen/auxie/reader"
)

const ErrStreamDoesNotSupportSeeking = "stream does not support seeking"

type ReaderSeeker struct {
	ir     reader.IReader
	seeker io.Seeker
}

func New(stream io.Reader) (rs *ReaderSeeker, err error) {
	_, ok := stream.(iors.ReaderSeeker)
	if !ok {
		return nil, errors.New(ErrStreamDoesNotSupportSeeking)
	}

	ir := reader.New(stream)

	rs = &ReaderSeeker{
		ir:     ir,
		seeker: ir.GetInternalReader().(io.Seeker),
	}

	return rs, nil
}

func (rs *ReaderSeeker) GetInternalReader() io.Reader {
	return rs.ir.GetInternalReader()
}

func (rs *ReaderSeeker) GetInternalSeeker() io.Seeker {
	return rs.ir.GetInternalReader().(io.Seeker)
}

func (rs *ReaderSeeker) GetInternalReaderSeeker() iors.ReaderSeeker {
	return rs.ir.GetInternalReader().(iors.ReaderSeeker)
}
