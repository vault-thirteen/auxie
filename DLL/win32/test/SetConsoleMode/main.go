package main

import (
	"time"

	dll "github.com/vault-thirteen/auxie/DLL/win32"
	"github.com/vault-thirteen/auxie/cio"
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
func prepareDlls() (dc *dll.DllController, err error) {
	dc = dll.NewDllController()

	err = dc.Init()
	if err != nil {
		return nil, err
	}

	return dc, nil
}
func printSomething() {
	printLine("\033[0;31m", "red", "\033[0;0m")
	printLine("\033[0;33m", "yellow", "\033[0;0m")
	printLine("\033[0;32m", "green", "\033[0;0m")
	printLine("\033[0;36m", "cyan", "\033[0;0m")
	printLine("\033[0;34m", "blue", "\033[0;0m")
	printLine("\033[0;35m", "purple", "\033[0;0m")
	printLine("\033[0;30m", "black", "\033[0;0m")
	printLine("\033[0;37m", "white", "\033[0;0m")
	cio.PrintMessage("Have a good day.")
}
func printLine(colOn string, colName string, colOff string) {
	msg := cio.FormatMessage(`This is a test of %s colour.`, colName)
	cio.PrintMessage(colOn + msg + colOff)
}
