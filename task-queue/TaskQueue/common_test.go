package tq

import (
	"time"

	tsk "github.com/vault-thirteen/auxie/task-queue/Task"
	ti "github.com/vault-thirteen/auxie/task-queue/TaskId"
	tqs "github.com/vault-thirteen/auxie/task-queue/TaskQueueSettings"
	tqsp "github.com/vault-thirteen/auxie/task-queue/TaskQueueSettingsPublic"
	ts "github.com/vault-thirteen/auxie/task-queue/TaskStatus"
	"github.com/vault-thirteen/auxie/tester"
)

// Helper data used in tests.

type TestRequest struct {
	a, b int
	fn   string
}

type TestResponse struct {
	c int
}

func test_getCommonQueueSettings() (tqSettings *tqs.TaskQueueSettings) {
	var err error
	tqSettings, err = tqs.NewTaskQueueSettings(&tqsp.TaskQueueSettingsPublic{InputQueueSizeLimit: 2, ErrorChannelSize: 8})
	if err != nil {
		panic(err)
	}

	return tqSettings
}
func test_createCommonTaskQueue(tqSettings *tqs.TaskQueueSettings) (tq *TaskQueue[TestRequest, TestResponse]) {
	var err error
	tq, err = NewTaskQueue[TestRequest, TestResponse](tqSettings)
	if err != nil {
		panic(err)
	}

	return tq
}
func test_createCommonRequest() (req *TestRequest) {
	return &TestRequest{a: 1, b: 2, fn: "x"}
}
func test_checkTaskIdAndStatus(
	aTest *tester.Test,
	tasks map[ti.TaskId]*tsk.Task[TestRequest, TestResponse],
	taskId ti.TaskId,
	taskStatus ts.TaskStatus,
) {
	task := tasks[taskId]
	aTest.MustBeEqual(task != nil, true)
	aTest.MustBeEqual(task.GetId(), taskId)
	aTest.MustBeEqual(task.GetStatus(), taskStatus)
}
func test_errorCatcher(aTest *tester.Test, errHandler func(error)) {
	// Init.
	tqCustomSettings, err := tqs.NewTaskQueueSettings(
		&tqsp.TaskQueueSettingsPublic{
			InputQueueSizeLimit: 2,
			ErrorChannelSize:    8,
			ErrorCatcherFn:      errHandler,
			CleanerSleepTimeSec: 1, // To see an error faster.
		},
	)
	aTest.MustBeNoError(err)
	tq := test_createCommonTaskQueue(tqCustomSettings)

	// Preparation.
	err = tq.Start()
	aTest.MustBeNoError(err)

	// Test #1. Default error handler for error catcher.
	// Start tampering with the registry to raise an error.
	tq.controls.lock.Lock()
	tq.tasks["fake-id"] = tsk.NewTask[TestRequest, TestResponse]("internal-id", nil)
	tq.controls.lock.Unlock()
	time.Sleep(2 * time.Second)

	// Cleaning.
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
