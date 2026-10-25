package stct

import (
	"sync"
	"sync/atomic"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	Err_Starting = "starting procedure is in progress"
	Err_Running  = "the process is already running"
	Err_Stopping = "stopping procedure is in progress"

	Err_NotStarting = "starting procedure is in not progress"
	Err_NotRunning  = "the process is not running"
	Err_NotStopping = "stopping procedure is not in progress"
)

type StartController struct {
	isStarting     *atomic.Bool
	isRunning      *atomic.Bool
	isStopping     *atomic.Bool
	startStopGuard *sync.Mutex
}

func NewStartController() (sc *StartController) {
	sc = &StartController{
		isStarting:     new(atomic.Bool),
		isRunning:      new(atomic.Bool),
		isStopping:     new(atomic.Bool),
		startStopGuard: new(sync.Mutex),
	}

	sc.isStarting.Store(false)
	sc.isRunning.Store(false)
	sc.isStopping.Store(false)

	return sc
}
func (sc *StartController) BeginStart() (err error) {
	sc.startStopGuard.Lock()
	defer sc.startStopGuard.Unlock()

	if sc.isStarting.Load() {
		return errors.ErrorFromText(Err_Starting)
	}
	if sc.isRunning.Load() {
		return errors.ErrorFromText(Err_Running)
	}
	if sc.isStopping.Load() {
		return errors.ErrorFromText(Err_Stopping)
	}

	sc.isStarting.Store(true)

	return nil
}
func (sc *StartController) CompleteStart() (err error) {
	sc.startStopGuard.Lock()
	defer sc.startStopGuard.Unlock()

	if !sc.isStarting.Load() {
		return errors.ErrorFromText(Err_NotStarting)
	}
	if sc.isRunning.Load() {
		return errors.ErrorFromText(Err_Starting)
	}
	if sc.isStopping.Load() {
		return errors.ErrorFromText(Err_Stopping)
	}

	sc.isStarting.Store(false)
	sc.isRunning.Store(true)

	return nil
}
func (sc *StartController) BeginStop() (err error) {
	sc.startStopGuard.Lock()
	defer sc.startStopGuard.Unlock()

	if sc.isStarting.Load() {
		return errors.ErrorFromText(Err_Starting)
	}
	if !sc.isRunning.Load() {
		return errors.ErrorFromText(Err_NotRunning)
	}
	if sc.isStopping.Load() {
		return errors.ErrorFromText(Err_Stopping)
	}

	sc.isRunning.Store(false)
	sc.isStopping.Store(true)

	return nil
}
func (sc *StartController) CompleteStop() (err error) {
	sc.startStopGuard.Lock()
	defer sc.startStopGuard.Unlock()

	if sc.isStarting.Load() {
		return errors.ErrorFromText(Err_Starting)
	}
	if sc.isRunning.Load() {
		return errors.ErrorFromText(Err_Running)
	}
	if !sc.isStopping.Load() {
		return errors.ErrorFromText(Err_NotStopping)
	}

	sc.isStopping.Store(false)

	return nil
}
