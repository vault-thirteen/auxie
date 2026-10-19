package as

import "math"

const (
	IndexNotFound = -1
)

// ftoi converts a float into an int using common sense, i.e. with mathematical
// rounding. Be aware that Go language by default truncates all the information
// about fraction when divides integers by integers. E.g., 199/100=1 in Golang.
func ftoi(f float64) int {
	return int(math.Round(f))
}

// SearchIBS is a search using the Interpolated Binary Search method.
// This is an implementation of the algorithm described at the source adapted
// to the Go programming language.
// Source:
// https://www.sciencedirect.com/science/article/pii/S221509862100046X
// Variables with small case letter are indices (int type).
// Variables with upper case letter are values (float type).
// X is a value to be found.
//
// P.S.
// Looks like this implementation has a bug. See tests for more details.
func SearchIBS(Arr []float64, leftIndex int, rightIndex int, X float64) (xIndex int) {
	var interIndex, midIndex int
	for leftIndex < rightIndex { // 1 Comparison.
		interIndex = ftoi(float64(leftIndex) + (X-Arr[leftIndex])*float64(rightIndex-leftIndex)/(Arr[rightIndex]-Arr[leftIndex]))

		if X > Arr[interIndex] { // 2 Comparisons.
			midIndex = (interIndex + rightIndex) / 2
			if X <= Arr[midIndex] { // 3 Comparisons.
				leftIndex = interIndex + 1
				rightIndex = midIndex
			} else {
				leftIndex = midIndex + 1
			}
		} else if X < Arr[interIndex] { // 3 Comparisons.
			midIndex = (leftIndex + interIndex) / 2
			if X >= Arr[midIndex] { // 4 Comparisons.
				leftIndex = midIndex
				rightIndex = interIndex - 1
			} else {
				rightIndex = midIndex - 1
			}
		} else {
			return interIndex
		}
	} // End of while.

	if X == Arr[leftIndex] {
		return leftIndex
	}

	return IndexNotFound
}

// SearchBS is a binary search.
// X is a value to be found.
func SearchBS(Arr []float64, leftIndex int, rightIndex int, X float64) (xIndex int) {
	var midIndex int
	for leftIndex <= rightIndex {
		midIndex = (leftIndex + rightIndex) / 2 // Go language does not round the result.
		if X < Arr[midIndex] {
			rightIndex = midIndex - 1
		} else if Arr[midIndex] < X {
			leftIndex = midIndex + 1
		} else {
			return midIndex
		}
	}

	return IndexNotFound
}

// SearchIS is an interpolated search.
// X is a value to be found.
func SearchIS(Arr []float64, leftIndex int, rightIndex int, X float64) (xIndex int) {
	var midIndex int
	for (Arr[leftIndex] < Arr[rightIndex]) && (Arr[leftIndex] <= X) && (X <= Arr[rightIndex]) {
		midIndex = ftoi(float64(leftIndex) + (X-Arr[leftIndex])*float64(rightIndex-leftIndex)/(Arr[rightIndex]-Arr[leftIndex]))

		if X < Arr[midIndex] {
			rightIndex = midIndex - 1
		} else if Arr[midIndex] < X {
			leftIndex = midIndex + 1
		} else {
			return midIndex
		}
	}

	return IndexNotFound
}

// FindNearestBS is a binary search modification to find an item with a value
// nearest to the specified value V. If two items are both nearest to the
// specified value, they are both returned.
func FindNearestBS(Arr []float64, V float64) (iIndices []int) {
	var l = 0
	var r = len(Arr) - 1

	if V < Arr[l] {
		return []int{l}
	}
	if Arr[r] < V {
		return []int{r}
	}

	var m int
	for l <= r {
		m = (l + r) / 2 // Go language does not round the result.
		if V < Arr[m] {
			r = m - 1
		} else if Arr[m] < V {
			l = m + 1
		} else {
			return []int{m}
		}
	}

	// l is r+1 after the swap.
	if (Arr[l] - V) < (V - Arr[r]) {
		return []int{l}
	} else if (Arr[l] - V) > (V - Arr[r]) {
		return []int{r}
	} else {
		return []int{r, l}
	}
}
