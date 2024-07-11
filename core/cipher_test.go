package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/the-code-genin/ges-cli/internal"
)

func Test_roundKey(t *testing.T) {
	data, err := internal.RandomBytes(16)
	assert.NoError(t, err)

	for i := uint8(0); i < 8; i++ {
		key, err := roundKey(i, data)
		assert.NoError(t, err)
		assert.Len(t, key, 8)
	}
}

func TestCipher(t *testing.T) {
	data, err := internal.RandomBytes(16)
	assert.NoError(t, err)

	key, err := internal.RandomBytes(16)
	assert.NoError(t, err)

	cipherBlock, err := Encrypt(data, key)
	assert.NoError(t, err)

	plainBlock, err := Decrypt(cipherBlock, key)
	assert.NoError(t, err)

	assert.Equal(t, data, plainBlock)
}
