package tqs

import (
	"reflect"
	"testing"

	"github.com/vault-thirteen/auxie/cio"
	tqsp "github.com/vault-thirteen/auxie/task-queue/TaskQueueSettingsPublic"
	ce "github.com/vault-thirteen/auxie/task-queue/common/error"
	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewTaskQueueSettings(t *testing.T) {
	aTest := tester.New(t)
	var source *tqsp.TaskQueueSettingsPublic
	var tqs *TaskQueueSettings
	var err error

	// Test #1. Null pointer.
	source = nil
	tqs, err = NewTaskQueueSettings(source)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), ce.Err_NullPointer)

	// Test #2. Invalid source.
	source = &tqsp.TaskQueueSettingsPublic{}
	tqs, err = NewTaskQueueSettings(source)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(tqs, (*TaskQueueSettings)(nil))

	// Test #3. Valid source.
	source = &tqsp.TaskQueueSettingsPublic{
		InputQueueSizeLimit: 1,
		ErrorChannelSize:    2,
		ErrorCatcherFn:      func(err error) { cio.PrintMessage("Hello world") },

		// Timings.
		CompletedTaskTtlSec:                      3,
		CleanerSleepTimeSec:                      4,
		ErrorCatcherSleepTimeSec:                 5,
		SleepTimeWaitingForZeroQueueSec:          6,
		SleepTimeWaitingForZeroTasksInWorkSec:    7,
		SleepTimeWaitingForZeroProcessedTasksSec: 8,
	}
	tqs, err = NewTaskQueueSettings(source)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(tqs != nil, true)
}
func Test_Validate(t *testing.T) {
	aTest := tester.New(t)
	var tqs *TaskQueueSettings
	var err error

	// Test #1. Invalid parameter: inputQueueSizeLimit.
	tqs = &TaskQueueSettings{
		inputQueueSizeLimit: 0,
	}
	err = tqs.Validate()
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), ce.Err_InputQueueSizeLimit)

	// Test #2. Invalid parameter: errorChannelSize.
	tqs = &TaskQueueSettings{
		inputQueueSizeLimit: 123,
		errorChannelSize:    0,
	}
	err = tqs.Validate()
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), ce.Err_ErrorChannelSize)

	// Test #3. Usage of all default timings.
	tqs = &TaskQueueSettings{
		inputQueueSizeLimit: 123,
		errorChannelSize:    456,
	}
	err = tqs.Validate()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(tqs.completedTaskTtlSec, Default_CompletedTaskTtlSec)
	aTest.MustBeEqual(tqs.cleanerSleepTimeSec, Default_CleanerSleepTimeSec)
	aTest.MustBeEqual(tqs.errorCatcherSleepTimeSec, Default_ErrorCatcherSleepTimeSec)
	aTest.MustBeEqual(tqs.sleepTimeWaitingForZeroQueueSec, Default_SleepTimeWaitingForZeroQueueSec)
	aTest.MustBeEqual(tqs.sleepTimeWaitingForZeroTasksInWorkSec, Default_SleepTimeWaitingForZeroTasksInWorkSec)
	aTest.MustBeEqual(tqs.sleepTimeWaitingForZeroProcessedTasksSec, Default_SleepTimeWaitingForZeroProcessedTasksSec)
}
func Test_InputQueueSizeLimit(t *testing.T) {
	aTest := tester.New(t)
	tqs := &TaskQueueSettings{}
	tqs.inputQueueSizeLimit = 1
	aTest.MustBeEqual(tqs.InputQueueSizeLimit(), 1)
}
func Test_ErrorChannelSize(t *testing.T) {
	aTest := tester.New(t)
	tqs := &TaskQueueSettings{}
	tqs.errorChannelSize = 1
	aTest.MustBeEqual(tqs.ErrorChannelSize(), 1)
}
func Test_ErrorCatcherFn(t *testing.T) {
	aTest := tester.New(t)
	tqs := &TaskQueueSettings{}
	fn := func(err error) { cio.PrintMessage("Hello world") }
	tqs.errorCatcherFn = fn
	fn(nil)
	tqs.errorCatcherFn(nil)

	// While we can not compare function pointers, we use a trick.
	// Further notes are below.
	fn1 := tqs.ErrorCatcherFn()
	fn2 := fn
	fn1Address := reflect.ValueOf(fn1).Pointer()
	fn2Address := reflect.ValueOf(fn2).Pointer()
	aTest.MustBeEqual(fn1Address, fn2Address)

	// N.B. This is a known bug of the Go language. Comparison of function
	// pointers with 'reflect.DeepEqual' always returns false. Even if they
	// have identical internals, they are not equal for Go language. What a
	// lame language !
	//
	// Quote: "Func values are deeply equal if both are nil; otherwise they
	// are not deeply equal.", Source: https://pkg.go.dev/reflect.
}
func Test_CompletedTaskTtlSec(t *testing.T) {
	aTest := tester.New(t)
	tqs := &TaskQueueSettings{}
	tqs.completedTaskTtlSec = 1
	aTest.MustBeEqual(tqs.CompletedTaskTtlSec(), 1)
}
func Test_CleanerSleepTimeSec(t *testing.T) {
	aTest := tester.New(t)
	tqs := &TaskQueueSettings{}
	tqs.cleanerSleepTimeSec = 1
	aTest.MustBeEqual(tqs.CleanerSleepTimeSec(), 1)
}
func Test_ErrorCatcherSleepTimeSec(t *testing.T) {
	aTest := tester.New(t)
	tqs := &TaskQueueSettings{}
	tqs.errorCatcherSleepTimeSec = 1
	aTest.MustBeEqual(tqs.ErrorCatcherSleepTimeSec(), 1)
}
func Test_SleepTimeWaitingForZeroQueueSec(t *testing.T) {
	aTest := tester.New(t)
	tqs := &TaskQueueSettings{}
	tqs.sleepTimeWaitingForZeroQueueSec = 1
	aTest.MustBeEqual(tqs.SleepTimeWaitingForZeroQueueSec(), 1)
}
func Test_SleepTimeWaitingForZeroTasksInWorkSec(t *testing.T) {
	aTest := tester.New(t)
	tqs := &TaskQueueSettings{}
	tqs.sleepTimeWaitingForZeroTasksInWorkSec = 1
	aTest.MustBeEqual(tqs.SleepTimeWaitingForZeroTasksInWorkSec(), 1)
}
func Test_SleepTimeWaitingForZeroProcessedTasksSec(t *testing.T) {
	aTest := tester.New(t)
	tqs := &TaskQueueSettings{}
	tqs.sleepTimeWaitingForZeroProcessedTasksSec = 1
	aTest.MustBeEqual(tqs.SleepTimeWaitingForZeroProcessedTasksSec(), 1)
}
