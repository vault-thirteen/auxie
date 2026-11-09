package list

import (
	"cmp"
	"slices"
)

// RemoveDuplicatesFromList removes duplicate items from a generic list.
// Order of items is preserved.
func RemoveDuplicatesFromList[T comparable](in []T) (out []T) {
	m := make(map[T]struct{})
	out = make([]T, 0, len(in))

	for _, x := range in {
		_, exists := m[x]
		if !exists {
			m[x] = struct{}{}
			out = append(out, x)
		}
	}

	return out
}

// RemoveDuplicatesFromListFast is a fast version of the
// RemoveDuplicatesFromList function. The order of items is not preserved. Do
// not use this function if the order of items is important.
func RemoveDuplicatesFromListFast[T cmp.Ordered](in []T) (out []T) {
	if len(in) == 0 {
		return in
	}

	slices.Sort(in)

	return slices.Compact(in)
}
