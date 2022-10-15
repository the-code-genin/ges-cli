package core

import (
	"bytes"
	"testing"	
)

func FuzzGESCipher(f *testing.F) {
	cipher, err := NewGESCipher()
	if err != nil {
		f.Error(err)
	}

	// Generate a random 64-bit key
	key, err := RandomBytes(cipher.blockSize / 16)
	if err != nil {
		f.Error(err)
	}

	f.Add([]byte("Hello world"))

	tinyBlock, err := RandomBytes(16)
	if err != nil {
		f.Error(err)
	}
	f.Add(tinyBlock)

	smallBlock, err := RandomBytes(48)
	if err != nil {
		f.Error(err)
	}
	f.Add(smallBlock)

	midBlock, err := RandomBytes(256)
	if err != nil {
		f.Error(err)
	}
	f.Add(midBlock)

	largeBlock, err := RandomBytes(512)
	if err != nil {
		f.Error(err)
	}
	f.Add(largeBlock)

	superLargeBlock, err := RandomBytes(2560)
	if err != nil {
		f.Error(err)
	}
	f.Add(superLargeBlock)

	f.Fuzz(func(t *testing.T, tc []byte) {
		cipherText, err := cipher.Encrypt(tc, key)
		if err != nil {
			t.Error(err)
		}

		plainText, err := cipher.Decrypt(cipherText, key)
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(plainText, tc) {
			t.Errorf("expected %v to match %v", plainText, tc)
		}
	})
}
