package main

import (
	"fmt"
	"log"

	ver "github.com/vault-thirteen/auxie/Versioneer/classes/Versioneer"
)

func main() {
	showIntro(false)
}

func showIntro(ignoreGolangDevelVersionBug bool) {
	versioneer, err := ver.New(ignoreGolangDevelVersionBug)
	mustBeNoError(err)
	versioneer.ShowIntroText("Server")
	versioneer.ShowComponentsInfoText()
	fmt.Println()
}

func mustBeNoError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
