package bom

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/vault-thirteen/auxie/reader"
	"github.com/vault-thirteen/tester"
)

func Test_BOMUTF8(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(BOMUTF8(), bomUTF8)
}

func Test_BOMUTF16BE(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(BOMUTF16BE(), bomUTF16BE)
}

func Test_BOMUTF16LE(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(BOMUTF16LE(), bomUTF16LE)
}

func Test_BOMUTF32BE(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(BOMUTF32BE(), bomUTF32BE)
}

func Test_BOMUTF32LE(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(BOMUTF32LE(), bomUTF32LE)
}

func Test_BOMUTF7(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(BOMUTF7(), bomUTF7)
}

func Test_BOMUTF1(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(BOMUTF1(), bomUTF1)
}

func Test_BOMUTF_EBCDIC(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(BOMUTF_EBCDIC(), bomUTF_EBCDIC)
}

func Test_BOMSCSU(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(BOMSCSU(), bomSCSU)
}

func Test_BOMBOCU1(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(BOMBOCU1(), bomBOCU1)
}

func Test_BOMGB18030(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(BOMGB18030(), bomGB18030)
}

func Test_BOMs(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(BOMs(), boms)
}

func Test_ReadBOMOfEncoding(t *testing.T) {
	tst := tester.New(t)

	type TestData struct {
		data           []byte
		encoding       Encoding
		expectedPrefix []byte
		expectedError  error
	}

	type Result struct {
		prefix []byte
		err    error
	}
	var result Result

	var tests = []TestData{
		{
			data:           []byte{},
			encoding:       Encoding(123),
			expectedPrefix: []byte(nil),
			expectedError:  errors.New("unknown encoding: 123"),
		},
		{
			data:           []byte{},
			encoding:       EncodingUTF8,
			expectedPrefix: []byte{},
			expectedError:  io.EOF,
		},
		{
			data:           []byte{'A', 'B'},
			encoding:       EncodingUTF8,
			expectedPrefix: []byte{'A', 'B'},
			expectedError:  errors.New("byte order mark is not found"),
		},
		{
			data:           []byte{0xEF, 0xBB},
			encoding:       EncodingUTF8,
			expectedPrefix: []byte{0xEF, 0xBB},
			expectedError:  errors.New("byte order mark is not found"),
		},
		{
			data:           []byte{'A', 'B', 'C'},
			encoding:       EncodingUTF8,
			expectedPrefix: []byte{'A', 'B', 'C'},
			expectedError:  errors.New("byte order mark is not found"),
		},
		{
			data:           []byte{0xEF, 0xBB, 0xBF},
			encoding:       EncodingUTF8,
			expectedPrefix: []byte{0xEF, 0xBB, 0xBF},
			expectedError:  nil,
		},
		{
			data:           []byte{'A', 'B', 'C', 'D'},
			encoding:       EncodingUTF8,
			expectedPrefix: []byte{'A', 'B', 'C'},
			expectedError:  errors.New("byte order mark is not found"),
		},
		{
			data:           []byte{0xEF, 0xBB, 0xBF, 'D'},
			encoding:       EncodingUTF8,
			expectedPrefix: []byte{0xEF, 0xBB, 0xBF},
			expectedError:  nil,
		},
	}

	// Run the tests.
	for i, test := range tests {
		fmt.Print("[", i+1, "] ")

		r := bytes.NewReader(test.data)
		result.prefix, result.err = ReadBOMOfEncoding(r, test.encoding)

		if test.expectedError == nil {
			tst.MustBeNoError(result.err)
		} else {
			tst.MustBeAnError(result.err)
			tst.MustBeEqual(result.err, test.expectedError)
		}
		tst.MustBeEqual(result.prefix, test.expectedPrefix)
	}
	fmt.Println()
}

