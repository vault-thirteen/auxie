package csv

import (
	"bytes"
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewWriter(t *testing.T) {
	aTest := tester.New(t)

	var bb bytes.Buffer
	w := NewWriter(&bb)
	wExpected := &Writer{
		iow:          &bb,
		columnsCount: -1,
		newLineBA:    []byte(NewLineStr),
	}
	aTest.MustBeEqual(w, wExpected)
}

func Test_Writer_WriteRow(t *testing.T) {
	aTest := tester.New(t)
	var err error
	var bb bytes.Buffer
	var w *Writer

	// Test #1. Bad input: no rows.
	{
		w = NewWriter(&bb)
		err = w.WriteRow([]any{})
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(bb.String(), ``)
		aTest.MustBeEqual(w.columnsCount, -1)
	}

	// Test #2. First input sets the number of columns.
	{
		err = w.WriteRow([]any{123, "abc", 4.56})
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(bb.String(), `123,"abc",4.56`+NewLineStr)
		aTest.MustBeEqual(w.columnsCount, 3)
	}

	// Test #3. Second input has wrong columns count.
	{
		err = w.WriteRow([]any{12345, "abcd"})
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(w.columnsCount, 3)
	}

	// Test #4. Third input has data of following types: uint, float32, []byte.
	{
		err = w.WriteRow([]any{uint(6), float32(0.7), []byte{123, 124, 125}})
		aTest.MustBeNoError(err)
	}

	// Test #5. Input has data of wrong type.
	{
		err = w.WriteRow([]any{"OK", "OK", time.Time{}})
		aTest.MustBeAnError(err)
	}

	// Final check.
	{
		textExpected := `123,"abc",4.56` + NewLineStr +
			`6,0.699999988079071,"7B7C7D"` + NewLineStr
		aTest.MustBeEqual(bb.String(), textExpected)
	}

	// Go language still has no built-in decimal type ! Version 1.0 of the
	// language was released on 28-th of March, year 2012 (Source:
	// https://go.dev/blog/go1). More than a dozen years have passed, but the
	// language is still very poor. What a shame !!!
}

func Test_escapeString(t *testing.T) {
	aTest := tester.New(t)

	aTest.MustBeEqual(escapeString(""), "")
	aTest.MustBeEqual(escapeString("abc"), "abc")
	aTest.MustBeEqual(escapeString(`"`), `""`)
	aTest.MustBeEqual(escapeString(`"a"`), `""a""`)
	aTest.MustBeEqual(escapeString(`""`), `""""`)
}

func Test_addEdgeQuotes(t *testing.T) {
	aTest := tester.New(t)

	aTest.MustBeEqual(addEdgeQuotes(``), `""`)
	aTest.MustBeEqual(addEdgeQuotes(`abc`), `"abc"`)
	aTest.MustBeEqual(addEdgeQuotes(`"`), `"""`)
	aTest.MustBeEqual(addEdgeQuotes(`"x`), `""x"`)
	aTest.MustBeEqual(addEdgeQuotes(`""`), `""""`)
}
