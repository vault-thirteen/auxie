package reader

import (
	"bytes"
	"testing"

	rs "github.com/vault-thirteen/auxie/ReaderSeeker"
	"github.com/vault-thirteen/tester"
)

func Test_NewReader(t *testing.T) {
	var reader rs.ReaderSeeker
	var result *Reader

	reader = bytes.NewReader([]byte{})
	result = NewReader(reader)
	if result.rs != reader {
		t.FailNow()
	}
}

func Test_GetInternalReader(t *testing.T) {
	tst := tester.New(t)
	internalReader := bytes.NewReader([]byte{1, 2, 3, 4, 5})
	r := NewReader(internalReader)

	// Simple check.
	x := r.GetInternalReader()
	tst.MustBeEqual(x, internalReader)

	// Try to move the cursor of internal reader-seeker.
	var threeBytes = make([]byte, 3)
	_, err := r.rs.Read(threeBytes)
	tst.MustBeNoError(err)
	tst.MustBeEqual(threeBytes, []byte{1, 2, 3})
	xx := r.GetInternalReader()
	var restBytes = make([]byte, 2)
	_, err = xx.Read(restBytes)
	tst.MustBeNoError(err)
	tst.MustBeEqual(restBytes, []byte{4, 5})
}
