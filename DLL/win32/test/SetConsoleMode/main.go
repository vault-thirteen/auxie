package main

import (
	"time"

	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/os/win32"
)

func main() {
	dc, err := prepareDlls()
	mustBeNoError(err)

	err = dc.EnableConsoleColours()
	mustBeNoError(err)

	printSomething()
	time.Sleep(time.Second * 10)
}

func mustBeNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func prepareDlls() (dc *win32.DllController, err error) {
	dc = win32.NewDllController()

	err = dc.Init()
	if err != nil {
		return nil, err
	}

	return dc, nil
}

func printSomething() {
	cio.PrintMessage("\033[0;31mThis is a test of red colour.\033[0;0m")
	cio.PrintMessage("\033[0;33mThis is a test of yellow colour.\033[0;0m")
	cio.PrintMessage("\033[0;32mThis is a test of green colour.\033[0;0m")
	cio.PrintMessage("\033[0;36mThis is a test of cyan colour.\033[0;0m")
	cio.PrintMessage("\033[0;34mThis is a test of blue colour.\033[0;0m")
	cio.PrintMessage("\033[0;35mThis is a test of purple colour.\033[0;0m")
	cio.PrintMessage("\033[0;30mThis is a test of black colour.\033[0;0m")
	cio.PrintMessage("\033[0;37mThis is a test of white colour.\033[0;0m")
	cio.PrintMessage("Have a good day.")
}
