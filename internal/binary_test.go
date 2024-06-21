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
	byteCXOR0  = byteC
	byteCNXOR0 = byte(0x67)
)

func TestXOR(t *testing.T) {
	t.Run("should return the XOR result of two equal length byte blocks", func(t *testing.T) {
		res := XOR([]byte{byteA}, []byte{byteB})

		assert.Equal(t, int(1), len(res))
		assert.Equal(t, byteAXORB, res[0])
	})

	t.Run("should return auto-padded XOR result of two unequal length byte blocks", func(t *testing.T) {
		res := XOR([]byte{byteA, byteC}, []byte{byteB})

		assert.Equal(t, int(2), len(res))
		assert.Equal(t, byteAXORB, res[0])
		assert.Equal(t, byteC, res[1])

		res = XOR([]byte{byteA}, []byte{byteB, byteC})

		assert.Equal(t, int(2), len(res))
		assert.Equal(t, byteAXORB, res[0])
		assert.Equal(t, byteCXOR0, res[1])
	})
}

func TestNXOR(t *testing.T) {
	t.Run("should return the NXOR result of two equal length byte blocks", func(t *testing.T) {
		res := NXOR([]byte{byteA}, []byte{byteB})

		assert.Equal(t, int(1), len(res))
		assert.Equal(t, byteANXORB, res[0])
	})

	t.Run("should return auto-padded NXOR result of two unequal length byte blocks", func(t *testing.T) {
		res := NXOR([]byte{byteA, byteC}, []byte{byteB})

		assert.Equal(t, int(2), len(res))
		assert.Equal(t, byteANXORB, res[0])
		assert.Equal(t, byteCNXOR0, res[1])

		res = NXOR([]byte{byteA}, []byte{byteB, byteC})

		assert.Equal(t, int(2), len(res))
		assert.Equal(t, byteANXORB, res[0])
		assert.Equal(t, byteCNXOR0, res[1])
	})
}
