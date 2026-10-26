package tqsp

// TaskQueueSettingsPublic is a public version of TaskQueueSettings type.
// This type is used for providing settings from external sources.
type TaskQueueSettingsPublic struct {
	InputQueueSizeLimit int             `json:"inputQueueSizeLimit"`
	ErrorChannelSize    int             `json:"errorChannelSize"`
	ErrorCatcherFn      func(err error) `json:"-"`

	// Timings.
	CompletedTaskTtlSec                      int `json:"completedTaskTtlSec"`
	CleanerSleepTimeSec                      int `json:"cleanerSleepTimeSec"`
	ErrorCatcherSleepTimeSec                 int `json:"errorCatcherSleepTimeSec"`
	SleepTimeWaitingForZeroQueueSec          int `json:"sleepTimeWaitingForZeroQueueSec"`
	SleepTimeWaitingForZeroTasksInWorkSec    int `json:"sleepTimeWaitingForZeroTasksInWorkSec"`
	SleepTimeWaitingForZeroProcessedTasksSec int `json:"sleepTimeWaitingForZeroProcessedTasksSec"`
}
