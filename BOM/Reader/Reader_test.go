package reader

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	bom "github.com/vault-thirteen/auxie/BOM"
	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewReader(t *testing.T) {
	aTest := *tester.New(t)
	var br *Reader
	var err error

	// Test #1. Error.
	r := bytes.NewReader([]byte{0xEF})
	br, err = NewReader(r, false)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err, io.EOF)
	aTest.MustBeEqual(br, (*Reader)(nil))

	// Test #2. Skipping is enabled, two encodings are detected.
	r = bytes.NewReader([]byte{0xFF, 0xFE, 0x00, 0x00, 'E'})
	br, err = NewReader(r, true)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(br, (*Reader)(nil))

	// Test #3. Skipping is enabled, one encoding is detected.
	r = bytes.NewReader([]byte{0xEF, 0xBB, 0xBF, 'D'})
	br, err = NewReader(r, true)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(br, &Reader{
		r:             r,
		shouldSkipBOM: true,
		firstBytes:    []byte{},
		encodings:     []bom.Encoding{bom.EncodingUTF8},
	})

	// Test #4. Skipping is enabled, no encoding is detected.
	r = bytes.NewReader([]byte{'A', 'B', 'C', 'D'})
	br, err = NewReader(r, true)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(br, &Reader{
		r:             r,
		shouldSkipBOM: true,
		firstBytes:    []byte{'A'},
		encodings:     []bom.Encoding{},
	})

	// Test #5. No error, skipping is disabled.
	r = bytes.NewReader([]byte{0xEF, 0xBB, 0xBF, 'D'})
	br, err = NewReader(r, false)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(br, &Reader{
		r:             r,
		shouldSkipBOM: false,
		firstBytes:    []byte{0xEF, 0xBB, 0xBF},
		encodings:     []bom.Encoding{bom.EncodingUTF8},
	})
}

func Test_detectEncoding(t *testing.T) {
	aTest := *tester.New(t)
	var err error
	var br *Reader

	// Test #1. No error.
	br = &Reader{
		r: bytes.NewReader([]byte{0xEF, 0xBB, 0xBF, 'D'}),
	}
	err = br.detectEncoding()
	aTest.MustBeNoError(err)

	// Test #2. Error.
	br = &Reader{
		r: bytes.NewReader([]byte{0xEF}),
	}
	err = br.detectEncoding()
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err, io.EOF)
}

func Test_GetEncodings(t *testing.T) {
	aTest := *tester.New(t)
	var br *Reader

	// Test.
	br = &Reader{
		encodings: []bom.Encoding{bom.Encoding(123)},
	}
	aTest.MustBeEqual(br.GetEncodings(), []bom.Encoding{bom.Encoding(123)})
}

func Test_Read(t *testing.T) {
	aTest := *tester.New(t)
	var br *Reader
	var err error
	var dst []byte

	type TestData struct {
		data           []byte
		shouldSkipBOM  bool
		dstSize        int
		expectedN      int
		expectedError  error
		expectedDst    []byte
		expectedBrFb   []byte
		expectedBrEncs []bom.Encoding
	}

	type Result struct {
		n   int
		err error
	}
	var result Result

	var tests = []TestData{
		{
			data:           []byte{'A', 'B', 'C'},
			shouldSkipBOM:  true,
			dstSize:        1,
			expectedN:      1,
			expectedError:  nil,
			expectedDst:    []byte{'A'},
			expectedBrFb:   []byte{},
			expectedBrEncs: []bom.Encoding{},
		},
		{
			data:           []byte{0xEF, 0xBB, 0xBF, 'D', 'E'},
			shouldSkipBOM:  true,
			dstSize:        1,
			expectedN:      1,
			expectedError:  nil,
			expectedDst:    []byte{'D'},
			expectedBrFb:   []byte{},
			expectedBrEncs: []bom.Encoding{bom.EncodingUTF8},
		},
		{
			data:           []byte{0xEF, 0xBB, 0xBF, 'D', 'E'},
			shouldSkipBOM:  false,
			dstSize:        2,
			expectedN:      2,
			expectedError:  nil,
			expectedDst:    []byte{0xEF, 0xBB},
			expectedBrFb:   []byte{0xBF},
			expectedBrEncs: []bom.Encoding{bom.EncodingUTF8},
		},
		{
			data:          []byte{0xEF, 0xBB, 0xBF, 'D'},
			shouldSkipBOM: false,
			dstSize:       5,
			expectedN:     4,
			// /!\ ACHTUNG /!\ Error should be io.EOF, but fucking stupid Google developer decided not to return an
			// error from the 'Read' method of the 'byte.Reader' object !!!
			expectedError:  nil, // io.EOF !!!
			expectedDst:    []byte{0xEF, 0xBB, 0xBF, 'D', 0},
			expectedBrFb:   []byte{},
			expectedBrEncs: []bom.Encoding{bom.EncodingUTF8},
		},
		{
			data:           []byte{0xEF, 0xBB, 0xBF, 'D', 'E', 'F', 'G'},
			shouldSkipBOM:  false,
			dstSize:        5,
			expectedN:      5,
			expectedError:  nil,
			expectedDst:    []byte{0xEF, 0xBB, 0xBF, 'D', 'E'},
			expectedBrFb:   []byte{},
			expectedBrEncs: []bom.Encoding{bom.EncodingUTF8},
		},
	}

	// Run the tests.
	for i, test := range tests {
		fmt.Print("[", i+1, "] ")

		br, err = NewReader(bytes.NewReader(test.data), test.shouldSkipBOM)
		aTest.MustBeNoError(err)
		dst = make([]byte, test.dstSize)
		result.n, result.err = br.Read(dst)

		if test.expectedError == nil {
			aTest.MustBeNoError(result.err)
		} else {
			aTest.MustBeAnError(result.err)
			aTest.MustBeEqual(result.err, test.expectedError)
		}
		aTest.MustBeEqual(result.n, test.expectedN)
		aTest.MustBeEqual(dst, test.expectedDst)
		aTest.MustBeEqual(br.firstBytes, test.expectedBrFb)
		aTest.MustBeEqual(br.encodings, test.expectedBrEncs)
	}
	fmt.Println()
}

func Test_Close(t *testing.T) {
	aTest := *tester.New(t)
	var br *Reader
	var err error

	// Test #1.
	br = &Reader{firstBytes: []byte{1}}
	err = br.Close()
	aTest.MustBeAnError(err)

	// Test #2.
	br = &Reader{firstBytes: []byte{}}
	err = br.Close()
	aTest.MustBeNoError(err)
}
