package to

import (
	"time"

	tp "github.com/vault-thirteen/auxie/task-queue/TimeStamp"
)

type TimeOf struct {
	creation   tp.TimeStamp
	workStart  tp.TimeStamp
	workEnd    tp.TimeStamp
	completion tp.TimeStamp
}

func NewTimeOf() (timeOf TimeOf) {
	timeOf = TimeOf{
		creation:   time.Now().UTC().Unix(),
		workStart:  tp.ZeroTimeStamp,
		workEnd:    tp.ZeroTimeStamp,
		completion: tp.ZeroTimeStamp,
	}

	return timeOf
}
func (to *TimeOf) Creation() tp.TimeStamp {
	return to.creation
}
func (to *TimeOf) WorkStart() tp.TimeStamp {
	return to.workStart
}
func (to *TimeOf) WorkEnd() tp.TimeStamp {
	return to.workEnd
}
func (to *TimeOf) Completion() tp.TimeStamp {
	return to.completion
}
func (to *TimeOf) SetWorkStart() {
	to.workStart = time.Now().UTC().Unix()
}
func (to *TimeOf) SetWorkEnd() {
	to.workEnd = time.Now().UTC().Unix()
}
func (to *TimeOf) SetCompletion() {
	to.completion = time.Now().UTC().Unix()
}
