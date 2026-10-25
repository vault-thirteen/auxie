package tq

import (
	"sync"
	"sync/atomic"

	stct "github.com/vault-thirteen/auxie/Start-Controller"
)

const (
	ErrorChannelSize = 8
)

type TaskQueueControls struct {
	ssc                  *stct.StartController
	lock                 *sync.Mutex
	sizeLimit            int
	receiveNewTasks      *atomic.Bool
	cleanerMustWork      *atomic.Bool
	cleanerWg            *sync.WaitGroup
	errorCatcherMustWork *atomic.Bool
	errorCatcherWg       *sync.WaitGroup
	errors               *chan error
}

func NewTaskQueueControls(sizeLimit int) (tqc *TaskQueueControls) {
	tqc = &TaskQueueControls{
		ssc:                  stct.NewStartController(),
		lock:                 new(sync.Mutex),
		sizeLimit:            sizeLimit,
		receiveNewTasks:      new(atomic.Bool),
		cleanerMustWork:      new(atomic.Bool),
		cleanerWg:            new(sync.WaitGroup),
		errorCatcherMustWork: new(atomic.Bool),
		errorCatcherWg:       new(sync.WaitGroup),
	}

	errChan := make(chan error, ErrorChannelSize)
	tqc.errors = &errChan

	return tqc
}
