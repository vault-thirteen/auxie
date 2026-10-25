package ts

import (
	"github.com/vault-thirteen/auxie/errors"
	ce "github.com/vault-thirteen/auxie/task-queue/common/error"
)

type TaskStatus byte

const (
	// TaskStatus_Unknown is an unknown TaskStatus.
	// Unknown TaskStatus is an un-initialised TaskStatus.
	TaskStatus_Unknown = TaskStatus(0)

	// TaskStatus_Created is a TaskStatus of a newly created Task.
	TaskStatus_Created = TaskStatus(1)

	// TaskStatus_Queued is a TaskStatus of a Task which is sitting inside a queue.
	TaskStatus_Queued = TaskStatus(2)

	// TaskStatus_InWork is a TaskStatus of a Task which is in work.
	TaskStatus_InWork = TaskStatus(3)

	// TaskStatus_Processed is a processed TaskStatus.
	// Processed TaskStatus means that Task's work has been done and the Task's
	// result is ready to be collected by the client.
	TaskStatus_Processed = TaskStatus(4)

	// TaskStatus_Completed is a completed TaskStatus.
	// Completed TaskStatus means that at least one client has taken the result of
	// the work, and the Task can be deleted from registry.
	TaskStatus_Completed = TaskStatus(5)
)

func NewTaskStatus() TaskStatus {
	return TaskStatus_Created
}

func (ts *TaskStatus) Change(newValue TaskStatus) (err error) {
	switch newValue {
	case TaskStatus_Unknown:
		return errors.ErrorFromText(ce.Err_UnknownStatusCanNotBeSet)

	case TaskStatus_Created:
		return errors.ErrorFromText(ce.Err_OnlyNewTasksCanHaveStatusCreated)

	case TaskStatus_Queued:
		if *ts != TaskStatus_Created {
			return errors.ErrorFromText(ce.Err_OnlyCreatedTasksCanBePutIntoQueue)
		}

	case TaskStatus_InWork:
		if *ts != TaskStatus_Queued {
			return errors.ErrorFromText(ce.Err_OnlyQueuedTasksCanBeTakenIntoWork)
		}

	case TaskStatus_Processed:
		if *ts != TaskStatus_InWork {
			return errors.ErrorFromText(ce.Err_OnlyTasksInWorkCanBeProcessed)
		}

	case TaskStatus_Completed:
		if *ts != TaskStatus_Processed {
			return errors.ErrorFromText(ce.Err_OnlyProcessedTasksCanBeCompleted)
		}

	default:
		return errors.ErrorFromFormattedText(ce.Errf_UnsupportedStatus, newValue)
	}

	*ts = newValue
	return nil
}
