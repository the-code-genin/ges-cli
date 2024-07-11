package core

import (

	"github.com/the-code-genin/ges-cli/internal"
)

// Extract the round key from the master key
func roundKey(round uint8, masterKey []byte) ([]byte, error) {
	// There can only be 8 rounds
	if int(round) > len(sBoxes)-1 {
		return nil, internal.ErrInvalidRound
	}

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

	// Final substituted key only has 64 bits
	substitutedKey := make([]byte, 8)

	// Apply S-Box to each set of 6 bits
	// There are 16 6-bit pairs in all
	sBox := sBoxes[round]
	for set := 0; set < 16; set++ {
		// Parse the s-box input
		input := byte(0)
		for bitIndex := 0; bitIndex < 6; bitIndex++ {
			bitPosition := (set * 6) + bitIndex

			// Determine which byte and position within the byte of the shiftedKey to extract the bit
			oldBytePosition := bitPosition / 8
			oldBitIndex := bitPosition - (oldBytePosition * 8)

			// Extract the bit from the master key
			oldBitMask := byte(1) << (7 - oldBitIndex)
			oldTargetBit := shiftedKey[oldBytePosition] & oldBitMask

			// Create the bit mask
			newBitMask := byte(0)
			if oldTargetBit != 0 {
				newBitMask = byte(1) << (5 - bitIndex)
			}

			input |= newBitMask
		}

		// Parse the row
		row := byte(0)

		if input&(byte(1)<<5) != 0 {
			row |= 2
		}

		if input&byte(1) != 0 {
			row |= 1
		}

		// Parse the column
		col := byte(0)

		if input&(byte(1)<<1) != 0 {
			col |= byte(1)
		}

		if input&(byte(1)<<2) != 0 {
			col |= byte(1) << 1
		}

		if input&(byte(1)<<3) != 0 {
			col |= byte(1) << 2
		}

		if input&(byte(1)<<4) != 0 {
			col |= byte(1) << 3
		}

		// compute and store the s-box output
		output := sBox[row][col]
		for bitIndex := 0; bitIndex < 4; bitIndex++ {
			// Create a bit mask for the output
			outputBitMask := output & (byte(1) << (3 - bitIndex))

			// Determine which byte and position within the byte of the substitutedKey to insert the bit
			bitPosition := (set * 4) + bitIndex
			newBytePosition := bitPosition / 8
			newBitIndex := bitPosition - (newBytePosition * 8)

			// Set the bit on the substituted key
			newBitMask := byte(0)
			if outputBitMask != 0 {
				newBitMask = byte(1) << (7 - newBitIndex)
			}

			substitutedKey[newBytePosition] |= newBitMask
		}
	}

	return substitutedKey, nil
}

// Run the encryption round operation on two blocks
func encryptRound(leftBlock, rightBlock, masterKey []byte, round uint8) ([]byte, []byte, error) {
	if len(leftBlock) != 8 || len(rightBlock) != 8 {
		return nil, nil, internal.ErrUnequalBlockLength
	}

	key, err := roundKey(round, masterKey)
	if err != nil {
		return nil, nil, err
	}

	keyXor, err := internal.XOR(rightBlock, key)
	if err != nil {
		return nil, nil, err
	}

	blockXor, err := internal.XOR(keyXor, leftBlock)
	if err != nil {
		return nil, nil, err
	}

	return rightBlock, blockXor, nil
}

// Encrypt a block of data
func Encrypt(block, key []byte) ([]byte, error) {
	if len(block) != 16 {
		return nil, internal.ErrUnequalBlockLength
	}

	leftBlock, rightBlock := block[:8], block[8:]
	var err error

	for i := uint8(0); i < 8; i++ {
		leftBlock, rightBlock, err = encryptRound(leftBlock, rightBlock, key, i)
		if err != nil {
			return nil, err
		}
	}

	output := make([]byte, 0)
	output = append(output, rightBlock...)
	output = append(output, leftBlock...)
	return output, nil
}

// Run the decryption round operation on two blocks
func decryptRound(leftBlock, rightBlock, masterKey []byte, round uint8) ([]byte, []byte, error) {
	if len(leftBlock) != 8 || len(rightBlock) != 8 {
		return nil, nil, internal.ErrUnequalBlockLength
	}

	key, err := roundKey(round, masterKey)
	if err != nil {
		return nil, nil, err
	}

	blockXOR, err := internal.XOR(rightBlock, key)
	if err != nil {
		return nil, nil, err
	}

	keyXor, err := internal.XOR(leftBlock, blockXOR)
	if err != nil {
		return nil, nil, err
	}

	return keyXor, leftBlock, nil
}

func Decrypt(block, key []byte) ([]byte, error) {
	if len(block) != 16 {
		return nil, internal.ErrUnequalBlockLength
	}

	leftBlock, rightBlock := block[8:], block[:8]
	var err error

	for i := uint8(0); i < 8; i++ {
		leftBlock, rightBlock, err = decryptRound(leftBlock, rightBlock, key, 7-i)
		if err != nil {
			return nil, err
		}
	}

	output := make([]byte, 0)
	output = append(output, leftBlock...)
	output = append(output, rightBlock...)
	return output, nil
}
