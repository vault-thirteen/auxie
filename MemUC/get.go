package muc

import (
	"errors"
	"runtime"

	"github.com/prometheus/procfs"
)

// GetMemoryUsage gets the memory usage amount. Result is set in MB.
func (muc *MemoryUsageController) GetMemoryUsage() (usageMb uint, err error) {
	return muc.getMemoryUsage()
}

// getMemoryUsage gets the memory usage amount. Result is set in MB.
func (muc *MemoryUsageController) getMemoryUsage() (usageMb uint, err error) {
	muc.usageLock.Lock()
	defer muc.usageLock.Unlock()

	switch muc.memoryUsageCriterion {

	case MemoryUsageCriterionWorkingSet:
		return muc.getMemoryUsageWorkingSet()

	case MemoryUsageCriterionResidentMemory:
		return muc.getMemoryUsageResident()

	default:
		return usageMb, errors.New(ErrMemoryUsageCriterionUnknown)
	}
}

// getMemoryUsageWorkingSet returns the memory usage as the "Working Set".
// This parameter is often used in Windows O.S.
func (muc *MemoryUsageController) getMemoryUsageWorkingSet() (usageMb uint, err error) {
	// Update the memory usage statistics using the built-in Go mechanism.
	var memoryUsageStatistics *runtime.MemStats
	memoryUsageStatistics = new(runtime.MemStats)
	runtime.ReadMemStats(memoryUsageStatistics)

	// Calculate the Working Set.
	return uint(memoryUsageStatistics.HeapInuse+
		memoryUsageStatistics.StackInuse+
		memoryUsageStatistics.MSpanInuse+
		memoryUsageStatistics.MCacheInuse+
		memoryUsageStatistics.BuckHashSys+
		memoryUsageStatistics.GCSys+
		memoryUsageStatistics.OtherSys) / MB, nil
}

// getMemoryUsageResident returns the memory usage as the "Resident Memory".
// This parameter is often used in Linux O.S.
func (muc *MemoryUsageController) getMemoryUsageResident() (usageMb uint, err error) {
	// Get the memory usage statistics using the external library which makes
	// a special call to the operating system. Such a call is not supported in
	// operating systems of the Windows family.
	var osProcess procfs.Proc
	osProcess, err = procfs.Self()
	if err != nil {
		return usageMb, err
	}

	var processStatistics procfs.ProcStat
	processStatistics, err = osProcess.Stat()
	if err != nil {
		return usageMb, err
	}

	// Calculate the resident Memory.
	return uint(processStatistics.ResidentMemory()) / MB, nil
}
