package bom

import (
	"bytes"
	"io"
	"testing"

	"github.com/vault-thirteen/auxie/reader"
	"github.com/vault-thirteen/tester"
)

func Test_GetEncoding(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result []Encoding
	var unknownEnc = []Encoding{EncodingUnknown}
	var r *bytes.Reader

	// Test #1. Unknown encoding.
	r = bytes.NewReader([]byte{})
	result, err = GetEncoding(r)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, unknownEnc)

	// Test #2. UTF-8.
	r = bytes.NewReader([]byte{0xEF, 0xBB, 0xBF})
	result, err = GetEncoding(r)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []Encoding{EncodingUTF8, EncodingUnknown})

	// Test #3. UTF-16 (BE).
	r = bytes.NewReader([]byte{0xFE, 0xFF})
	result, err = GetEncoding(r)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []Encoding{EncodingUTF16BE, EncodingUnknown})

	// Test #4. UTF-16 (LE).
	r = bytes.NewReader([]byte{0xFF, 0xFE})
	result, err = GetEncoding(r)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []Encoding{EncodingUTF16LE, EncodingUnknown})

	// Test #5. UTF-32 (BE).
	r = bytes.NewReader([]byte{0x00, 0x00, 0xFE, 0xFF})
	result, err = GetEncoding(r)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []Encoding{EncodingUTF32BE, EncodingUnknown})

	// Test #6. UTF-32 (LE).
	r = bytes.NewReader([]byte{0xFF, 0xFE, 0x00, 0x00})
	result, err = GetEncoding(r)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []Encoding{EncodingUTF16LE, EncodingUTF32LE, EncodingUnknown})

	// Test #7. UTF-7.
	r = bytes.NewReader([]byte{0x2B, 0x2F, 0x76})
	result, err = GetEncoding(r)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []Encoding{EncodingUTF7, EncodingUnknown})

	// Test #8. UTF-1.
	r = bytes.NewReader([]byte{0xF7, 0x64, 0x4C})
	result, err = GetEncoding(r)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []Encoding{EncodingUTF1, EncodingUnknown})

	// Test #9. UTF-EBCDIC.
	r = bytes.NewReader([]byte{0xDD, 0x73, 0x66, 0x73})
	result, err = GetEncoding(r)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []Encoding{EncodingUTF_EBCDIC, EncodingUnknown})

	// Test #10. SCSU.
	r = bytes.NewReader([]byte{0x0E, 0xFE, 0xFF})
	result, err = GetEncoding(r)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []Encoding{EncodingSCSU, EncodingUnknown})

	// Test #11. BOCU-1.
	r = bytes.NewReader([]byte{0xFB, 0xEE, 0x28})
	result, err = GetEncoding(r)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []Encoding{EncodingBOCU1, EncodingUnknown})

	// Test #12. GB18030.
	r = bytes.NewReader([]byte{0x84, 0x31, 0x95, 0x33})
	result, err = GetEncoding(r)
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, []Encoding{EncodingGB18030, EncodingUnknown})
}

