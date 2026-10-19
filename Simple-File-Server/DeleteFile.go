package sfs

import (
	"os"

	"github.com/vault-thirteen/auxie/errors"
)

func (sfs *SimpleFileServer) DeleteFile(relPath string) (err error) {
	if !IsPathValid(relPath) {
		return errors.ErrorFromText(Err_PathIsNotValid)
	}

	if !sfs.isCachingEnabled {
		return sfs.deleteFileFromStorage(relPath)
	}

	sfs.deleteFileFromCache(relPath)

	err = sfs.deleteFileFromStorage(relPath)
	if err != nil {
		return err
	}

	return nil
}

func (sfs *SimpleFileServer) deleteFileFromCache(relPath string) {
	sfs.cache.RemoveRecord(relPath)
}

func (sfs *SimpleFileServer) deleteFileFromStorage(relPath string) (err error) {
	return os.Remove(sfs.GetAbsolutePath(relPath))
}
