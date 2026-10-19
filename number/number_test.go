package number

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_ParseUint(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result uint

	// Test #1.
	result, err = ParseUint("123")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, uint(123))

	// Test #2. Overflow.
	result, err = ParseUint("18446744073709551616")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, uint(0))

	// Test #3. Not a number.
	result, err = ParseUint("not a number")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, uint(0))
}

func Test_ParseUint64(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result uint64

	// Test #1.
	result, err = ParseUint64("123")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, uint64(123))

	// Test #2. Overflow.
	result, err = ParseUint64("18446744073709551616")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, uint64(0xffffffffffffffff))

	// Test #3. Not a number.
	result, err = ParseUint64("not a number")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, uint64(0))
}

func Test_ParseUint32(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result uint32

	// Test #1.
	result, err = ParseUint32("123")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, uint32(123))

	// Test #2. Overflow.
	result, err = ParseUint32("4294967296")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, uint32(0))

	// Test #3. Not a number.
	result, err = ParseUint32("not a number")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, uint32(0))
}

func Test_ParseUint16(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result uint16

	// Test #1.
	result, err = ParseUint16("123")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, uint16(123))

	// Test #2. Overflow.
	result, err = ParseUint16("65536")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, uint16(0))

	// Test #3. Not a number.
	result, err = ParseUint16("not a number")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, uint16(0))
}

func Test_ParseUint8(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result uint8

	// Test #1.
	result, err = ParseUint8("123")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, uint8(123))

	// Test #2. Overflow.
	result, err = ParseUint8("256")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, uint8(0))

	// Test #3. Not a number.
	result, err = ParseUint8("not a number")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, uint8(0))
}

func Test_ParseInt(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result int

	// Test #1.
	result, err = ParseInt("123")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, int(123))

	// Test #2. Overflow.
	result, err = ParseInt("9223372036854775808")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, int(0))

	// Test #3. Not a number.
	result, err = ParseInt("not a number")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, int(0))
}

func Test_ParseInt64(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result int64

	// Test #1.
	result, err = ParseInt64("123")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, int64(123))

	// Test #2. Overflow.
	result, err = ParseInt64("9223372036854775808")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, int64(9223372036854775807))

	// Test #3. Not a number.
	result, err = ParseInt64("not a number")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, int64(0))
}

func Test_ParseInt32(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result int32

	// Test #1.
	result, err = ParseInt32("123")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, int32(123))

	// Test #2. Overflow.
	result, err = ParseInt32("2147483648")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, int32(0))

	// Test #3. Not a number.
	result, err = ParseInt32("not a number")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, int32(0))
}

func Test_ParseInt16(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result int16

	// Test #1.
	result, err = ParseInt16("123")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, int16(123))

	// Test #2. Overflow.
	result, err = ParseInt16("32768")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, int16(0))

	// Test #3. Not a number.
	result, err = ParseInt16("not a number")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, int16(0))
}

func Test_ParseInt8(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result int8

	// Test #1.
	result, err = ParseInt8("123")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, int8(123))

	// Test #2. Overflow.
	result, err = ParseInt8("128")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, int8(0))

	// Test #3. Not a number.
	result, err = ParseInt8("not a number")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, int8(0))
}

func Test_ParseFloat64(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result float64

	// Test #1.
	result, err = ParseFloat64("1.0")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, float64(1.0))

	// Test #2. Overflow.
	result, err = ParseFloat64("1.8e+308")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, float64(0))

	// Test #3. Not a number.
	result, err = ParseFloat64("not a number")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, float64(0))
}

func Test_ParseFloat32(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result float32

	// Test #1.
	result, err = ParseFloat32("1.0")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, float32(1.0))

	// Test #2. Overflow.
	result, err = ParseFloat32("3.5e+38")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, float32(0))

	// Test #3. Not a number.
	result, err = ParseFloat32("not a number")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, float32(0))
}
