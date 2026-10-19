package ssc

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewSSC(t *testing.T) {
	aTest := tester.New(t)
	var s *SSC[string, bool]

	// Test #1.
	s = NewSSC[string, bool](0)
	aTest.MustBeEqual(s, (*SSC[string, bool])(nil))

	// Test #2.
	s = NewSSC[string, bool](1)
	aTest.MustBeDifferent(s, (*SSC[string, bool])(nil))
}

func Test_GetSize(t *testing.T) {
	aTest := tester.New(t)
	var s *SSC[string, bool]

	// Test.
	s = NewSSC[string, bool](10)
	s.SetValue("a", true)
	aTest.MustBeEqual(s.GetSize(), 1)
	s.SetValue("b", true)
	aTest.MustBeEqual(s.GetSize(), 2)
	s.SetValue("b", false)
	aTest.MustBeEqual(s.GetSize(), 2)
	s.SetValue("a", false)
	aTest.MustBeEqual(s.GetSize(), 2)
}

func Test_GetValue(t *testing.T) {
	aTest := tester.New(t)
	var s *SSC[string, bool]
	var value bool
	var recordExists bool

	// Test.
	s = NewSSC[string, bool](10)
	s.SetValue("a", true)

	value, recordExists = s.GetValue("a")
	aTest.MustBeEqual(value, true)
	aTest.MustBeEqual(recordExists, true)

	value, recordExists = s.GetValue("b")
	aTest.MustBeEqual(recordExists, false)
}

func Test_SetValue(t *testing.T) {
	aTest := tester.New(t)
	var s *SSC[string, int]
	var value int
	var recordExists bool

	// Test.
	s = NewSSC[string, int](3)

	s.SetValue("a", 111)
	value, recordExists = s.GetValue("a")
	aTest.MustBeEqual(value, 111)
	aTest.MustBeEqual(recordExists, true)
	aTest.MustBeEqual(s.GetSize(), 1)

	s.SetValue("b", 222)
	value, recordExists = s.GetValue("b")
	aTest.MustBeEqual(value, 222)
	aTest.MustBeEqual(recordExists, true)
	aTest.MustBeEqual(s.GetSize(), 2)

	s.SetValue("a", 11)
	value, recordExists = s.GetValue("a")
	aTest.MustBeEqual(value, 11)
	aTest.MustBeEqual(recordExists, true)
	aTest.MustBeEqual(s.GetSize(), 2)

	s.SetValue("c", 333)
	value, recordExists = s.GetValue("c")
	aTest.MustBeEqual(value, 333)
	aTest.MustBeEqual(recordExists, true)
	aTest.MustBeEqual(s.GetSize(), 3)

	s.SetValue("d", 444)
	value, recordExists = s.GetValue("d")
	aTest.MustBeEqual(value, 444)
	aTest.MustBeEqual(recordExists, true)
	aTest.MustBeEqual(s.GetSize(), 3)

	// Corner case for the 'removeRandomRecord' function.
	// This case is not possible in real life and is used only for increased
	// test coverage.
	s = &SSC[string, int]{
		values:  map[string]int{},
		maxSize: 0,
	}
	s.SetValue("a", 111)

	// Another corner case to test extreme usage.
	s = NewSSC[string, int](1)
	s.SetValue("a", 111)
	s.SetValue("b", 222)

	value, recordExists = s.GetValue("b")
	aTest.MustBeEqual(value, 222)
	aTest.MustBeEqual(recordExists, true)

	value, recordExists = s.GetValue("a")
	aTest.MustBeEqual(recordExists, false)
}

func Test_Reset(t *testing.T) {
	aTest := tester.New(t)
	var s *SSC[string, int]

	// Test.
	s = NewSSC[string, int](2)

	s.SetValue("a", 123)
	aTest.MustBeEqual(s.GetSize(), 1)

	s.Reset()
	aTest.MustBeEqual(s.GetSize(), 0)
}

func Test_Delete(t *testing.T) {
	aTest := tester.New(t)
	var s *SSC[string, int]

	// Test.
	s = NewSSC[string, int](2)

	s.SetValue("a", 123)
	s.SetValue("b", 456)
	aTest.MustBeEqual(s.GetSize(), 2)

	s.Delete("a")
	aTest.MustBeEqual(s.GetSize(), 1)
	s.Delete("?")
	aTest.MustBeEqual(s.GetSize(), 1)
}
