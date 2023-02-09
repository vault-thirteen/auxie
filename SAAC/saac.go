package saac

import (
	"math/big"
)

// ValueType is the type of elements used for calculation of Calculator.
type ValueType = float64

// Calculator is a simple arithmetic average calculator.
type Calculator struct {
	// Number of elements in the series.
	n *big.Float

	// Previous simple arithmetic average value.
	previousValue *big.Float

	// A cached value for fast increments.
	one *big.Float
}

// New is the constructor of the Calculator.
func New() (c *Calculator) {
	c = &Calculator{
		n:             big.NewFloat(float64(0)),
		previousValue: nil,
		one:           big.NewFloat(float64(1)),
	}

	return c
}

// AddItemAndGetAverage inserts a new item into the data series and returns a
// new simple arithmetic average value.
func (c *Calculator) AddItemAndGetAverage(item string) (avg string, err error) {
	// avg = ((c.previousValue * n) + x) / (n + 1).

	var newItem *big.Float
	newItem, _, err = new(big.Float).Parse(item, 10)
	if err != nil {
		return "", err
	}

	nPlusOne := new(big.Float)
	nPlusOne.Add(c.n, c.one)

	defer func() {
		c.n = nPlusOne
	}()

	// On a cold start, initialize the calculator.
	if c.previousValue == nil {
		c.previousValue = newItem

		return newItem.String(), nil
	}

	// Normal calculation.
	tmp := new(big.Float)
	tmp.Mul(c.previousValue, c.n)
	tmp.Add(tmp, newItem)
	tmp.Quo(tmp, nPlusOne)

	c.previousValue = tmp

	return tmp.String(), nil
}
