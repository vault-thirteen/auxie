package reader

import (
	"bytes"
	"io"
	"testing"

	bt "github.com/vault-thirteen/auxie/BasicTypes"
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
	resultExpected = []byte(nil)

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = NewReader(reader1)
	result, err = reader2.ReadBytes(100)
	tst.MustBeAnError(err)
	tst.MustBeEqual(err.Error(), "unexpected data size: 100 vs 7")
	tst.MustBeEqual(result, resultExpected)

	// Test #3. Empty Data.

	// Prepare the Data.
	data = []byte{}
	resultExpected = []byte(nil)

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = NewReader(reader1)
	result, err = reader2.ReadBytes(3)
	tst.MustBeAnError(err)
	tst.MustBeEqual(err.Error(), "unexpected data size: 3 vs 0")
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

func Test_ReadByte(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result byte

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{'A', 'B'}))
	result, err = r.ReadByte()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, byte('A'))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadByte()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, byte(0))
}

func Test_ReadSByte(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result int8

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{127, 'Q'}))
	result, err = r.ReadSByte()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, int8(127))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadSByte()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, int8(0))
}

func Test_Read2Bytes(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result []byte

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{'A', 'B'}))
	result, err = r.Read2Bytes()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []byte{'A', 'B'})

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.Read2Bytes()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, []byte(nil))
}

func Test_Read4Bytes(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result []byte

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{'A', 'B', 'C', 'D'}))
	result, err = r.Read4Bytes()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []byte{'A', 'B', 'C', 'D'})

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.Read4Bytes()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, []byte(nil))
}

func Test_Read8Bytes(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result []byte

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}))
	result, err = r.Read8Bytes()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'})

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.Read8Bytes()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, []byte(nil))
}

func Test_ReadWord_BE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result bt.Word

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0xFF}))
	result, err = r.ReadWord_BE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, bt.Word(255))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadWord_BE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, bt.Word(0))
}

func Test_ReadWord_LE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result bt.Word

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0xFF}))
	result, err = r.ReadWord_LE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, bt.Word(65280))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadWord_LE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, bt.Word(0))
}

func Test_ReadDWord_BE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result bt.DWord

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0xFF}))
	result, err = r.ReadDWord_BE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, bt.DWord(255))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadDWord_BE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, bt.DWord(0))
}

func Test_ReadDWord_LE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result bt.DWord

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0xFF}))
	result, err = r.ReadDWord_LE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, bt.DWord(4278190080))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadDWord_LE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, bt.DWord(0))
}

func Test_ReadUShort_BE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result bt.Word

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0xFF}))
	result, err = r.ReadUShort_BE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, bt.Word(255))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadUShort_BE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, bt.Word(0))
}

func Test_ReadUShort_LE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result bt.Word

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0xFF}))
	result, err = r.ReadUShort_LE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, bt.Word(65280))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadUShort_LE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, bt.Word(0))
}

func Test_ReadULong_BE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result bt.DWord

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0xFF}))
	result, err = r.ReadULong_BE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, bt.DWord(255))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadULong_BE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, bt.DWord(0))
}

func Test_ReadULong_LE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result bt.DWord

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0xFF}))
	result, err = r.ReadULong_LE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, bt.DWord(4278190080))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadULong_LE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, bt.DWord(0))
}

func Test_ReadSShort_BE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result int16

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0xFF}))
	result, err = r.ReadSShort_BE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, int16(255))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadSShort_BE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, int16(0))
}

func Test_ReadSShort_LE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result int16

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0xFF}))
	result, err = r.ReadSShort_LE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, int16(-256))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadSShort_LE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, int16(0))
}

func Test_ReadSLong_BE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result int32

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0xFF}))
	result, err = r.ReadSLong_BE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, int32(255))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadSLong_BE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, int32(0))
}

func Test_ReadSLong_LE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result int32

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0xFF}))
	result, err = r.ReadSLong_LE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, int32(-16777216))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadSLong_LE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, int32(0))
}

func Test_ReadFloat_BE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result float32

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0xFF}))
	result, err = r.ReadFloat_BE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, float32(3.573311e-43))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadFloat_BE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, float32(0))
}

func Test_ReadFloat_LE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result float32

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0xFF}))
	result, err = r.ReadFloat_LE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, float32(-1.70141183e+38))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadFloat_LE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, float32(0))
}

func Test_ReadDouble_BE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result float64

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF}))
	result, err = r.ReadDouble_BE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, float64(1.26e-321))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadDouble_BE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, float64(0))
}

func Test_ReadDouble_LE(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result float64

	// Test #1. Normal Data.
	r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF}))
	result, err = r.ReadDouble_LE()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, float64(-5.486124068793689e+303))

	// Test #2. Bad Data.
	r = NewReader(bytes.NewReader([]byte{}))
	result, err = r.ReadDouble_LE()
	tst.MustBeAnError(err)
	tst.MustBeEqual(result, float64(0))
}
