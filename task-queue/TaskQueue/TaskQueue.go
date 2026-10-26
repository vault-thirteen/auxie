package tq

import (
	"time"

	"github.com/vault-thirteen/auxie/cio"
	c "github.com/vault-thirteen/auxie/counter"
	"github.com/vault-thirteen/auxie/errors"
	q "github.com/vault-thirteen/auxie/queue"
	t "github.com/vault-thirteen/auxie/task-queue/Task"
	ti "github.com/vault-thirteen/auxie/task-queue/TaskId"
	tqs "github.com/vault-thirteen/auxie/task-queue/TaskQueueSettings"
	ts "github.com/vault-thirteen/auxie/task-queue/TaskStatus"
	to "github.com/vault-thirteen/auxie/task-queue/TimeOf"
	tp "github.com/vault-thirteen/auxie/task-queue/TimeStamp"
	ce "github.com/vault-thirteen/auxie/task-queue/common/error"
	cm "github.com/vault-thirteen/auxie/task-queue/common/message"
)

type TaskQueue[T1 any, T2 any] struct {
	settings   *tqs.TaskQueueSettings
	inQueue    *q.Queue[t.Task[T1, T2]]
	tasks      map[ti.TaskId]*t.Task[T1, T2]
	lastTaskId *c.Counter
	controls   *TaskQueueControls
}

func NewTaskQueue[T1 any, T2 any](settings *tqs.TaskQueueSettings) (tq *TaskQueue[T1, T2], err error) {
	if settings == nil {
		return nil, errors.ErrorFromText(ce.Err_NullPointer)
	}

	err = settings.Validate()
	if err != nil {
		return nil, err
	}

	tq = &TaskQueue[T1, T2]{
		settings:   settings,
		inQueue:    q.NewQueue[t.Task[T1, T2]](),
		tasks:      make(map[string]*t.Task[T1, T2]),
		lastTaskId: c.NewCounter(),
		controls:   NewTaskQueueControls(settings),
	}

	return tq, nil
}

func (tq *TaskQueue[T1, T2]) startErrorCatcher() {
	tq.controls.errorCatcherMustWork.Store(true)
	tq.controls.errorCatcherWg.Add(1)
	go tq.errorCatcher()
}
func (tq *TaskQueue[T1, T2]) stopErrorCatcher() {
	tq.controls.errorCatcherMustWork.Store(false)
	tq.controls.errorCatcherWg.Wait()
}
func (tq *TaskQueue[T1, T2]) errorCatcher() {
	defer tq.controls.errorCatcherWg.Done()

	var ecFn = tq.defaultErrorCatcherFn
	if tq.controls.errorCatcherFn != nil {
		ecFn = tq.controls.errorCatcherFn
	}

	var sleepTime = time.Duration(tq.settings.ErrorCatcherSleepTimeSec()) * time.Second
	var loop = true
	var err error

	for loop {
		select {
		case err = <-*tq.controls.errors:
			{
				ecFn(err)
			}

		default:
			{
				if !tq.controls.errorCatcherMustWork.Load() {
					loop = false
					break
				}
				time.Sleep(sleepTime)
			}
		}
	}
}
func (tq *TaskQueue[T1, T2]) defaultErrorCatcherFn(err error) {
	errors.LogError(err)
}
func (tq *TaskQueue[T1, T2]) startCleaner() {
	tq.controls.cleanerMustWork.Store(true)
	tq.controls.cleanerWg.Add(1)
	go tq.cleaner()
}
func (tq *TaskQueue[T1, T2]) stopCleaner() {
	tq.controls.cleanerMustWork.Store(false)
	tq.controls.cleanerWg.Wait()
}
func (tq *TaskQueue[T1, T2]) cleaner() {
	defer tq.controls.cleanerWg.Done()

	var sleepTimeSec = tq.settings.CleanerSleepTimeSec()
	var loop = true

	for loop {
		tq.clean()

		// Break sleep time into several parts in order not to miss signals.
		for i := 0; i < sleepTimeSec; i++ {
			time.Sleep(time.Second)
			if !tq.controls.cleanerMustWork.Load() {
				loop = false
				break
			}
		}
	}
}
func (tq *TaskQueue[T1, T2]) clean() {
	tq.controls.lock.Lock()
	defer tq.controls.lock.Unlock()

	var timeNowTS = time.Now().UTC().Unix()
	var completedTaskTtlSec = tp.TimeStamp(tq.settings.CompletedTaskTtlSec())

	// Delete old completed tasks.
	var err error
	var tOf to.TimeOf
	var tasksIdsForRemoval []ti.TaskId

	// Collect IDs.
	tasksIdsForRemoval = make([]ti.TaskId, 0, len(tq.tasks))
	for taskId, task := range tq.tasks {
		// Check data consistency.
		if task.GetId() != taskId {
			err = errors.ErrorFromFormattedText(ce.Errf_BrokenTaskInRegistry, taskId)
			*tq.controls.errors <- err
		}

		// Find an old completed task.
		if task.GetStatus() == ts.TaskStatus_Completed {
			tOf = task.GetTimeOf()
			if tOf.Completion()+completedTaskTtlSec < timeNowTS {
				tasksIdsForRemoval = append(tasksIdsForRemoval, taskId)
			}
		}
	}

	// Remove old completed tasks from registry.
	for _, taskId := range tasksIdsForRemoval {
		err = tq.deleteTask(taskId)
		if err != nil {
			*tq.controls.errors <- err
		}
	}
}

