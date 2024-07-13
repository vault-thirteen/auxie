package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"hash/crc32"
	"strings"
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

// CalculateMd5 calculates the MD-5 check sum and returns it as a
// hexadecimal text and byte array.
func CalculateMd5(data []byte) (resultAsBytes Md5Sum, resultAsText string) {
	resultAsBytes = md5.Sum(data)
	resultAsText = strings.ToUpper(hex.EncodeToString(resultAsBytes[:]))
	return resultAsBytes, resultAsText
}

// CalculateSha1 calculates the SHA-1 check sum and returns it as a hexadecimal
// text and byte array.
func CalculateSha1(data []byte) (resultAsBytes Sha1Sum, resultAsText string) {
	resultAsBytes = sha1.Sum(data)
	resultAsText = strings.ToUpper(hex.EncodeToString(resultAsBytes[:]))
	return resultAsBytes, resultAsText
}

// CalculateSha256 calculates the SHA-256 check sum and returns it as a
// hexadecimal text and byte array.
func CalculateSha256(data []byte) (resultAsBytes Sha256Sum, resultAsText string) {
	resultAsBytes = sha256.Sum256(data)
	resultAsText = strings.ToUpper(hex.EncodeToString(resultAsBytes[:]))
	return resultAsBytes, resultAsText
}
