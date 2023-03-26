package iors

import "io"

type ReaderSeeker interface {
	io.Reader
	io.Seeker
}
