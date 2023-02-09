// convertBytesToBits_test.go.

package bit

import (
	"reflect"
	"testing"
)

type TestDataForConvertBytesToBits struct {

	// Name of the Test.
	testName string

	// Input.
	bytes []byte

	// Output.
	bitsExpected []Bit
	bitsGot      []Bit
}

type TestRunnerForConvertBytesToBits = func(
	t *testing.T,
	testData *TestDataForConvertBytesToBits,
)

func Test_ConvertBytesToBits(t *testing.T) {

	var testData TestDataForConvertBytesToBits
	var testRunner TestRunnerForConvertBytesToBits
	var testsData []TestDataForConvertBytesToBits

	// Test Runner.
	testRunner = func(
		t *testing.T,
		testData *TestDataForConvertBytesToBits,
	) {
		// Fool Check.
		if testData == nil {
			t.Error(ErrNullPointer)
			t.FailNow()
		}

		// Run the tested Function.
		testData.bitsGot = ConvertBytesToBits(testData.bytes)

		// Check the Results.
		if !reflect.DeepEqual(testData.bitsGot, testData.bitsExpected) {
			t.Error(testData.testName, TestNameErrDelimiter, ErrBitsMismatch)
			t.FailNow()
		}
	}

	// Test #1: No Bytes.
	testData = TestDataForConvertBytesToBits{
		testName:     "Test #1",
		bytes:        []byte{},
		bitsExpected: []Bit{},
	}
	testsData = append(testsData, testData)

	// Test #2: One Byte 0101 0101.
	testData = TestDataForConvertBytesToBits{
		testName: "Test #2",
		bytes:    []byte{85},
		bitsExpected: []Bit{
			One,  // 1.
			Zero, // 2.
			One,  // 3.
			Zero, // 4.
			One,  // 5.
			Zero, // 6.
			One,  // 7.
			Zero, // 8.
		},
	}
	testsData = append(testsData, testData)

	// Test #3: Two Bytes 0101 0101; 1010 1010.
	testData = TestDataForConvertBytesToBits{
		testName: "Test #3",
		bytes: []byte{
			85,
			170,
		},
		bitsExpected: []Bit{
			// Byte #1.
			One,  // 1.
			Zero, // 2.
			One,  // 3.
			Zero, // 4.
			One,  // 5.
			Zero, // 6.
			One,  // 7.
			Zero, // 8.
			// Byte #2.
			Zero, // 1.
			One,  // 2.
			Zero, // 3.
			One,  // 4.
			Zero, // 5.
			One,  // 6.
			Zero, // 7.
			One,  // 8.
		},
	}
	testsData = append(testsData, testData)

	// Run the Tests.
	for _, testData = range testsData {
		testRunner(t, &testData)
	}
}
