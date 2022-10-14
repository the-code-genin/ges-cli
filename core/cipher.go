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
	} else if len(leftBlock) != 4 {
		return nil, nil, fmt.Errorf("left and right blocks must be 32 bits long")
	} else if len(leftBlock) != len(rightBlock) {
		return nil, nil, fmt.Errorf("left and right blocks must be 32 bits long")
	} else if len(key) != 4 {
		return nil, nil, fmt.Errorf("key must 32 bits long")
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
	leftBlock, rightBlock, err := c.runEncryption(block[:4], block[4:], key, 2)
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
	} else if len(leftBlock) != 4 {
		return nil, nil, fmt.Errorf("left and right blocks must be 32 bits long")
	} else if len(leftBlock) != len(rightBlock) {
		return nil, nil, fmt.Errorf("left and right blocks must be 32 bits long")
	} else if len(key) != 4 {
		return nil, nil, fmt.Errorf("key must 32 bits long")
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
	leftBlock, rightBlock, err := c.runDecryption(block[4:], block[:4], key, 2)
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
