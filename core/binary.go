package core

import (
	"fmt"
	"math"
)

type Bit uint8

type Binary struct{}

// Run XOR on the bits in blockA and blockB.
// Both blocks must be of the same size.
func (b Binary) RunXOR(blockA []byte, blockB []byte) ([]byte, error) {
	if len(blockA) != len(blockB) {
		return nil, fmt.Errorf("size of blocks to be XOR do not match")
	}

	output := make([]byte, len(blockA))
	for i := 0; i < len(blockA); i++ {
		output[i] = blockA[i] ^ blockB[i]
	}

	return output, nil
}

// Convert a byte into a bit array for easier manipulation
func (b Binary) ByteToBitArray(data byte) []Bit {
	output := make([]Bit, 8)
	for i := 7; i >= 0; i-- {
		if data&(1<<i) != 0 {
			output[7-i] = 1
		}
	}

	return output
}

// Convert a bit array to a byte for compactness
func (b Binary) BitArrayToByte(data []Bit) (byte, error) {
	if len(data) > 8 {
		return 0, fmt.Errorf("a byte can only be represented by 8 bits")
	}

	output := uint8(0)
	for i := 0; i < len(data); i++ {
		if data[i] > 1 {
			return 0, fmt.Errorf("invalid bit")
		} else if data[i] == 1 {
			output += uint8(math.Pow(2, float64(7-i)))
		}
	}

	return output, nil
}
