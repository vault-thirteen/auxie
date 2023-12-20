package ver

import (
	"fmt"
	"log"
	"runtime"

	"github.com/kr/pretty"
	ver "github.com/vault-thirteen/auxie/VCS/common/Version"
	pi "github.com/vault-thirteen/auxie/Versioneer/ProgramInfo"
)

const (
	IntroTextShort       = "%s, ver. %s. Go language: %s."
	IntroTextFull        = "%s %s, ver. %s. Go language: %s."
	MsgUpdateIsAvailable = "An update is available to version %v."
)

// Versioneer is an extended version of the ProgramInfo class.
type Versioneer struct {
	programInfo *pi.ProgramInfo
}

func New() (v *Versioneer, err error) {
	v = new(Versioneer)

	v.programInfo, err = pi.NewProgramInfo()
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ShowIntroText shows introductory text about the program.
// The 'product' parameter is optional. If it is set, it is printed after the
// program name. It is used mostly for showing server and client variants of a
// product.
func (v *Versioneer) ShowIntroText(product string) {
	if len(product) == 0 {
		fmt.Println(
			fmt.Sprintf(IntroTextShort,
				v.programInfo.ProgramName(),
				v.programInfo.ProgramVersionNumber(),
				runtime.Version(),
			),
		)
	} else {
		fmt.Println(
			fmt.Sprintf(IntroTextFull,
				v.programInfo.ProgramName(),
				product,
				v.programInfo.ProgramVersionNumber(),
				runtime.Version(),
			),
		)
	}

	// Information about a new version.
	if v.programInfo.IsUpdateAvailable() {
		fmt.Println(
			fmt.Sprintf(MsgUpdateIsAvailable,
				v.programInfo.LatestVersion().ToString(),
			),
		)
	}
}

func (v *Versioneer) ShowComponentsInfoText() {
	fmt.Println(v.programInfo.DependenciesText())
}

func (v *Versioneer) ShowComponentsInfoList() {
	_, err := pretty.Println(v.programInfo.DependenciesList())
	if err != nil {
		log.Println(err)
	}
}

func (v *Versioneer) ProgramName() (programName string) {
	return v.programInfo.ProgramName()
}

func (v *Versioneer) ProgramVersionString() (programVersion string) {
	return v.programInfo.ProgramVersionString()
}

func (v *Versioneer) ProgramVcsVersion() *ver.Version {
	return v.programInfo.ProgramVcsVersion()
}

func (v *Versioneer) IsUpdateAvailable() bool {
	return v.programInfo.IsUpdateAvailable()
}

func (v *Versioneer) LatestVersion() *ver.Version {
	return v.programInfo.LatestVersion()
}

func (v *Versioneer) DependenciesList() (list []*pi.Dependency) {
	return v.programInfo.DependenciesList()
}

func (v *Versioneer) DependenciesText() (txt string) {
	return v.programInfo.DependenciesText()
}