func (tq *TaskQueue[T1, T2]) enableReceivingNewTasks() {
	tq.controls.receiveNewTasks.Store(true)
}
func (tq *TaskQueue[T1, T2]) disableReceivingNewTasks() {
	tq.controls.receiveNewTasks.Store(false)
}
func (tq *TaskQueue[T1, T2]) waitForEmptyQueue() {
	var sleepTime = time.Duration(tq.settings.SleepTimeWaitingForZeroQueueSec()) * time.Second

	var qSize int
	for {
		qSize = tq.getQueueSize()

		if qSize != 0 {
			time.Sleep(sleepTime)
		} else {
			return
		}
	}
}
func (tq *TaskQueue[T1, T2]) getQueueSize() (n int) {
	tq.controls.lock.Lock()
	defer tq.controls.lock.Unlock()

	return tq.inQueue.Size()
}
func (tq *TaskQueue[T1, T2]) waitForZeroTasksInWork() {
	var sleepTime = time.Duration(tq.settings.SleepTimeWaitingForZeroTasksInWorkSec()) * time.Second

	var n int
	for {
		n = tq.countTasksWithStatus(ts.TaskStatus_InWork)

		if n != 0 {
			time.Sleep(sleepTime)
		} else {
			return
		}
	}
}
func (tq *TaskQueue[T1, T2]) waitForZeroProcessedTasks() {
	var sleepTime = time.Duration(tq.settings.SleepTimeWaitingForZeroProcessedTasksSec()) * time.Second

	var n int
	for {
		n = tq.countTasksWithStatus(ts.TaskStatus_Processed)

		if n != 0 {
			time.Sleep(sleepTime)
		} else {
			return
		}
	}
}
func (tq *TaskQueue[T1, T2]) countTasksWithStatus(status ts.TaskStatus) (n int) {
	tq.controls.lock.Lock()
	defer tq.controls.lock.Unlock()

	n = 0
	for _, task := range tq.tasks {
		if task.GetStatus() == status {
			n++
		}
	}

	return n
}
func (tq *TaskQueue[T1, T2]) registerTask(taskId ti.TaskId, task *t.Task[T1, T2]) (err error) {
	var taskAlreadyExists bool
	_, taskAlreadyExists = tq.tasks[taskId]
	if taskAlreadyExists {
		return errors.ErrorFromFormattedText(ce.Errf_TaskAlreadyExists, taskId)
	}

	tq.tasks[taskId] = task

	return nil
}
func (tq *TaskQueue[T1, T2]) getTask(taskId ti.TaskId) (task *t.Task[T1, T2], err error) {
	var ok bool
	task, ok = tq.tasks[taskId]
	if !ok {
		return nil, errors.ErrorFromFormattedText(ce.Errf_TaskIsNotFound, taskId)
	}

	return task, nil
}
func (tq *TaskQueue[T1, T2]) deleteTask(taskId ti.TaskId) (err error) {
	task, taskExists := tq.tasks[taskId]
	if !taskExists {
		return errors.ErrorFromFormattedText(ce.Errf_TaskIsNotFound, taskId)
	}

	err = task.Destroy()
	if err != nil {
		return err
	}

	delete(tq.tasks, taskId)

	return nil
}

func (tq *TaskQueue[T1, T2]) Start() (err error) {
	err = tq.controls.ssc.BeginStart()
	if err != nil {
		return err
	}

	// 1. Error Catcher.
	tq.startErrorCatcher()

	// 2. Cleaner.
	tq.startCleaner()

	// 3. Receiving Tasks.
	tq.enableReceivingNewTasks()

	err = tq.controls.ssc.CompleteStart()
	if err != nil {
		return err
	}

	return nil
}
func (tq *TaskQueue[T1, T2]) Stop() (err error) {
	err = tq.controls.ssc.BeginStop()
	if err != nil {
		return err
	}

	// 1. Receiving Tasks.
	tq.disableReceivingNewTasks()
	cio.PrintMessage(cm.Msg_ReceivingNewTasksIsDisabled)

	// 2. Wait for queue to become empty.
	cio.Print(cm.Msg_WaitingForEmptyQueue)
	tq.waitForEmptyQueue()
	cio.PrintMessage(cm.Msg_OK)

	// 3. Wait for all tasks in work to change their status.
	cio.Print(cm.Msg_WaitingForTasksInWork)
	tq.waitForZeroTasksInWork()
	cio.PrintMessage(cm.Msg_OK)

	// 4. Wait for all processed tasks to change their status.
	cio.Print(cm.Msg_WaitingForProcessedTasks)
	tq.waitForZeroProcessedTasks()
	cio.PrintMessage(cm.Msg_OK)

	// 5. Cleaner.
	cio.Print(cm.Msg_StoppingCleaner)
	tq.stopCleaner()
	cio.PrintMessage(cm.Msg_OK)

	// 6. Error Catcher.
	cio.Print(cm.Msg_StoppingErrorCatcher)
	tq.stopErrorCatcher()
	cio.PrintMessage(cm.Msg_OK)

	// 7. Error channel.
	close(*tq.controls.errors)

	err = tq.controls.ssc.CompleteStop()
	if err != nil {
		return err
	}

	return nil
}

