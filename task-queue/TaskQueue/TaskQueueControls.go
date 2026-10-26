package tq

import (
	"sync"
	"sync/atomic"

	stct "github.com/vault-thirteen/auxie/Start-Controller"
	tqs "github.com/vault-thirteen/auxie/task-queue/TaskQueueSettings"
)

type TaskQueueControls struct {
	ssc                  *stct.StartController
	lock                 *sync.Mutex
	receiveNewTasks      *atomic.Bool
	cleanerMustWork      *atomic.Bool
	cleanerWg            *sync.WaitGroup
	errorCatcherMustWork *atomic.Bool
	errorCatcherWg       *sync.WaitGroup
	errors               *chan error
	errorCatcherFn       func(err error)
}

func NewTaskQueueControls(settings *tqs.TaskQueueSettings) (tqc *TaskQueueControls) {
	tqc = &TaskQueueControls{
		ssc:                  stct.NewStartController(),
		lock:                 new(sync.Mutex),
		receiveNewTasks:      new(atomic.Bool),
		cleanerMustWork:      new(atomic.Bool),
		cleanerWg:            new(sync.WaitGroup),
		errorCatcherMustWork: new(atomic.Bool),
		errorCatcherWg:       new(sync.WaitGroup),
		errorCatcherFn:       settings.ErrorCatcherFn(),
	}

	errChan := make(chan error, settings.ErrorChannelSize())
	tqc.errors = &errChan

	return tqc
}
