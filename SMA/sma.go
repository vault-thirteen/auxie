package sma

import (
	"container/list"
	"errors"
)

// Error messages.
const (
	ErrWindowsSize   = "window size is wrong"
	ErrDataSetIsCold = "data set is cold"
	ErrNotEnoughData = "not enough data"
)

// ValueType is the type of elements used for calculation of Calculator.
type ValueType = float64

// Calculator is a simple moving average (SMA) calculator.
//
// If a data set does not have enough values to calculate the SMA, the
// calculator is considered to be cold. If a data set contains enough values to
// calculate the SMA, the calculator is considered to be hot. A newly created
// calculator is cold.
type Calculator struct {
	// A time frame, size of the window for calculating average values, N.
	size int

	// Items of the current timeframe (window), last N items.
	items *list.List

	// Previous SMA value. It is used for accelerating calculations.
	previousValue *ValueType
}

// New is the constructor of the Calculator.
func New(windowSize int) (c *Calculator, err error) {
	if windowSize < 1 {
		return nil, errors.New(ErrWindowsSize)
	}

	c = &Calculator{
		size:          windowSize,
		items:         list.New(),
		previousValue: nil,
	}

	return c, nil
}

// IsCold checks whether the data set is cold.
func (c Calculator) IsCold() bool {
	return c.items.Len() < c.size
}

// AddItemAndGetSMA inserts a new item into the data series and returns a new
// SMA value.
func (c *Calculator) AddItemAndGetSMA(p ValueType) (sma ValueType, err error) {
	c.items.PushFront(p)

	if c.IsCold() {
		return 0, errors.New(ErrDataSetIsCold)
	}

	// If no previous SMA value is cached.
	if c.previousValue == nil {
		// Perform a full calculation.
		sma, err = c.getItemsAverage()
		if err != nil {
			return 0, err
		}

		// Save the last SMA value.
		c.previousValue = new(ValueType)
		*c.previousValue = sma

		return sma, nil
	}

	// A previous SMA value is cached.
	// Perform a fast calculation.
	oldestItem := c.items.Back()
	sma = *c.previousValue + ((p - oldestItem.Value.(ValueType)) / ValueType(c.size))

	// Remove the oldest item.
	c.items.Remove(oldestItem)

	// Save the last SMA value.
	*c.previousValue = sma

	return sma, nil
}

// Returns an average value of the stored items.
func (c Calculator) getItemsAverage() (average ValueType, err error) {
	// Fool check.
	if c.items.Len() != c.size {
		return 0, errors.New(ErrNotEnoughData)
	}

	// Calculate the average value.
	for item := c.items.Front(); item != nil; item = item.Next() {
		average = average + item.Value.(ValueType)
	}

	average = average / ValueType(c.size)

	return average, nil
}
