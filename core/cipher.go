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
	} else if len(leftBlock) != len(rightBlock) {
		return nil, nil, fmt.Errorf("left and right half blocks must be of the same size")
	} else if len(key) != len(leftBlock) {
		return nil, nil, fmt.Errorf("key size must be the same size as a half block")
	}

	roundFuncOutput, err := c.runXOR(rightBlock, key)
	if err != nil {
		return nil, nil, err
	}

	outputRightBlock, err := c.runXOR(leftBlock, roundFuncOutput)
	if err != nil {
		return nil, nil, err
	}

	return c.runEncryption(rightBlock, outputRightBlock, key, round-1)
}

func (c *GESCipher) Encrypt(block []byte, key []byte) ([]byte, error) {
	blockSize := len(block)
	if blockSize % 2 != 0 {
		block = append(block, 0)
	}

	halfBlockSize := blockSize/2

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
	} else if len(leftBlock) != len(rightBlock) {
		return nil, nil, fmt.Errorf("left and right half blocks must be of the same size")
	} else if len(key) != len(leftBlock) {
		return nil, nil, fmt.Errorf("key size must be the same size as a half block")
	}

	roundFuncOutput, err := c.runXOR(leftBlock, key)
	if err != nil {
		return nil, nil, err
	}

	outputLeftBlock, err := c.runXOR(rightBlock, roundFuncOutput)
	if err != nil {
		return nil, nil, err
	}

	return c.runEncryption(outputLeftBlock, rightBlock, key, round-1)
}

func (c *GESCipher) Decrypt(block []byte, key []byte) ([]byte, error) {
	blockSize := len(block)
	if blockSize % 2 != 0 {
		return nil, fmt.Errorf("block size must be even")
	}

	halfBlockSize := blockSize/2

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
