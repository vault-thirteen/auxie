package mcci

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_CloneArrayOfNumbers(t *testing.T) {
	aTest := tester.New(t)

	a := []uint8{1, 2, 3}
	b := []uint16{}
	b = CloneArrayOfNumbers[uint8, uint16](a)
	aTest.MustBeEqual(len(b), len(a))
	aTest.MustBeEqual(b[0], uint16(1))
	aTest.MustBeEqual(b[1], uint16(2))
	aTest.MustBeEqual(b[2], uint16(3))
}
