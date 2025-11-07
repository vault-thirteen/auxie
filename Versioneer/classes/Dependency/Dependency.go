package dependency

import (
	"errors"
	"path"
	"runtime/debug"
	"strings"

	"github.com/vault-thirteen/auxie/VCS/common/Version"
	"github.com/vault-thirteen/auxie/number"
)

const (
	NameUnknown    = "???"
	VersionUnknown = "v??"
)

const (
	ErrDependencyInfoIsNotAvailable = "dependency info is not available"
)

// Dependency is information about program's dependency.
type Dependency struct {
	name    string
	version string
}

func New(m *debug.Module) (dep *Dependency, err error) {
	dep = &Dependency{}

	if m == nil {
		return dep, errors.New(ErrDependencyInfoIsNotAvailable)
	}

	// Some repositories with code written in Go language are using an ugly
	// postfix "hack" with version in it. We need to read that shit too.
	var tmp string
	tmp = path.Base(m.Path)
	if isStringAVersionPostfix(tmp) {
		tmp = getLastTwoPathParts(m.Path)
	}

	dep.name = tmp
	if len(dep.name) == 0 {
		dep.name = NameUnknown
	}

	dep.version = path.Base(m.Version)
	if len(dep.version) == 0 {
		dep.version = VersionUnknown
	}

	return dep, nil
}

func (d *Dependency) Name() string { return d.name }

func (d *Dependency) Version() string { return d.version }

func isStringAVersionPostfix(s string) (isVersionPostfix bool) {
	symbols := []rune(s)

	if len(symbols) <= 1 {
		return false
	}

	if symbols[0] != version.GolangVersionMark {
		return false
	}

	nonMarkString := string(symbols[1:])

	versionNumber, err := number.ParseUint(nonMarkString)
	if err != nil {
		return false
	}

	if versionNumber < 1 {
		return false
	}

	return true
}

func getLastTwoPathParts(p string) (ltp string) {
	var lastPart, preLastPart, tmp string
	tmp, lastPart = path.Split(p)
	tmp = strings.TrimSuffix(tmp, `\`)
	tmp = strings.TrimSuffix(tmp, `/`)
	tmp, preLastPart = path.Split(tmp)

	return path.Join(preLastPart, lastPart)
}
