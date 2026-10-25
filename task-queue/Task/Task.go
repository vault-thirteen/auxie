package t

import (
	"github.com/vault-thirteen/auxie/errors"
	"github.com/vault-thirteen/auxie/task-queue/TaskId"
	ts "github.com/vault-thirteen/auxie/task-queue/TaskStatus"
	to "github.com/vault-thirteen/auxie/task-queue/TimeOf"
	ce "github.com/vault-thirteen/auxie/task-queue/common/error"
)

type Task[T1 any, T2 any] struct {
	id       ti.TaskId
	request  *T1
	response *T2
	status   ts.TaskStatus
	timeOf   to.TimeOf
}

func NewTask[T1 any, T2 any](id ti.TaskId, request *T1) (task *Task[T1, T2]) {
	task = &Task[T1, T2]{
		id:       id,
		request:  request,
		response: nil,
		status:   ts.NewTaskStatus(),
		timeOf:   to.NewTimeOf(),
	}

	return task
}

func (t *Task[T1, T2]) ChangeStatus(newStatus ts.TaskStatus) (err error) {
	return t.status.Change(newStatus)
}
func (t *Task[T1, T2]) Destroy() (err error) {
	if t.status != ts.TaskStatus_Completed {
		return errors.ErrorFromText(ce.Err_OnlyCompletedTasksCanBeDestroyed)
	}

	t.id = ti.ZeroTaskId
	t.request = nil
	t.response = nil

	return nil
}
func (t *Task[T1, T2]) GetStatus() ts.TaskStatus {
	return t.status
}
func (t *Task[T1, T2]) GetTimeOf() to.TimeOf {
	return t.timeOf
}
func (t *Task[T1, T2]) GetId() ti.TaskId {
	return t.id
}
func (t *Task[T1, T2]) SetTimeOfWorkStart() {
	t.timeOf.SetWorkStart()
}
func (t *Task[T1, T2]) SaveResponse(response *T2) (err error) {
	t.response = response

	err = t.ChangeStatus(ts.TaskStatus_Processed)
	if err != nil {
		return err
	}

	t.timeOf.SetWorkEnd()

	return nil
}
func (t *Task[T1, T2]) Complete() (err error) {
	err = t.ChangeStatus(ts.TaskStatus_Completed)
	if err != nil {
		return err
	}

	t.timeOf.SetCompletion()

	return nil
}