func Test_SkipBOM(t *testing.T) {
	tst := tester.New(t)

	type TestData struct {
		data             []byte
		encoding         Encoding
		expectedNextByte byte
		expectedError    error
	}

	type Result struct {
		prefix []byte
		err    error
		b      byte
	}
	var result Result

	var tests = []TestData{
		{
			data:             []byte{'A', 'B', 'C', 'D', 'E'},
			encoding:         EncodingUTF8,
			expectedNextByte: 'D',
			expectedError:    errors.New("byte order mark is not found"),
		},
		{
			data:             []byte{0xEF, 0xBB, 0xBF, 'D'},
			encoding:         EncodingUTF8,
			expectedNextByte: 'D',
			expectedError:    nil,
		},
	}

	// Run the tests.
	for i, test := range tests {
		fmt.Print("[", i+1, "] ")

		r := bytes.NewReader(test.data)
		result.err = SkipBOM(r, test.encoding)

		if test.expectedError == nil {
			tst.MustBeNoError(result.err)
		} else {
			tst.MustBeAnError(result.err)
			tst.MustBeEqual(result.err, test.expectedError)
		}

		result.b, result.err = reader.ReadByte(r)
		tst.MustBeNoError(result.err)
		tst.MustBeEqual(result.b, test.expectedNextByte)
	}
	fmt.Println()
}

func Test_SearchForBOM(t *testing.T) {
	tst := tester.New(t)

	type TestData struct {
		data              []byte
		expectedEncodings []Encoding
		expectedN         int
		expectedError     error
	}

	type Result struct {
		encodings []Encoding
		n         int
		err       error
	}
	var result Result

	var tests = []TestData{
		// Test #1. Empty data.
		{
			data:              []byte{},
			expectedEncodings: []Encoding{},
			expectedN:         0,
			expectedError:     io.EOF,
		},

		// Test #2. One byte. Not a BOM.
		{
			data:              []byte{'A'}, // Not a BOM, 100%.
			expectedEncodings: []Encoding{},
			expectedN:         1,
			expectedError:     nil,
		},
		// Test #3. One byte. Possibly a BOM.
		{
			data:              []byte{0xDD}, // It may be the UTF-EBCDIC BOM.
			expectedEncodings: []Encoding{},
			expectedN:         1,
			expectedError:     io.EOF,
		},

		// Test #4. Two bytes. Not a BOM from the first byte.
		{
			data:              []byte{'A', 'B'}, // Not a BOM, 100%.
			expectedEncodings: []Encoding{},
			expectedN:         1,
			expectedError:     nil,
		},
		// Test #5. Two bytes. Not a BOM from the second byte.
		{
			data:              []byte{0xDD, 'B'}, // Not a BOM, 100%.
			expectedEncodings: []Encoding{},
			expectedN:         2,
			expectedError:     nil,
		},
		// Test #6. Two bytes. Mystery of colliding BOMs (0xFF 0xFE).
		{
			data:              []byte{0xFF, 0xFE}, // Mystery.
			expectedEncodings: []Encoding{},
			expectedN:         2,
			expectedError:     io.EOF,
		},

		// Test #7. Three bytes. Not a BOM from the first byte.
		{
			data:              []byte{'A', 'B', 'C'}, // Not a BOM, 100%.
			expectedEncodings: []Encoding{},
			expectedN:         1,
			expectedError:     nil,
		},
		// Test #8. Three bytes. Not a BOM from the second byte.
		{
			data:              []byte{0xDD, 'B', 'C'}, // Not a BOM, 100%.
			expectedEncodings: []Encoding{},
			expectedN:         2,
			expectedError:     nil,
		},
		// Test #9. Three bytes. Two-byte BOM.
		{
			data:              []byte{0xFE, 0xFF, 'C'}, // BOM.
			expectedEncodings: []Encoding{EncodingUTF16BE},
			expectedN:         2,
			expectedError:     nil,
		},
		// Test #10. Three bytes. Three-byte BOM.
		{
			data:              []byte{0xEF, 0xBB, 0xBF}, // BOM.
			expectedEncodings: []Encoding{EncodingUTF8},
			expectedN:         3,
			expectedError:     nil,
		},
		// Test #11. Three bytes. Possibly a BOM.
		{
			data:              []byte{0x00, 0x00, 0xFE}, // It may be UTF-32 BE.
			expectedEncodings: []Encoding{},
			expectedN:         3,
			expectedError:     io.EOF,
		},

		// Test #12. Four bytes. Not a BOM from the first byte.
		{
			data:              []byte{'A', 'B', 'C', 'D'}, // Not a BOM, 100%.
			expectedEncodings: []Encoding{},
			expectedN:         1,
			expectedError:     nil,
		},
		// Test #13. Four bytes. Not a BOM from the second byte.
		{
			data:              []byte{0xDD, 'B', 'C', 'D'}, // Not a BOM, 100%.
			expectedEncodings: []Encoding{},
			expectedN:         2,
			expectedError:     nil,
		},
		// Test #14. Four bytes. Not a BOM from the third byte.
		{
			data:              []byte{0xEF, 0xBB, 'C', 'D'}, // Not a BOM, 100%.
			expectedEncodings: []Encoding{},
			expectedN:         3,
			expectedError:     nil,
		},
		// Test #15. Four bytes. Not a BOM from the fourth byte.
		{
			data:              []byte{0x00, 0x00, 0xFE, 'D'}, // Not a BOM, 100%.
			expectedEncodings: []Encoding{},
			expectedN:         4,
			expectedError:     nil,
		},
		// Test #16. Four bytes. Two-byte BOM.
		{
			data:              []byte{0xFF, 0xFE, 'C', 'D'}, // BOM.
			expectedEncodings: []Encoding{EncodingUTF16LE},
			expectedN:         3,
			expectedError:     nil,
		},
		// Test #17. Four bytes. Mystery.
		{
			data:              []byte{0xFF, 0xFE, 0x00, 0x00}, // Two colliding BOMs.
			expectedEncodings: []Encoding{EncodingUTF16LE, EncodingUTF32LE},
			expectedN:         4,
			expectedError:     nil,
		},
		// Test #18. Four bytes. Three-byte BOM.
		{
			data:              []byte{0xEF, 0xBB, 0xBF, 'D'}, // BOM.
			expectedEncodings: []Encoding{EncodingUTF8},
			expectedN:         3,
			expectedError:     nil,
		},
		// Test #18. Four bytes. Four-byte BOM.
		{
			data:              []byte{0x00, 0x00, 0xFE, 0xFF}, // BOM.
			expectedEncodings: []Encoding{EncodingUTF32BE},
			expectedN:         4,
			expectedError:     nil,
		},
	}

	// Run the tests.
	for i, test := range tests {
		fmt.Print("[", i+1, "] ")

		r := bytes.NewReader(test.data)
		result.encodings, result.n, result.err = SearchForBOM(r)

		if test.expectedError == nil {
			tst.MustBeNoError(result.err)
		} else {
			tst.MustBeAnError(result.err)
			tst.MustBeEqual(result.err, test.expectedError)
		}
		tst.MustBeEqual(result.n, test.expectedN)

		// Array items are randomly placed because they are taken from the map.
		// We need to sort the result before checking it.
		sortEncodings(result.encodings)
		tst.MustBeEqual(result.encodings, test.expectedEncodings)
	}
	fmt.Println()
}

