package muc

import (
	"runtime/debug"
)

// freeMemory tries to free the unused Memory.
func (muc *MemoryUsageController) freeMemory() {
	muc.freeLock.Lock()
	defer muc.freeLock.Unlock()

	debug.FreeOSMemory()
}
