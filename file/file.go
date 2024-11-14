package file

import (
	"errors"
	"io"
	"os"

	ae "github.com/vault-thirteen/auxie/errors"
)

const (
	ErrObjectIsNotFile   = "object is not a file"
	ErrObjectIsNotFolder = "object is not a folder"
)

// CreateFolderSafely tries to create a folder ignoring an error if the folder
// already exists.
func CreateFolderSafely(path string, permissions os.FileMode) (err error) {
	err = os.Mkdir(path, permissions)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	return nil
}

// Exists tries to get access to an object at the specified path – it may be a
// file or a folder.
func Exists(path string) (exists bool, err error) {
	_, err = os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	// We can not say whether it exists or not while it is not accessible.
	// So, we return an error.
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

// GetFileContents gets file's contents.
func GetFileContents(filePath string) (contents []byte, err error) {
	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		derr := f.Close()
		if derr != nil {
			err = ae.Combine(err, derr)
		}
	}()

	return io.ReadAll(f)
}

// ListFileNames lists names of all files in the folder.
// Sub-folders are not used.
func ListFileNames(folderPath string) (fileNames []string, err error) {
	var des []os.DirEntry
	des, err = os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	fileNames = []string{}
	for _, de := range des {
		if de.IsDir() {
			continue
		}

		fileNames = append(fileNames, de.Name())
	}

	return fileNames, nil
}
