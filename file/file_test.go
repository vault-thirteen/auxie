package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/tester"
)

const (
	TestFile                   = "_test-file.txt"
	TestFileThatDoesNotExist   = "_test-file-that-does-not-exist.txt"
	TestFolder                 = "_test-temporary-xyz-folder"
	TestFolderThatDoesNotExist = "_test-temporary-xyz-folder-junky-town"
)

const (
	TestFolder2          = "test"
	TestFile2            = "Test.txt"
	NonExistentTestFile2 = "NonExistentTest.txt"
)

func Test_CreateFolderSafely(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result os.FileInfo

	// Prepare the Environment.
	testFolder := filepath.Join(TestFolder, "")
	err = os.Mkdir(testFolder, 0755)
	aTest.MustBeNoError(err)
	folderX := filepath.Join(testFolder, "x")

	defer func() {
		// Clear the Environment.
		err = os.RemoveAll(testFolder)
		aTest.MustBeNoError(err)
	}()

	// Test #1. Create a Folder for the first Time.
	err = CreateFolderSafely(
		folderX,
		0755,
	)
	aTest.MustBeNoError(err)
	result, err = os.Stat(folderX)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result.IsDir(), true)

	// Test #2. Create a Folder for the second Time.
	err = CreateFolderSafely(
		folderX,
		0755,
	)
	aTest.MustBeNoError(err)
	result, err = os.Stat(folderX)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result.IsDir(), true)
}

func Test_Exists(t *testing.T) {
	var aTest = tester.New(t)
	var exists bool

	// Prepare the Environment.
	file, err := os.Create(TestFile)
	aTest.MustBeNoError(err)
	err = file.Close()
	aTest.MustBeNoError(err)

	defer func() {
		// Clear the Environment.
		err = os.Remove(TestFile)
		aTest.MustBeNoError(err)
	}()

	// Test #1. Existing file.
	exists, err = Exists(TestFile)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(exists, true)

	// Test #2. File does not exist.
	exists, err = Exists(TestFileThatDoesNotExist)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(exists, false)
}

func Test_FileExists(t *testing.T) {
	var aTest = tester.New(t)
	var exists bool

	// Prepare the Environment.
	file, err := os.Create(TestFile)
	aTest.MustBeNoError(err)
	err = file.Close()
	aTest.MustBeNoError(err)
	testFolder := filepath.Join(TestFolder, "")
	err = os.Mkdir(testFolder, 0755)
	aTest.MustBeNoError(err)

	defer func() {
		// Clear the Environment.
		err = os.Remove(TestFile)
		aTest.MustBeNoError(err)
		err = os.RemoveAll(testFolder)
		aTest.MustBeNoError(err)
	}()

	// Test #1. File exists.
	exists, err = FileExists(TestFile)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(exists, true)

	// Test #2. File does not exist.
	exists, err = FileExists(TestFileThatDoesNotExist)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(exists, false)

	// Test #3. Folder exists.
	exists, err = FileExists(TestFolder)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), "object is not a file")
	aTest.MustBeEqual(exists, false)
}

func Test_FolderExists(t *testing.T) {
	var aTest = tester.New(t)
	var exists bool

	// Prepare the Environment.
	file, err := os.Create(TestFile)
	aTest.MustBeNoError(err)
	err = file.Close()
	aTest.MustBeNoError(err)
	testFolder := filepath.Join(TestFolder, "")
	err = os.Mkdir(testFolder, 0755)
	aTest.MustBeNoError(err)

	defer func() {
		// Clear the Environment.
		err = os.Remove(TestFile)
		aTest.MustBeNoError(err)
		err = os.RemoveAll(testFolder)
		aTest.MustBeNoError(err)
	}()

	// Test #1. Folder exists.
	exists, err = FolderExists(TestFolder)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(exists, true)

	// Test #2. Folder does not exist.
	exists, err = FolderExists(TestFolderThatDoesNotExist)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(exists, false)

	// Test #3. File exists.
	exists, err = FolderExists(TestFile)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), "object is not a folder")
	aTest.MustBeEqual(exists, false)
}

func Test_GetFileContents(t *testing.T) {
	aTest := tester.New(t)
	filePath := filepath.Join(TestFolder2, TestFile2)
	var output []byte
	var err error

	// Test #1. Existing file.
	output, err = GetFileContents(filePath)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(output, []byte("Test.\r\n"))

	// Test #2. Non-existent file.
	output, err = GetFileContents(NonExistentTestFile2)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(output, ([]byte)(nil))
}
