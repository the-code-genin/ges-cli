package core

import (
	"fmt"
	"math"
)

type GESCipher struct {
	binary    Binary
	blockSize uint64
}

func (c *GESCipher) runEncryption(
	leftBlock []byte,
	rightBlock []byte,
	key []byte,
	round uint64,
) ([]byte, []byte, error) {
	if round <= 0 {
		return leftBlock, rightBlock, nil
	}

	roundFuncOutput, err := c.binary.RunXOR(rightBlock, key)
	if err != nil {
		return nil, nil, err
	}

	outputRightBlock, err := c.binary.RunXOR(leftBlock, roundFuncOutput)
	if err != nil {
		return nil, nil, err
	}

	return c.runEncryption(rightBlock, outputRightBlock, key, round-1)
}

func (c *GESCipher) encrypt(block []byte, key []byte) ([]byte, error) {
	halfBlockSize := c.blockSize / 16
	leftBlock, rightBlock, err := c.runEncryption(block[:halfBlockSize], block[halfBlockSize:], key, 2)
	if err != nil {
		return nil, err
	}

	output := make([]byte, 0)
	output = append(output, rightBlock...)
	output = append(output, leftBlock...)

	return output, nil
}

func (c *GESCipher) Encrypt(data []byte, key []byte) ([]byte, error) {
	// Validate key size
	if requiredKeySize := int(c.blockSize/2) / 8; len(key) != requiredKeySize {
		return nil, fmt.Errorf("key size must be %v bits", requiredKeySize*8)
	}

	// Pad the input data
	paddedData, err := c.binary.PadBytes(data, c.blockSize)
	if err != nil {
		return nil, err
	}

	// Run the ECB method
	output := make([]byte, 0)
	blockBytes := int(c.blockSize / 8)
	noBlocks := len(paddedData) / blockBytes
	for i := 0; i < noBlocks; i++ {
		offset := i * blockBytes
		cipherText, err := c.encrypt(data[offset:offset+blockBytes], key)
		if err != nil {
			return nil, err
		}

		output = append(output, cipherText...)
	}

	return output, nil
}

func (c *GESCipher) runDecryption(
	leftBlock []byte,
	rightBlock []byte,
	key []byte,
	round uint64,
) ([]byte, []byte, error) {
	if round <= 0 {
		return leftBlock, rightBlock, nil
	}

	roundFuncOutput, err := c.binary.RunXOR(leftBlock, key)
	if err != nil {
		return nil, nil, err
	}

	outputLeftBlock, err := c.binary.RunXOR(rightBlock, roundFuncOutput)
	if err != nil {
		return nil, nil, err
	}

	return c.runDecryption(outputLeftBlock, leftBlock, key, round-1)
}

func (c *GESCipher) decrypt(block []byte, key []byte) ([]byte, error) {
	halfBlockSize := c.blockSize / 16
	leftBlock, rightBlock, err := c.runDecryption(block[halfBlockSize:], block[:halfBlockSize], key, 2)
	if err != nil {
		return nil, err
	}

	output := make([]byte, 0)
	output = append(output, leftBlock...)
	output = append(output, rightBlock...)

	return output, nil
}

func (c *GESCipher) Decrypt(data []byte, key []byte) ([]byte, error) {
	// Validate key size
	if requiredKeySize := int(c.blockSize/2) / 8; len(key) != requiredKeySize {
		return nil, fmt.Errorf("key size must be %v bits", requiredKeySize*8)
	}

	buffer := make([]byte, 0)
	buffer = append(buffer, data...)

	// Pad the input data with null bytes if needed
	requiredBlockBytes := int(c.blockSize / 8)
	inputDataBytes := len(data)
	noBlocks := int(math.Ceil(float64(inputDataBytes) / float64(requiredBlockBytes)))
	if inputDataBytes%requiredBlockBytes != 0 {
		fillableBytes := make([]byte, (noBlocks*requiredBlockBytes)-inputDataBytes)
		buffer = append(buffer, fillableBytes...)
	}

	// Run the ECB method
	output := make([]byte, 0)
	for i := 0; i < noBlocks; i++ {
		offset := i * requiredBlockBytes
		plainText, err := c.decrypt(buffer[offset:offset+requiredBlockBytes], key)
		if err != nil {
			return nil, err
		}

		output = append(output, plainText...)
	}

	return output, nil
}

func NewGESCipher(blockSize uint64) (*GESCipher, error) {
	if blockSize%8 != 0 {
		return nil, fmt.Errorf("block size must be a multiple of 8")
	}

	return &GESCipher{Binary{}, blockSize}, nil
}