func (tq *TaskQueue[T1, T2]) AddTask(request *T1) (taskId ti.TaskId, err error) {
	tq.controls.lock.Lock()
	defer tq.controls.lock.Unlock()

	if !tq.controls.receiveNewTasks.Load() {
		return ti.ZeroTaskId, errors.ErrorFromText(ce.Err_NotReceivingNewTasks)
	}

	if tq.inQueue.Size() >= tq.settings.InputQueueSizeLimit() {
		return ti.ZeroTaskId, errors.ErrorFromText(ce.Err_QueueIsFull)
	}

	// Task ID.
	tq.lastTaskId.Inc()
	taskId = tq.lastTaskId.ToString()
	task := t.NewTask[T1, T2](taskId, request)

	err = tq.registerTask(taskId, task)
	if err != nil {
		return ti.ZeroTaskId, err
	}

	tq.inQueue.Push(task)

	err = task.ChangeStatus(ts.TaskStatus_Queued)
	if err != nil {
		return ti.ZeroTaskId, err
	}

	return task.GetId(), nil
}
func (tq *TaskQueue[T1, T2]) TakeTaskIntoWork() (task *t.Task[T1, T2], err error) {
	tq.controls.lock.Lock()
	defer tq.controls.lock.Unlock()

	isRunningOrStopping := tq.controls.ssc.IsRunning() || tq.controls.ssc.IsStopping()
	if !isRunningOrStopping {
		return nil, errors.ErrorFromText(ce.Err_NotRunningOrStopping)
	}

	task = tq.inQueue.Pop()
	if task == nil {
		return nil, nil
	}

	task.SetTimeOfWorkStart()

	err = task.ChangeStatus(ts.TaskStatus_InWork)
	if err != nil {
		return nil, err
	}

	return task, nil
}
func (tq *TaskQueue[T1, T2]) MarkTaskAsProcessed(taskId string, response *T2) (err error) {
	tq.controls.lock.Lock()
	defer tq.controls.lock.Unlock()

	isRunningOrStopping := tq.controls.ssc.IsRunning() || tq.controls.ssc.IsStopping()
	if !isRunningOrStopping {
		return errors.ErrorFromText(ce.Err_NotRunningOrStopping)
	}

	var task *t.Task[T1, T2]
	task, err = tq.getTask(taskId)
	if err != nil {
		return err
	}

	err = task.SaveResponse(response)
	if err != nil {
		return err
	}

	return nil
}
func (tq *TaskQueue[T1, T2]) CompleteTask(taskId string) (task *t.Task[T1, T2], err error) {
	tq.controls.lock.Lock()
	defer tq.controls.lock.Unlock()

	isRunningOrStopping := tq.controls.ssc.IsRunning() || tq.controls.ssc.IsStopping()
	if !isRunningOrStopping {
		return nil, errors.ErrorFromText(ce.Err_NotRunningOrStopping)
	}

	task, err = tq.getTask(taskId)
	if err != nil {
		return nil, err
	}

	err = task.Complete()
	if err != nil {
		return nil, err
	}

	return task, nil
}
func (tq *TaskQueue[T1, T2]) GetTaskStatus(taskId string) (taskStatus ts.TaskStatus, err error) {
	tq.controls.lock.Lock()
	defer tq.controls.lock.Unlock()

	isRunningOrStopping := tq.controls.ssc.IsRunning() || tq.controls.ssc.IsStopping()
	if !isRunningOrStopping {
		return ts.TaskStatus_Unknown, errors.ErrorFromText(ce.Err_NotRunningOrStopping)
	}

	var task *t.Task[T1, T2]
	task, err = tq.getTask(taskId)
	if err != nil {
		return ts.TaskStatus_Unknown, err
	}

	return task.GetStatus(), nil
}
func (tq *TaskQueue[T1, T2]) GetQueueSize() (qSize int, err error) {
	tq.controls.lock.Lock()
	defer tq.controls.lock.Unlock()

	isRunningOrStopping := tq.controls.ssc.IsRunning() || tq.controls.ssc.IsStopping()
	if !isRunningOrStopping {
		return 0, errors.ErrorFromText(ce.Err_NotRunningOrStopping)
	}

	return tq.inQueue.Size(), nil
}
