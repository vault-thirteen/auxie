package tq

import (
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"
	q "github.com/vault-thirteen/auxie/queue"
)

const (
	CleanerSleepTimeSec = 60

	// CompletedTaskTtlSec is time period after which completed task may be
	// deleted.
	CompletedTaskTtlSec = 1 * 60
)

type TaskQueue[T1 any, T2 any] struct {
	lock            *sync.Mutex
	inQueue         *q.Queue[Task[T1, T2]]
	tasks           map[string]*Task[T1, T2]
	sizeLimit       int
	lastTaskId      TaskId
	cleanerWg       *sync.WaitGroup
	receiveNewTasks *atomic.Bool
	cleanerMustWork *atomic.Bool
}

func NewTaskQueue[T1 any, T2 any](sizeLimit int) (tq *TaskQueue[T1, T2]) {
	tq = &TaskQueue[T1, T2]{
		lock:            new(sync.Mutex),
		inQueue:         q.NewQueue[Task[T1, T2]](),
		tasks:           make(map[string]*Task[T1, T2]),
		sizeLimit:       sizeLimit,
		lastTaskId:      big.NewInt(0),
		cleanerWg:       new(sync.WaitGroup),
		receiveNewTasks: new(atomic.Bool),
		cleanerMustWork: new(atomic.Bool),
	}

	tq.receiveNewTasks.Store(true)
	tq.cleanerMustWork.Store(true)

	tq.cleanerWg.Add(1)
	go tq.cleaner()

	return tq
}

func (tq *TaskQueue[T1, T2]) AddTask(request *T1) (taskId TaskId, err error) {
	tq.lock.Lock()
	defer tq.lock.Unlock()

	if !tq.receiveNewTasks.Load() {
		return nil, errors.ErrorFromText(Err_NotReceivingNewTasks)
	}

	if tq.inQueue.Size() >= tq.sizeLimit {
		return nil, errors.ErrorFromText(Err_QueueIsFull)
	}

	// Task ID.
	tq.lastTaskId.Add(tq.lastTaskId, big.NewInt(1))
	taskId = new(big.Int).Set(tq.lastTaskId)

	task := NewTask[T1, T2](request, taskId)

	tq.inQueue.Push(task)
	task.Status = TaskStatus_Queued

	tq.tasks[task.Id.String()] = task

	return task.Id, nil
}
func (tq *TaskQueue[T1, T2]) GetTaskStatus(taskId string) (taskStatus TaskStatus, err error) {
	tq.lock.Lock()
	defer tq.lock.Unlock()

	task, ok := tq.tasks[taskId]
	if !ok {
		return TaskStatus_Unknown, errors.ErrorFromText(Err_TaskIsNotFound)
	}

	return task.Status, nil
}
func (tq *TaskQueue[T1, T2]) GetTaskToWork() (task *Task[T1, T2]) {
	tq.lock.Lock()
	defer tq.lock.Unlock()

	task = tq.inQueue.Pop()
	if task == nil {
		return nil
	}

	task.WorkStart = time.Now().UTC().Unix()
	task.Status = TaskStatus_InWork

	return task
}
func (tq *TaskQueue[T1, T2]) MarkTaskAsProcessed(taskId string, response *T2) (err error) {
	tq.lock.Lock()
	defer tq.lock.Unlock()

	task, ok := tq.tasks[taskId]
	if !ok {
		return errors.ErrorFromText(Err_TaskIsNotFound)
	}

	task.Response = response
	task.Processed = time.Now().UTC().Unix()
	task.Status = TaskStatus_Processed

	return nil
}
func (tq *TaskQueue[T1, T2]) ShowTaskToClient(taskId string) (task *Task[T1, T2], err error) {
	tq.lock.Lock()
	defer tq.lock.Unlock()

	var ok bool
	task, ok = tq.tasks[taskId]
	if !ok {
		return nil, errors.ErrorFromText(Err_TaskIsNotFound)
	}

	task.Completed = time.Now().UTC().Unix()
	task.Status = TaskStatus_Completed

	return task, nil
}
func (tq *TaskQueue[T1, T2]) Shutdown() (err error) {
	// 1. Stop receiving new tasks.
	{
		tq.receiveNewTasks.Store(false)
	}

	// 2. Wait for queue to become empty.
	{
		cio.Print(Msg_WaitingForEmptyQueue)
		var qs int
		for {
			tq.lock.Lock()
			qs = tq.inQueue.Size()
			tq.lock.Unlock()

			if qs > 0 {
				time.Sleep(time.Second)
			} else {
				break
			}
		}
		cio.PrintMessage(Msg_OK)
	}

	// 3. Wait for all tasks in work to change status.
	{
		cio.Print(Msg_WaitingForTasksInWork)
		var n int
		for {
			n = tq.countTasksWithStatus(TaskStatus_InWork)
			if n == 0 {
				break
			}

			time.Sleep(time.Second)
		}
		cio.PrintMessage(Msg_OK)
	}

	// 4. Wait for all processed tasks to change status.
	{
		cio.Print(Msg_WaitingForProcessedTasks)
		var n int
		for {
			n = tq.countTasksWithStatus(TaskStatus_Processed)
			if n == 0 {
				break
			}

			time.Sleep(time.Second)
		}
		cio.PrintMessage(Msg_OK)
	}

	// 4. Stop the cleaner.
	{
		cio.Print(Msg_WaitingForCleaner)
		tq.cleanerMustWork.Store(false)
		tq.cleanerWg.Wait()
		cio.PrintMessage(Msg_OK)
	}

	return nil
}
func (tq *TaskQueue[T1, T2]) Size() int {
	tq.lock.Lock()
	defer tq.lock.Unlock()

	return tq.inQueue.Size()
}
func (tq *TaskQueue[T1, T2]) cleaner() {
	defer tq.cleanerWg.Done()

	loop := true
	for loop {
		tq.clean()

		for i := 0; i < CleanerSleepTimeSec; i++ {
			time.Sleep(time.Second)
			if !tq.cleanerMustWork.Load() {
				loop = false
				break
			}
		}
	}
}
func (tq *TaskQueue[T1, T2]) clean() {
	tq.lock.Lock()
	defer tq.lock.Unlock()

	timeNowTS := time.Now().UTC().Unix()

	for k, task := range tq.tasks {
		if task.Status == TaskStatus_Completed {
			if task.Completed+CompletedTaskTtlSec < timeNowTS {
				delete(tq.tasks, k)
			}
		}
	}
}
func (tq *TaskQueue[T1, T2]) countTasksWithStatus(status TaskStatus) (n int) {
	tq.lock.Lock()
	defer tq.lock.Unlock()

	n = 0
	for _, task := range tq.tasks {
		if task.Status == status {
			n++
		}
	}

	return n
}
