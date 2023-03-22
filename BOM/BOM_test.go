package bom

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/vault-thirteen/auxie/reader"
	"github.com/vault-thirteen/tester"
)

func Test_GetEncoding(t *testing.T) {
	tst := tester.New(t)
	var err error
	var result []Encoding
	var resultEncodingUnknown = []Encoding{EncodingUnknown}
	var r *bytes.Reader

	type TestData struct {
		data           []byte
		isPedanticTest bool
		expectedResult []Encoding
	}

	var testsv1 = []TestData{
		// Test #1. Unknown encoding.
		{
			data:           []byte{},
			isPedanticTest: true,
			expectedResult: resultEncodingUnknown,
		},

		// Test #2. UTF-8.
		{
			data:           []byte{0xEF, 0xBB, 0xBF},
			isPedanticTest: true,
			expectedResult: []Encoding{EncodingUTF8, EncodingUnknown},
		},

		// Test #3. UTF-16 (BE).
		{
			data:           []byte{0xFE, 0xFF},
			isPedanticTest: true,
			expectedResult: []Encoding{EncodingUTF16BE, EncodingUnknown},
		},

		// Test #4. UTF-16 (LE).
		{
			data:           []byte{0xFF, 0xFE},
			isPedanticTest: true,
			expectedResult: []Encoding{EncodingUTF16LE, EncodingUnknown},
		},

		// Test #5. UTF-32 (BE).
		{
			data:           []byte{0x00, 0x00, 0xFE, 0xFF},
			isPedanticTest: true,
			expectedResult: []Encoding{EncodingUTF32BE, EncodingUnknown},
		},

		// Test #6. UTF-32 (LE).
		{
			data:           []byte{0xFF, 0xFE, 0x00, 0x00},
			isPedanticTest: true,
			expectedResult: []Encoding{EncodingUTF16LE, EncodingUTF32LE, EncodingUnknown},
		},

		// Test #7. UTF-7.
		{
			data:           []byte{0x2B, 0x2F, 0x76},
			isPedanticTest: true,
			expectedResult: []Encoding{EncodingUTF7, EncodingUnknown},
		},

		// Test #8. UTF-1.
		{
			data:           []byte{0xF7, 0x64, 0x4C},
			isPedanticTest: true,
			expectedResult: []Encoding{EncodingUTF1, EncodingUnknown},
		},

		// Test #9. UTF-EBCDIC.
		{
			data:           []byte{0xDD, 0x73, 0x66, 0x73},
			isPedanticTest: true,
			expectedResult: []Encoding{EncodingUTF_EBCDIC, EncodingUnknown},
		},

		// Test #10. SCSU.
		{
			data:           []byte{0x0E, 0xFE, 0xFF},
			isPedanticTest: true,
			expectedResult: []Encoding{EncodingSCSU, EncodingUnknown},
		},

		// Test #11. BOCU-1.
		{
			data:           []byte{0xFB, 0xEE, 0x28},
			isPedanticTest: true,
			expectedResult: []Encoding{EncodingBOCU1, EncodingUnknown},
		},

		// Test #12. GB18030.
		{
			data:           []byte{0x84, 0x31, 0x95, 0x33},
			isPedanticTest: true,
			expectedResult: []Encoding{EncodingGB18030, EncodingUnknown},
		},
	}

	// Add a second version of each test.
	testsv2 := make([]TestData, 0, len(testsv1))

	for _, test := range testsv1 {
		testsv2 = append(testsv2, TestData{
			data:           test.data,
			isPedanticTest: false,
			expectedResult: test_withoutEncoding(test.expectedResult, EncodingUnknown),
		})
	}

	// Combine the tests.
	tests := make([]TestData, 0, len(testsv1)+len(testsv2))
	for _, tv1 := range testsv1 {
		tests = append(tests, tv1)
	}
	for _, tv2 := range testsv2 {
		tests = append(tests, tv2)
	}

	// Run the tests.
	for i, test := range tests {
		fmt.Print("[", i+1, "] ")

		r = bytes.NewReader(test.data)
		result, err = GetEncoding(r, test.isPedanticTest)
		tst.MustBeNoError(err)
		tst.MustBeEqual(result, test.expectedResult)
	}
	fmt.Println()
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

	var testsv1 = []TestData{
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

	// Add a second version of each test.
	testsv2 := make([]TestData, 0, len(testsv1))
	deltaBytes := []byte{'A', 'B', 'C'}
	for _, test := range testsv1 {
		testsv2 = append(testsv2, TestData{
			Bytes:        append(test.Bytes, deltaBytes...),
			Encoding:     test.Encoding,
			ExpectedData: deltaBytes,
		})
	}

	// Combine the tests.
	tests := make([]TestData, 0, len(testsv1)+len(testsv2))
	for _, tv1 := range testsv1 {
		tests = append(tests, tv1)
	}
	for _, tv2 := range testsv2 {
		tests = append(tests, tv2)
	}

	// Run the tests.
	for i, theTest := range tests {
		fmt.Print("[", i+1, "] ")

		r = bytes.NewReader(theTest.Bytes)
		newRS, err = SkipBOMPrefix(r, theTest.Encoding)
		tst.MustBeNoError(err)
		restBytes, err = io.ReadAll(newRS)
		tst.MustBeNoError(err)
		tst.MustBeEqual(restBytes, theTest.ExpectedData)
	}
	fmt.Println()
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

// test_withoutEncoding deletes an encoding from an array.
func test_withoutEncoding(old []Encoding, del Encoding) (new []Encoding) {
	new = make([]Encoding, 0, len(old))
	for _, o := range old {
		if o == del {
			continue
		}
		new = append(new, o)
	}
	return new
}
