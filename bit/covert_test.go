package bit

import (
	"fmt"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_ConvertBitsToBytes(t *testing.T) {
	aTest := tester.New(t)

	type TestData struct {
		bits              []Bit
		expectedBytes     []byte
		expectedBitsCount int
	}

	tests := []TestData{
		{
			bits:              []Bit{},
			expectedBitsCount: 0,
			expectedBytes:     []byte{},
		},
		{
			bits:              []Bit{Zero},
			expectedBitsCount: 1,
			expectedBytes:     []byte{0},
		},
		{
			bits:              []Bit{One},
			expectedBitsCount: 1,
			expectedBytes:     []byte{1},
		},
		{
			bits:              []Bit{Zero, One},
			expectedBitsCount: 2,
			expectedBytes:     []byte{2},
		},
		{
			bits:              []Bit{Zero, Zero, One},
			expectedBitsCount: 3,
			expectedBytes:     []byte{4},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, One},
			expectedBitsCount: 4,
			expectedBytes:     []byte{8},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, Zero, One},
			expectedBitsCount: 5,
			expectedBytes:     []byte{16},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, Zero, Zero, One},
			expectedBitsCount: 6,
			expectedBytes:     []byte{32},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, Zero, Zero, Zero, One},
			expectedBitsCount: 7,
			expectedBytes:     []byte{64},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, One},
			expectedBitsCount: 8,
			expectedBytes:     []byte{128},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One},
			expectedBitsCount: 9,
			expectedBytes:     []byte{0, 1},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One},
			expectedBitsCount: 10,
			expectedBytes:     []byte{0, 2},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One},
			expectedBitsCount: 11,
			expectedBytes:     []byte{0, 4},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One},
			expectedBitsCount: 12,
			expectedBytes:     []byte{0, 8},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One},
			expectedBitsCount: 13,
			expectedBytes:     []byte{0, 16},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One},
			expectedBitsCount: 14,
			expectedBytes:     []byte{0, 32},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One},
			expectedBitsCount: 15,
			expectedBytes:     []byte{0, 64},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One},
			expectedBitsCount: 16,
			expectedBytes:     []byte{0, 128},
		},
		{
			bits:              []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One},
			expectedBitsCount: 17,
			expectedBytes:     []byte{0, 0, 1},
		},
		{
			bits:              []Bit{One, One},
			expectedBitsCount: 2,
			expectedBytes:     []byte{3},
		},
		{
			bits:              []Bit{One, One, One},
			expectedBitsCount: 3,
			expectedBytes:     []byte{7},
		},
		{
			bits:              []Bit{One, One, One, One},
			expectedBitsCount: 4,
			expectedBytes:     []byte{15},
		},
		{
			bits:              []Bit{One, One, One, One, One},
			expectedBitsCount: 5,
			expectedBytes:     []byte{31},
		},
		{
			bits:              []Bit{One, One, One, One, One, One},
			expectedBitsCount: 6,
			expectedBytes:     []byte{63},
		},
		{
			bits:              []Bit{One, One, One, One, One, One, One},
			expectedBitsCount: 7,
			expectedBytes:     []byte{127},
		},
		{
			bits:              []Bit{One, One, One, One, One, One, One, One},
			expectedBitsCount: 8,
			expectedBytes:     []byte{255},
		},
		{
			bits:              []Bit{One, One, One, One, One, One, One, One, One},
			expectedBitsCount: 9,
			expectedBytes:     []byte{255, 1},
		},
		{
			bits:              []Bit{One, One, One, One, One, One, One, One, One, One},
			expectedBitsCount: 10,
			expectedBytes:     []byte{255, 3},
		},
		{
			bits:              []Bit{One, One, One, One, One, One, One, One, One, One, One},
			expectedBitsCount: 11,
			expectedBytes:     []byte{255, 7},
		},
		{
			bits:              []Bit{One, One, One, One, One, One, One, One, One, One, One, One},
			expectedBitsCount: 12,
			expectedBytes:     []byte{255, 15},
		},
		{
			bits:              []Bit{One, One, One, One, One, One, One, One, One, One, One, One, One},
			expectedBitsCount: 13,
			expectedBytes:     []byte{255, 31},
		},
		{
			bits:              []Bit{One, One, One, One, One, One, One, One, One, One, One, One, One, One},
			expectedBitsCount: 14,
			expectedBytes:     []byte{255, 63},
		},
		{
			bits:              []Bit{One, One, One, One, One, One, One, One, One, One, One, One, One, One, One},
			expectedBitsCount: 15,
			expectedBytes:     []byte{255, 127},
		},
		{
			bits:              []Bit{One, One, One, One, One, One, One, One, One, One, One, One, One, One, One, One},
			expectedBitsCount: 16,
			expectedBytes:     []byte{255, 255},
		},
		{
			bits:              []Bit{One, One, One, One, One, One, One, One, One, One, One, One, One, One, One, One, One},
			expectedBitsCount: 17,
			expectedBytes:     []byte{255, 255, 1},
		},
	}

	for i, test := range tests {
		fmt.Print("[", i+1, "] ")
		result, bitsCount := ConvertBitsToBytes(test.bits)
		aTest.MustBeEqual(result, test.expectedBytes)
		aTest.MustBeEqual(bitsCount, test.expectedBitsCount)
	}
	fmt.Println()
}

