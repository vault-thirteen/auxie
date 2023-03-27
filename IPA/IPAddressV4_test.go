package ipa

import (
	"math"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_NewFromBytes(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(NewFromBytes(0, 0, 0, 0), IPAddressV4(0))
	aTest.MustBeEqual(NewFromBytes(1, 1, 1, 1), IPAddressV4(1+256+65_536+16_777_216))
	aTest.MustBeEqual(NewFromBytes(0xFF, 0xFF, 0xFF, 0xFF), IPAddressV4(math.MaxUint32))
}

func Test_NewFromString(t *testing.T) {
	aTest := tester.New(t)
	var err error
	var ipaddr IPAddressV4

	// Test #1. Bad string.
	ipaddr, err = NewFromString("1.2.3")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(ipaddr, IPAddressV4(0))

	// Test #2. Normal string.
	ipaddr, err = NewFromString("0.0.0.0")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(ipaddr, IPAddressV4(0))

	// Test #3. Normal string.
	ipaddr, err = NewFromString("1.1.1.1")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(ipaddr, IPAddressV4(1+256+65_536+16_777_216))

	// Test #4. Normal string.
	ipaddr, err = NewFromString("255.255.255.255")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(ipaddr, IPAddressV4(math.MaxUint32))
}

func Test_NewFromUintString(t *testing.T) {
	aTest := tester.New(t)
	var err error
	var ipaddr IPAddressV4

	// Test #1. Bad string.
	ipaddr, err = NewFromUintString("-999") // Not an uint.
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(ipaddr, IPAddressV4(0))

	// Test #2. Normal string.
	ipaddr, err = NewFromUintString("0")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(ipaddr, IPAddressV4(0))

	// Test #3. Normal string.
	ipaddr, err = NewFromUintString("16843009")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(ipaddr, IPAddressV4(1+256+65_536+16_777_216))

	// Test #4. Normal string.
	ipaddr, err = NewFromUintString("4294967295")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(ipaddr, IPAddressV4(math.MaxUint32))

	// Test #5. Overflow.
	ipaddr, err = NewFromUintString("4294967296")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(ipaddr, IPAddressV4(0))
}
