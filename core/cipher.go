package core

import "fmt"

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

func NewGESCipher(blockSize uint64) (*GESCipher, error) {
	if blockSize%8 != 0 {
		return nil, fmt.Errorf("block size must be a multiple of 8")
	}

	return &GESCipher{Binary{}, blockSize}, nil
}
