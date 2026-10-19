package ver

import (
	"runtime"

	"github.com/kr/pretty"

	"github.com/vault-thirteen/auxie/VCS/common/Version"
	"github.com/vault-thirteen/auxie/Versioneer/classes/Dependency"
	"github.com/vault-thirteen/auxie/Versioneer/classes/ProgramInfo"
	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"
)

const (
	Msgf_IntroTextShort    = "%s, ver. %s. Go language: %s."
	Msgf_IntroTextFull     = "%s %s, ver. %s. Go language: %s."
	Msgf_UpdateIsAvailable = "An update is available to version %v."
)

const (
	GolangIsShit = true
)

// Versioneer is an extended version of the ProgramInfo class.
type Versioneer struct {
	programInfo *pi.ProgramInfo
}

func New(ignoreGolangDevelVersionBug bool) (v *Versioneer, err error) {
	v = new(Versioneer)

	v.programInfo, err = pi.New(ignoreGolangDevelVersionBug)
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
		cio.PrintMessage(
			cio.FormatMessage(Msgf_IntroTextShort,
				v.programInfo.ProgramName(),
				v.programInfo.ProgramVersionNumber(),
				runtime.Version(),
			),
		)
	} else {
		cio.PrintMessage(
			cio.FormatMessage(Msgf_IntroTextFull,
				v.programInfo.ProgramName(),
				product,
				v.programInfo.ProgramVersionNumber(),
				runtime.Version(),
			),
		)
	}

	// Information about a new version.
	if v.programInfo.IsUpdateAvailable() {
		cio.PrintMessage(
			cio.FormatMessage(Msgf_UpdateIsAvailable,
				v.programInfo.LatestVersion().ToString(),
			),
		)
	}
}

func (v *Versioneer) ShowComponentsInfoText() {
	cio.PrintMessage(v.programInfo.DependenciesText())
}

func (v *Versioneer) ShowComponentsInfoList() {
	_, err := pretty.Println(v.programInfo.DependenciesList())
	if err != nil {
		errors.LogError(err)
	}
}

func (v *Versioneer) ProgramName() (programName string) {
	return v.programInfo.ProgramName()
}

func (v *Versioneer) ProgramVersionString() (programVersion string) {
	return v.programInfo.ProgramVersionString()
}

func (v *Versioneer) ProgramVcsVersion() *version.Version {
	return v.programInfo.ProgramVersion()
}

func (v *Versioneer) IsUpdateAvailable() bool {
	return v.programInfo.IsUpdateAvailable()
}

func (v *Versioneer) LatestVersion() *version.Version {
	return v.programInfo.LatestVersion()
}

func (v *Versioneer) DependenciesList() (list []*dependency.Dependency) {
	return v.programInfo.DependenciesList()
}

func (v *Versioneer) DependenciesText() (txt string) {
	return v.programInfo.DependenciesText()
}

func (v *Versioneer) GoVersion() string {
	return v.programInfo.GoVersion()
}