func Test_IsEncoding(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result bool
	var r *bytes.Reader

	type TestData struct {
		Bytes          []byte
		Encoding       Encoding
		ExpectedResult bool
	}

	var tests = []TestData{
		// Positive tests.
		{
			Bytes:          []byte{0xEF, 0xBB, 0xBF},
			Encoding:       EncodingUTF8,
			ExpectedResult: true,
		},
		{
			Bytes:          []byte{0xFE, 0xFF},
			Encoding:       EncodingUTF16BE,
			ExpectedResult: true,
		},
		{
			Bytes:          []byte{0xFF, 0xFE},
			Encoding:       EncodingUTF16LE,
			ExpectedResult: true,
		},
		{
			Bytes:          []byte{0x00, 0x00, 0xFE, 0xFF},
			Encoding:       EncodingUTF32BE,
			ExpectedResult: true,
		},
		{
			Bytes:          []byte{0xFF, 0xFE, 0x00, 0x00},
			Encoding:       EncodingUTF32LE,
			ExpectedResult: true,
		},
		{
			Bytes:          []byte{0x2B, 0x2F, 0x76},
			Encoding:       EncodingUTF7,
			ExpectedResult: true,
		},
		{
			Bytes:          []byte{0xF7, 0x64, 0x4C},
			Encoding:       EncodingUTF1,
			ExpectedResult: true,
		},
		{
			Bytes:          []byte{0xDD, 0x73, 0x66, 0x73},
			Encoding:       EncodingUTF_EBCDIC,
			ExpectedResult: true,
		},
		{
			Bytes:          []byte{0x0E, 0xFE, 0xFF},
			Encoding:       EncodingSCSU,
			ExpectedResult: true,
		},
		{
			Bytes:          []byte{0xFB, 0xEE, 0x28},
			Encoding:       EncodingBOCU1,
			ExpectedResult: true,
		},
		{
			Bytes:          []byte{0x84, 0x31, 0x95, 0x33},
			Encoding:       EncodingGB18030,
			ExpectedResult: true,
		},

		// Negative tests.
		{
			Bytes:          []byte{},
			Encoding:       EncodingUTF8,
			ExpectedResult: false,
		},
		{
			Bytes:          []byte{},
			Encoding:       EncodingUTF16BE,
			ExpectedResult: false,
		},
		{
			Bytes:          []byte{},
			Encoding:       EncodingUTF16LE,
			ExpectedResult: false,
		},
		{
			Bytes:          []byte{},
			Encoding:       EncodingUTF32BE,
			ExpectedResult: false,
		},
		{
			Bytes:          []byte{},
			Encoding:       EncodingUTF32LE,
			ExpectedResult: false,
		},
		{
			Bytes:          []byte{},
			Encoding:       EncodingUTF7,
			ExpectedResult: false,
		},
		{
			Bytes:          []byte{},
			Encoding:       EncodingUTF1,
			ExpectedResult: false,
		},
		{
			Bytes:          []byte{},
			Encoding:       EncodingUTF_EBCDIC,
			ExpectedResult: false,
		},
		{
			Bytes:          []byte{},
			Encoding:       EncodingSCSU,
			ExpectedResult: false,
		},
		{
			Bytes:          []byte{},
			Encoding:       EncodingBOCU1,
			ExpectedResult: false,
		},
		{
			Bytes:          []byte{},
			Encoding:       EncodingGB18030,
			ExpectedResult: false,
		},
	}

	for _, theTest := range tests {
		r = bytes.NewReader(theTest.Bytes)
		result, err = IsEncoding(r, theTest.Encoding)
		tst.MustBeNoError(err)
		tst.MustBeEqual(result, theTest.ExpectedResult)
	}
}

