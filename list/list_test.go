package list

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_RemoveDuplicatesFromList(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(RemoveDuplicatesFromList([]int{1, 3, 5, 3, 7, 1}), []int{1, 3, 5, 7})
	aTest.MustBeEqual(RemoveDuplicatesFromList([]int{8, 0, 8}), []int{8, 0})
	aTest.MustBeEqual(RemoveDuplicatesFromList([]int{0}), []int{0})
	aTest.MustBeEqual(RemoveDuplicatesFromList([]int{}), []int{})
}

func Test_RemoveDuplicatesFromListFast(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(RemoveDuplicatesFromListFast([]int{}), []int{})
	aTest.MustBeEqual(RemoveDuplicatesFromListFast([]int{0}), []int{0})
	aTest.MustBeEqual(RemoveDuplicatesFromListFast([]int{8, 0, 8}), []int{0, 8})
	aTest.MustBeEqual(RemoveDuplicatesFromListFast([]int{7, 5, 3, 5}), []int{3, 5, 7})
}
