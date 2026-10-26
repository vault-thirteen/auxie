package tqs

import (
	"github.com/vault-thirteen/auxie/errors"
	tqsp "github.com/vault-thirteen/auxie/task-queue/TaskQueueSettingsPublic"
	ce "github.com/vault-thirteen/auxie/task-queue/common/error"
)

const (
	// Default_CompletedTaskTtlSec is the default time period after which a
	// completed Task may be deleted from registry.
	Default_CompletedTaskTtlSec = 1 * 60

	// Default_CleanerSleepTimeSec is the default delay between cleans made by
	// cleaner.
	Default_CleanerSleepTimeSec = 60

	Default_ErrorCatcherSleepTimeSec                 = 1
	Default_SleepTimeWaitingForZeroQueueSec          = 1
	Default_SleepTimeWaitingForZeroTasksInWorkSec    = 1
	Default_SleepTimeWaitingForZeroProcessedTasksSec = 1
)

// TaskQueueSettings is an object for settings of a task queue.
// In order to make all the settings read-only, we separate this type in its
// own package and make all the fields of the structure unexported.
// Unfortunately, Go language does not support private visibility which is
// available in such mature languages as C++, C# and Java.
type TaskQueueSettings struct {
	inputQueueSizeLimit int
	errorChannelSize    int
	errorCatcherFn      func(err error)

	// Timings.
	completedTaskTtlSec                      int
	cleanerSleepTimeSec                      int
	errorCatcherSleepTimeSec                 int
	sleepTimeWaitingForZeroQueueSec          int
	sleepTimeWaitingForZeroTasksInWorkSec    int
	sleepTimeWaitingForZeroProcessedTasksSec int
}

func NewTaskQueueSettings(source *tqsp.TaskQueueSettingsPublic) (tqs *TaskQueueSettings, err error) {
	if source == nil {
		return nil, errors.ErrorFromText(ce.Err_NullPointer)
	}

	tqs = &TaskQueueSettings{
		inputQueueSizeLimit: source.InputQueueSizeLimit,
		errorChannelSize:    source.ErrorChannelSize,
		errorCatcherFn:      source.ErrorCatcherFn,

		// Timings.
		completedTaskTtlSec:                      source.CompletedTaskTtlSec,
		cleanerSleepTimeSec:                      source.CleanerSleepTimeSec,
		errorCatcherSleepTimeSec:                 source.ErrorCatcherSleepTimeSec,
		sleepTimeWaitingForZeroQueueSec:          source.SleepTimeWaitingForZeroQueueSec,
		sleepTimeWaitingForZeroTasksInWorkSec:    source.SleepTimeWaitingForZeroTasksInWorkSec,
		sleepTimeWaitingForZeroProcessedTasksSec: source.SleepTimeWaitingForZeroProcessedTasksSec,
	}

	err = tqs.Validate()
	if err != nil {
		return nil, err
	}

	return tqs, nil
}
func (tqs *TaskQueueSettings) Validate() (err error) {
	if tqs.inputQueueSizeLimit <= 0 {
		return errors.ErrorFromText(ce.Err_InputQueueSizeLimit)
	}
	if tqs.errorChannelSize <= 0 {
		return errors.ErrorFromText(ce.Err_ErrorChannelSize)
	}

	// Default timings.
	if tqs.completedTaskTtlSec <= 0 {
		tqs.completedTaskTtlSec = Default_CompletedTaskTtlSec
	}
	if tqs.cleanerSleepTimeSec <= 0 {
		tqs.cleanerSleepTimeSec = Default_CleanerSleepTimeSec
	}
	if tqs.errorCatcherSleepTimeSec <= 0 {
		tqs.errorCatcherSleepTimeSec = Default_ErrorCatcherSleepTimeSec
	}
	if tqs.sleepTimeWaitingForZeroQueueSec <= 0 {
		tqs.sleepTimeWaitingForZeroQueueSec = Default_SleepTimeWaitingForZeroQueueSec
	}
	if tqs.sleepTimeWaitingForZeroTasksInWorkSec <= 0 {
		tqs.sleepTimeWaitingForZeroTasksInWorkSec = Default_SleepTimeWaitingForZeroTasksInWorkSec
	}
	if tqs.sleepTimeWaitingForZeroProcessedTasksSec <= 0 {
		tqs.sleepTimeWaitingForZeroProcessedTasksSec = Default_SleepTimeWaitingForZeroProcessedTasksSec
	}

	return nil
}
func (tqs *TaskQueueSettings) InputQueueSizeLimit() int {
	return tqs.inputQueueSizeLimit
}
func (tqs *TaskQueueSettings) ErrorChannelSize() int {
	return tqs.errorChannelSize
}
func (tqs *TaskQueueSettings) ErrorCatcherFn() func(err error) {
	return tqs.errorCatcherFn
}
func (tqs *TaskQueueSettings) CompletedTaskTtlSec() int {
	return tqs.completedTaskTtlSec
}
func (tqs *TaskQueueSettings) CleanerSleepTimeSec() int {
	return tqs.cleanerSleepTimeSec
}
func (tqs *TaskQueueSettings) ErrorCatcherSleepTimeSec() int {
	return tqs.errorCatcherSleepTimeSec
}
func (tqs *TaskQueueSettings) SleepTimeWaitingForZeroQueueSec() int {
	return tqs.sleepTimeWaitingForZeroQueueSec
}
func (tqs *TaskQueueSettings) SleepTimeWaitingForZeroTasksInWorkSec() int {
	return tqs.sleepTimeWaitingForZeroTasksInWorkSec
}
func (tqs *TaskQueueSettings) SleepTimeWaitingForZeroProcessedTasksSec() int {
	return tqs.sleepTimeWaitingForZeroProcessedTasksSec
}
