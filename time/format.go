package time

import (
	"github.com/vault-thirteen/auxie/cio"
)

// NewTimeStringRFC3339 creates a new time string using the RFC3339.
func NewTimeStringRFC3339(year uint, month uint, day uint, hour uint, minute uint, second uint) string {
	return cio.FormatMessage("%04d-%02d-%02dT%02d:%02d:%02dZ", year, month, day, hour, minute, second)
}
