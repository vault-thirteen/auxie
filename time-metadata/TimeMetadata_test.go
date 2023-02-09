// TimeMetadata_test.go.

package tm

import (
	"math"
	"testing"
	"time"

	"github.com/vault-thirteen/tester"
)

func Test_New(t *testing.T) {

	var aTester *tester.Test
	var tmd *TimeMetadata
	var timeNow time.Time

	aTester = tester.New(t)
	waitForApproximateSecondStart()

	// Test.
	timeNow = time.Now()
	tmd = New()
	aTester.MustBeEqual(
		tmd.creationTime-timeNow.Unix(),
		int64(0),
	)
	aTester.MustBeEqual(
		tmd.updateTime,
		int64(0),
	)
}

func Test_Update(t *testing.T) {

	var aTester *tester.Test
	var tmd *TimeMetadata
	var timeNow time.Time

	aTester = tester.New(t)
	tmd = new(TimeMetadata)
	waitForApproximateSecondStart()

	// Test.
	timeNow = time.Now()
	tmd.Update()
	aTester.MustBeEqual(
		tmd.creationTime,
		int64(0),
	)
	aTester.MustBeEqual(
		tmd.updateTime-timeNow.Unix(),
		int64(0),
	)
}

func Test_GetCreationTime(t *testing.T) {

	var creationTime int64
	var creationTimeExpected int64
	var aTester *tester.Test
	var tmd *TimeMetadata

	aTester = tester.New(t)

	// Test.
	tmd = new(TimeMetadata)
	creationTimeExpected = math.MaxInt64
	tmd.creationTime = creationTimeExpected
	creationTime = tmd.GetCreationTime()
	aTester.MustBeEqual(creationTime, creationTimeExpected)
}

func Test_GetUpdateTime(t *testing.T) {

	var aTester *tester.Test
	var tmd *TimeMetadata
	var updateTime int64
	var updateTimeExpected int64

	aTester = tester.New(t)

	// Test.
	tmd = new(TimeMetadata)
	updateTimeExpected = math.MaxInt64
	tmd.updateTime = updateTimeExpected
	updateTime = tmd.GetUpdateTime()
	aTester.MustBeEqual(updateTime, updateTimeExpected)
}

func waitForApproximateSecondStart() {

	// Attention!
	// This Parameter may depend on the System!
	const TimeSecondNanosecondMaxThreshold = 1000 * 1000

	var loop bool
	var timeNow time.Time

	loop = true
	for loop {
		timeNow = time.Now()
		if timeNow.Nanosecond() < TimeSecondNanosecondMaxThreshold {
			loop = false
		}
	}
}
