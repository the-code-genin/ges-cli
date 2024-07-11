package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// You might need to do some manual binary arthmetic to verify these values.
const (
	byteA = byte(0xAD)
	byteB = byte(0xC7)
	byteC = byte(0x98)

	byteAXORB  = byte(0x6A)
	byteANXORB = byte(0x95)
)

func TestXOR(t *testing.T) {
	t.Run("should return the XOR result of two equal length byte blocks", func(t *testing.T) {
		res, err := XOR([]byte{byteA}, []byte{byteB})

		assert.NoError(t, err)
		assert.Equal(t, int(1), len(res))
		assert.Equal(t, byteAXORB, res[0])
	})

	t.Run("should fail if byte blocks are of unequal length", func(t *testing.T) {
		res, err := XOR([]byte{byteA, byteC}, []byte{byteB})

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
