package tq

import (
	"math/big"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/tester"
)

type Request struct {
	a, b int
	fn   string
}

type Response struct {
	c int
}

func Test_all(t *testing.T) {
	aTest := tester.New(t)
	tq := NewTaskQueue[Request, Response](2)

	req1 := &Request{a: 101, b: 1001, fn: "sum"}
	taskId, err := tq.AddTask(req1)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(taskId.String(), "1")
	aTest.MustBeEqual(tq.inQueue.Size(), 1)
	var status TaskStatus
	status, err = tq.GetTaskStatus("1")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(status, TaskStatus_Queued)

	req2 := &Request{a: 202, b: 2002, fn: "sum"}
	taskId, err = tq.AddTask(req2)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(taskId.String(), "2")
	aTest.MustBeEqual(tq.inQueue.Size(), 2)
	aTest.MustBeEqual(tq.tasks["2"].Status, TaskStatus_Queued)

	req3 := &Request{a: 303, b: 3003, fn: "sum"}
	taskId, err = tq.AddTask(req3)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), Err_QueueIsFull)
	aTest.MustBeEqual(taskId, (*big.Int)(nil))
	aTest.MustBeEqual(tq.inQueue.Size(), 2)
	status, err = tq.GetTaskStatus("3")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), Err_TaskIsNotFound)

	time.Sleep(time.Second)
	workedTask := tq.GetTaskToWork()
	aTest.MustBeEqual(workedTask.Status, TaskStatus_InWork)
	aTest.MustBeEqual(workedTask.WorkStart > 0, true)
	aTest.MustBeEqual(tq.inQueue.Size(), 1)
	time.Sleep(time.Second)
	resp := &Response{c: 1102}
	err = tq.MarkTaskAsProcessed("1", resp)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(workedTask.Status, TaskStatus_Processed)
	aTest.MustBeEqual(workedTask.Processed > 0, true)
	time.Sleep(time.Second)
	var task *Task[Request, Response]
	task, err = tq.ShowTaskToClient("1")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(workedTask.Status, TaskStatus_Completed)
	aTest.MustBeEqual(task.Completed > 0, true)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		gerr := tq.Shutdown()
		aTest.MustBeNoError(gerr)
		wg.Done()
	}()

	time.Sleep(time.Second)
	workedTask = tq.GetTaskToWork()
	resp = &Response{c: 2204}
	err = tq.MarkTaskAsProcessed("2", resp)
	task, err = tq.ShowTaskToClient("99")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), Err_TaskIsNotFound)

	time.Sleep(time.Second)
	req4 := &Request{a: 4, b: 4, fn: "none"}
	taskId, err = tq.AddTask(req4)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), Err_NotReceivingNewTasks)
	time.Sleep(time.Second)
	cio.PrintMessage(err)

	aTest.MustBeEqual(tq.Size(), 0)
	time.Sleep(time.Second)

	aTest.MustBeEqual(tq.Size(), 0)
	aTest.MustBeEqual(len(tq.tasks), 2)
	aTest.MustBeEqual(tq.countTasksWithStatus(TaskStatus_Processed), 1)
	aTest.MustBeEqual(tq.countTasksWithStatus(TaskStatus_Completed), 1)

	task, err = tq.ShowTaskToClient("2")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(tq.countTasksWithStatus(TaskStatus_Processed), 0)
	aTest.MustBeEqual(tq.countTasksWithStatus(TaskStatus_Completed), 2)

	wg.Wait()
}

func Test_stress(t *testing.T) {
	aTest := tester.New(t)
	cio.PrintMessage("This test can take up to several minutes to pass. Please, wait.")

	tq := NewTaskQueue[Request, Response](1_000)

	iMax := 25_000
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		for i := 0; i < iMax; i++ {
			req := &Request{a: i, b: i, fn: ""}
			sb := strings.Builder{}
			for j := 0; j < 1000; j++ {
				sb.WriteString("QWERTY123")
			}
			req.fn = sb.String()

			var err error
			for {
				_, err = tq.AddTask(req)
				if err == nil {
					break
				}
				time.Sleep(time.Millisecond)
			}
			aTest.MustBeNoError(err)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < iMax; i++ {
			var wt *Task[Request, Response]
			for {
				wt = tq.GetTaskToWork()
				if wt != nil {
					break
				}
				time.Sleep(time.Millisecond)
			}

			// Simulate some work.
			time.Sleep(time.Millisecond * 10)

			resp := &Response{c: 2 * i}
			taskId := wt.Id.String()
			err := tq.MarkTaskAsProcessed(taskId, resp)
			aTest.MustBeNoError(err)

			_, err = tq.ShowTaskToClient(taskId)
			aTest.MustBeNoError(err)
		}
		wg.Done()
	}()

	wg.Wait()
	err := tq.Shutdown()
	aTest.MustBeNoError(err)
}
