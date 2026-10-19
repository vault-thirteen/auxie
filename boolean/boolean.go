package boolean

import (
	"strings"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	TrueLc  = "true"
	FalseLc = "false"
	YesLc   = "yes"
	NoLc    = "no"
	One     = "1"
	Zero    = "0"
)

const (
	Errf_BadBoolean = "bad boolean value: '%v'"
)

func FromString(s string) (b bool, err error) {
	if len(s) == 0 {
		return false, errors.ErrorFromFormattedText(Errf_BadBoolean, s)
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

	return false, errors.ErrorFromFormattedText(Errf_BadBoolean, s)
}

func FromStringP(s string) (b bool) {
	var err error
	b, err = FromString(s)
	if err != nil {
		panic(err)
		return
	}
	return b
}

func FromStringM(s string) *bool {
	b, err := FromString(s)
	if err != nil {
		return nil
	}
	return &b
}
