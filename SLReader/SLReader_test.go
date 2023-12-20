package slreader

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/tester"
	"golang.org/x/time/rate"
)

func Test_NewReader(t *testing.T) {
	aTest := tester.New(t)

	type TestData struct {
		r           io.Reader
		normalLimit float64
		burstLimit  int
		bnr         float64
		//
		expectedSLR   *SLReader
		expectedError error
	}

	type Result struct {
		slr *SLReader
		err error
	}
	var result Result

	var tests = []TestData{
		{
			r:           bytes.NewReader([]byte{}),
			normalLimit: 0,
			burstLimit:  1,
			bnr:         10,
			//
			expectedSLR:   (*SLReader)(nil),
			expectedError: errors.New(ErrZeroSpeed),
		},
		{
			r:           bytes.NewReader([]byte{}),
			normalLimit: 10,
			burstLimit:  90,
			bnr:         2,
			//
			expectedSLR:   (*SLReader)(nil),
			expectedError: fmt.Errorf(ErrBurstToNormalRatioOverflow, 2, 9),
		},
		{
			r:           bytes.NewReader([]byte{1, 2, 3}),
			normalLimit: 10,
			burstLimit:  20,
			bnr:         2,
			//
			expectedSLR:   &SLReader{}, // Non-nil, see below.
			expectedError: nil,
		},
	}

	// Run the tests.
	for i, test := range tests {
		fmt.Print("[", i+1, "] ")

		result.slr, result.err = NewReader(test.r, test.normalLimit, test.burstLimit, test.bnr)

		if test.expectedError == nil {
			aTest.MustBeNoError(result.err)
		} else {
			aTest.MustBeAnError(result.err)
			aTest.MustBeEqual(result.err, test.expectedError)
		}

		if test.expectedSLR == nil {
			aTest.MustBeEqual(result.slr, test.expectedSLR)
		} else {
			aTest.MustBeDifferent(result.slr, (*SLReader)(nil))
		}
	}
	fmt.Println()
}

func Test_ChangeLimits(t *testing.T) {
	aTest := tester.New(t)

	type TestData struct {
		limiterNormalLimit float64
		limiterBurstLimit  float64

		// This parameter controls how many tokens to subtract in order to
		// emulate a working limiter at some point in time.
		tokensToSpare float64

		// New limits.
		normalLimit float64
		burstLimit  int

		// Expected results.
		expectedError         error
		expectedLimiterTokens float64
	}

	type Result struct {
		err error
	}
	var result Result

	var tests = []TestData{
		{
			limiterNormalLimit: 10,
			limiterBurstLimit:  20, // 20.
			tokensToSpare:      15, // 15.

			normalLimit: 8,
			burstLimit:  16,

			expectedError:         nil,
			expectedLimiterTokens: 5, // min(5; 16) = 5.
		},
		{
			limiterNormalLimit: 10,
			limiterBurstLimit:  20, // 20.
			tokensToSpare:      15, // 15.

			normalLimit: 31,
			burstLimit:  32,

			expectedError:         nil,
			expectedLimiterTokens: 5, // min(5; 32) = 5.
		},
		{
			limiterNormalLimit: 10,
			limiterBurstLimit:  20, // 20.
			tokensToSpare:      15, // 15.

			normalLimit: 1,
			burstLimit:  2,

			expectedError:         nil,
			expectedLimiterTokens: 2, // min(5; 2) = 2.
		},
	}

	// Run the tests.
	for i, test := range tests {
		fmt.Print("[", i+1, "] ")

		slr := &SLReader{
			limiter: rate.NewLimiter(rate.Limit(test.limiterNormalLimit), int(test.limiterBurstLimit)),
			bnr:     999,
		}
		slr.limiter.AllowN(time.Now(), int(test.tokensToSpare))

		result.err = slr.ChangeLimits(test.normalLimit, test.burstLimit)
		if test.expectedError != nil {
			aTest.MustBeAnError(result.err)
			aTest.MustBeEqual(result.err, test.expectedError)
		} else {
			aTest.MustBeNoError(result.err)

			aTest.MustBeEqual(slr.limiter.Limit(), rate.Limit(test.normalLimit))
			aTest.MustBeEqual(slr.limiter.Burst(), test.burstLimit)
			aTest.MustBeEqual(slr.limiter.Tokens(), test.expectedLimiterTokens)
		}
	}
	fmt.Println()
}

func Test_Read(t *testing.T) {
	aTest := tester.New(t)

	type TestData struct {
		r           io.Reader
		normalLimit float64
		burstLimit  int
		bnr         float64
		dstSize     int
		//
		expectedN           int
		expectedError       error
		expectedDurationSec time.Duration
	}

	type Result struct {
		n int
	}
	var result Result

	var tests = []TestData{
		{
			r:           bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			normalLimit: 1,
			burstLimit:  5,
			bnr:         5,
			dstSize:     5,
			//
			expectedN:           5,
			expectedDurationSec: 5,
		},
	}

	// Run the tests.
	for i, test := range tests {
		fmt.Print("[", i+1, "] ")

		t1 := time.Now()
		slr, err := NewReader(test.r, test.normalLimit, test.burstLimit, test.bnr)
		aTest.MustBeNoError(err)

		dst := make([]byte, test.dstSize)
		result.n, err = slr.Read(dst)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result.n, test.expectedN)

		durationSec := time.Now().Sub(t1).Seconds()
		aTest.MustBeEqual(math.Round(durationSec), math.Round(float64(test.expectedDurationSec)))
	}
	fmt.Println()
}

func Test_Close(t *testing.T) {
	aTest := *tester.New(t)
	var err error
	var slr = &SLReader{}
	err = slr.Close()
	aTest.MustBeNoError(err)
}
