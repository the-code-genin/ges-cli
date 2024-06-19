package internal

import "math"

// Expects blocks to be in little endian format.
// The blocks are automatically padded as required.
func XOR(blockA []byte, blockB []byte) []byte {
	lengthDiff := int(math.Abs(float64(len(blockA) - len(blockB))))
	if lengthDiff != 0 && len(blockA) < len(blockB) { // length of A < B
		blockA = append(blockA, make([]byte, lengthDiff)...)
	} else if lengthDiff != 0 { // length of A > B
		blockB = append(blockB, make([]byte, lengthDiff)...)
	}

	output := make([]byte, 0)
	for i := 0; i < len(blockA); i++ {
		output = append(output, blockA[i] ^ blockB[i])
	}

	return output
}