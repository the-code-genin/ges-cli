package core

import "fmt"

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
