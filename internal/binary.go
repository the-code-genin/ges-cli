package internal

import (
	"crypto/rand"
	"math"
)

// Automatically pads both byte blocks so that they each have equal length.
func autoPadBytes(blockA, blockB []byte) ([]byte, []byte) {
	lengthDiff := int(math.Abs(float64(len(blockA) - len(blockB))))

	if lengthDiff != 0 {
		if len(blockA) < len(blockB) { // length of A < B
			blockA = append(blockA, make([]byte, lengthDiff)...)
		} else { // length of A > B
			blockB = append(blockB, make([]byte, lengthDiff)...)
		}
	}

	return blockA, blockB
}

// Expects blocks to be in little endian format.
// The blocks are automatically padded as required.
func XOR(blockA, blockB []byte) []byte {
	blockA, blockB = autoPadBytes(blockA, blockB)

	output := make([]byte, 0)
	for i := 0; i < len(blockA); i++ {
		output = append(output, blockA[i]^blockB[i])
	}

	return output
}

// Expects blocks to be in little endian format.
// The blocks are automatically padded as required.
func NXOR(blockA, blockB []byte) []byte {
	blockA, blockB = autoPadBytes(blockA, blockB)

	output := make([]byte, 0)
	for i := 0; i < len(blockA); i++ {
		output = append(output, (blockA[i]^blockB[i])^0xff)
	}

	return output
}

// Generate random bytes of fixed length
func RandomBytes(length uint64) ([]byte, error) {
	key := make([]byte, length)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}
