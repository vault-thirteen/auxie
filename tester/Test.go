package tester

import (
	"testing"
)

// Test object helps to make tests.
type Test struct {
	t *testing.T
}

// New creates a test.
func New(t *testing.T) *Test {
	return &Test{t: t}
}
