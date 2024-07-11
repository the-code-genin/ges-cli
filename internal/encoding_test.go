package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	randBytes, err := RandomBytes(4)
	assert.NoError(t, err)

	t.Run("should fail if attempting to encode with unknown format", func(t *testing.T) {
		res, err := EncodeBytes(EncodingFormat("random"), randBytes)

		assert.Error(t, err)
		assert.Empty(t, res)
		assert.Equal(t, ErrUnknownEncodingFormat, err)
	})

	t.Run("should fail if attempting to decode with unknown format", func(t *testing.T) {
		res, err := DecodeBytes(EncodingFormat("random"), "randBytes")

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, ErrUnknownEncodingFormat, err)
	})

	t.Run("should encode and decode hexadecimal successfully", func(t *testing.T) {
		encodedBytes, err := EncodeBytes(EncodingFormatHex, randBytes)
		assert.NoError(t, err)

		decodedBytes, err := DecodeBytes(EncodingFormatHex, encodedBytes)
		assert.NoError(t, err)

		assert.Equal(t, decodedBytes, randBytes)
	})

	t.Run("should encode and decode base64 successfully", func(t *testing.T) {
		encodedBytes, err := EncodeBytes(EncodingFormatBase64, randBytes)
		assert.NoError(t, err)

		decodedBytes, err := DecodeBytes(EncodingFormatBase64, encodedBytes)
		assert.NoError(t, err)

		assert.Equal(t, decodedBytes, randBytes)
	})
}
