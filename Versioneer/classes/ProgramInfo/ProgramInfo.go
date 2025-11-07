package pi

import (
	"errors"
	"fmt"
	"log"
	"path"
	"runtime/debug"
	"strings"

	"github.com/vault-thirteen/auxie/VCS/common/Version"
	"github.com/vault-thirteen/auxie/Versioneer/classes/Dependency"
)

const (
	DependenciesTextPrefix = "Dependencies: "
	ProjectUnknown         = "???"
	VersionUnknown         = "v??"
)

const (
	ErrUnsupportedMainPathFormat = "unsupported main path format: %v"
	ErrBuildInfoIsNotAvailable   = "build info is not available"
)

// ProgramInfo is information about program and its dependencies.
type ProgramInfo struct {
	buildInfo *debug.BuildInfo

	platform string
	account  string
	project  string

	version      *version.Version
	versionText1 string // Long string with prefix.
	versionText2 string // Short string without prefix.

	dependencies     []*dependency.Dependency
	dependenciesText string

	latestVersion     *version.Version
	isUpdateAvailable bool
}

// New provides convenient access to program's name, program's version and
// information about dependencies. Program's version is taken from both runtime
// and VCS. Dependencies are taken from runtime.
func New() (info *ProgramInfo, err error) {
	info = &ProgramInfo{
		dependencies: make([]*dependency.Dependency, 0),
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return info, errors.New(ErrBuildInfoIsNotAvailable)
	}

	info.buildInfo = bi

	// Projects hosted on GitHub often provide a "Main.Path" similar to
	// following example: "github.com/account/project/...".
	//
	// Notes.
	//
	// 1. Be aware that when a project represents a collection of multiple
	// projects inside one, then this scheme will not work.
	//
	// 2. When a project is hosted on a platform other than GitHub, this scheme
	// may not work due to many reasons. We can not guarantee anything on any
	// platform except the GitHub while there is no official universal standard
	// for storing projects.

	mainPathParts := strings.Split(bi.Main.Path, "/")
	if len(mainPathParts) < 3 {
		return info, fmt.Errorf(ErrUnsupportedMainPathFormat, bi.Main.Path)
	}

	info.platform = mainPathParts[0]

	info.account = mainPathParts[1]

	info.project = mainPathParts[2]
	if len(info.project) == 0 {
		info.project = ProjectUnknown
	}

	info.versionText1 = path.Base(bi.Main.Version)
	if len(info.versionText1) == 0 {
		info.versionText1 = VersionUnknown
	}

	info.version, err = version.New(info.versionText1)
	if err != nil {
		return nil, err
	}

	info.versionText2 = info.version.ShortString()

	var dep *dependency.Dependency
	for _, d := range bi.Deps {
		dep, err = dependency.New(d)
		if err != nil {
			return nil, err
		}

		info.dependencies = append(info.dependencies, dep)
	}

	info.initDependenciesText()

	//_, _ = pretty.Println(bi) //DEBUG.

	err = info.checkForUpdates()
	if err != nil {
		log.Println(err)
	}

	return info, nil
}

func (pi *ProgramInfo) initDependenciesText() {
	var sb = new(strings.Builder)
	sb.WriteString(DependenciesTextPrefix)

	for _, dep := range pi.dependencies {
		sb.WriteString("[" + dep.Name() + " " + dep.Version() + "] ")
	}

	pi.dependenciesText = sb.String()
}

// ProgramName returns program's name provided by runtime.
func (pi *ProgramInfo) ProgramName() (name string) {
	return pi.project
}

// ProgramVersionString returns version string provided by runtime.
func (pi *ProgramInfo) ProgramVersionString() (version string) {
	return pi.versionText1
}

// ProgramVersionNumber returns version string provided by runtime without
// prefix.
func (pi *ProgramInfo) ProgramVersionNumber() (versionNumber string) {
	return pi.versionText2
}

// ProgramVersion returns a version object provided by runtime.
func (pi *ProgramInfo) ProgramVersion() (v *version.Version) {
	return pi.version
}

// IsUpdateAvailable tells if there is an updated versionText of the program
// available. This information is provided by VCS.
func (pi *ProgramInfo) IsUpdateAvailable() bool {
	return pi.isUpdateAvailable
}

// LatestVersion return latest versionText provided by VCS.
func (pi *ProgramInfo) LatestVersion() *version.Version {
	return pi.latestVersion
}

// DependenciesList returns a list of dependencies provided by runtime.
func (pi *ProgramInfo) DependenciesList() (list []*dependency.Dependency) {
	return pi.dependencies
}

// DependenciesText returns a textual list of dependencies provided by runtime.
func (pi *ProgramInfo) DependenciesText() (txt string) {
	return pi.dependenciesText
}
