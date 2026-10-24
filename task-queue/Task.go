package tq

import "time"

type Task[T1 any, T2 any] struct {
	Id        TaskId
	Request   *T1
	Response  *T2
	Status    TaskStatus
	Created   TimeStamp
	WorkStart TimeStamp
	Processed TimeStamp
	Completed TimeStamp
}

func NewTask[T1 any, T2 any](request *T1, taskId TaskId) (task *Task[T1, T2]) {
	task = &Task[T1, T2]{
		Id:        taskId,
		Request:   request,
		Response:  nil,
		Status:    TaskStatus_Created,
		Created:   time.Now().UTC().Unix(),
		WorkStart: 0,
		Processed: 0,
		Completed: 0,
	}

	return task
}