func Test_getMaximumBOMSize(t *testing.T) {
	tst := tester.New(t)

	// 2.
	tst.MustBeEqual(getMaximumBOMSize([]Encoding{EncodingUTF16BE}), 2)
	tst.MustBeEqual(getMaximumBOMSize([]Encoding{EncodingUTF16LE}), 2)

	// 3.
	tst.MustBeEqual(getMaximumBOMSize([]Encoding{EncodingUTF8}), 3)
	tst.MustBeEqual(getMaximumBOMSize([]Encoding{EncodingUTF7}), 3)
	tst.MustBeEqual(getMaximumBOMSize([]Encoding{EncodingUTF1}), 3)
	tst.MustBeEqual(getMaximumBOMSize([]Encoding{EncodingSCSU}), 3)
	tst.MustBeEqual(getMaximumBOMSize([]Encoding{EncodingBOCU1}), 3)

	// 4.
	tst.MustBeEqual(getMaximumBOMSize([]Encoding{EncodingUTF32BE}), 4)
	tst.MustBeEqual(getMaximumBOMSize([]Encoding{EncodingUTF32LE}), 4)
	tst.MustBeEqual(getMaximumBOMSize([]Encoding{EncodingUTF_EBCDIC}), 4)
	tst.MustBeEqual(getMaximumBOMSize([]Encoding{EncodingGB18030}), 4)

	// Mix.
	tst.MustBeEqual(getMaximumBOMSize([]Encoding{EncodingUTF16BE, EncodingUTF8, EncodingUTF32BE}), 4)
}
