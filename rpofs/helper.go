// helper.go.

package rpofs

import (
	"slices"
)

func isSetOfUniqueSymbols(set []rune) (ok bool) {
	if len(set) < 1 {
		return false
	}

	slices.Sort(set)

	for i := 0; i < len(set)-1; i++ {
		if set[i] == set[i+1] {
			return false
		}
	}

	return true
}
