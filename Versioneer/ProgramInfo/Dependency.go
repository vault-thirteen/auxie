package pi

import (
	"errors"
	"path"
	"runtime/debug"
	"strings"

	"github.com/vault-thirteen/auxie/number"
)

const VersionMark = 'v'

// Dependency is information about program's dependency.
type Dependency struct {
	name    string
	version string
}

func NewDependency(m *debug.Module) (dep *Dependency, err error) {
	dep = &Dependency{
		name:    NameUnknown,
		version: VersionUnknown,
	}

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

	if len(tmp) == 0 {
		dep.name = NameUnknown
	} else {
		dep.name = tmp
	}

	tmp = path.Base(m.Version)
	if len(tmp) == 0 {
		dep.version = VersionUnknown
	} else {
		dep.version = tmp
	}

	return dep, nil
}

func isStringAVersionPostfix(s string) (isVersionPostfix bool) {
	symbols := []rune(s)

	if len(symbols) <= 1 {
		return false
	}

	if symbols[0] != VersionMark {
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
