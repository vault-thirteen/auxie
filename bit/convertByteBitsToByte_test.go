// convertByteBitsToByte_test.go.

package bit

import (
	"bytes"
	"testing"
)

type TestDataForConvertByteBitsToByte struct {

	// Name of the Test.
	testName string

	// Input.
	bits [BitsPerByte]Bit

	// Output.
	byteExpected byte
	byteGot      byte
}

type TestRunnerForconvertByteBitsToByte = func(
	t *testing.T,
	testData *TestDataForConvertByteBitsToByte,
)

func Test_ConvertByteBitsToByte(t *testing.T) {

	var testData TestDataForConvertByteBitsToByte
	var testRunner TestRunnerForconvertByteBitsToByte
	var testsData []TestDataForConvertByteBitsToByte

	// Test Runner.
	testRunner = func(
		t *testing.T,
		testData *TestDataForConvertByteBitsToByte,
	) {
		// Fool Check.
		if testData == nil {
			t.Error(ErrNullPointer)
			t.FailNow()
		}

		// Run the tested Function.
		testData.byteGot = ConvertByteBitsToByte(
			testData.bits,
		)

		// Check the Results.
		if !bytes.Equal(
			[]byte{testData.byteGot},
			[]byte{testData.byteExpected},
		) {
			t.Error(testData.testName, TestNameErrDelimiter, ErrBytesMismatch)
			t.FailNow()
		}
	}

	// Test #1: No Bits.
	testData = TestDataForConvertByteBitsToByte{
		testName: "Test #1",
		bits: [BitsPerByte]Bit{
			Zero, // 1.
			Zero, // 2.
			Zero, // 3.
			Zero, // 4.
			Zero, // 5.
			Zero, // 6.
			Zero, // 7.
			Zero, // 8.
		},
		byteExpected: 0,
	}
	testsData = append(testsData, testData)

	// Test #2. 0000 0001.
	testData = TestDataForConvertByteBitsToByte{
		testName: "Test #2",
		bits: [BitsPerByte]Bit{
			One,  // 1.
			Zero, // 2.
			Zero, // 3.
			Zero, // 4.
			Zero, // 5.
			Zero, // 6.
			Zero, // 7.
			Zero, // 8.
		},
		byteExpected: 1,
	}
	testsData = append(testsData, testData)

	// Test #3. 0000 0010.
	testData = TestDataForConvertByteBitsToByte{
		testName: "Test #3",
		bits: [BitsPerByte]Bit{
			Zero, // 1
			One,  // 2.
			Zero, // 3.
			Zero, // 4.
			Zero, // 5.
			Zero, // 6.
			Zero, // 7.
			Zero, // 8.
		},
		byteExpected: 2,
	}
	testsData = append(testsData, testData)

	// Test #4. 0000 0100.
	testData = TestDataForConvertByteBitsToByte{
		testName: "Test #4",
		bits: [BitsPerByte]Bit{
			Zero, // 1
			Zero, // 2.
			One,  // 3.
			Zero, // 4.
			Zero, // 5.
			Zero, // 6.
			Zero, // 7.
			Zero, // 8.
		},
		byteExpected: 4,
	}
	testsData = append(testsData, testData)

	// Test #5. 0000 1000.
	testData = TestDataForConvertByteBitsToByte{
		testName: "Test #5",
		bits: [BitsPerByte]Bit{
			Zero, // 1
			Zero, // 2.
			Zero, // 3.
			One,  // 4.
			Zero, // 5.
			Zero, // 6.
			Zero, // 7.
			Zero, // 8.
		},
		byteExpected: 8,
	}
	testsData = append(testsData, testData)

	// Test #6. 0001 0000.
	testData = TestDataForConvertByteBitsToByte{
		testName: "Test #6",
		bits: [BitsPerByte]Bit{
			Zero, // 1
			Zero, // 2.
			Zero, // 3.
			Zero, // 4.
			One,  // 5.
			Zero, // 6.
			Zero, // 7.
			Zero, // 8.
		},
		byteExpected: 16,
	}
	testsData = append(testsData, testData)

	// Test #7. 0010 0000.
	testData = TestDataForConvertByteBitsToByte{
		testName: "Test #7",
		bits: [BitsPerByte]Bit{
			Zero, // 1
			Zero, // 2.
			Zero, // 3.
			Zero, // 4.
			Zero, // 5.
			One,  // 6.
			Zero, // 7.
			Zero, // 8.
		},
		byteExpected: 32,
	}
	testsData = append(testsData, testData)

	// Test #8. 0100 0000.
	testData = TestDataForConvertByteBitsToByte{
		testName: "Test #8",
		bits: [BitsPerByte]Bit{
			Zero, // 1
			Zero, // 2.
			Zero, // 3.
			Zero, // 4.
			Zero, // 5.
			Zero, // 6.
			One,  // 7.
			Zero, // 8.
		},
		byteExpected: 64,
	}
	testsData = append(testsData, testData)

	// Test #9. 1000 0000.
	testData = TestDataForConvertByteBitsToByte{
		testName: "Test #9",
		bits: [BitsPerByte]Bit{
			Zero, // 1
			Zero, // 2.
			Zero, // 3.
			Zero, // 4.
			Zero, // 5.
			Zero, // 6.
			Zero, // 7.
			One,  // 8.
		},
		byteExpected: 128,
	}
	testsData = append(testsData, testData)

	// Test #10. 1111 1111.
	testData = TestDataForConvertByteBitsToByte{
		testName: "Test #10",
		bits: [BitsPerByte]Bit{
			One, // 1
			One, // 2.
			One, // 3.
			One, // 4.
			One, // 5.
			One, // 6.
			One, // 7.
			One, // 8.
		},
		byteExpected: 255,
	}
	testsData = append(testsData, testData)

	// Run the Tests.
	for _, testData = range testsData {
		testRunner(t, &testData)
	}
}
