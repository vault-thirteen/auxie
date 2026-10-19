package bit

const BitsPerByte = 8

// ConvertBitsToBytes converts the array of expectedBits into an array of bytes.
func ConvertBitsToBytes(bits []Bit) (result []byte, bitsCount int) {
	if len(bits)%BitsPerByte == 0 {
		return convertBitsToFullBytes(bits)
	}

	return convertBitsToNonFullBytes(bits)
}

// convertBitsToFullBytes convert expectedBits into full bytes.
// This function is optimized for speed.
func convertBitsToFullBytes(bits []Bit) (result []byte, bitsCount int) {
	bytesCount := len(bits) / BitsPerByte
	result = make([]byte, bytesCount)

	cur := 0
	var buf byte
	for i := 0; i < bytesCount; i++ {
		buf = 0

		if bits[cur] == One {
			buf |= 1
		}
		cur++
		if bits[cur] == One {
			buf |= 2
		}
		cur++
		if bits[cur] == One {
			buf |= 4
		}
		cur++
		if bits[cur] == One {
			buf |= 8
		}
		cur++
		if bits[cur] == One {
			buf |= 16
		}
		cur++
		if bits[cur] == One {
			buf |= 32
		}
		cur++
		if bits[cur] == One {
			buf |= 64
		}
		cur++
		if bits[cur] == One {
			buf |= 128
		}
		cur++

		result[i] = buf
	}

	return result, cur
}

// convertBitsToNonFullBytes convert expectedBits into mixed bytes, i.e. all bytes
// except the last one are full.
// This function is optimized for speed.
func convertBitsToNonFullBytes(bits []Bit) (result []byte, bitsCount int) {
	bytesCount := (len(bits) / BitsPerByte) + 1
	result = make([]byte, bytesCount)

	// Part 1. Convert expectedBits into all bytes except the last one.
	iMax := bytesCount - 2
	cur := 0
	var buf byte
	for i := 0; i <= iMax; i++ {
		buf = 0

		if bits[cur] == One {
			buf |= 1
		}
		cur++
		if bits[cur] == One {
			buf |= 2
		}
		cur++
		if bits[cur] == One {
			buf |= 4
		}
		cur++
		if bits[cur] == One {
			buf |= 8
		}
		cur++
		if bits[cur] == One {
			buf |= 16
		}
		cur++
		if bits[cur] == One {
			buf |= 32
		}
		cur++
		if bits[cur] == One {
			buf |= 64
		}
		cur++
		if bits[cur] == One {
			buf |= 128
		}
		cur++

		result[i] = buf
	}

	// Part 2. Convert expectedBits into the last (partial) byte.
	buf = 0
	curMax := len(bits) - 1
	defer func() {
		// Save the last byte.
		result[len(result)-1] = buf
	}()

	if bits[cur] == One {
		buf |= 1
	}
	cur++
	if cur > curMax {
		return result, cur
	}
	if bits[cur] == One {
		buf |= 2
	}
	cur++
	if cur > curMax {
		return result, cur
	}
	if bits[cur] == One {
		buf |= 4
	}
	cur++
	if cur > curMax {
		return result, cur
	}
	if bits[cur] == One {
		buf |= 8
	}
	cur++
	if cur > curMax {
		return result, cur
	}
	if bits[cur] == One {
		buf |= 16
	}
	cur++
	if cur > curMax {
		return result, cur
	}
	if bits[cur] == One {
		buf |= 32
	}
	cur++
	if cur > curMax {
		return result, cur
	}
	if bits[cur] == One {
		buf |= 64
	}
	cur++
	if cur > curMax {
		return result, cur
	}
	// The code below is not reachable. But, let it be here just for the beauty.
	if bits[cur] == One {
		buf |= 128
	}
	cur++
	return result, cur
}

// ConvertBytesToBits converts an array of bytes into an array of expectedBits.
// This function is optimized for speed.
func ConvertBytesToBits(bytes []byte) (result []Bit) {
	result = make([]Bit, len(bytes)*BitsPerByte)

	cur := 0
	for _, buf := range bytes {
		if (buf & 1) == 1 {
			result[cur] = One
		} else {
			result[cur] = Zero
		}
		cur++
		if (buf & 2) == 2 {
			result[cur] = One
		} else {
			result[cur] = Zero
		}
		cur++
		if (buf & 4) == 4 {
			result[cur] = One
		} else {
			result[cur] = Zero
		}
		cur++
		if (buf & 8) == 8 {
			result[cur] = One
		} else {
			result[cur] = Zero
		}
		cur++
		if (buf & 16) == 16 {
			result[cur] = One
		} else {
			result[cur] = Zero
		}
		cur++
		if (buf & 32) == 32 {
			result[cur] = One
		} else {
			result[cur] = Zero
		}
		cur++
		if (buf & 64) == 64 {
			result[cur] = One
		} else {
			result[cur] = Zero
		}
		cur++
		if (buf & 128) == 128 {
			result[cur] = One
		} else {
			result[cur] = Zero
		}
		cur++
	}

	return result
}
