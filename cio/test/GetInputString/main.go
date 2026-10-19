package main

import (
	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"
)

func main() {
	cio.PrintMessage("GetInputString")

	s, err := cio.GetInputString(true)
	if err != nil {
		errors.LogError(err)
		return
	}

	cio.PrintMessage("String entered:", s)
}
