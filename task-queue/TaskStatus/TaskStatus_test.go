package ts

import (
	"testing"

	"github.com/vault-thirteen/auxie/cio"
	ce "github.com/vault-thirteen/auxie/task-queue/common/error"
	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewTaskStatus(t *testing.T) {
	aTest := tester.New(t)

	x := NewTaskStatus()
	aTest.MustBeEqual(x, TaskStatus_Created)
}

func Test_Change(t *testing.T) {
	type TestData struct {
		ts1               TaskStatus
		arg               TaskStatus
		ts2               TaskStatus
		isErrorExpected   bool
		expectedErrorText string
	}

	aTest := tester.New(t)
	var err error
	testData := []TestData{
		{
			ts1:               123,
			arg:               TaskStatus_Unknown,
			ts2:               123,
			isErrorExpected:   true,
			expectedErrorText: ce.Err_UnknownStatusCanNotBeSet,
		},
		{
			ts1:               123,
			arg:               TaskStatus_Created,
			ts2:               123,
			isErrorExpected:   true,
			expectedErrorText: ce.Err_OnlyNewTasksCanHaveStatusCreated,
		},
		{
			ts1:               123,
			arg:               TaskStatus_Queued,
			ts2:               123,
			isErrorExpected:   true,
			expectedErrorText: ce.Err_OnlyCreatedTasksCanBePutIntoQueue,
		},
		{
			ts1:               TaskStatus_Created,
			arg:               TaskStatus_Queued,
			ts2:               TaskStatus_Queued,
			isErrorExpected:   false,
			expectedErrorText: "",
		},
		{
			ts1:               123,
			arg:               TaskStatus_InWork,
			ts2:               123,
			isErrorExpected:   true,
			expectedErrorText: ce.Err_OnlyQueuedTasksCanBeTakenIntoWork,
		},
		{
			ts1:               TaskStatus_Queued,
			arg:               TaskStatus_InWork,
			ts2:               TaskStatus_InWork,
			isErrorExpected:   false,
			expectedErrorText: "",
		},
		{
			ts1:               123,
			arg:               TaskStatus_Processed,
			ts2:               123,
			isErrorExpected:   true,
			expectedErrorText: ce.Err_OnlyTasksInWorkCanBeProcessed,
		},
		{
			ts1:               TaskStatus_InWork,
			arg:               TaskStatus_Processed,
			ts2:               TaskStatus_Processed,
			isErrorExpected:   false,
			expectedErrorText: "",
		},
		{
			ts1:               123,
			arg:               TaskStatus_Completed,
			ts2:               123,
			isErrorExpected:   true,
			expectedErrorText: ce.Err_OnlyProcessedTasksCanBeCompleted,
		},
		{
			ts1:               TaskStatus_Processed,
			arg:               TaskStatus_Completed,
			ts2:               TaskStatus_Completed,
			isErrorExpected:   false,
			expectedErrorText: "",
		},
		{
			ts1:               TaskStatus_Unknown,
			arg:               123,
			ts2:               TaskStatus_Unknown,
			isErrorExpected:   true,
			expectedErrorText: cio.FormatMessage(ce.Errf_UnsupportedStatus, "123"),
		},
	}

	var status *TaskStatus

	for i, td := range testData {
		cio.Print(cio.FormatMessage("[%d]", i+1))

		status = new(TaskStatus)
		*status = td.ts1

		err = status.Change(td.arg)
		if td.isErrorExpected {
			aTest.MustBeAnError(err)
			aTest.MustBeEqual(err.Error(), td.expectedErrorText)
		} else {
			aTest.MustBeNoError(err)
		}

		aTest.MustBeEqual(*status, td.ts2)
	}
	cio.PrintMessage()
}
