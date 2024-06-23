package internal

import (
	"crypto/rand"
)

// The blocks are expected to be of the same length
func XOR(blockA, blockB []byte) ([]byte, error) {
	if len(blockA) != len(blockB) {
		return nil, ErrUnequalBlockLength
	}

	output := make([]byte, 0)
	for i := 0; i < len(blockA); i++ {
		output = append(output, blockA[i]^blockB[i])
	}

	return output, nil
}

// The blocks are expected to be of the same length
func NXOR(blockA, blockB []byte) ([]byte, error) {
	if len(blockA) != len(blockB) {
		return nil, ErrUnequalBlockLength
	}

	output := make([]byte, 0)
	for i := 0; i < len(blockA); i++ {
		output = append(output, ^(blockA[i] ^ blockB[i]))
	}

	return output, nil
}

// Generate random bytes of fixed length
func RandomBytes(length uint64) ([]byte, error) {
	key := make([]byte, length)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}
