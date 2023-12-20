package rs

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"

	iors "github.com/vault-thirteen/auxie/ReaderSeeker"
	"github.com/vault-thirteen/auxie/reader"
	"github.com/vault-thirteen/auxie/tester"
)

func Test_New(t *testing.T) {
	aTest := tester.New(t)
	var rs *ReaderSeeker
	var err error

	// Test #1. Read-only stream.
	r := rand.Reader
	rs, err = New(r)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), ErrStreamDoesNotSupportSeeking)
	aTest.MustBeEqual(rs, (*ReaderSeeker)(nil))

	// Test #2. Seek-able stream.
	r = bytes.NewReader([]byte{})
	rs, err = New(r)
	aTest.MustBeNoError(err)
	ir := reader.New(r.(iors.ReaderSeeker))
	aTest.MustBeEqual(rs, &ReaderSeeker{
		ir:     ir,
		seeker: ir.GetInternalReader().(io.Seeker),
	})
}

func Test_GetInternalReader(t *testing.T) {
	aTest := tester.New(t)
	r := bytes.NewReader([]byte{1, 2, 3, 4, 5})
	rs, err := New(r)
	aTest.MustBeNoError(err)

	// Simple check.
	x := rs.GetInternalReader()
	aTest.MustBeEqual(x, r)

	// Try to move the cursor of internal reader.
	var threeBytes = make([]byte, 3)
	_, err = rs.ir.Read(threeBytes)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(threeBytes, []byte{1, 2, 3})
	xx := rs.GetInternalReader()
	var restBytes = make([]byte, 2)
	_, err = xx.Read(restBytes)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(restBytes, []byte{4, 5})
}

func Test_GetInternalSeeker(t *testing.T) {
	aTest := tester.New(t)
	r := bytes.NewReader([]byte{1, 2, 3, 4, 5})
	rs, err := New(r)
	aTest.MustBeNoError(err)

	// Simple check.
	x := rs.GetInternalSeeker()
	aTest.MustBeEqual(x, r)

	// Try to move the cursor of internal seeker.
	_, err = x.Seek(3, io.SeekStart)
	aTest.MustBeNoError(err)

	// Check the rest bytes.
	xx := rs.GetInternalReader()
	var restBytes = make([]byte, 2)
	_, err = xx.Read(restBytes)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(restBytes, []byte{4, 5})
}

func Test_GetInternalReaderSeeker(t *testing.T) {
	aTest := tester.New(t)
	r := bytes.NewReader([]byte{1, 2, 3, 4, 5})
	rs, err := New(r)
	aTest.MustBeNoError(err)

	// Simple check.
	x := rs.GetInternalReaderSeeker()
	aTest.MustBeEqual(x, r)

	// Try to move the cursor using a reader.
	var buf = make([]byte, 1)
	_, err = rs.ir.Read(buf)
	aTest.MustBeNoError(err)

	// Try to move the cursor using a seeker.
	_, err = rs.seeker.Seek(1, io.SeekCurrent)
	aTest.MustBeNoError(err)

	// Try to move the cursor using a seeker-reader.
	_, err = x.Seek(1, io.SeekCurrent)
	aTest.MustBeNoError(err)

	// Check the rest bytes using a reader-seeker.
	var restBytes = make([]byte, 2)
	_, err = x.Read(restBytes)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(restBytes, []byte{4, 5})
}
