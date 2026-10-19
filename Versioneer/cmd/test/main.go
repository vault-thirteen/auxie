package main

import (
	"os"

	ver "github.com/vault-thirteen/auxie/Versioneer/classes/Versioneer"
	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"
)

func main() {
	showIntro(false)
}

func showIntro(ignoreGolangDevelVersionBug bool) {
	versioneer, err := ver.New(ignoreGolangDevelVersionBug)
	mustBeNoError(err)
	versioneer.ShowIntroText("Server")
	versioneer.ShowComponentsInfoText()
	cio.PrintMessage()
}

func mustBeNoError(err error) {
	if err != nil {
		errors.LogError(err)
		os.Exit(1)
	}
}
