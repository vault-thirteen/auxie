package nts

import (
	"bytes"
	"errors"
)

// Package for processing null-terminated strings.

const NUL = 0

const ErrNTSIsNotValid = "null-terminated string is not valid"

// ByteArrayToStrings converts a byte array of null-terminated strings into a
// slice of Golang's strings.
func ByteArrayToStrings(ba []byte) (ss []string, err error) {
	if len(ba) == 0 {
		return ss, nil
	}

	// If the last byte is not a NUL byte,
	// then this null-terminated string is not valid.
	if ba[len(ba)-1] != NUL {
		return nil, errors.New(ErrNTSIsNotValid)
	}

	// Split the byte array into strings.
	ss = make([]string, 0)
	window := ba[:]
	sepIdx := bytes.IndexByte(window, NUL)
	for sepIdx >= 0 {
		ss = append(ss, string(window[:sepIdx]))

		// Next.
		if sepIdx == len(window)-1 {
			break
		}
		window = window[sepIdx+1:]
		sepIdx = bytes.IndexByte(window, NUL)
	}

	return ss, nil
}
