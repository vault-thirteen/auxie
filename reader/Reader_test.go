package reader

import (
	"bytes"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_New(t *testing.T) {
	var result *Reader

	r := bytes.NewReader([]byte{})
	result = New(r)
	if result.r != r {
		t.FailNow()
	}
}

func Test_GetInternalReader(t *testing.T) {
	aTest := tester.New(t)
	internalReader := bytes.NewReader([]byte{1, 2, 3, 4, 5})
	r := New(internalReader)

	// Simple check.
	x := r.GetInternalReader()
	aTest.MustBeEqual(x, internalReader)

	// Try to move the cursor of internal reader.
	var threeBytes = make([]byte, 3)
	_, err := r.r.Read(threeBytes)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(threeBytes, []byte{1, 2, 3})

	// Check the rest bytes.
	xx := r.GetInternalReader()
	var restBytes = make([]byte, 2)
	_, err = xx.Read(restBytes)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(restBytes, []byte{4, 5})
}
