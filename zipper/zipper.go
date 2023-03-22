package zipper

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/vault-thirteen/errorz"
)

const (
	ExtZip  = "zip"
	Ext7Zip = "7z" //TODO: Golang does not have a built-in library for this, as usual ...
)

// CompressFileAsZip creates a single-file archive of ZIP format and puts the
// specified file into it.
func CompressFileAsZip(srcFilePath string, dstFolderPath string) (dstFilePath string, err error) {
	srcFileName := filepath.Base(srcFilePath)
	dstFileName := srcFileName + "." + ExtZip
	dstFilePath = filepath.Join(dstFolderPath, dstFileName)

	var srcFile *os.File
	srcFile, err = os.Open(srcFilePath)
	if err != nil {
		return "", err
	}
	defer func() {
		derr := srcFile.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	var archive *os.File
	archive, err = os.Create(dstFilePath)
	if err != nil {
		return "", err
	}
	defer func() {
		derr := archive.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	zipWriter := zip.NewWriter(archive)
	defer func() {
		derr := zipWriter.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	var wrt io.Writer
	wrt, err = zipWriter.Create(srcFileName)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(wrt, srcFile)
	if err != nil {
		return "", err
	}

	return dstFilePath, nil
}

// UnpackZipFile unpacks a single-file archive of ZIP format.
func UnpackZipFile(zipFilePath string, dstFolderPath string) (dstFilePath string, err error) {
	zipFileName := filepath.Base(zipFilePath)
	dstFileName := strings.TrimSuffix(zipFileName, filepath.Ext(zipFileName))
	dstFilePath = filepath.Join(dstFolderPath, dstFileName)

	var archive *zip.ReadCloser
	archive, err = zip.OpenReader(zipFilePath)
	if err != nil {
		return "", err
	}
	defer func() {
		derr := archive.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	for _, af := range archive.File {
		if af.Name != dstFileName {
			continue
		}

		err = extractZipFile(af, dstFilePath)
		if err != nil {
			return "", err
		}
	}

	return dstFilePath, nil
}

// extractZipFile extracts a file from Zip archive.
// Without usage of the function deferred statements are not really possible.
func extractZipFile(af *zip.File, dstFilePath string) (err error) {
	var dstFile *os.File
	dstFile, err = os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, af.Mode())
	if err != nil {
		return err
	}
	defer func() {
		derr := dstFile.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	var fileInArchive io.ReadCloser
	fileInArchive, err = af.Open()
	if err != nil {
		return err
	}
	defer func() {
		derr := fileInArchive.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	_, err = io.Copy(dstFile, fileInArchive)
	if err != nil {
		return err
	}

	return nil
}