func Test_SkipBOMPrefix(t *testing.T) {
	tst := tester.New(t)
	var err error
	var newRS io.Reader
	var r *bytes.Reader
	var restBytes []byte

	type TestData struct {
		Bytes        []byte
		Encoding     Encoding
		ExpectedData []byte
	}

	var tests = []TestData{
		{
			Bytes:        []byte{0xEF, 0xBB, 0xBF},
			Encoding:     EncodingUTF8,
			ExpectedData: []byte{},
		},
		{
			Bytes:        []byte{0xFE, 0xFF},
			Encoding:     EncodingUTF16BE,
			ExpectedData: []byte{},
		},
		{
			Bytes:        []byte{0xFF, 0xFE},
			Encoding:     EncodingUTF16LE,
			ExpectedData: []byte{},
		},
		{
			Bytes:        []byte{0x00, 0x00, 0xFE, 0xFF},
			Encoding:     EncodingUTF32BE,
			ExpectedData: []byte{},
		},
		{
			Bytes:        []byte{0xFF, 0xFE, 0x00, 0x00},
			Encoding:     EncodingUTF32LE,
			ExpectedData: []byte{},
		},
		{
			Bytes:        []byte{0x2B, 0x2F, 0x76},
			Encoding:     EncodingUTF7,
			ExpectedData: []byte{},
		},
		{
			Bytes:        []byte{0xF7, 0x64, 0x4C},
			Encoding:     EncodingUTF1,
			ExpectedData: []byte{},
		},
		{
			Bytes:        []byte{0xDD, 0x73, 0x66, 0x73},
			Encoding:     EncodingUTF_EBCDIC,
			ExpectedData: []byte{},
		},
		{
			Bytes:        []byte{0x0E, 0xFE, 0xFF},
			Encoding:     EncodingSCSU,
			ExpectedData: []byte{},
		},
		{
			Bytes:        []byte{0xFB, 0xEE, 0x28},
			Encoding:     EncodingBOCU1,
			ExpectedData: []byte{},
		},
		{
			Bytes:        []byte{0x84, 0x31, 0x95, 0x33},
			Encoding:     EncodingGB18030,
			ExpectedData: []byte{},
		},
	}

	// Add second version of each test.
	testsv2 := make([]TestData, 0, len(tests))
	deltaBytes := []byte{'A', 'B', 'C'}
	for _, test := range tests {
		testsv2 = append(testsv2, TestData{
			Bytes:        append(test.Bytes, deltaBytes...),
			Encoding:     test.Encoding,
			ExpectedData: deltaBytes,
		})
	}

	for _, theTest := range tests {
		r = bytes.NewReader(theTest.Bytes)
		newRS, err = SkipBOMPrefix(r, theTest.Encoding)
		tst.MustBeNoError(err)
		restBytes, err = io.ReadAll(newRS)
		tst.MustBeNoError(err)
		tst.MustBeEqual(restBytes, theTest.ExpectedData)
	}
}

