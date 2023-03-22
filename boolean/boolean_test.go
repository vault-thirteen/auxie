package boolean

import (
	"fmt"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_FromString(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result bool

	type Test struct {
		input                 string
		isErrorExpected       bool
		outputErrorExpected   string
		outputBooleanExpected bool
	}

	var tests = []Test{
		{
			input:                 "",
			isErrorExpected:       true,
			outputErrorExpected:   "bad boolean value: ''",
			outputBooleanExpected: false,
		},
		{
			input:                 "Unreal Tournament 2004",
			isErrorExpected:       true,
			outputErrorExpected:   "bad boolean value: 'Unreal Tournament 2004'",
			outputBooleanExpected: false,
		},
		// 1.
		{
			input:                 "true",
			isErrorExpected:       false,
			outputBooleanExpected: true,
		},
		{
			input:                 "false",
			isErrorExpected:       false,
			outputBooleanExpected: false,
		},
		{
			input:                 "yes",
			isErrorExpected:       false,
			outputBooleanExpected: true,
		},
		{
			input:                 "no",
			isErrorExpected:       false,
			outputBooleanExpected: false,
		},
		{
			input:                 "1",
			isErrorExpected:       false,
			outputBooleanExpected: true,
		},
		{
			input:                 "0",
			isErrorExpected:       false,
			outputBooleanExpected: false,
		},
		// 2.
		{
			input:                 " TruE   ",
			isErrorExpected:       false,
			outputBooleanExpected: true,
		},
		{
			input:                 "  fAlSe ",
			isErrorExpected:       false,
			outputBooleanExpected: false,
		},
		{
			input:                 " YES ",
			isErrorExpected:       false,
			outputBooleanExpected: true,
		},
		{
			input:                 "  NO  ",
			isErrorExpected:       false,
			outputBooleanExpected: false,
		},
		{
			input:                 "  1  ",
			isErrorExpected:       false,
			outputBooleanExpected: true,
		},
		{
			input:                 "  0  ",
			isErrorExpected:       false,
			outputBooleanExpected: false,
		},
	}

	for i, tst := range tests {
		fmt.Printf("[%d]", i+1)
		result, err = FromString(tst.input)
		if tst.isErrorExpected {
			aTest.MustBeAnError(err)
			aTest.MustBeEqual(err.Error(), tst.outputErrorExpected)
		} else {
			aTest.MustBeNoError(err)
		}
		aTest.MustBeEqual(result, tst.outputBooleanExpected)
	}
	fmt.Println()
}
