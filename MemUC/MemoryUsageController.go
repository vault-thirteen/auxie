package muc

import (
	"errors"
	"sync"
)

const (
	MB = 1 * 1000 * 1000 // 1 M.
)

// Errors.
const (
	ErrUsageLimitError             = "usage limit error"
	ErrUsageRatioThresholdError    = "usage ratio threshold error"
	ErrUsageIsBeyondLimits         = "memory usage is beyond the limits"
	ErrMemoryUsageCriterionInvalid = "memory usage criterion is not valid"
	ErrMemoryUsageCriterionUnknown = "memory usage criterion is unknown"
)

const (
	fReportVerbose = "memory usage has been minimized: %d -> %d MB.\r\n"
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
		return nil, errors.New(ErrUsageLimitError)
	}
	if memoryUsedToLimitRatioThreshold <= 0 {
		return nil, errors.New(ErrUsageRatioThresholdError)
	}

	muc = new(MemoryUsageController)

	// Set the memory usage criterion.
	muc.memoryUsageCriterion = NewMemoryUsageCriterion(memoryUsageCriterion)
	if !muc.memoryUsageCriterion.IsValid() {
		return nil, errors.New(ErrMemoryUsageCriterionInvalid)
	}

	// Set other fields.
	muc.memUsageLimitMb = memUsageLimitMb
	muc.memoryUsedToLimitRatioThreshold = memoryUsedToLimitRatioThreshold
	muc.verboseMode = verboseMode

	return muc, nil
}
