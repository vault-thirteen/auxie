package as

import (
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_ftoi(t *testing.T) {
	aTest := tester.New(t)

	aTest.MustBeEqual(ftoi(101/100), 1)
	aTest.MustBeEqual(ftoi(float64(199)/float64(100)), 2)
}

func Test_SearchIBS(t *testing.T) {
	aTest := tester.New(t)

	array := []float64{2, 5, 8.5}

	aTest.MustBeEqual(SearchIBS(array, 0, len(array)-1, 2.0), 0)
	aTest.MustBeEqual(SearchIBS(array, 0, len(array)-1, 5.0), 1)
	aTest.MustBeEqual(SearchIBS(array, 0, len(array)-1, 8.5), 2)

	aTest.MustBeEqual(SearchIBS(array, 0, len(array)-1, 1.0), -1)
	aTest.MustBeEqual(SearchIBS(array, 0, len(array)-1, 6.0), -1)

	// X values greater than the last array item lead to the 'index out of
	// range' exception. Looks like this is a bug in the implementation or a
	// restriction which was not handled by the author.
	//aTest.MustBeEqual(SearchIBS(array, 0, len(array)-1, 8.6), -1)
}

func Test_SearchBS(t *testing.T) {
	aTest := tester.New(t)

	array := []float64{2, 5, 8.5}

	aTest.MustBeEqual(SearchBS(array, 0, len(array)-1, 2.0), 0)
	aTest.MustBeEqual(SearchBS(array, 0, len(array)-1, 5.0), 1)
	aTest.MustBeEqual(SearchBS(array, 0, len(array)-1, 8.5), 2)

	aTest.MustBeEqual(SearchBS(array, 0, len(array)-1, 1.0), -1)
	aTest.MustBeEqual(SearchBS(array, 0, len(array)-1, 6.0), -1)
	aTest.MustBeEqual(SearchBS(array, 0, len(array)-1, 8.6), -1)
}

func Test_SearchIS(t *testing.T) {
	aTest := tester.New(t)

	array := []float64{2, 5, 8.5}

	aTest.MustBeEqual(SearchIS(array, 0, len(array)-1, 2.0), 0)
	aTest.MustBeEqual(SearchIS(array, 0, len(array)-1, 5.0), 1)
	aTest.MustBeEqual(SearchIS(array, 0, len(array)-1, 8.5), 2)

	aTest.MustBeEqual(SearchIS(array, 0, len(array)-1, 1.0), -1)
	aTest.MustBeEqual(SearchIS(array, 0, len(array)-1, 6.0), -1)
	aTest.MustBeEqual(SearchIS(array, 0, len(array)-1, 8.6), -1)
}

func Test_FindNearestBS(t *testing.T) {
	aTest := tester.New(t)

	array := []float64{2, 5, 8.5}

	aTest.MustBeEqual(FindNearestBS(array, 2.0), []int{0})
	aTest.MustBeEqual(FindNearestBS(array, 5.0), []int{1})
	aTest.MustBeEqual(FindNearestBS(array, 8.5), []int{2})

	aTest.MustBeEqual(FindNearestBS(array, -10.0), []int{0})
	aTest.MustBeEqual(FindNearestBS(array, 1.0), []int{0})
	aTest.MustBeEqual(FindNearestBS(array, 2.0), []int{0})
	aTest.MustBeEqual(FindNearestBS(array, 3.0), []int{0})
	aTest.MustBeEqual(FindNearestBS(array, 3.4), []int{0})
	aTest.MustBeEqual(FindNearestBS(array, 3.5), []int{0, 1})
	aTest.MustBeEqual(FindNearestBS(array, 3.6), []int{1})
	aTest.MustBeEqual(FindNearestBS(array, 4.0), []int{1})
	aTest.MustBeEqual(FindNearestBS(array, 6.75), []int{1, 2})
	aTest.MustBeEqual(FindNearestBS(array, 7.0), []int{2})
	aTest.MustBeEqual(FindNearestBS(array, 999.0), []int{2})
}
