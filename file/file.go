// file.go.

package file

import (
	"errors"
	"os"
)

const (
	ErrObjectIsNotFile   = "object is not a file"
	ErrObjectIsNotFolder = "object is not a folder"
)

// CreateFolderSafely tries to create a folder ignoring an error if the folder
// already exists.
func CreateFolderSafely(
	path string,
	permissions os.FileMode,
) (err error) {

	err = os.Mkdir(path, permissions)
	if err != nil {
		if os.IsExist(err) {
			err = nil
			return
		}
		return
	}

	return
}

// Exists tries to get access to an object at the specified path – it may be a
// file or a folder.
func Exists(
	path string,
) (exists bool, err error) {
	_, err = os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	// We can not say whether it exists or not while it is not accessible.
	// So, we return an Error.
	return exists, err
}

// FileExists tells if a file exists.
func FileExists(filePath string) (ok bool, err error) {
	var info os.FileInfo
	info, err = os.Stat(filePath)
	if (err == nil) && (info != nil) && (!info.IsDir()) {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	if (info != nil) && (info.IsDir()) {
		return false, errors.New(ErrObjectIsNotFile)
	}
	return false, err
}

// FolderExists tells if a folder exists.
func FolderExists(folderPath string) (ok bool, err error) {
	var info os.FileInfo
	info, err = os.Stat(folderPath)
	if (err == nil) && (info != nil) && (info.IsDir()) {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	if (info != nil) && (!info.IsDir()) {
		return false, errors.New(ErrObjectIsNotFolder)
	}
	return false, err
}
