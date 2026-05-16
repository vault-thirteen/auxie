package dependency

import (
	"errors"
	"path"
	"runtime/debug"
	"strings"
)

const (
	ErrDependencyInfoIsNotAvailable = "dependency info is not available"
	ErrSyntaxError                  = "syntax error"
)

const (
	DomainGithub = "github.com"
	DomainGolang = "golang.org"
)

// Dependency is information about program's dependency.
type Dependency struct {
	// Raw data.
	fullPath   string
	domain     string
	path       string
	account    string
	repository string
	postfix    string

	// Public data.
	name    string
	version string
}

func New(m *debug.Module) (dep *Dependency, err error) {
	if m == nil {
		return nil, errors.New(ErrDependencyInfoIsNotAvailable)
	}

	dep = &Dependency{
		fullPath: m.Path,
	}

	err = dep.parseModulePath()
	if err != nil {
		return nil, err
	}

	dep.version = path.Base(m.Version)

	return dep, nil
}
func (d *Dependency) parseModulePath() (err error) {
	if d == nil {
		return errors.New(ErrDependencyInfoIsNotAvailable)
	}

	// 1. Get domain & path.
	var ok bool
	d.domain, d.path, ok = strings.Cut(d.fullPath, "/")
	if !ok {
		return errors.New(ErrSyntaxError)
	}

	switch d.domain {
	case DomainGithub:
		{
			var buf string
			d.account, buf, ok = strings.Cut(d.path, "/")
			if !ok {
				return errors.New(ErrSyntaxError)
			}

			var hasPostfix bool
			d.repository, d.postfix, hasPostfix = strings.Cut(buf, "/")

			if !hasPostfix {
				d.name = d.repository
			} else {
				switch strings.ToLower(d.postfix) {
				case "src", "source":
					d.name = d.repository
				default:
					d.name = d.repository + "/" + d.postfix
				}
			}
		}
	case DomainGolang:
		{
			d.name = d.path
		}
	default:
		{
			d.name = d.fullPath
		}
	}

	return nil
}

func (d *Dependency) Name() string    { return d.name }
func (d *Dependency) Version() string { return d.version }
