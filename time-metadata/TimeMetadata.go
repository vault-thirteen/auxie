// TimeMetadata.go.

package tm

import (
	"time"
)

type TimeMetadata struct {

	// UNIX Timestamp of Creation Time.
	creationTime int64

	// UNIX Timestamp of Update Time.
	updateTime int64
}

// New creates a new Time Meta-Data with 'Creation Time' Field filled with the
// current Time.
func New() *TimeMetadata {
	return &TimeMetadata{
		creationTime: time.Now().Unix(),
	}
}

// Update updates the 'Update Time' Field of the Time Meta-Data Object using
// the current Time.
func (tmd *TimeMetadata) Update() {
	tmd.updateTime = time.Now().Unix()
}

// GetCreationTime returns the 'Creation Time' Field of the Time Meta-Data
// Object.
func (tmd TimeMetadata) GetCreationTime() int64 {
	return tmd.creationTime
}

// GetUpdateTime returns the 'Update Time' Field of the Time Meta-Data Object.
func (tmd TimeMetadata) GetUpdateTime() int64 {
	return tmd.updateTime
}
