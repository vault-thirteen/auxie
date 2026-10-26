package tq

import (
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/cio"
	tsk "github.com/vault-thirteen/auxie/task-queue/Task"
	ti "github.com/vault-thirteen/auxie/task-queue/TaskId"
	tqs "github.com/vault-thirteen/auxie/task-queue/TaskQueueSettings"
	tqsp "github.com/vault-thirteen/auxie/task-queue/TaskQueueSettingsPublic"
	ts "github.com/vault-thirteen/auxie/task-queue/TaskStatus"
	ce "github.com/vault-thirteen/auxie/task-queue/common/error"
	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewTaskQueue(t *testing.T) {
	aTest := tester.New(t)
	var settings *tqs.TaskQueueSettings
	var tq *TaskQueue[TestRequest, TestResponse]
	var err error

	// Test #1. No settings.
	settings = nil
	tq, err = NewTaskQueue[TestRequest, TestResponse](settings)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), ce.Err_NullPointer)

	// Test #2. Invalid settings.
	// We can not provide invalid settings while they use private fields.

	// Test #3. Normal test.
	settings = test_getCommonQueueSettings()
	tq, err = NewTaskQueue[TestRequest, TestResponse](settings)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(tq.settings, settings)
	aTest.MustBeEqual(tq.inQueue != nil, true)
	aTest.MustBeEqual(tq.tasks != nil, true)
	aTest.MustBeEqual(tq.lastTaskId != nil, true)
	aTest.MustBeEqual(tq.controls != nil, true)
}

