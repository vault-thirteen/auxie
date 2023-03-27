package slreader

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"time"

	"golang.org/x/time/rate"
)

const (
	ErrBurstToNormalRatioOverflow = "overflow in burst-to-normal ratio: %v vs %v"
	ErrZeroSpeed                  = "zero speed is not allowed"
)

// SLReader is a speed-limited reader.
type SLReader struct {
	// Burst-to-normal ratio.
	bnr float64

	// Context for the limiter.
	ctx context.Context

	// Limiter.
	limiter *rate.Limiter

	// Reader.
	r io.Reader
}

// NewReader creates a new SLReader.
// 'normalLimit' and 'burstLimit' are speed limits set in bytes per second.
// 'bnr' is the maximum burst-to-normal ratio allowed, i.e. max(burst/normal).
func NewReader(r io.Reader, normalLimit float64, burstLimit int, bnr float64) (slr *SLReader, err error) {
	if normalLimit == 0 {
		return nil, errors.New(ErrZeroSpeed)
	}

	bnrActual := float64(burstLimit) / normalLimit
	if bnrActual > bnr {
		return nil, fmt.Errorf(ErrBurstToNormalRatioOverflow, bnr, bnrActual)
	}

	slr = &SLReader{
		bnr: bnr,

		// Instead of using a context, we limit the burst speed indirectly with
		// a burst-to-normal ratio.
		ctx: context.Background(),

		// Looks like a fresh limiter has the number of tokens equal to its
		// burst limit. Also, the limiter does not allow to change its settings.
		// To change the limiter's settings we need to create a new limiter.
		// Well, this is Golang. I am not surprised.
		limiter: rate.NewLimiter(rate.Limit(normalLimit), burstLimit),

		r: r,
	}

	// Spend all the tokens, i.e. reset the tokens count to zero.
	slr.limiter.AllowN(time.Now(), burstLimit)

	return slr, nil
}

// ChangeLimits changes the speed limit settings.
func (slr *SLReader) ChangeLimits(normalLimit float64, burstLimit int) (err error) {
	if normalLimit == 0 {
		return errors.New(ErrZeroSpeed)
	}

	bnrActual := float64(burstLimit) / normalLimit
	if bnrActual > slr.bnr {
		return fmt.Errorf(ErrBurstToNormalRatioOverflow, slr.bnr, bnrActual)
	}

	// Create a new limiter and transfer non-spent tickets to it.

	// If new speed limits are greater than the old ones, we can simply copy
	// all the non-spent tokens from an old limiter. If the settings are lesser
	// than the old ones, we need to copy the minimum of two values: non-spent
	// tokens from the old limiter and the limit of the new limiter.
	newLimiter := rate.NewLimiter(rate.Limit(normalLimit), burstLimit)
	newLimiterTokens := newLimiter.Tokens()
	nonSpentTokens := slr.limiter.Tokens()
	tokensToTransfer := math.Min(newLimiterTokens, nonSpentTokens)

	// Tokens to spend from new limiter to get a new value equal to the
	// 'tokensToTransfer' variable.
	deltaTokens := newLimiterTokens - tokensToTransfer

	// Built-in library does not care about differences in types and their
	// extremal values. Well, this is Golang, and I am not surprised again.
	// LOL.
	newLimiter.AllowN(time.Now(), int(deltaTokens))

	slr.limiter = newLimiter

	return nil
}

// Read tries to read bytes into the destination (dst).
// For more information see the io.Reader interface.
func (slr *SLReader) Read(dst []byte) (n int, err error) {
	n, err = slr.r.Read(dst)
	if err != nil {
		return n, err
	}

	err = slr.limiter.WaitN(slr.ctx, n)
	if err != nil {
		return n, err
	}

	return n, nil
}

// Close closes the reader.
// It is a standard method for the 'io.Closer' interface.
func (slr *SLReader) Close() (err error) {
	return nil
}
