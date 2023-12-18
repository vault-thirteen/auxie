package httphelper

import (
	"errors"
	"sort"
	"strings"
)

// Parsing of an 'Accept' HTTP header according to the documentation:
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept

const (
	ErrNoRecords           = "no records"
	ErrSyntaxErrorInRecord = "syntax error in record: %s"
	ErrSyntaxErrorInWeight = "syntax error in weight: %s"
	ErrEndOfList           = "EOL"
)

// AcceptedMimeTypes provides convenient access to a list of MIME types
// received from the 'Accept' HTTP header. The struct has a simple iterator
// which allows to list values with the 'Next' method. The first returned value
// is the most acceptable, i.e. a value with the greatest weight. The last
// returned value is the least acceptable, i.e. a value with the lowest weight.
// Please, note that the iterator is not protected by a mutex against
// simultaneous access for performance reasons.
type AcceptedMimeTypes struct {
	// Values that can be iterated.
	values []*AcceptedMimeType

	// Cursor which shows index of the last retrieved element.
	cursor int

	// Last iterable index in the sequence of stored values.
	lastIdx int
}

// NewAcceptedMimeTypesFromHeader parses the provided 'Accept' HTTP header and
// returns an AcceptedMimeType object.
func NewAcceptedMimeTypesFromHeader(hdr string) (amts *AcceptedMimeTypes, err error) {
	var values []*AcceptedMimeType
	values, err = parseAcceptHttpHeader(hdr)
	if err != nil {
		return nil, err
	}

	// It is important to save the original order of items with equal weights,
	// that is why the stable sorting is used.
	sort.SliceStable(values, func(i, j int) bool { return values[i].Weight > values[j].Weight })

	return newAcceptedMimeTypes(values), nil
}

// newAcceptedMimeTypes is a simple constructor of the AcceptedMimeTypes
// object.
func newAcceptedMimeTypes(values []*AcceptedMimeType) (amts *AcceptedMimeTypes) {
	amts = &AcceptedMimeTypes{
		values:  values,
		lastIdx: len(values) - 1,
	}

	amts.resetIterator()

	return amts
}

// resetIterator resets the iterator to its initial state.
func (amts *AcceptedMimeTypes) resetIterator() {
	amts.cursor = -1
}

// parseAcceptHttpHeader tries to parse the 'Accept' HTTP header.
func parseAcceptHttpHeader(hdr string) (types []*AcceptedMimeType, err error) {
	records := strings.Split(hdr, ",")

	if len(records) == 0 {
		return nil, errors.New(ErrNoRecords)
	}

	types = make([]*AcceptedMimeType, 0)
	var t *AcceptedMimeType
	for _, r := range records {
		t, err = ParseRecord(r)
		if err != nil {
			return nil, err
		}

		types = append(types, t)
	}

	return types, nil
}

// Next gets the next MIME type in the list. If no more items are available,
// an error is returned.
func (amts *AcceptedMimeTypes) Next() (amt *AcceptedMimeType, err error) {
	amts.cursor++

	if amts.cursor > amts.lastIdx {
		return nil, errors.New(ErrEndOfList)
	}

	return amts.values[amts.cursor], nil
}

// Reset resets the iterator. Normally, such a method is not used in iterable
// objects, but since the Go programming language is very unusual and unique,
// for those who for some reason want to iterate values several times, this
// method may be useful.
func (amts *AcceptedMimeTypes) Reset() {
	amts.resetIterator()
}
