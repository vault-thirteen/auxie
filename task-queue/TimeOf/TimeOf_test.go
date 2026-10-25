package to

import (
	"testing"

	tp "github.com/vault-thirteen/auxie/task-queue/TimeStamp"
	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewTimeOf(t *testing.T) {
	aTest := tester.New(t)

	x := NewTimeOf()
	aTest.MustBeEqual(x.creation > tp.ZeroTimeStamp, true)
	aTest.MustBeEqual(x.workStart == tp.ZeroTimeStamp, true)
	aTest.MustBeEqual(x.workEnd == tp.ZeroTimeStamp, true)
	aTest.MustBeEqual(x.completion == tp.ZeroTimeStamp, true)
}
func Test_Creation(t *testing.T) {
	aTest := tester.New(t)

	x := TimeOf{creation: 123}
	aTest.MustBeEqual(x.Creation(), tp.TimeStamp(123))
}
func Test_WorkStart(t *testing.T) {
	aTest := tester.New(t)

	x := TimeOf{workStart: 123}
	aTest.MustBeEqual(x.WorkStart(), tp.TimeStamp(123))
}
func Test_WorkEnd(t *testing.T) {
	aTest := tester.New(t)

	x := TimeOf{workEnd: 123}
	aTest.MustBeEqual(x.WorkEnd(), tp.TimeStamp(123))
}
func Test_Completion(t *testing.T) {
	aTest := tester.New(t)

	x := TimeOf{completion: 123}
	aTest.MustBeEqual(x.Completion(), tp.TimeStamp(123))
}
func Test_SetWorkStart(t *testing.T) {
	aTest := tester.New(t)

	x := TimeOf{}
	x.SetWorkStart()
	aTest.MustBeEqual(x.workStart > tp.ZeroTimeStamp, true)
}
func Test_SetWorkEnd(t *testing.T) {
	aTest := tester.New(t)

	x := TimeOf{}
	x.SetWorkEnd()
	aTest.MustBeEqual(x.workEnd > tp.ZeroTimeStamp, true)
}
func Test_SetCompletion(t *testing.T) {
	aTest := tester.New(t)

	x := TimeOf{}
	x.SetCompletion()
	aTest.MustBeEqual(x.completion > tp.ZeroTimeStamp, true)
}
