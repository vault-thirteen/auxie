package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"strings"
)

const (
	ErrF_StreamCopyError    = "stream copy error: %v"
	ErrF_Internal           = "internal error: %v"
	Err_HashSize            = "hash size"
	ErrF_ResultSizeMismatch = "result size mismatch: %v vs %v"
)

// CalculateCrc32 calculates the CRC-32 check sum and returns it as a
// hexadecimal text and byte array.
func CalculateCrc32(data []byte) (resultAsBytes Crc32Sum, resultAsText string) {
	var buf1 = crc32.ChecksumIEEE(data)
	var buf2 = make([]byte, 4)
	binary.BigEndian.PutUint32(buf2, buf1)
	resultAsBytes = [4]byte(buf2)
	resultAsText = strings.ToUpper(hex.EncodeToString(resultAsBytes[:]))
	return resultAsBytes, resultAsText
}

// CalculateCrc32S is a stream-variant of the 'CalculateCrc32' function.
func CalculateCrc32S(stream io.Reader) (resultAsBytes Crc32Sum, resultAsText string, err error) {
	hasher := crc32.NewIEEE()

	_, err = io.Copy(hasher, stream)
	if err != nil {
		return [4]byte{}, "", fmt.Errorf(ErrF_StreamCopyError, err)
	}

	if hasher.Size() != 4 {
		return [4]byte{}, "", fmt.Errorf(ErrF_Internal, Err_HashSize)
	}

	var resultBA []byte
	resultBA = hasher.Sum(resultBA)

	if len(resultBA) != hasher.Size() {
		return [4]byte{}, "", fmt.Errorf(ErrF_ResultSizeMismatch, len(resultBA), hasher.Size())
	}

	resultAsBytes = [4]byte(resultBA)
	resultAsText = strings.ToUpper(hex.EncodeToString(resultAsBytes[:]))
	return resultAsBytes, resultAsText, nil
}

// CalculateMd5 calculates the MD-5 check sum and returns it as a
// hexadecimal text and byte array.
func CalculateMd5(data []byte) (resultAsBytes Md5Sum, resultAsText string) {
	resultAsBytes = md5.Sum(data)
	resultAsText = strings.ToUpper(hex.EncodeToString(resultAsBytes[:]))
	return resultAsBytes, resultAsText
}

// CalculateMd5S is a stream-variant of the 'CalculateMd5' function.
func CalculateMd5S(stream io.Reader) (resultAsBytes Md5Sum, resultAsText string, err error) {
	hasher := md5.New()

	_, err = io.Copy(hasher, stream)
	if err != nil {
		return [16]byte{}, "", fmt.Errorf(ErrF_StreamCopyError, err)
	}

	if hasher.Size() != 16 {
		return [16]byte{}, "", fmt.Errorf(ErrF_Internal, Err_HashSize)
	}

	var resultBA []byte
	resultBA = hasher.Sum(resultBA)

	if len(resultBA) != hasher.Size() {
		return [16]byte{}, "", fmt.Errorf(ErrF_ResultSizeMismatch, len(resultBA), hasher.Size())
	}

	resultAsBytes = [16]byte(resultBA)
	resultAsText = strings.ToUpper(hex.EncodeToString(resultAsBytes[:]))
	return resultAsBytes, resultAsText, nil
}

// CalculateSha1 calculates the SHA-1 check sum and returns it as a hexadecimal
// text and byte array.
func CalculateSha1(data []byte) (resultAsBytes Sha1Sum, resultAsText string) {
	resultAsBytes = sha1.Sum(data)
	resultAsText = strings.ToUpper(hex.EncodeToString(resultAsBytes[:]))
	return resultAsBytes, resultAsText
}

// CalculateSha1S is a stream-variant of the 'CalculateSha1' function.
func CalculateSha1S(stream io.Reader) (resultAsBytes Sha1Sum, resultAsText string, err error) {
	hasher := sha1.New()

	_, err = io.Copy(hasher, stream)
	if err != nil {
		return [20]byte{}, "", fmt.Errorf(ErrF_StreamCopyError, err)
	}

	if hasher.Size() != 20 {
		return [20]byte{}, "", fmt.Errorf(ErrF_Internal, Err_HashSize)
	}

	var resultBA []byte
	resultBA = hasher.Sum(resultBA)

	if len(resultBA) != hasher.Size() {
		return [20]byte{}, "", fmt.Errorf(ErrF_ResultSizeMismatch, len(resultBA), hasher.Size())
	}

	resultAsBytes = [20]byte(resultBA)
	resultAsText = strings.ToUpper(hex.EncodeToString(resultAsBytes[:]))
	return resultAsBytes, resultAsText, nil
}

// CalculateSha256 calculates the SHA-256 check sum and returns it as a
// hexadecimal text and byte array.
func CalculateSha256(data []byte) (resultAsBytes Sha256Sum, resultAsText string) {
	resultAsBytes = sha256.Sum256(data)
	resultAsText = strings.ToUpper(hex.EncodeToString(resultAsBytes[:]))
	return resultAsBytes, resultAsText
}

// CalculateSha256S is a stream-variant of the 'CalculateSha256' function.
func CalculateSha256S(stream io.Reader) (resultAsBytes Sha256Sum, resultAsText string, err error) {
	hasher := sha256.New()

	_, err = io.Copy(hasher, stream)
	if err != nil {
		return [32]byte{}, "", fmt.Errorf(ErrF_StreamCopyError, err)
	}

	if hasher.Size() != 32 {
		return [32]byte{}, "", fmt.Errorf(ErrF_Internal, Err_HashSize)
	}

	var resultBA []byte
	resultBA = hasher.Sum(resultBA)

	if len(resultBA) != hasher.Size() {
		return [32]byte{}, "", fmt.Errorf(ErrF_ResultSizeMismatch, len(resultBA), hasher.Size())
	}

	resultAsBytes = [32]byte(resultBA)
	resultAsText = strings.ToUpper(hex.EncodeToString(resultAsBytes[:]))
	return resultAsBytes, resultAsText, nil
}
