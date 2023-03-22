package tm

import (
	"math"
	"testing"
	"time"

	bt "github.com/vault-thirteen/auxie/BasicTypes"
	"github.com/vault-thirteen/tester"
)

func Test_New(t *testing.T) {
	var aTester = tester.New(t)
	waitForApproximateSecondStart()

	// Test.
	timeNow := time.Now()
	tmd := New()
	aTester.MustBeEqual(
		tmd.creationTime-timeNow.Unix(),
		bt.UnixTimeStamp(0),
	)
	aTester.MustBeEqual(
		tmd.updateTime,
		bt.UnixTimeStamp(0),
	)
}

func Test_Update(t *testing.T) {
	var aTester = tester.New(t)
	tmd := new(TimeMetadata)
	waitForApproximateSecondStart()

	// Test.
	timeNow := time.Now()
	tmd.Update()
	aTester.MustBeEqual(
		tmd.creationTime,
		bt.UnixTimeStamp(0),
	)
	aTester.MustBeEqual(
		tmd.updateTime-timeNow.Unix(),
		bt.UnixTimeStamp(0),
	)
}

func Test_GetCreationTime(t *testing.T) {
	var aTester = tester.New(t)

	// Test.
	tmd := new(TimeMetadata)
	creationTimeExpected := bt.UnixTimeStamp(math.MaxInt64)
	tmd.creationTime = creationTimeExpected
	creationTime := tmd.GetCreationTime()
	aTester.MustBeEqual(creationTime, creationTimeExpected)
}

func Test_GetUpdateTime(t *testing.T) {
	var aTester = tester.New(t)

	// Test.
	tmd := new(TimeMetadata)
	updateTimeExpected := bt.UnixTimeStamp(math.MaxInt64)
	tmd.updateTime = updateTimeExpected
	updateTime := tmd.GetUpdateTime()
	aTester.MustBeEqual(updateTime, updateTimeExpected)
}

// waitForApproximateSecondStart waits for the start of a new second.
func waitForApproximateSecondStart() {
	// Attention !
	// This parameter may depend on the system !
	const TimeSecondNanosecondMaxThreshold = 1000 * 1000

	var timeNow time.Time
	var loop = true
	for loop {
		timeNow = time.Now()
		if timeNow.Nanosecond() < TimeSecondNanosecondMaxThreshold {
			loop = false
		}
	}
}
