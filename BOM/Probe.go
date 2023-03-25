package bom

import (
	"bytes"
	"errors"
	"fmt"

	tsb "github.com/vault-thirteen/auxie/TSB"
)

const (
	ErrArraysHaveDifferentLengths = "arrays have different lengths: %v vs %v"
	ErrNoData                     = "no data"
)

// Probe stores the result of probing the text for a specified encoding.
// In other words, it stores the probability of the probes to be of the
// specified encoding.
type Probe struct {
	// Encoding is the specified encoding which is searched in the probes.
	Encoding Encoding

	// Probability is the probability of the encoding to be used in the probes.
	Probability tsb.TSB

	// ReadBytesCount is the number of bytes which were read to get the probe.
	ReadBytesCount int
}

// IsAccurate tells whether the probe results are accurate or not.
// Here by accuracy we mean the exact 'yes' or 'no' probability.
func (p *Probe) IsAccurate() bool {
	if p.Probability.IsYes() || p.Probability.IsNo() {
		return true
	}
	if p.Probability.IsMaybe() {
		return false
	}

	panic(p.Probability)
}

// ProbeForEncoding tries to search the probes for the specified encoding.
func ProbeForEncoding(data []byte, enc Encoding) (probe *Probe, err error) {
	bom, ok := boms[enc]
	if !ok {
		return nil, fmt.Errorf(ErrUnknownEncoding, enc)
	}

	// Is there enough data ?
	if len(data) >= len(bom) {
		// Data is enough.
		sample := data[:len(bom)]
		if bytes.Equal(sample, bom) {
			return &Probe{
				Encoding:       enc,
				Probability:    tsb.Yes,
				ReadBytesCount: len(bom),
			}, nil
		} else {
			return &Probe{
				Encoding:       enc,
				Probability:    tsb.No,
				ReadBytesCount: len(bom),
			}, nil
		}
	}

	// Data is not enough.
	bomPart := bom[:len(data)]

	var ecb int
	ecb, err = countEqualConsecutiveBytes(bomPart, data)
	if err != nil {
		return nil, err
	}

	if ecb < len(data) {
		// No match.
		return &Probe{
			Encoding:       enc,
			Probability:    tsb.No,
			ReadBytesCount: len(data),
		}, nil
	} else {
		// Match.
		return &Probe{
			Encoding:       enc,
			Probability:    tsb.Maybe,
			ReadBytesCount: len(data),
		}, nil
	}
}

// countEqualConsecutiveBytes counts the number of equal consecutive bytes from
// the beginning of arrays (slices).
func countEqualConsecutiveBytes(s1, s2 []byte) (n int, err error) {
	if len(s1) != len(s2) {
		return n, fmt.Errorf(ErrArraysHaveDifferentLengths, len(s1), len(s2))
	}
	if len(s1) == 0 {
		return n, errors.New(ErrNoData)
	}

	for i, b1 := range s1 {
		if b1 != s2[i] {
			return i, nil
		}
	}

	return len(s1), nil
}
