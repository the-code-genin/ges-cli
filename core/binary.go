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

// Pad a byte sequence to the required number of bits per block.
// Always use this before encryption
func (b Binary) PadBytes(data []byte, blockSize uint64) ([]byte, error) {
	if blockSize%8 != 0 {
		return nil, fmt.Errorf("block size must be a multiple of 8 bits")
	}

	dataByteSize := len(data)
	blockByteSize := int(blockSize / 8)
	noBlocks := int(math.Ceil(float64(dataByteSize) / float64(blockByteSize)))
	fillableBytes := make([]byte, (noBlocks*blockByteSize)-dataByteSize)
	if len(fillableBytes) > 0 {
		fillableBytes[0] = uint8(1 << 7)
	}

	output := make([]byte, 0)
	output = append(output, data...)
	output = append(output, fillableBytes...)

	return output, nil
}

// Unpad a padded byte sequence to retrieve the unpadded sequence.
// Always use this after decryption.
func (b Binary) UnpadBytes(data []byte) ([]byte, error) {
	output := make([]byte, 0)
	output = append(output, data...)

	for i := len(data) - 1; i > 0; i-- {
		lastByte := data[i]
		if lastByte == 0 { // Strip 0 padding byte
			output = data[0:i]
		} else {
			if lastByte == uint8(1<<7) { // Strip the terminator byte
				output = data[0:i]
			} else { // Strip the terminator bit
				bitArray := b.ByteToBitArray(lastByte)
				for j := len(bitArray) - 1; j > 0; j-- {
					if bitArray[j] != 0 {
						bitArray[j] = 0
						break
					}
				}

				res, err := b.BitArrayToByte(bitArray)
				if err != nil {
					return nil, err
				}

				output[i] = res
			}
			break
		}
	}

	return output, nil
}
