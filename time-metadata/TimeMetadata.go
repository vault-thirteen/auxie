package tm

import (
	"time"

	bt "github.com/vault-thirteen/auxie/BasicTypes"
)

type TimeMetadata struct {

	// UNIX timestamp of creation time.
	creationTime bt.UnixTimeStamp

	// UNIX timestamp of update time.
	updateTime bt.UnixTimeStamp
}

// New creates a new Time Meta-Data with 'Creation Time' field filled with the
// current Time.
func New() *TimeMetadata {
	return &TimeMetadata{
		creationTime: time.Now().Unix(),
	}
}

// Update updates the 'Update Time' field of the Time Meta-Data object using
// the current time.
func (tmd *TimeMetadata) Update() {
	tmd.updateTime = time.Now().Unix()
}

// GetCreationTime returns the 'Creation Time' field of the Time Meta-Data
// object.
func (tmd *TimeMetadata) GetCreationTime() bt.UnixTimeStamp {
	return tmd.creationTime
}

// GetUpdateTime returns the 'Update Time' field of the Time Meta-Data object.
func (tmd *TimeMetadata) GetUpdateTime() bt.UnixTimeStamp {
	return tmd.updateTime
}
