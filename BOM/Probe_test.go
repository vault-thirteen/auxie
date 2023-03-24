package bom

import (
	"errors"
	"fmt"
	"testing"

	tsb "github.com/vault-thirteen/auxie/TSB"
	"github.com/vault-thirteen/tester"
)

func Test_IsAccurate(t *testing.T) {
	tst := tester.New(t)
	var probe *Probe

	// Test #1.
	probe = &Probe{
		Probability: tsb.Yes,
	}
	tst.MustBeEqual(probe.IsAccurate(), true)

	// Test #2.
	probe = &Probe{
		Probability: tsb.No,
	}
	tst.MustBeEqual(probe.IsAccurate(), true)

	// Test #3.
	probe = &Probe{
		Probability: tsb.Maybe,
	}
	tst.MustBeEqual(probe.IsAccurate(), false)
}

func Test_ProbeForEncoding(t *testing.T) {
	tst := tester.New(t)

	type TestData struct {
		data          []byte
		encoding      Encoding
		expectedProbe *Probe
		expectedError error
	}

	type Result struct {
		probe *Probe
		err   error
	}
	var result Result

	var tests = []TestData{
		{
			data:          []byte{},
			encoding:      Encoding(123),
			expectedProbe: nil,
			expectedError: errors.New("unknown encoding: 123"),
		},
		{
			data:     []byte{1, 2, 3, 4, 5},
			encoding: EncodingUTF8,
			expectedProbe: &Probe{
				Encoding:       EncodingUTF8,
				Probability:    tsb.No,
				ReadBytesCount: 3,
			},
			expectedError: nil,
		},
		{
			data:     []byte{0xEF, 0xBB, 0xBF, 4, 5},
			encoding: EncodingUTF8,
			expectedProbe: &Probe{
				Encoding:       EncodingUTF8,
				Probability:    tsb.Yes,
				ReadBytesCount: 3,
			},
			expectedError: nil,
		},
		{
			data:     []byte{0xEF},
			encoding: EncodingUTF8,
			expectedProbe: &Probe{
				Encoding:       EncodingUTF8,
				Probability:    tsb.Maybe,
				ReadBytesCount: 1,
			},
			expectedError: nil,
		},
		{
			data:     []byte{0xEF, 0xBB},
			encoding: EncodingUTF8,
			expectedProbe: &Probe{
				Encoding:       EncodingUTF8,
				Probability:    tsb.Maybe,
				ReadBytesCount: 2,
			},
			expectedError: nil,
		},
		{
			data:     []byte{1},
			encoding: EncodingUTF8,
			expectedProbe: &Probe{
				Encoding:       EncodingUTF8,
				Probability:    tsb.No,
				ReadBytesCount: 1,
			},
			expectedError: nil,
		},
		{
			data:     []byte{1, 2},
			encoding: EncodingUTF8,
			expectedProbe: &Probe{
				Encoding:       EncodingUTF8,
				Probability:    tsb.No,
				ReadBytesCount: 2,
			},
			expectedError: nil,
		},
		{
			data:     []byte{0xEF, 2},
			encoding: EncodingUTF8,
			expectedProbe: &Probe{
				Encoding:       EncodingUTF8,
				Probability:    tsb.No,
				ReadBytesCount: 2,
			},
			expectedError: nil,
		},
	}

	// Run the tests.
	for i, test := range tests {
		fmt.Print("[", i+1, "] ")

		result.probe, result.err = ProbeForEncoding(test.data, test.encoding)

		if test.expectedError == nil {
			tst.MustBeNoError(result.err)
		} else {
			tst.MustBeAnError(result.err)
			tst.MustBeEqual(result.err, test.expectedError)
		}
		tst.MustBeEqual(result.probe, test.expectedProbe)
	}
	fmt.Println()
}

func Test_countEqualConsecutiveBytes(t *testing.T) {
	tst := tester.New(t)

	type TestData struct {
		data1         []byte
		data2         []byte
		expectedN     int
		expectedError error
	}

	type Result struct {
		n   int
		err error
	}
	var result Result

	var tests = []TestData{
		{
			data1:         []byte{1},
			data2:         []byte{1, 2},
			expectedN:     0,
			expectedError: errors.New("arrays have different lengths: 1 vs 2"),
		},
		{
			data1:         []byte{},
			data2:         []byte{},
			expectedN:     0,
			expectedError: errors.New("no data"),
		},
		{
			data1:         []byte{1, 2, 3},
			data2:         []byte{1, 2, 4},
			expectedN:     2,
			expectedError: nil,
		},
		{
			data1:         []byte{1, 2, 3},
			data2:         []byte{1, 9, 8},
			expectedN:     1,
			expectedError: nil,
		},
		{
			data1:         []byte{1, 2, 3},
			data2:         []byte{9, 8, 7},
			expectedN:     0,
			expectedError: nil,
		},
	}

	// Run the tests.
	for i, test := range tests {
		fmt.Print("[", i+1, "] ")

		result.n, result.err = countEqualConsecutiveBytes(test.data1, test.data2)

		if test.expectedError == nil {
			tst.MustBeNoError(result.err)
		} else {
			tst.MustBeAnError(result.err)
			tst.MustBeEqual(result.err, test.expectedError)
		}
		tst.MustBeEqual(result.n, test.expectedN)
	}
	fmt.Println()
}
