package core

import "fmt"

type GESCipher struct{}

func (c *GESCipher) runXOR(blockA []byte, blockB []byte) ([]byte, error) {
	if len(blockA) != len(blockB) {
		return nil, fmt.Errorf("size of blocks to be XOR do not match")
	}

	output := make([]byte, len(blockA))
	for i := 0; i < len(blockA); i++ {
		output[i] = blockA[i] ^ blockB[i]
	}

	return output, nil
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
		return nil, fmt.Errorf("block size must be even")
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

	return output, nil
}

func NewGESCipher() *GESCipher {
	return &GESCipher{}
}
