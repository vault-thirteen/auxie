// convertBitsToBytes_test.go.

package bit

import (
	"bytes"
	"testing"
)

type TestDataForConvertBitsToBytes struct {

	// Name of the Test.
	testName string

	// Input.
	bits []Bit

	// Output.
	bitsCountExpected int
	bitsCountGot      int
	bytesExpected     []byte
	bytesGot          []byte
}

type TestRunnerForConvertBitsToBytes = func(
	t *testing.T,
	testData *TestDataForConvertBitsToBytes,
)

func Test_ConvertBitsToBytes(t *testing.T) {

	var testData TestDataForConvertBitsToBytes
	var testRunner TestRunnerForConvertBitsToBytes
	var testsData []TestDataForConvertBitsToBytes

	// Test Runner.
	testRunner = func(
		t *testing.T,
		testData *TestDataForConvertBitsToBytes,
	) {
		// Fool Check.
		if testData == nil {
			t.Error(ErrNullPointer)
			t.FailNow()
		}

		// Run the tested Function.
		testData.bytesGot, testData.bitsCountGot = ConvertBitsToBytes(
			testData.bits,
		)

		// Check the Results.
		if testData.bitsCountGot != testData.bitsCountExpected {
			t.Error(testData.testName, TestNameErrDelimiter, ErrBitsCountMismatch)
			t.FailNow()
		}
		if !bytes.Equal(testData.bytesGot, testData.bytesExpected) {
			t.Error(testData.testName, TestNameErrDelimiter, ErrBytesMismatch)
			t.FailNow()
		}
	}

	// Test #1: No Bits.
	testData = TestDataForConvertBitsToBytes{
		testName:          "Test #1",
		bits:              []Bit{},
		bytesExpected:     []byte{},
		bitsCountExpected: 0,
	}
	testsData = append(testsData, testData)

	// Test #2. Three Bits => One partial Byte.
	testData = TestDataForConvertBitsToBytes{
		testName: "Test #2",
		bits: []Bit{ // 0000 0101.
			Zero,
			One,
			Zero,
		},
		bytesExpected:     []byte{2},
		bitsCountExpected: 3,
	}
	testsData = append(testsData, testData)

	// Test #3. Eight Bits => One full Byte.
	testData = TestDataForConvertBitsToBytes{
		testName: "Test #3",
		bits: []Bit{ // 1000 0010.
			Zero,
			One,
			Zero,
			Zero,
			Zero,
			Zero,
			Zero,
			One,
		},
		bytesExpected:     []byte{128 + 2},
		bitsCountExpected: 8,
	}
	testsData = append(testsData, testData)

	// Test #4. Nine Bits => One full Byte, one partial Byte.
	testData = TestDataForConvertBitsToBytes{
		testName: "Test #4",
		bits: []Bit{ // 1000 0000; 0000 0001.
			// Byte #1.
			Zero,
			Zero,
			Zero,
			Zero,
			Zero,
			Zero,
			Zero,
			One,
			// Byte #2.
			One,
		},
		bytesExpected: []byte{
			128,
			1,
		},
		bitsCountExpected: 9,
	}
	testsData = append(testsData, testData)

	// Test #5. 16 Bits => Two full Bytes.
	testData = TestDataForConvertBitsToBytes{
		testName: "Test #5",
		bits: []Bit{ // 1000 0000; 1111 1111.
			// Byte #1.
			Zero,
			Zero,
			Zero,
			Zero,
			Zero,
			Zero,
			Zero,
			One,
			// Byte #2.
			One,
			One,
			One,
			One,
			One,
			One,
			One,
			One,
		},
		bytesExpected: []byte{
			128,
			255,
		},
		bitsCountExpected: 16,
	}
	testsData = append(testsData, testData)

	// Test #6. 17 Bits => Two full Bytes, one partial Byte.
	testData = TestDataForConvertBitsToBytes{
		testName: "Test #6",
		bits: []Bit{ // 1000 0000; 1111 1111; 0000 0001.
			// Byte #1.
			Zero,
			Zero,
			Zero,
			Zero,
			Zero,
			Zero,
			Zero,
			One,
			// Byte #2.
			One,
			One,
			One,
			One,
			One,
			One,
			One,
			One,
			// Byte #3.
			One,
		},
		bytesExpected: []byte{
			128,
			255,
			1,
		},
		bitsCountExpected: 17,
	}
	testsData = append(testsData, testData)

	// Run the Tests.
	for _, testData = range testsData {
		testRunner(t, &testData)
	}
}
