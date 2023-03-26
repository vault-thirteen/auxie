package rs

import (
	"bytes"
	"io"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_Seek(t *testing.T) {
	aTest := tester.New(t)
	r := bytes.NewReader([]byte{1, 2, 3, 4, 5})
	rs, err := New(r)
	aTest.MustBeNoError(err)

	// Try to move the cursor using a seeker.
	_, err = rs.Seek(2, io.SeekCurrent)
	aTest.MustBeNoError(err)

	// Check the rest bytes.
	var restBytes = make([]byte, 3)
	_, err = rs.Read(restBytes)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(restBytes, []byte{3, 4, 5})
}
