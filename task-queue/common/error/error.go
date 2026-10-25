package ce

const (
	Err_QueueIsFull                       = "queue is full"
	Err_NotReceivingNewTasks              = "not receiving new tasks"
	Err_UnknownStatusCanNotBeSet          = "unknown TaskStatus can not be set"
	Err_OnlyNewTasksCanHaveStatusCreated  = "only new tasks can have 'created' TaskStatus"
	Err_OnlyCreatedTasksCanBePutIntoQueue = "only created tasks can be put into queue"
	Err_OnlyQueuedTasksCanBeTakenIntoWork = "only queued tasks can be taken into work"
	Err_OnlyTasksInWorkCanBeProcessed     = "only tasks in work can be processed"
	Err_OnlyProcessedTasksCanBeCompleted  = "only processed tasks can be completed"
	Err_OnlyCompletedTasksCanBeDestroyed  = "only completed tasks can be destroyed"
	Err_NotRunningOrStopping              = "not running or stopping"

	Errf_UnsupportedStatus = "unsupported TaskStatus: %v"
	Errf_TaskIsNotFound    = "task is not found, id: %v"
	Errf_TaskAlreadyExists = "task already exists, id: %v"
)