func Test_ConvertBytesToBits(t *testing.T) {
	aTest := tester.New(t)

	type TestData struct {
		bytes        []byte
		expectedBits []Bit
	}

	tests := []TestData{
		{
			bytes:        []byte{},
			expectedBits: []Bit{},
		},
		{
			bytes:        []byte{0},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{1},
			expectedBits: []Bit{One, Zero, Zero, Zero, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{2},
			expectedBits: []Bit{Zero, One, Zero, Zero, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{4},
			expectedBits: []Bit{Zero, Zero, One, Zero, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{8},
			expectedBits: []Bit{Zero, Zero, Zero, One, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{16},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, One, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{32},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, Zero, One, Zero, Zero},
		},
		{
			bytes:        []byte{64},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, Zero, Zero, One, Zero},
		},
		{
			bytes:        []byte{128},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, One},
		},
		{
			bytes:        []byte{0, 1},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One, Zero, Zero, Zero, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{0, 2},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One, Zero, Zero, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{0, 4},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One, Zero, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{0, 8},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{0, 16},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{0, 32},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One, Zero, Zero},
		},
		{
			bytes:        []byte{0, 64},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One, Zero},
		},
		{
			bytes:        []byte{0, 128},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One},
		},
		{
			bytes:        []byte{0, 0, 1},
			expectedBits: []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero, One, Zero, Zero, Zero, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{3},
			expectedBits: []Bit{One, One, Zero, Zero, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{7},
			expectedBits: []Bit{One, One, One, Zero, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{15},
			expectedBits: []Bit{One, One, One, One, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{31},
			expectedBits: []Bit{One, One, One, One, One, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{63},
			expectedBits: []Bit{One, One, One, One, One, One, Zero, Zero},
		},
		{
			bytes:        []byte{127},
			expectedBits: []Bit{One, One, One, One, One, One, One, Zero},
		},
		{
			bytes:        []byte{255},
			expectedBits: []Bit{One, One, One, One, One, One, One, One},
		},
		{
			bytes:        []byte{255, 1},
			expectedBits: []Bit{One, One, One, One, One, One, One, One, One, Zero, Zero, Zero, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{255, 3},
			expectedBits: []Bit{One, One, One, One, One, One, One, One, One, One, Zero, Zero, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{255, 7},
			expectedBits: []Bit{One, One, One, One, One, One, One, One, One, One, One, Zero, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{255, 15},
			expectedBits: []Bit{One, One, One, One, One, One, One, One, One, One, One, One, Zero, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{255, 31},
			expectedBits: []Bit{One, One, One, One, One, One, One, One, One, One, One, One, One, Zero, Zero, Zero},
		},
		{
			bytes:        []byte{255, 63},
			expectedBits: []Bit{One, One, One, One, One, One, One, One, One, One, One, One, One, One, Zero, Zero},
		},
		{
			bytes:        []byte{255, 127},
			expectedBits: []Bit{One, One, One, One, One, One, One, One, One, One, One, One, One, One, One, Zero},
		},
		{
			bytes:        []byte{255, 255},
			expectedBits: []Bit{One, One, One, One, One, One, One, One, One, One, One, One, One, One, One, One},
		},
		{
			bytes:        []byte{255, 255, 1},
			expectedBits: []Bit{One, One, One, One, One, One, One, One, One, One, One, One, One, One, One, One, One, Zero, Zero, Zero, Zero, Zero, Zero, Zero},
		},
	}

	for i, test := range tests {
		fmt.Print("[", i+1, "] ")
		result := ConvertBytesToBits(test.bytes)
		aTest.MustBeEqual(result, test.expectedBits)
	}
	fmt.Println()
}
