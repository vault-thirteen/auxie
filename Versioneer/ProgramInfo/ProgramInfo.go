package pi

import (
	"errors"
	"log"
	"path"
	"runtime/debug"
	"strings"

	ver "github.com/vault-thirteen/auxie/VCS/common/Version"
)

const (
	DependenciesTextPrefix = "Dependencies: "
)

// ProgramInfo is information about program and its dependencies.
type ProgramInfo struct {
	name             string
	version          string
	dependencies     []*Dependency
	dependenciesText string

	buildInfo         *debug.BuildInfo
	vcsVersion        *ver.Version
	latestVersion     *ver.Version
	isUpdateAvailable bool
}

// NewProgramInfo provides convenient access to program's name, program's
// version and information about dependencies. Program's version is taken from
// both runtime and VCS. Dependencies are taken from runtime.
func NewProgramInfo() (info *ProgramInfo, err error) {
	info = &ProgramInfo{
		name:         NameUnknown,
		version:      VersionUnknown,
		dependencies: make([]*Dependency, 0),
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return info, errors.New(ErrBuildInfoIsNotAvailable)
	}

	var tmp string
	tmp = path.Base(bi.Main.Path)
	if len(tmp) == 0 {
		info.name = NameUnknown
	} else {
		info.name = tmp
	}

	tmp = path.Base(bi.Main.Version)
	if len(tmp) == 0 {
		info.version = VersionUnknown
	} else {
		info.version = tmp
	}

	var dep *Dependency
	for _, d := range bi.Deps {
		dep, err = NewDependency(d)
		if err != nil {
			return nil, err
		}

		info.dependencies = append(info.dependencies, dep)
	}

	info.initDependenciesText()

	info.buildInfo = bi
	//_, _ = pretty.Println(bi) //DEBUG.

	var vcsErr error
	vcsErr = info.parseVCSVersion()
	if vcsErr != nil {
		log.Println(vcsErr)
	}

	vcsErr = info.checkForUpdates()
	if vcsErr != nil {
		log.Println(vcsErr)
	}

	return info, nil
}

func (pi *ProgramInfo) initDependenciesText() {
	var sb = new(strings.Builder)
	sb.WriteString(DependenciesTextPrefix)

	for _, dep := range pi.dependencies {
		sb.WriteString("[" + dep.name + " " + dep.version + "] ")
	}

	pi.dependenciesText = sb.String()
}

// ProgramName returns program's name provided by runtime.
func (pi *ProgramInfo) ProgramName() (name string) {
	return pi.name
}

// ProgramVersionString returns version string provided by runtime.
func (pi *ProgramInfo) ProgramVersionString() (version string) {
	return pi.version
}

// ProgramVersionNumber returns version string provided by runtime with prefix
// trimmed.
func (pi *ProgramInfo) ProgramVersionNumber() (versionNumber string) {
	return strings.TrimPrefix(strings.TrimPrefix(pi.version, "v"), "ver")
}

// ProgramVcsVersion returns version string provided by VCS.
func (pi *ProgramInfo) ProgramVcsVersion() *ver.Version {
	return pi.vcsVersion
}

// IsUpdateAvailable tells if there is an updated version of the program
// available. This information is provided by VCS.
func (pi *ProgramInfo) IsUpdateAvailable() bool {
	return pi.isUpdateAvailable
}

// LatestVersion return latest version provided by VCS.
func (pi *ProgramInfo) LatestVersion() *ver.Version {
	return pi.latestVersion
}

// DependenciesList returns a list of dependencies provided by runtime.
func (pi *ProgramInfo) DependenciesList() (list []*Dependency) {
	return pi.dependencies
}

// DependenciesText returns a textual list of dependencies provided by runtime.
func (pi *ProgramInfo) DependenciesText() (txt string) {
	return pi.dependenciesText
}
