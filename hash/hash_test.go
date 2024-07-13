package hash

import (
	"encoding/hex"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_CalculateCrc32(t *testing.T) {
	const (
		Data        string = "Just a Test."
		HashSumText string = "4D87E232"
	)

	var aTest = tester.New(t)

	var (
		ba                    []byte
		data                  []byte
		err                   error
		expectedResultAsBytes Crc32Sum
		expectedResultAsText  string
		resultAsBytes         Crc32Sum
		resultAsText          string
	)

	// Test #1.
	{
		data = []byte(Data)
		expectedResultAsText = HashSumText
		ba, err = hex.DecodeString(HashSumText)
		aTest.MustBeNoError(err)
		copy(expectedResultAsBytes[:], ba)
		resultAsBytes, resultAsText = CalculateCrc32(data)
		aTest.MustBeEqual(resultAsText, expectedResultAsText)
		aTest.MustBeEqual(resultAsBytes, expectedResultAsBytes)
	}
}

func Test_CalculateMd5(t *testing.T) {
	const (
		Data        string = "Just a Test."
		HashSumText string = "E2DCA918A6E9C485830E6D969F4D7858"
	)

	var aTest = tester.New(t)

	var (
		ba                    []byte
		data                  []byte
		err                   error
		expectedResultAsBytes Md5Sum
		expectedResultAsText  string
		resultAsBytes         Md5Sum
		resultAsText          string
	)

	// Test #1.
	{
		data = []byte(Data)
		expectedResultAsText = HashSumText
		ba, err = hex.DecodeString(HashSumText)
		aTest.MustBeNoError(err)
		copy(expectedResultAsBytes[:], ba)
		resultAsBytes, resultAsText = CalculateMd5(data)
		aTest.MustBeEqual(resultAsText, expectedResultAsText)
		aTest.MustBeEqual(resultAsBytes, expectedResultAsBytes)
	}
}

func Test_CalculateSha1(t *testing.T) {
	const (
		Data        string = "Just a Test."
		HashSumText string = "7B708EF0A8EFED41F005C67546A9467BF612A145"
	)

	var aTest = tester.New(t)

	var (
		ba                    []byte
		data                  []byte
		err                   error
		expectedResultAsBytes Sha1Sum
		expectedResultAsText  string
		resultAsBytes         Sha1Sum
		resultAsText          string
	)

	// Test #1.
	{
		data = []byte(Data)
		expectedResultAsText = HashSumText
		ba, err = hex.DecodeString(HashSumText)
		aTest.MustBeNoError(err)
		copy(expectedResultAsBytes[:], ba)
		resultAsBytes, resultAsText = CalculateSha1(data)
		aTest.MustBeEqual(resultAsText, expectedResultAsText)
		aTest.MustBeEqual(resultAsBytes, expectedResultAsBytes)
	}
}

func Test_CalculateSha256(t *testing.T) {
	const (
		Data        string = "Just a Test."
		HashSumText string = "83A5404730CBE8DB7281806C1344955629AFD84E20C39A8DDAB33159798DC148"
	)

	var aTest = tester.New(t)

	var (
		ba                    []byte
		data                  []byte
		err                   error
		expectedResultAsBytes Sha256Sum
		expectedResultAsText  string
		resultAsBytes         Sha256Sum
		resultAsText          string
	)

	// Test #1.
	{
		data = []byte(Data)
		expectedResultAsText = HashSumText
		ba, err = hex.DecodeString(HashSumText)
		aTest.MustBeNoError(err)
		copy(expectedResultAsBytes[:], ba)
		resultAsBytes, resultAsText = CalculateSha256(data)
		aTest.MustBeEqual(resultAsText, expectedResultAsText)
		aTest.MustBeEqual(resultAsBytes, expectedResultAsBytes)
	}
}
