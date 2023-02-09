package sms

import (
	"errors"
	"sort"
)

// SimpleMergeSorter is the simple merge sorter.
type SimpleMergeSorter struct {
	data    [][]ValueType
	cursors []*int
}

// ValueType is the type of the sorted data.
// If you change this, you will also need to make changes to the code.
//type ValueType = int
type ValueType = float64

// New is a constructor of a simple merge sorter.
func New(data ...[]ValueType) (sms *SimpleMergeSorter, err error) {
	// Prepare the raw data.
	// Remove the empty arrays (slices).
	for i := range data {
		//sort.Ints(data[i])
		sort.Float64s(data[i])
	}

	nonEmptyData := getNonEmptyArrays(data...)

	if len(nonEmptyData) < 1 {
		return nil, errors.New("no data to sort")
	}

	sms = &SimpleMergeSorter{
		data: nonEmptyData,
	}

	return sms, nil
}

// getNonEmptyArrays transforms several slices into a slice of non-empty slices.
func getNonEmptyArrays(data ...[]ValueType) (nonEmptyData [][]ValueType) {
	nonEmptyData = make([][]ValueType, 0, len(data))

	for _, items := range data {
		if len(items) != 0 {
			nonEmptyData = append(nonEmptyData, items)
		}
	}

	return nonEmptyData
}

// Sort sorts the data.
func (s *SimpleMergeSorter) Sort() (result []ValueType) {
	// Initialize the cursors.
	s.cursors = make([]*int, len(s.data))
	for i := range s.cursors {
		s.cursors[i] = newIntPointer(0)
	}

	result = make([]ValueType, 0)

	var (
		err  error
		lvci int // A least value cursor index.
	)

	for {
		lvci, err = s.findCursorWithLeastValue()
		if err != nil {
			break
		}

		result = append(result, s.data[lvci][*s.cursors[lvci]])

		// Move the cursor of the current data set.
		*s.cursors[lvci]++
		if *s.cursors[lvci] >= len(s.data[lvci]) {
			s.cursors[lvci] = nil
		}
	}

	return result
}

// newIntPointer creates a new pointer to int.
func newIntPointer(i int) (p *int) {
	p = new(int)
	*p = i

	return p
}

// findCursorWithLeastValue finds a cursor pointing to the least value. It
// returns an error if the least value is not found.
func (s *SimpleMergeSorter) findCursorWithLeastValue() (cursorIndex int, err error) {
	// Find the first non-null value and use it as the first least value.
	// This saves our time, so that we do not check for null least value in
	// every iteration.
	var leastValue ValueType
	leastValueExists := false
	for i, cursor := range s.cursors { // i is a data series index.
		if cursor == nil {
			continue
		}

		// Least value exists.
		leastValueExists = true
		leastValue = s.data[i][*cursor]

		// Remember the last scanned cursor.
		cursorIndex = i
		break
	}

	if !leastValueExists {
		return -1, errors.New("least value is not found")
	}

	// Scan other cursors comparing with the current non-null minimum value.
	iMax := len(s.cursors) - 1
	for i := cursorIndex + 1; i <= iMax; i++ {
		cursor := s.cursors[i]

		if cursor == nil {
			continue
		}

		if s.data[i][*cursor] < leastValue {
			leastValue = s.data[i][*cursor]
			cursorIndex = i
		}
	}

	return cursorIndex, nil
}
