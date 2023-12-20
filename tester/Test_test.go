package tester

import (
	"testing"
)

func Test_New(t *testing.T) {
	var aTestingT *testing.T
	var result *Test

	aTestingT = new(testing.T)
	result = New(aTestingT)
	if result.t != aTestingT {
		t.FailNow()
	}
}
