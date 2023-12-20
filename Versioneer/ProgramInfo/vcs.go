package pi

import (
	"errors"
	"fmt"
	"strings"

	ver "github.com/vault-thirteen/auxie/VCS/common/Version"
	api "github.com/vault-thirteen/auxie/VCS/git/GitHub"
)

const (
	VcsGitHub = "github.com"
)

const (
	ErrUnsupportedVCS        = "unsupported VCS: %v"
	ErrGitHubRepositoryOwner = "GitHub repository owner is not found"
)

func (pi *ProgramInfo) parseVCSVersion() (err error) {
	pi.vcsVersion, err = ver.New(pi.version)
	if err != nil {
		return err
	}

	return nil
}

func (pi *ProgramInfo) getVCSType() (vcs string) {
	// Normally 'Main.Path' field should look like:
	// "github.com/vault-thirteen/Versioneer".
	parts := strings.Split(pi.buildInfo.Main.Path, `/`)
	if len(parts) < 1 {
		return vcs
	}

	return parts[0]
}

func (pi *ProgramInfo) checkForUpdates() (err error) {
	vcsType := pi.getVCSType()
	switch vcsType {
	case VcsGitHub:
		return pi.checkForUpdatesOnGitHub()
	default:
		return fmt.Errorf(ErrUnsupportedVCS, vcsType)
	}
}

func (pi *ProgramInfo) getGitHubRepositoryOwner() (owner, repo string, err error) {
	// Normally 'Main.Path' field should look like:
	// "github.com/vault-thirteen/Versioneer".
	parts := strings.Split(pi.buildInfo.Main.Path, `/`)
	if len(parts) < 3 {
		return owner, repo, errors.New(ErrGitHubRepositoryOwner)
	}

	return parts[1], parts[2], nil
}

func (pi *ProgramInfo) checkForUpdatesOnGitHub() (err error) {
	pi.isUpdateAvailable = false

	var owner, repo string
	owner, repo, err = pi.getGitHubRepositoryOwner()
	if err != nil {
		return err
	}

	var r *api.Repository
	r, err = api.NewRepository(owner, repo)
	if err != nil {
		return err
	}

	var versions []*ver.Version
	versions, err = r.ListCleanVersions()
	if err != nil {
		return err
	}

	pi.latestVersion = ver.LatestVersion(versions)
	if pi.latestVersion == nil {
		return nil
	}

	if pi.vcsVersion == nil {
		pi.isUpdateAvailable = true
		return nil
	}

	var isGreater bool
	isGreater, err = pi.latestVersion.IsGreaterThan(pi.vcsVersion)
	if err != nil {
		return err
	}

	if isGreater {
		pi.isUpdateAvailable = true
	}

	return nil
}
