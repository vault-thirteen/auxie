package reader

import (
	"bytes"
	"io"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_ReadByte_simple(t *testing.T) {
	aTest := tester.New(t)
	var err error
	var result byte

	// Test #1.
	src := bytes.NewBuffer([]byte{255})
	result, err = ReadByte(src)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, byte(255))

	// Test #2.
	src = bytes.NewBuffer([]byte{})
	result, err = ReadByte(src)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err, io.EOF)
	aTest.MustBeEqual(result, byte(0))
}
