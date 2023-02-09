// convertByteToBits_test.go.

package bit

import (
	"reflect"
	"testing"
)

type TestDataForConvertByteToBits struct {

	// Name of the Test.
	testName string

	// Input.
	byte byte

	// Output.
	bitsExpected []Bit
	bitsGot      []Bit
}

type TestRunnerForConvertByteToBits = func(
	t *testing.T,
	testData *TestDataForConvertByteToBits,
)

func Test_ConvertByteToBits(t *testing.T) {

	var testData TestDataForConvertByteToBits
	var testRunner TestRunnerForConvertByteToBits
	var testsData []TestDataForConvertByteToBits

	// Test Runner.
	testRunner = func(
		t *testing.T,
		testData *TestDataForConvertByteToBits,
	) {
		// Fool Check.
		if testData == nil {
			t.Error(ErrNullPointer)
			t.FailNow()
		}

		// Run the tested Function.
		testData.bitsGot = ConvertByteToBits(
			testData.byte,
		)

		// Check the Results.
		if !reflect.DeepEqual(testData.bitsGot, testData.bitsExpected) {
			t.Error(testData.testName, TestNameErrDelimiter, ErrBitsMismatch)
			t.FailNow()
		}
	}

	// Test #1: Byte 0000 0000.
	testData = TestDataForConvertByteToBits{
		testName: "Test #1",
		byte:     0,
		bitsExpected: []Bit{
			Zero, // 1.
			Zero, // 2.
			Zero, // 3.
			Zero, // 4.
			Zero, // 5.
			Zero, // 6.
			Zero, // 7.
			Zero, // 8.
		},
	}
	testsData = append(testsData, testData)

	// Test #2: Byte 0000 0001.
	testData = TestDataForConvertByteToBits{
		testName: "Test #2",
		byte:     1,
		bitsExpected: []Bit{
			One,  // 1.
			Zero, // 2.
			Zero, // 3.
			Zero, // 4.
			Zero, // 5.
			Zero, // 6.
			Zero, // 7.
			Zero, // 8.
		},
	}
	testsData = append(testsData, testData)

	// Test #3: Byte 0000 0010.
	testData = TestDataForConvertByteToBits{
		testName: "Test #3",
		byte:     2,
		bitsExpected: []Bit{
			Zero, // 1.
			One,  // 2.
			Zero, // 3.
			Zero, // 4.
			Zero, // 5.
			Zero, // 6.
			Zero, // 7.
			Zero, // 8.
		},
	}
	testsData = append(testsData, testData)

	// Test #4: Byte 0000 0100.
	testData = TestDataForConvertByteToBits{
		testName: "Test #4",
		byte:     4,
		bitsExpected: []Bit{
			Zero, // 1.
			Zero, // 2.
			One,  // 3.
			Zero, // 4.
			Zero, // 5.
			Zero, // 6.
			Zero, // 7.
			Zero, // 8.
		},
	}
	testsData = append(testsData, testData)

	// Test #5: Byte 0000 1000.
	testData = TestDataForConvertByteToBits{
		testName: "Test #5",
		byte:     8,
		bitsExpected: []Bit{
			Zero, // 1.
			Zero, // 2.
			Zero, // 3.
			One,  // 4.
			Zero, // 5.
			Zero, // 6.
			Zero, // 7.
			Zero, // 8.
		},
	}
	testsData = append(testsData, testData)

	// Test #6: Byte 0001 0000.
	testData = TestDataForConvertByteToBits{
		testName: "Test #6",
		byte:     16,
		bitsExpected: []Bit{
			Zero, // 1.
			Zero, // 2.
			Zero, // 3.
			Zero, // 4.
			One,  // 5.
			Zero, // 6.
			Zero, // 7.
			Zero, // 8.
		},
	}
	testsData = append(testsData, testData)

	// Test #7: Byte 0010 0000.
	testData = TestDataForConvertByteToBits{
		testName: "Test #7",
		byte:     32,
		bitsExpected: []Bit{
			Zero, // 1.
			Zero, // 2.
			Zero, // 3.
			Zero, // 4.
			Zero, // 5.
			One,  // 6.
			Zero, // 7.
			Zero, // 8.
		},
	}
	testsData = append(testsData, testData)

	// Test #8: Byte 0100 0000.
	testData = TestDataForConvertByteToBits{
		testName: "Test #8",
		byte:     64,
		bitsExpected: []Bit{
			Zero, // 1.
			Zero, // 2.
			Zero, // 3.
			Zero, // 4.
			Zero, // 5.
			Zero, // 6.
			One,  // 7.
			Zero, // 8.
		},
	}
	testsData = append(testsData, testData)

	// Test #9: Byte 1000 0000.
	testData = TestDataForConvertByteToBits{
		testName: "Test #9",
		byte:     128,
		bitsExpected: []Bit{
			Zero, // 1.
			Zero, // 2.
			Zero, // 3.
			Zero, // 4.
			Zero, // 5.
			Zero, // 6.
			Zero, // 7.
			One,  // 8.
		},
	}
	testsData = append(testsData, testData)

	// Test #10: Byte 1111 1111.
	testData = TestDataForConvertByteToBits{
		testName: "Test #9",
		byte:     255,
		bitsExpected: []Bit{
			One, // 1.
			One, // 2.
			One, // 3.
			One, // 4.
			One, // 5.
			One, // 6.
			One, // 7.
			One, // 8.
		},
	}
	testsData = append(testsData, testData)

	// Run the Tests.
	for _, testData = range testsData {
		testRunner(t, &testData)
	}
}
