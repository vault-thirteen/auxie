package rs

import "io"

type ReaderSeeker interface {
	io.Reader
	io.Seeker
}
