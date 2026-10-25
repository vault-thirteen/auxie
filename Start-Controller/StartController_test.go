package stct

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_BeginStart(t *testing.T) {
	aTest := tester.New(t)
	var sc *StartController
	var err error

	// Test #1. Normal start.
	sc = NewStartController()
	err = sc.BeginStart()
	aTest.MustBeNoError(err)

	// Test #2. Starting an already starting process.
	sc = NewStartController()
	sc.isStarting.Store(true)
	err = sc.BeginStart()
	aTest.MustBeAnError(err)

	// Test #3. Starting an already running process.
	sc = NewStartController()
	sc.isRunning.Store(true)
	err = sc.BeginStart()
	aTest.MustBeAnError(err)

	// Test #4. Starting an already stopping process.
	sc = NewStartController()
	sc.isStopping.Store(true)
	err = sc.BeginStart()
	aTest.MustBeAnError(err)
}
func Test_CompleteStart(t *testing.T) {
	aTest := tester.New(t)
	var sc *StartController
	var err error

	// Test #1. Completing start of a stopped process.
	sc = NewStartController()
	err = sc.CompleteStart()
	aTest.MustBeAnError(err)

	// Test #2. Normal start completion.
	sc = NewStartController()
	sc.isStarting.Store(true)
	err = sc.CompleteStart()
	aTest.MustBeNoError(err)

	// Test #3. Completing start of a running process.
	sc = NewStartController()
	sc.isRunning.Store(true)
	err = sc.CompleteStart()
	aTest.MustBeAnError(err)

	// Test #4. Completing start of a stopping process.
	sc = NewStartController()
	sc.isStopping.Store(true)
	err = sc.CompleteStart()
	aTest.MustBeAnError(err)
}
func Test_BeginStop(t *testing.T) {
	aTest := tester.New(t)
	var sc *StartController
	var err error

	// Test #1. Stopping a stopped process.
	sc = NewStartController()
	err = sc.BeginStop()
	aTest.MustBeAnError(err)

	// Test #2. Stopping a starting process.
	sc = NewStartController()
	sc.isStarting.Store(true)
	err = sc.BeginStop()
	aTest.MustBeAnError(err)

	// Test #3. Normal stopping.
	sc = NewStartController()
	sc.isRunning.Store(true)
	err = sc.BeginStop()
	aTest.MustBeNoError(err)

	// Test #4. Stopping a stopping process.
	sc = NewStartController()
	sc.isStopping.Store(true)
	err = sc.BeginStop()
	aTest.MustBeAnError(err)
}
func Test_CompleteStop(t *testing.T) {
	aTest := tester.New(t)
	var sc *StartController
	var err error

	// Test #1. Completing stop of a stopped process.
	sc = NewStartController()
	err = sc.CompleteStop()
	aTest.MustBeAnError(err)

	// Test #2. Completing stop of a starting process.
	sc = NewStartController()
	sc.isStarting.Store(true)
	err = sc.CompleteStop()
	aTest.MustBeAnError(err)

	// Test #3. Completing stop of a running process.
	sc = NewStartController()
	sc.isRunning.Store(true)
	err = sc.CompleteStop()
	aTest.MustBeAnError(err)

	// Test #4. Normal stop completion.
	sc = NewStartController()
	sc.isStopping.Store(true)
	err = sc.CompleteStop()
	aTest.MustBeNoError(err)
}
func Test_IsStarting(t *testing.T) {
	aTest := tester.New(t)
	var sc *StartController
	sc = NewStartController()

	// 1.
	sc.isStarting.Store(true)
	aTest.MustBeEqual(sc.IsStarting(), true)

	// 2.
	sc.isStarting.Store(false)
	aTest.MustBeEqual(sc.IsStarting(), false)
}
func Test_IsRunning(t *testing.T) {
	aTest := tester.New(t)
	var sc *StartController
	sc = NewStartController()

	// 1.
	sc.isRunning.Store(true)
	aTest.MustBeEqual(sc.IsRunning(), true)

	// 2.
	sc.isRunning.Store(false)
	aTest.MustBeEqual(sc.IsRunning(), false)
}
func Test_IsStopping(t *testing.T) {
	aTest := tester.New(t)
	var sc *StartController
	sc = NewStartController()

	// 1.
	sc.isStopping.Store(true)
	aTest.MustBeEqual(sc.IsStopping(), true)

	// 2.
	sc.isStopping.Store(false)
	aTest.MustBeEqual(sc.IsStopping(), false)
}
