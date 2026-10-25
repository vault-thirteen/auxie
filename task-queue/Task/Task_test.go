package t

import (
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/cio"
	ti "github.com/vault-thirteen/auxie/task-queue/TaskId"
	ts "github.com/vault-thirteen/auxie/task-queue/TaskStatus"
	to "github.com/vault-thirteen/auxie/task-queue/TimeOf"
	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewTask(t *testing.T) {
	aTest := tester.New(t)

	var req = "Feed the cat"
	var task *Task[string, string]
	nowTS := time.Now().UTC().Unix()
	task = NewTask[string, string]("T1", &req)
	aTest.MustBeEqual(task.id, "T1")
	aTest.MustBeEqual(task.request, &req)
	aTest.MustBeEqual(task.response, (*string)(nil))
	aTest.MustBeEqual(task.status, ts.TaskStatus_Created)
	aTest.MustBeEqual(task.timeOf.Creation(), nowTS)
}

func Test_ChangeStatus(t *testing.T) {
	aTest := tester.New(t)
	var err error

	var req = "Feed the cat"
	var task *Task[string, string]
	task = NewTask[string, string]("T1", &req)

	err = task.ChangeStatus(ts.TaskStatus_Queued)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(task.status, ts.TaskStatus_Queued)
}
func Test_Destroy(t *testing.T) {
	aTest := tester.New(t)
	var err error
	var req = "Feed the cat"
	var task *Task[string, string]

	// Test #1. Status is not completed.
	task = NewTask[string, string]("T1", &req)
	err = task.Destroy()
	aTest.MustBeAnError(err)

	// Test #2. Status is completed.
	task = NewTask[string, string]("T1", &req)
	task.status = ts.TaskStatus_Completed
	err = task.Destroy()
	aTest.MustBeNoError(err)

	aTest.MustBeEqual(task.id, ti.ZeroTaskId)
	aTest.MustBeEqual(task.request, (*string)(nil))
	aTest.MustBeEqual(task.response, (*string)(nil))
}
func Test_GetStatus(t *testing.T) {
	aTest := tester.New(t)
	var req = "Feed the cat"
	var task *Task[string, string]

	task = NewTask[string, string]("T1", &req)
	task.status = 123
	aTest.MustBeEqual(task.GetStatus(), ts.TaskStatus(123))
}
func Test_GetTimeOf(t *testing.T) {
	aTest := tester.New(t)
	var req = "Feed the cat"
	var task *Task[string, string]

	task = NewTask[string, string]("T1", &req)
	tOf := to.NewTimeOf()
	task.timeOf = tOf
	aTest.MustBeEqual(task.GetTimeOf(), tOf)
}
func Test_GetId(t *testing.T) {
	aTest := tester.New(t)
	var req = "Feed the cat"
	var task *Task[string, string]

	task = NewTask[string, string]("T1", &req)
	aTest.MustBeEqual(task.GetId(), "T1")
}
func Test_SetTimeOfWorkStart(t *testing.T) {
	//aTest := tester.New(t)
	var req = "Feed the cat"
	var task *Task[string, string]

	task = NewTask[string, string]("T1", &req)

	// We can not check private fields from here, so
	// we simply print the object. Hello, Mister Trololo.
	cio.PrintMessage(task)
	task.SetTimeOfWorkStart()
	cio.PrintMessage(task)
}
func Test_SaveResponse(t *testing.T) {
	aTest := tester.New(t)
	var err error
	var req = "Feed the cat"
	var resp = "The cat has been fed"
	var task *Task[string, string]

	task = NewTask[string, string]("T1", &req)
	task.status = ts.TaskStatus_InWork
	cio.PrintMessage(task)

	err = task.SaveResponse(&resp)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(task.response, &resp)
	aTest.MustBeEqual(task.status, ts.TaskStatus_Processed)

	// We can not check private fields from here, so
	// we simply print the object. Hello, Mister Trololo.
	cio.PrintMessage(task)
}
func Test_Complete(t *testing.T) {
	aTest := tester.New(t)
	var err error
	var req = "Feed the cat"
	var task *Task[string, string]

	task = NewTask[string, string]("T1", &req)
	task.status = ts.TaskStatus_Processed
	cio.PrintMessage(task)

	err = task.Complete()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(task.status, ts.TaskStatus_Completed)

	// We can not check private fields from here, so
	// we simply print the object. Hello, Mister Trololo.
	cio.PrintMessage(task)
}
