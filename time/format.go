package time

import "fmt"

// NewTimeStringRFC3339 creates a new time string using the RFC3339.
func NewTimeStringRFC3339(year uint, month uint, day uint, hour uint, minute uint, second uint) string {
	return fmt.Sprintf(
		"%s-%s-%sT%s:%s:%sZ",
		fmt.Sprintf("%04d", year),
		fmt.Sprintf("%02d", month),
		fmt.Sprintf("%02d", day),
		fmt.Sprintf("%02d", hour),
		fmt.Sprintf("%02d", minute),
		fmt.Sprintf("%02d", second),
	)
}
