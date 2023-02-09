// convert.go.

package bit

const BitsPerByte = 8

// ConvertBitsToBytes converts the Array of Bits into an Array of Bytes.
func ConvertBitsToBytes(
	bits []Bit,
) (result []byte, bitsCount int) {

	var b byte
	var byteBits [BitsPerByte]Bit
	var bytesCount int
	var collectorIdx int
	var collectorIdxMax int
	var idx int
	var idxMax int
	var lastByteBitIdx int
	var lastByteBits [BitsPerByte]Bit
	var lastByteIsFull bool

	bitsCount = len(bits)
	bytesCount = bitsCount / BitsPerByte
	if bitsCount%BitsPerByte == 0 {
		lastByteIsFull = true
	} else {
		bytesCount++
	}
	result = make([]byte, bytesCount)

	// If we have all the Bytes full, use a simple Algorithm.
	if lastByteIsFull {
		// Simply convert all Bits to all Bytes.
		idxMax = bytesCount - 1
		for idx = 0; idx <= idxMax; idx++ {
			// Collect next 8 Bits and convert them into a Byte.
			collectorIdx = idx * BitsPerByte
			byteBits[0] = bits[collectorIdx]
			collectorIdx++
			byteBits[1] = bits[collectorIdx]
			collectorIdx++
			byteBits[2] = bits[collectorIdx]
			collectorIdx++
			byteBits[3] = bits[collectorIdx]
			collectorIdx++
			byteBits[4] = bits[collectorIdx]
			collectorIdx++
			byteBits[5] = bits[collectorIdx]
			collectorIdx++
			byteBits[6] = bits[collectorIdx]
			collectorIdx++
			byteBits[7] = bits[collectorIdx]
			b = ConvertByteBitsToByte(byteBits)

			// Save the converted Byte into the Result.
			result[idx] = b
		}
		return
	}

	// The last Byte in not full. Use a wise Algorithm.

	// Part 1. Convert Bits to all Bytes except the last One.
	idxMax = bytesCount - 2
	for idx = 0; idx <= idxMax; idx++ {
		// Collect next 8 Bits and convert them into a Byte.
		collectorIdx = idx * BitsPerByte
		byteBits[0] = bits[collectorIdx]
		collectorIdx++
		byteBits[1] = bits[collectorIdx]
		collectorIdx++
		byteBits[2] = bits[collectorIdx]
		collectorIdx++
		byteBits[3] = bits[collectorIdx]
		collectorIdx++
		byteBits[4] = bits[collectorIdx]
		collectorIdx++
		byteBits[5] = bits[collectorIdx]
		collectorIdx++
		byteBits[6] = bits[collectorIdx]
		collectorIdx++
		byteBits[7] = bits[collectorIdx]
		b = ConvertByteBitsToByte(byteBits)

		// Save the converted Byte into the Result.
		result[idx] = b
	}

	// Part 2. Convert Bits to the last (partial) Byte.
	idx = idxMax + 1
	collectorIdx = idx * BitsPerByte
	collectorIdxMax = bitsCount - 1
	// Collect next available Bits.
	lastByteBitIdx = 0
	for collectorIdx <= collectorIdxMax {
		// Collect a Bit.
		lastByteBits[lastByteBitIdx] = bits[collectorIdx]
		lastByteBitIdx++

		// Take the next Bit.
		collectorIdx++
	}

	// Collect next available Bits and convert them into a Byte.
	b = ConvertByteBitsToByte(lastByteBits)

	// Save the converted Byte into the Result.
	result[idx] = b

	return
}

// ConvertByteBitsToByte converts the eight Bits into a Byte.
func ConvertByteBitsToByte(
	bits [BitsPerByte]Bit,
) (result byte) {

	if bits[0] == One {
		result = result | 1
	}
	if bits[1] == One {
		result = result | 2
	}
	if bits[2] == One {
		result = result | 4
	}
	if bits[3] == One {
		result = result | 8
	}
	if bits[4] == One {
		result = result | 16
	}
	if bits[5] == One {
		result = result | 32
	}
	if bits[6] == One {
		result = result | 64
	}
	if bits[7] == One {
		result = result | 128
	}

	return
}

// ConvertBytesToBits converts the Array of Bytes into an Array of Bits.
func ConvertBytesToBits(
	bytes []byte,
) (result []Bit) {

	var b byte
	var byteBits []Bit
	var bytesCount int

	bytesCount = len(bytes)
	result = make([]Bit, 0, bytesCount*BitsPerByte)

	for _, b = range bytes {
		byteBits = ConvertByteToBits(b)
		result = append(result, byteBits...)
	}

	return
}

// ConvertByteToBits converts the Byte into an Array of Bits.
func ConvertByteToBits(
	b byte,
) (result []Bit) {

	result = make([]Bit, BitsPerByte)

	// Byte #1.
	if (b & 1) == 1 {
		result[0] = One
	} else {
		result[0] = Zero
	}

	// Byte #2.
	if (b & 2) == 2 {
		result[1] = One
	} else {
		result[1] = Zero
	}

	// Byte #3.
	if (b & 4) == 4 {
		result[2] = One
	} else {
		result[2] = Zero
	}

	// Byte #4.
	if (b & 8) == 8 {
		result[3] = One
	} else {
		result[3] = Zero
	}

	// Byte #5.
	if (b & 16) == 16 {
		result[4] = One
	} else {
		result[4] = Zero
	}

	// Byte #6.
	if (b & 32) == 32 {
		result[5] = One
	} else {
		result[5] = Zero
	}

	// Byte #7.
	if (b & 64) == 64 {
		result[6] = One
	} else {
		result[6] = Zero
	}

	// Byte #8.
	if (b & 128) == 128 {
		result[7] = One
	} else {
		result[7] = Zero
	}

	return
}
