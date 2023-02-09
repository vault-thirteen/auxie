// random.go.

package random

// Random Number Generator.
// This Packages offers a convenient Generator of random Integer Numbers.

import (
	crand "crypto/rand"
	"errors"
	"math"
	"math/big"
)

// Errors.
const (
	ErrLimits   = "limits Error"
	ErrOverflow = "overflow"
)

// Uint creates a new random unsigned Integer Number in the [min;max] Interval.
func Uint(
	min uint,
	max uint,
) (result uint, err error) {
	// Fool Check.
	if min >= max {
		return 0, errors.New(ErrLimits)
	}

	// Unfortunately, the 'big' Library does not accept unsigned Integer Values.
	// Check the Limits.
	var ránge uint = max - min
	if ránge > math.MaxInt64-1 {
		return 0, errors.New(ErrOverflow)
	}

	var crandMax = big.NewInt(int64(ránge) + 1)

	// Create a uniform random Value in [0; crandMax).
	var crandRandomValue *big.Int
	crandRandomValue, err = crand.Int(crand.Reader, crandMax)
	if err != nil {
		return 0, err
	}

	if !crandRandomValue.IsUint64() {
		return 0, errors.New(ErrOverflow)
	}

	var offset uint = min
	result = uint(crandRandomValue.Uint64()) + offset

	return result, nil
}

// GenerateRandomBytes generates random bytes.
func GenerateRandomBytes(
	bytesCount int,
) (bytes []byte, err error) {
	bytes = make([]byte, bytesCount)

	_, err = crand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// GenerateRandomBytesA1 generates random bytes using an alternative way.
func GenerateRandomBytesA1(
	bytesCount int,
) (bytes []byte, err error) {
	var tmp = make([]byte, bytesCount*2)

	_, err = crand.Read(tmp)
	if err != nil {
		return nil, err
	}

	bytes = make([]byte, bytesCount)

	var i, j = 0, 0
	for i < bytesCount {
		bytes[i] = tmp[j]
		i++
		j += 2
	}

	return bytes, nil
}