func Test_StartStop(t *testing.T) {
	aTest := tester.New(t)
	var err error
	tqSettings := test_getCommonQueueSettings()
	tq := test_createCommonTaskQueue(tqSettings)

	err = tq.Start()
	aTest.MustBeNoError(err)
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
func Test_AddTask(t *testing.T) {
	// Init.
	aTest := tester.New(t)
	tqSettings := test_getCommonQueueSettings()
	tq := test_createCommonTaskQueue(tqSettings)
	req := test_createCommonRequest()
	var task *tsk.Task[TestRequest, TestResponse]
	var taskId ti.TaskId
	var err error

	// Preparation.
	err = tq.Start()
	aTest.MustBeNoError(err)

	// Test.
	taskId, err = tq.AddTask(req)
	aTest.MustBeNoError(err)
	test_checkTaskIdAndStatus(aTest, tq.tasks, taskId, ts.TaskStatus_Queued)

	// Cleaning.
	task, err = tq.TakeTaskIntoWork()
	aTest.MustBeNoError(err)
	err = tq.MarkTaskAsProcessed(task.GetId(), nil)
	aTest.MustBeNoError(err)
	task, err = tq.CompleteTask(task.GetId())
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(task.GetId(), taskId)
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
func Test_TakeTaskIntoWork(t *testing.T) {
	// Init.
	aTest := tester.New(t)
	tqSettings := test_getCommonQueueSettings()
	tq := test_createCommonTaskQueue(tqSettings)
	req := test_createCommonRequest()
	var task *tsk.Task[TestRequest, TestResponse]
	var taskId ti.TaskId
	var err error

	// Preparation.
	err = tq.Start()
	aTest.MustBeNoError(err)
	taskId, err = tq.AddTask(req)
	aTest.MustBeNoError(err)

	// Test.
	task, err = tq.TakeTaskIntoWork()
	aTest.MustBeNoError(err)
	test_checkTaskIdAndStatus(aTest, tq.tasks, taskId, ts.TaskStatus_InWork)

	// Cleaning.
	err = tq.MarkTaskAsProcessed(task.GetId(), nil)
	aTest.MustBeNoError(err)
	task, err = tq.CompleteTask(task.GetId())
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(task.GetId(), taskId)
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
func Test_MarkTaskAsProcessed(t *testing.T) {
	// Init.
	aTest := tester.New(t)
	tqSettings := test_getCommonQueueSettings()
	tq := test_createCommonTaskQueue(tqSettings)
	req := test_createCommonRequest()
	var task *tsk.Task[TestRequest, TestResponse]
	var taskId ti.TaskId
	var err error

	// Preparation.
	err = tq.Start()
	aTest.MustBeNoError(err)
	taskId, err = tq.AddTask(req)
	aTest.MustBeNoError(err)
	task, err = tq.TakeTaskIntoWork()
	aTest.MustBeNoError(err)

	// Test.
	err = tq.MarkTaskAsProcessed(task.GetId(), nil)
	aTest.MustBeNoError(err)
	test_checkTaskIdAndStatus(aTest, tq.tasks, taskId, ts.TaskStatus_Processed)

	// Cleaning.
	task, err = tq.CompleteTask(task.GetId())
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(task.GetId(), taskId)
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
func Test_CompleteTask(t *testing.T) {
	// Init.
	aTest := tester.New(t)
	tqSettings := test_getCommonQueueSettings()
	tq := test_createCommonTaskQueue(tqSettings)
	req := test_createCommonRequest()
	var task *tsk.Task[TestRequest, TestResponse]
	var taskId ti.TaskId
	var err error

	// Preparation.
	err = tq.Start()
	aTest.MustBeNoError(err)
	taskId, err = tq.AddTask(req)
	aTest.MustBeNoError(err)
	task, err = tq.TakeTaskIntoWork()
	aTest.MustBeNoError(err)
	err = tq.MarkTaskAsProcessed(task.GetId(), nil)
	aTest.MustBeNoError(err)

	// Test.
	task, err = tq.CompleteTask(task.GetId())
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(task.GetId(), taskId)
	test_checkTaskIdAndStatus(aTest, tq.tasks, taskId, ts.TaskStatus_Completed)

	// Cleaning.
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
func Test_GetTaskStatus(t *testing.T) {
	// Init.
	aTest := tester.New(t)
	tqSettings := test_getCommonQueueSettings()
	tq := test_createCommonTaskQueue(tqSettings)
	req := test_createCommonRequest()
	var task *tsk.Task[TestRequest, TestResponse]
	var taskId ti.TaskId
	var err error

	// Preparation.
	err = tq.Start()
	aTest.MustBeNoError(err)
	taskId, err = tq.AddTask(req)
	aTest.MustBeNoError(err)
	task, err = tq.TakeTaskIntoWork()
	aTest.MustBeNoError(err)
	err = tq.MarkTaskAsProcessed(task.GetId(), nil)
	aTest.MustBeNoError(err)
	task, err = tq.CompleteTask(task.GetId())
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(task.GetId(), taskId)

	// Test.
	var status ts.TaskStatus
	status, err = tq.GetTaskStatus(taskId)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(status, ts.TaskStatus_Completed)
	test_checkTaskIdAndStatus(aTest, tq.tasks, taskId, ts.TaskStatus_Completed)

	// Cleaning.
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
func Test_GetQueueSize(t *testing.T) {
	// Init.
	aTest := tester.New(t)
	tqSettings := test_getCommonQueueSettings()
	tq := test_createCommonTaskQueue(tqSettings)
	req := test_createCommonRequest()
	var task *tsk.Task[TestRequest, TestResponse]
	var taskId ti.TaskId
	var err error

	// Preparation.
	err = tq.Start()
	aTest.MustBeNoError(err)
	taskId, err = tq.AddTask(req)
	aTest.MustBeNoError(err)

	// Test #1.
	var queueSize int
	queueSize, err = tq.GetQueueSize()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(queueSize, 1)

	// Test #2.
	// Change the queue size by taking a task from it.
	task, err = tq.TakeTaskIntoWork()
	aTest.MustBeNoError(err)
	queueSize, err = tq.GetQueueSize()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(queueSize, 0)

	// Cleaning.
	err = tq.MarkTaskAsProcessed(task.GetId(), nil)
	aTest.MustBeNoError(err)
	task, err = tq.CompleteTask(task.GetId())
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(task.GetId(), taskId)
	err = tq.Stop()
	aTest.MustBeNoError(err)
}

func Test_errorCatcher(t *testing.T) {
	aTest := tester.New(t)

	// Test #1.
	test_errorCatcher(aTest, nil)

	// Test #2.
	customErrorHandler := func(err error) {
		cio.PrintMessage("in custom error handler, error:", err)
	}
	test_errorCatcher(aTest, customErrorHandler)
}
func Test_stress(t *testing.T) {
	cio.PrintMessage("This test can take some time. Please, wait.")

	// Init.
	aTest := tester.New(t)
	tqCustomSettings, err := tqs.NewTaskQueueSettings(
		&tqsp.TaskQueueSettingsPublic{
			InputQueueSizeLimit: 100,
			ErrorChannelSize:    8,
		},
	)
	aTest.MustBeNoError(err)
	tq := test_createCommonTaskQueue(tqCustomSettings)

	// Preparation.
	err = tq.Start()
	aTest.MustBeNoError(err)

	// Test.
	iMax := 1000
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		for i := 0; i < iMax; i++ {
			req := &TestRequest{a: i, b: i, fn: ""}
			sb := strings.Builder{}
			for j := 0; j < 1000; j++ {
				sb.WriteString("QWERTY123")
			}
			req.fn = sb.String()

			for {
				_, errA := tq.AddTask(req)
				if errA == nil {
					break
				}
				time.Sleep(time.Millisecond)
			}
		}
		wg.Done()
	}()

	go func() {
		var errB error
		for i := 0; i < iMax; i++ {
			var wt *tsk.Task[TestRequest, TestResponse]
			for {
				wt, errB = tq.TakeTaskIntoWork()
				aTest.MustBeNoError(errB)
				if wt != nil {
					break
				}
				time.Sleep(time.Millisecond)
			}

			// Simulate some work.
			time.Sleep(time.Millisecond * 10)

			resp := &TestResponse{c: 2 * i}
			taskId := wt.GetId()
			errB = tq.MarkTaskAsProcessed(taskId, resp)
			aTest.MustBeNoError(errB)

			_, errB = tq.CompleteTask(taskId)
			aTest.MustBeNoError(errB)
		}
		wg.Done()
	}()

	wg.Wait()

	var qSize int
	qSize, err = tq.GetQueueSize()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(qSize, 0)

	// Cleaning.
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
