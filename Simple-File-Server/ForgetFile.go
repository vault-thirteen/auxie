package sfs

import (
	"github.com/vault-thirteen/auxie/errors"
)

func (sfs *SimpleFileServer) ForgetFile(relPath string) (err error) {
	if !IsPathValid(relPath) {
		return errors.ErrorFromText(Err_PathIsNotValid)
	}

	if !sfs.isCachingEnabled {
		return errors.ErrorFromText(Err_CacheIsDisabled)
	}

	sfs.deleteFileFromCache(relPath)
	return nil
}
