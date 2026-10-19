package main

import (
	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"
)

func main() {
	cio.PrintMessage("GetInputKey")
	var loop = true
	for loop {
		key, err := cio.GetInputKey()
		if err != nil {
			errors.LogError(err)
			continue
		}

		cio.PrintMessage("Pressed key: ", key.Name())
		if key.Code() == cio.VK_ESCAPE {
			loop = false
		}
	}
}