// Test_ALL_GetBOM tests all 11 functions that get BOM, i.e.
// all the functions with name similar to GetBOMForEncoding...
func Test_ALL_GetBOM(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result []byte
	var r *reader.Reader

	type TestData struct {
		Bytes           []byte
		MethodToCheck   func(r *reader.Reader) (prefix []byte, err error)
		ExpectedPrefix  []byte
		IsErrorExpected bool
	}

	var tests = []TestData{
		// Positive tests.
		{
			Bytes:           []byte{0xEF, 0xBB, 0xBF},
			ExpectedPrefix:  []byte{0xEF, 0xBB, 0xBF},
			MethodToCheck:   GetBOMForEncodingUTF8,
			IsErrorExpected: false,
		},
		{
			Bytes:           []byte{0xFE, 0xFF},
			ExpectedPrefix:  []byte{0xFE, 0xFF},
			MethodToCheck:   GetBOMForEncodingUTF16BE,
			IsErrorExpected: false,
		},
		{
			Bytes:           []byte{0xFF, 0xFE},
			ExpectedPrefix:  []byte{0xFF, 0xFE},
			MethodToCheck:   GetBOMForEncodingUTF16LE,
			IsErrorExpected: false,
		},
		{
			Bytes:           []byte{0x00, 0x00, 0xFE, 0xFF},
			ExpectedPrefix:  []byte{0x00, 0x00, 0xFE, 0xFF},
			MethodToCheck:   GetBOMForEncodingUTF32BE,
			IsErrorExpected: false,
		},
		{
			Bytes:           []byte{0xFF, 0xFE, 0x00, 0x00},
			ExpectedPrefix:  []byte{0xFF, 0xFE, 0x00, 0x00},
			MethodToCheck:   GetBOMForEncodingUTF32LE,
			IsErrorExpected: false,
		},
		{
			Bytes:           []byte{0x2B, 0x2F, 0x76},
			ExpectedPrefix:  []byte{0x2B, 0x2F, 0x76},
			MethodToCheck:   GetBOMForEncodingUTF7,
			IsErrorExpected: false,
		},
		{
			Bytes:           []byte{0xF7, 0x64, 0x4C},
			ExpectedPrefix:  []byte{0xF7, 0x64, 0x4C},
			MethodToCheck:   GetBOMForEncodingUTF1,
			IsErrorExpected: false,
		},
		{
			Bytes:           []byte{0xDD, 0x73, 0x66, 0x73},
			ExpectedPrefix:  []byte{0xDD, 0x73, 0x66, 0x73},
			MethodToCheck:   GetBOMForEncodingUTF_EBCDIC,
			IsErrorExpected: false,
		},
		{
			Bytes:           []byte{0x0E, 0xFE, 0xFF},
			ExpectedPrefix:  []byte{0x0E, 0xFE, 0xFF},
			MethodToCheck:   GetBOMForEncodingSCSU,
			IsErrorExpected: false,
		},
		{
			Bytes:           []byte{0xFB, 0xEE, 0x28},
			ExpectedPrefix:  []byte{0xFB, 0xEE, 0x28},
			MethodToCheck:   GetBOMForEncodingBOCU1,
			IsErrorExpected: false,
		},
		{
			Bytes:           []byte{0x84, 0x31, 0x95, 0x33},
			ExpectedPrefix:  []byte{0x84, 0x31, 0x95, 0x33},
			MethodToCheck:   GetBOMForEncodingGB18030,
			IsErrorExpected: false,
		},

		// Negative tests.
		{
			Bytes:           []byte{},
			ExpectedPrefix:  []byte(nil),
			MethodToCheck:   GetBOMForEncodingUTF8,
			IsErrorExpected: true,
		},
		{
			Bytes:           []byte{},
			ExpectedPrefix:  []byte(nil),
			MethodToCheck:   GetBOMForEncodingUTF16BE,
			IsErrorExpected: true,
		},
		{
			Bytes:           []byte{},
			ExpectedPrefix:  []byte(nil),
			MethodToCheck:   GetBOMForEncodingUTF16LE,
			IsErrorExpected: true,
		},
		{
			Bytes:           []byte{},
			ExpectedPrefix:  []byte(nil),
			MethodToCheck:   GetBOMForEncodingUTF32BE,
			IsErrorExpected: true,
		},
		{
			Bytes:           []byte{},
			ExpectedPrefix:  []byte(nil),
			MethodToCheck:   GetBOMForEncodingUTF32LE,
			IsErrorExpected: true,
		},
		{
			Bytes:           []byte{},
			ExpectedPrefix:  []byte(nil),
			MethodToCheck:   GetBOMForEncodingUTF7,
			IsErrorExpected: true,
		},
		{
			Bytes:           []byte{},
			ExpectedPrefix:  []byte(nil),
			MethodToCheck:   GetBOMForEncodingUTF1,
			IsErrorExpected: true,
		},
		{
			Bytes:           []byte{},
			ExpectedPrefix:  []byte(nil),
			MethodToCheck:   GetBOMForEncodingUTF_EBCDIC,
			IsErrorExpected: true,
		},
		{
			Bytes:           []byte{},
			ExpectedPrefix:  []byte(nil),
			MethodToCheck:   GetBOMForEncodingSCSU,
			IsErrorExpected: true,
		},
		{
			Bytes:           []byte{},
			ExpectedPrefix:  []byte(nil),
			MethodToCheck:   GetBOMForEncodingBOCU1,
			IsErrorExpected: true,
		},
		{
			Bytes:           []byte{},
			ExpectedPrefix:  []byte(nil),
			MethodToCheck:   GetBOMForEncodingGB18030,
			IsErrorExpected: true,
		},
	}

	for _, theTest := range tests {
		r = reader.NewReader(bytes.NewReader(theTest.Bytes))
		result, err = theTest.MethodToCheck(r)
		if theTest.IsErrorExpected {
			tst.MustBeAnError(err)
		} else {
			tst.MustBeNoError(err)
		}
		tst.MustBeEqual(result, theTest.ExpectedPrefix)
	}
}
