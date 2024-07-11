package core

import (
	"fmt"
	"math"

	"github.com/the-code-genin/ges-cli/internal"
)

type GESCipher struct {
	binary    Binary
	blockSize uint64
	rounds    uint8
}

// Extract the round key from the master key
func roundKey(round uint8, masterKey []byte) ([]byte, error) {
	// Ensure master key has 128 bits
	if len(masterKey) != 16 {
		return nil, internal.ErrInvalidKeyLength
	}

	// Parity dropped key is 2 bytes lesser
	strippedKey := make([]byte, 14)

	// Bit position to insert the next bit
	bitPosition := 0

	// Perform parity dropping for every 8 bits
	for i := 0; i < 128; i++ {
		// Skip every 8-th bit
		if (i+1)%8 == 0 {
			continue
		}

		// Determine which byte and position within the byte of the masterkey to extract the bit
		oldBytePosition := i / 8
		oldBitIndex := i - (oldBytePosition * 8)

		// Extract the bit from the master key
		oldBitMask := byte(1) << (7 - oldBitIndex)
		oldTargetBit := masterKey[oldBytePosition] & oldBitMask

		// Determine which byte and position within the byte of the strippedKey to insert the bit
		newBytePosition := bitPosition / 8
		newBitIndex := bitPosition - (newBytePosition * 8)

		// Create the bit mask
		newBitMask := byte(0)
		if oldTargetBit != 0 {
			newBitMask = byte(1) << (7 - newBitIndex)
		}

		// Turn the bit on if it exists
		strippedKey[newBytePosition] |= newBitMask

		bitPosition++
	}

	// Bit shifted key is 2 bytes lesser than stripped key
	shiftedKey := make([]byte, 12)

	// Reset the bit position in the stripped key
	bitPosition = 0

	// Bit index should only go from 0 to 6
	// -1 here is an initial value
	bitIndex := -1

	// Perform bit shifting which converts every 7 bits to 6 bits
	// Shift formula: 1234567 => 132657
	for i := 0; i < 112; i++ {
		// Ensure bit index is never equal to -1 and never exceeds 6
		bitIndex++
		bitIndex %= 7

		// Skip the 4th bit
		if bitIndex == 3 {
			continue
		}

		// Determine which byte and position within the byte of the strippedKey to extract the bit
		oldBytePosition := i / 8
		oldBitIndex := i - (oldBytePosition * 8)

		// Swap the 2nd and 5th bits with their next bits
		if bitIndex == 1 || bitIndex == 4 {
			oldBitIndex++
		}

		// Swap the 3rd and 6th bits with their previous bits
		if bitIndex == 2 || bitIndex == 5 {
			oldBitIndex--
		}

		// Flows into the next byte
		if oldBitIndex > 7 {
			oldBitIndex = 0
			oldBytePosition++
		}

		// Flows into the previous byte
		if oldBitIndex < 0 {
			oldBitIndex = 7
			oldBytePosition--
		}

		// Extract the bit from the master key
		oldBitMask := byte(1) << (7 - oldBitIndex)
		oldTargetBit := strippedKey[oldBytePosition] & oldBitMask

		// Determine which byte and position within the byte of the shiftedKey to insert the bit
		newBytePosition := bitPosition / 8
		newBitIndex := bitPosition - (newBytePosition * 8)

		// Create the bit mask
		newBitMask := byte(0)
		if oldTargetBit != 0 {
			newBitMask = byte(1) << (7 - newBitIndex)
		}

		// Turn the bit on if it exists
		shiftedKey[newBytePosition] |= newBitMask

		bitPosition++
	}

	return shiftedKey, nil
}

func (c *GESCipher) runRoundFunc(block []byte, key []byte, round uint8) ([]byte, error) {
	if len(key) != int(c.rounds) {
		return nil, fmt.Errorf("keys must be 64 bits long")
	}

	roundByteIndex := c.rounds - round
	roundByte := key[roundByteIndex]

	// Create the round key by flipping every odd-indexed bit of the round byte
	roundKey := make([]byte, len(key))
	roundKey = append(roundKey, key...)
	for i := 1; i < 8; i += 2 {
		roundByte = roundByte ^ (1 << i)
	}
	roundKey[roundByteIndex] = roundByte

	return internal.XOR(block, key)
}

func (c *GESCipher) runEncryption(
	leftBlock []byte,
	rightBlock []byte,
	key []byte,
	round uint8,
) ([]byte, []byte, error) {
	if round <= 0 {
		return leftBlock, rightBlock, nil
	}

	roundFuncOutput, err := c.runRoundFunc(rightBlock, key, round)
	if err != nil {
		return nil, nil, err
	}

	outputRightBlock, err := internal.XOR(leftBlock, roundFuncOutput)
	if err != nil {
		return nil, nil, err
	}

	return c.runEncryption(rightBlock, outputRightBlock, key, round-1)
}

func (c *GESCipher) encrypt(block []byte, key []byte) ([]byte, error) {
	halfBlockSize := c.blockSize / 16

	// Do Initial scrabling
	inputLeftBlock, err := internal.NXOR(block[:halfBlockSize], key)
	if err != nil {
		return nil, err
	}

	inputRightBlock, err := internal.NXOR(block[halfBlockSize:], key)
	if err != nil {
		return nil, err
	}

	leftBlock, rightBlock, err := c.runEncryption(inputLeftBlock, inputRightBlock, key, c.rounds)
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
		cipherText, err := c.encrypt(paddedData[offset:offset+blockBytes], key)
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
	round uint8,
) ([]byte, []byte, error) {
	if round <= 0 {
		return leftBlock, rightBlock, nil
	}

	roundFuncOutput, err := c.runRoundFunc(leftBlock, key, round)
	if err != nil {
		return nil, nil, err
	}

	outputLeftBlock, err := internal.XOR(rightBlock, roundFuncOutput)
	if err != nil {
		return nil, nil, err
	}

	return c.runDecryption(outputLeftBlock, leftBlock, key, round-1)
}

func (c *GESCipher) decrypt(block []byte, key []byte) ([]byte, error) {
	halfBlockSize := c.blockSize / 16
	leftBlock, rightBlock, err := c.runDecryption(block[halfBlockSize:], block[:halfBlockSize], key, c.rounds)
	if err != nil {
		return nil, err
	}

	// Undo Initial scrabling
	outputLeftBlock, err := internal.NXOR(leftBlock, key)
	if err != nil {
		return nil, err
	}

	inputRightBlock, err := internal.NXOR(rightBlock, key)
	if err != nil {
		return nil, err
	}

	output := make([]byte, 0)
	output = append(output, outputLeftBlock...)
	output = append(output, inputRightBlock...)

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

	unpaddedOutput, err := c.binary.UnpadBytes(output)
	if err != nil {
		return nil, err
	}

	return unpaddedOutput, nil
}

func NewGESCipher() (*GESCipher, error) {
	return &GESCipher{Binary{}, 128, 8}, nil
}
