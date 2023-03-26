package rs

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/vault-thirteen/tester"
)

// Notes.
//
// All these methods are just wrappers. Here we can test that the main method
// works and this will mean that all other methods work as well.

func Test_ReadBytes(t *testing.T) {
	var tst = tester.New(t)
	var err error
	var rs *ReaderSeeker
	var result []byte

	type TestData struct {
		Data                []byte
		NumberOFBytesToRead int
		ExpectedResult      []byte
		ExpectedError       error
	}

	tests := []TestData{
		// Test #1. Normal Data.
		{
			Data:                []byte("ABCDEFG"),
			NumberOFBytesToRead: 3,
			ExpectedResult:      []byte("ABC"),
			ExpectedError:       nil,
		},

		// Test #2. Data is not enough.
		{
			Data:                []byte("ABCDEFG"),
			NumberOFBytesToRead: 10,
			ExpectedResult:      []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G'},
			ExpectedError:       io.ErrUnexpectedEOF,
		},

		// Test #3. Empty Data.
		{
			Data:                []byte{},
			NumberOFBytesToRead: 3,
			ExpectedResult:      []byte{},
			ExpectedError:       io.EOF,
		},
	}

	n := 1
	for _, test := range tests {
		fmt.Print("[", n, "] ")

		rs, err = New(bytes.NewReader(test.Data))
		tst.MustBeNoError(err)

		result, err = rs.ReadBytes(test.NumberOFBytesToRead)
		if test.ExpectedError == nil {
			tst.MustBeNoError(err)
		} else {
			tst.MustBeAnError(err)
			tst.MustBeEqual(err, test.ExpectedError)
		}
		tst.MustBeEqual(result, test.ExpectedResult)

		n++
	}
	fmt.Println()
}
