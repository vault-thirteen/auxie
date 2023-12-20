package zipper

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/file"
	"github.com/vault-thirteen/auxie/tester"
)

const (
	TestFolder          = "test"
	TestFile            = "Test.txt"
	NonExistentTestFile = "NonExistentTest.txt"
	TempFolder          = "tmp"
)

func Test_CompressFileAsZip(t *testing.T) {
	aTest := tester.New(t)
	srcFilePath := filepath.Join(TestFolder, TestFile)
	dstFolderPath := TestFolder
	dstFileName := TestFile + "." + ExtZip
	dstFilePath := filepath.Join(dstFolderPath, dstFileName)
	var output string
	var err error

	// Test #1. Existing file.
	output, err = CompressFileAsZip(srcFilePath, dstFolderPath)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(output, dstFilePath)
	// Clean the folder.
	log.Printf("deleting a temporary file: %v.\r\n", dstFilePath)
	err = os.Remove(dstFilePath)
	aTest.MustBeNoError(err)

	// Test #2. Non-existent file.
	output, err = CompressFileAsZip(NonExistentTestFile, dstFolderPath)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(output, "")
}

func Test_UnpackZipFile(t *testing.T) {
	aTest := tester.New(t)
	dstFolderPath := filepath.Join(TestFolder, TempFolder)

	// Preparations.
	var zipFilePath string
	var err error
	zipFilePath, err = _test_prepareData1()
	aTest.MustBeNoError(err)

	// Test.
	var dstFilePath string
	dstFilePath, err = UnpackZipFile(zipFilePath, dstFolderPath)
	aTest.MustBeNoError(err)
	dstFilePathExpected := filepath.Join(TestFolder, TempFolder, TestFile)
	aTest.MustBeEqual(dstFilePath, dstFilePathExpected)

	var contents []byte
	contents, err = file.GetFileContents(dstFilePath)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(contents, []byte("Test.\r\n"))

	// Clean the folder.
	log.Printf("deleting a temporary folder: %v.\r\n", dstFolderPath)
	err = os.RemoveAll(dstFolderPath)
	aTest.MustBeNoError(err)
}

func _test_prepareData1() (zipFilePath string, err error) {
	srcFilePath := filepath.Join(TestFolder, TestFile)
	tmpFolderPath := filepath.Join(TestFolder, TempFolder)

	err = os.MkdirAll(tmpFolderPath, 0777)
	if err != nil {
		return "", err
	}

	return CompressFileAsZip(srcFilePath, tmpFolderPath)
}
