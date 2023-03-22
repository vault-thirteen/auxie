package time

import "time"

// IsEmpty tells whether the time is empty or not.
func IsEmpty(t time.Time) bool {
	return t == emptyTime
}
