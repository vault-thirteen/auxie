package main

import (
	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"
)

func main() {
	cio.PrintMessage("GetInputPassword")

	pwd, err := cio.GetInputPassword()
	if err != nil {
		errors.LogError(err)
		return
	}

	cio.PrintMessage("Password entered:", pwd)
}
