package pi

import (
	"fmt"

	"github.com/vault-thirteen/auxie/VCS/common/Version"
	api "github.com/vault-thirteen/auxie/VCS/git/GitHub"
)

const (
	Platform_GitHub = "github.com"
)

const (
	ErrUnsupportedPlatform = "unsupported platform: %v"
)

func (pi *ProgramInfo) checkForUpdates() (err error) {
	switch pi.platform {
	case Platform_GitHub:
		return pi.checkForUpdatesOnGitHub()
	default:
		return fmt.Errorf(ErrUnsupportedPlatform, pi.platform)
	}
}

func (pi *ProgramInfo) checkForUpdatesOnGitHub() (err error) {
	pi.isUpdateAvailable = false

	var r *api.Repository
	r, err = api.NewRepository(pi.account, pi.project)
	if err != nil {
		return err
	}

	var versions []*version.Version
	versions, err = r.ListCleanVersions()
	if err != nil {
		return err
	}

	pi.latestVersion = version.LatestVersion(versions)
	if pi.latestVersion == nil {
		return nil
	}

	if pi.version == nil {
		pi.isUpdateAvailable = true
		return nil
	}

	var isGreater bool
	isGreater, err = pi.latestVersion.IsGreaterThan(pi.version)
	if err != nil {
		return err
	}

	if isGreater {
		pi.isUpdateAvailable = true
	}

	return nil
}
