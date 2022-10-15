package core

import (
	"bytes"
	"testing"	
)

func FuzzGESCipher(f *testing.F) {
	// We are working with a 64 bit cipher
	cipher, err := NewGESCipher(64)
	if err != nil {
		f.Error(err)
	}

	// Generate a random 32-bit key
	key, err := RandomBytes(cipher.blockSize / 16)
	if err != nil {
		f.Error(err)
	}

	f.Add([]byte("Hello world"))
	f.Add([]byte("foo bar"))
	f.Add([]byte{6, 2, 3})

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
