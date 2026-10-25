package tq

import (
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/cio"
	tsk "github.com/vault-thirteen/auxie/task-queue/Task"
	ti "github.com/vault-thirteen/auxie/task-queue/TaskId"
	ts "github.com/vault-thirteen/auxie/task-queue/TaskStatus"
	"github.com/vault-thirteen/auxie/tester"
)

type TestRequest struct {
	a, b int
	fn   string
}

type TestResponse struct {
	c int
}

func Test_StartStop(t *testing.T) {
	aTest := tester.New(t)
	var err error
	tq := NewTaskQueue[TestRequest, TestResponse](2)

	err = tq.Start()
	aTest.MustBeNoError(err)
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
func Test_AddTask(t *testing.T) {
	// Init.
	aTest := tester.New(t)
	var err error
	tq := NewTaskQueue[TestRequest, TestResponse](2)
	err = tq.Start()
	aTest.MustBeNoError(err)
	var task *tsk.Task[TestRequest, TestResponse]
	var taskId ti.TaskId

	// Test.
	req := &TestRequest{a: 1, b: 2, fn: "hello"}
	taskId, err = tq.AddTask(req)
	aTest.MustBeNoError(err)

	// Check.
	task = tq.tasks[taskId]
	aTest.MustBeDifferent(task, (*tsk.Task[TestRequest, TestResponse])(nil))
	aTest.MustBeEqual(task.GetId(), taskId)
	aTest.MustBeEqual(task.GetStatus(), ts.TaskStatus_Queued)

	// Un-queue and process the task to unblock the shutdown.
	task, err = tq.TakeTaskIntoWork()
	aTest.MustBeNoError(err)
	err = tq.MarkTaskAsProcessed(task.GetId(), nil)
	aTest.MustBeNoError(err)
	task, err = tq.CompleteTask(task.GetId())
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(task.GetId(), taskId)

	// Fin.
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
func Test_TakeTaskIntoWork(t *testing.T) {
	// Init.
	aTest := tester.New(t)
	var err error
	tq := NewTaskQueue[TestRequest, TestResponse](2)
	err = tq.Start()
	aTest.MustBeNoError(err)
	var task *tsk.Task[TestRequest, TestResponse]
	var taskId ti.TaskId
	req := &TestRequest{a: 1, b: 2, fn: "hello"}
	taskId, err = tq.AddTask(req)
	aTest.MustBeNoError(err)

	// Test.
	task, err = tq.TakeTaskIntoWork()
	aTest.MustBeNoError(err)

	// Check.
	task = tq.tasks[taskId]
	aTest.MustBeDifferent(task, (*tsk.Task[TestRequest, TestResponse])(nil))
	aTest.MustBeEqual(task.GetId(), taskId)
	aTest.MustBeEqual(task.GetStatus(), ts.TaskStatus_InWork)

	// Process the task to unblock the shutdown.
	err = tq.MarkTaskAsProcessed(task.GetId(), nil)
	aTest.MustBeNoError(err)
	task, err = tq.CompleteTask(task.GetId())
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(task.GetId(), taskId)

	// Fin.
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
func Test_MarkTaskAsProcessed(t *testing.T) {
	// Init.
	aTest := tester.New(t)
	var err error
	tq := NewTaskQueue[TestRequest, TestResponse](2)
	err = tq.Start()
	aTest.MustBeNoError(err)
	var task *tsk.Task[TestRequest, TestResponse]
	var taskId ti.TaskId
	req := &TestRequest{a: 1, b: 2, fn: "hello"}
	taskId, err = tq.AddTask(req)
	aTest.MustBeNoError(err)
	task, err = tq.TakeTaskIntoWork()
	aTest.MustBeNoError(err)

	// Test.
	err = tq.MarkTaskAsProcessed(task.GetId(), nil)
	aTest.MustBeNoError(err)

	// Check.
	task = tq.tasks[taskId]
	aTest.MustBeDifferent(task, (*tsk.Task[TestRequest, TestResponse])(nil))
	aTest.MustBeEqual(task.GetId(), taskId)
	aTest.MustBeEqual(task.GetStatus(), ts.TaskStatus_Processed)

	// Process the task to unblock the shutdown.
	task, err = tq.CompleteTask(task.GetId())
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(task.GetId(), taskId)

	// Fin.
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
func Test_CompleteTask(t *testing.T) {
	// Init.
	aTest := tester.New(t)
	var err error
	tq := NewTaskQueue[TestRequest, TestResponse](2)
	err = tq.Start()
	aTest.MustBeNoError(err)
	var task *tsk.Task[TestRequest, TestResponse]
	var taskId ti.TaskId
	req := &TestRequest{a: 1, b: 2, fn: "hello"}
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

	// Check.
	task = tq.tasks[taskId]
	aTest.MustBeDifferent(task, (*tsk.Task[TestRequest, TestResponse])(nil))
	aTest.MustBeEqual(task.GetId(), taskId)
	aTest.MustBeEqual(task.GetStatus(), ts.TaskStatus_Completed)

	// Fin.
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
func Test_GetTaskStatus(t *testing.T) {
	// Init.
	aTest := tester.New(t)
	var err error
	tq := NewTaskQueue[TestRequest, TestResponse](2)
	err = tq.Start()
	aTest.MustBeNoError(err)
	var task *tsk.Task[TestRequest, TestResponse]
	var taskId ti.TaskId
	req := &TestRequest{a: 1, b: 2, fn: "hello"}
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

	// Check.
	task = tq.tasks[taskId]
	aTest.MustBeDifferent(task, (*tsk.Task[TestRequest, TestResponse])(nil))
	aTest.MustBeEqual(task.GetId(), taskId)
	aTest.MustBeEqual(task.GetStatus(), ts.TaskStatus_Completed)

	// Fin.
	err = tq.Stop()
	aTest.MustBeNoError(err)
}
func Test_GetQueueSize(t *testing.T) {
	// Init.
	aTest := tester.New(t)
	var err error
	tq := NewTaskQueue[TestRequest, TestResponse](2)
	err = tq.Start()
	aTest.MustBeNoError(err)
	var task *tsk.Task[TestRequest, TestResponse]
	var taskId ti.TaskId
	req := &TestRequest{a: 1, b: 2, fn: "hello"}
	taskId, err = tq.AddTask(req)
	aTest.MustBeNoError(err)

	// Test #1.
	var queueSize int
	queueSize, err = tq.GetQueueSize()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(queueSize, 1)

	// Change the queue size by taking a task from it.
	task, err = tq.TakeTaskIntoWork()
	aTest.MustBeNoError(err)

	// Test #2.
	queueSize, err = tq.GetQueueSize()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(queueSize, 0)

	// Process the task to unblock the shutdown.
	err = tq.MarkTaskAsProcessed(task.GetId(), nil)
	aTest.MustBeNoError(err)
	task, err = tq.CompleteTask(task.GetId())
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(task.GetId(), taskId)

	// Fin.
	err = tq.Stop()
	aTest.MustBeNoError(err)
}

func Test_stress(t *testing.T) {
	aTest := tester.New(t)
	cio.PrintMessage("This test can take some time. Please, wait.")

	tq := NewTaskQueue[TestRequest, TestResponse](100)

	err := tq.Start()
	aTest.MustBeNoError(err)

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

	err = tq.Stop()
	aTest.MustBeNoError(err)
}
