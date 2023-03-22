// helper.go.

package rpofs

import "sort"

func isSetOfUniqueSymbols(set []rune) (ok bool) {
	if len(set) < 1 {
		return false
	}

	lessFn := func(i, j int) bool {
		return set[i] < set[j]
	}

	sort.Slice(set, lessFn)

	for i := 0; i < len(set)-1; i++ {
		if set[i] == set[i+1] {
			return false
		}
	}

	return true
}
