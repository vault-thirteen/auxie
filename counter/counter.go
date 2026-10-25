package c

import (
	"math/big"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	Err_Load = "load error"
)

type Counter struct {
	bic *big.Int
}

func NewCounter() (c *Counter) {
	c = &Counter{
		bic: big.NewInt(0),
	}

	return c
}

func (c *Counter) RawValue() *big.Int {
	return c.bic
}

func (c *Counter) ToString() string {
	return c.bic.String()
}

func (c *Counter) Load(value string) (err error) {
	_, ok := c.bic.SetString(value, 10)
	if !ok {
		return errors.ErrorFromText(Err_Load)
	}

	return nil
}

func (c *Counter) Inc() {
	c.bic.Add(c.bic, big.NewInt(1))
}
