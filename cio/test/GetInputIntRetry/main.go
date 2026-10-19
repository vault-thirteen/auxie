package main

import (
	"github.com/vault-thirteen/auxie/cio"
)

func main() {
	cio.PrintMessage("GetInputIntRetry")

	i := cio.GetInputIntRetry()
	cio.PrintMessage("Integer number entered:", i)
}
