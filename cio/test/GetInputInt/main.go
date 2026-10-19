package main

import (
	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"
)

func main() {
	cio.PrintMessage("GetInputInt")

	i, err := cio.GetInputInt()
	if err != nil {
		errors.LogError(err)
		return
	}

	cio.PrintMessage("Integer number entered:", i)
}
