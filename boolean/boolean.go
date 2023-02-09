// boolean.go.

package boolean

import (
	"fmt"
	"strings"
)

const (
	TrueLc  = "true"
	FalseLc = "false"
	YesLc   = "yes"
	NoLc    = "no"
	One     = "1"
	Zero    = "0"
)

const ErrfBadBoolean = "bad boolean value: '%v'"

func FromString(
	s string,
) (b bool, err error) {
	if len(s) == 0 {
		return false, fmt.Errorf(ErrfBadBoolean, s)
	}

	// 1. Try without conversion to lower case.
	switch s {
	case TrueLc:
		return true, nil
	case FalseLc:
		return false, nil
	case YesLc:
		return true, nil
	case NoLc:
		return false, nil
	case One:
		return true, nil
	case Zero:
		return false, nil
	}

	// 2. Try with conversion to lower case and space trimming.
	sLc := strings.ToLower(strings.TrimSpace(s))
	switch sLc {
	case TrueLc:
		return true, nil
	case FalseLc:
		return false, nil
	case YesLc:
		return true, nil
	case NoLc:
		return false, nil
	case One:
		return true, nil
	case Zero:
		return false, nil
	}

	return false, fmt.Errorf(ErrfBadBoolean, s)
}
