package main

import (
	"fmt"
	"log"

	ver "github.com/vault-thirteen/auxie/Versioneer/classes/Versioneer"
)

func main() {
	showIntro()
}

func showIntro() {
	versioneer, err := ver.New()
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
