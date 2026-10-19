package muc

import (
	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"
)

// CheckFreeMemory checks free memory.
func (muc *MemoryUsageController) CheckFreeMemory() (err error) {
	// Get the memory usage.
	var memUsageAMb uint
	memUsageAMb, err = muc.getMemoryUsage()
	if err != nil {
		return err
	}

	// Perform a check.
	if float64(muc.memUsageLimitMb)/float64(memUsageAMb) < muc.memoryUsedToLimitRatioThreshold {
		// Memory usage is below the threshold.
		return nil
	}

	// Try to free some memory.
	muc.freeMemory()

	// Get the memory usage again.
	var memUsageBMb uint
	memUsageBMb, err = muc.getMemoryUsage()
	if err != nil {
		return err
	}

	// Verbose memory usage change report.
	if muc.verboseMode {
		cio.LogMessage(cio.FormatMessage(Msgf_ReportVerbose, memUsageAMb, memUsageBMb))
	}

	// Have we succeeded in freeing the memory ?
	if memUsageBMb < muc.memUsageLimitMb {
		return nil
	}

	// Memory usage breaks the limit.
	return errors.ErrorFromText(Err_UsageIsBeyondLimits)
}
