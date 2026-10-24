package tq

type TaskStatus byte

const (
	TaskStatus_Unknown = TaskStatus(0)
	TaskStatus_Created = TaskStatus(1)
	TaskStatus_Queued  = TaskStatus(2)
	TaskStatus_InWork  = TaskStatus(3)

	// TaskStatus_Processed is a processed status.
	// Processed status means that task's work has been done and the task's
	// result is ready to be collected by the client.
	TaskStatus_Processed = TaskStatus(4)

	// TaskStatus_Completed is a completed status.
	// Completed status means that at least one client has taken the result of
	// the work, and the task can be deleted from registry.
	TaskStatus_Completed = TaskStatus(5)
)
