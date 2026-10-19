package muc

import (
	"sync"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	MB = 1 * 1000 * 1000 // 1 M.
)

// Errors.
const (
	Err_UsageLimitError             = "usage limit error"
	Err_UsageRatioThresholdError    = "usage ratio threshold error"
	Err_UsageIsBeyondLimits         = "memory usage is beyond the limits"
	Err_MemoryUsageCriterionInvalid = "memory usage criterion is not valid"
	Err_MemoryUsageCriterionUnknown = "memory usage criterion is unknown"
)

const (
	Msgf_ReportVerbose = "memory usage has been minimized: %d -> %d MB."
)

// MemoryUsageController is the memory usage controller.
type MemoryUsageController struct {
	memoryUsageCriterion            MemoryUsageCriterion
	memUsageLimitMb                 uint
	memoryUsedToLimitRatioThreshold float64
	verboseMode                     bool

	// Locks.
	usageLock sync.Mutex
	freeLock  sync.Mutex
}

func NewMemoryUsageController(
	memUsageLimitMb uint,
	memoryUsedToLimitRatioThreshold float64,
	memoryUsageCriterion byte,
	verboseMode bool,
) (muc *MemoryUsageController, err error) {
	if memUsageLimitMb == 0 {
		return nil, errors.ErrorFromText(Err_UsageLimitError)
	}
	if memoryUsedToLimitRatioThreshold <= 0 {
		return nil, errors.ErrorFromText(Err_UsageRatioThresholdError)
	}

	muc = new(MemoryUsageController)

	// Set the memory usage criterion.
	muc.memoryUsageCriterion = NewMemoryUsageCriterion(memoryUsageCriterion)
	if !muc.memoryUsageCriterion.IsValid() {
		return nil, errors.ErrorFromText(Err_MemoryUsageCriterionInvalid)
	}

	// Set other fields.
	muc.memUsageLimitMb = memUsageLimitMb
	muc.memoryUsedToLimitRatioThreshold = memoryUsedToLimitRatioThreshold
	muc.verboseMode = verboseMode

	return muc, nil
}
