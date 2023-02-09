// read_test.go.

package reader

import (
	"bytes"
	"io"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_ReadLineEndingWithCRLF(t *testing.T) {

	var data []byte
	var err error
	var reader1 io.Reader
	var reader2 *Reader
	var result []byte
	var resultExpected []byte
	var tst *tester.Test

	tst = tester.New(t)

	// Test #1. Normal Data.

	// Prepare the Data.
	data = []byte("123")
	data = append(data, CR)
	data = append(data, []byte("456")...)
	data = append(data, LF)
	data = append(data, []byte("789")...)
	data = append(data, CR, LF)
	data = append(data, []byte("AB")...)
	data = append(data, CR, LF)

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = NewReader(reader1)
	resultExpected = data[0:13]
	result, err = reader2.ReadLineEndingWithCRLF()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, resultExpected)
	resultExpected = []byte("AB\r\n")
	result, err = reader2.ReadLineEndingWithCRLF()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, resultExpected)

	// Test #2. No CR+LF.

	// Prepare the Data.
	data = []byte("12")
	data = append(data, CR)
	data = append(data, []byte("34")...)
	data = append(data, LF)
	data = append(data, []byte("5")...)
	data = append(data, LF, CR)
	data = append(data, []byte("67")...)
	resultExpected = []byte{}

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = NewReader(reader1)
	result, err = reader2.ReadLineEndingWithCRLF()
	tst.MustBeAnError(err)
	tst.MustBeEqual(err.Error(), io.EOF.Error())
	tst.MustBeEqual(result, resultExpected)

	// Test #3. Empty Data.

	// Prepare the Data.
	data = []byte{}
	resultExpected = []byte{}

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = NewReader(reader1)
	result, err = reader2.ReadLineEndingWithCRLF()
	tst.MustBeAnError(err)
	tst.MustBeEqual(err.Error(), io.EOF.Error())
	tst.MustBeEqual(result, resultExpected)

	// Test #4. Normal Data.

	// Prepare the Data.
	data = []byte("A\rB\nC\n\rD\r\n")
	resultExpected = data

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = NewReader(reader1)
	result, err = reader2.ReadLineEndingWithCRLF()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, resultExpected)
}

func Test_ReadBytes(t *testing.T) {

	var data []byte
	var err error
	var reader1 io.Reader
	var reader2 *Reader
	var result []byte
	var resultExpected []byte
	var tst *tester.Test

	tst = tester.New(t)

	// Test #1. Normal Data.

	// Prepare the Data.
	data = []byte("ABCDEFG")
	resultExpected = []byte("ABC")

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = NewReader(reader1)
	result, err = reader2.ReadBytes(3)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, resultExpected)

	// Test #2. Data is not enough.

	// Prepare the Data.
	data = []byte("ABCDEFG")
	resultExpected = []byte{}

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = NewReader(reader1)
	result, err = reader2.ReadBytes(100)
	tst.MustBeAnError(err)
	tst.MustBeEqual(err.Error(), io.ErrUnexpectedEOF.Error())
	tst.MustBeEqual(result, resultExpected)

	// Test #3. Empty Data.

	// Prepare the Data.
	data = []byte{}
	resultExpected = []byte{}

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = NewReader(reader1)
	result, err = reader2.ReadBytes(3)
	tst.MustBeAnError(err)
	tst.MustBeEqual(err.Error(), io.EOF.Error())
	tst.MustBeEqual(result, resultExpected)

	// Test #4. Normal Data. Combined Read.

	// Prepare the Data.
	data = []byte("ABC")
	data = append(data, CR, LF)
	data = append(data, []byte("1234567")...)

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = NewReader(reader1)

	// Part 1.
	resultExpected = []byte("ABC\r\n")
	result, err = reader2.ReadLineEndingWithCRLF()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, resultExpected)

	// Part 2.
	resultExpected = []byte("123")
	result, err = reader2.ReadBytes(3)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, resultExpected)
}
