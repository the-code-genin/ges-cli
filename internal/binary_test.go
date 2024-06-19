package internal

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXOR(t *testing.T) {
	t.Run("should return the XOR result of two equal length byte blocks", func(t *testing.T) {
		byteA := byte(gofakeit.Int8())
		byteB := byte(gofakeit.Int8())
		expectedRes := byteA ^ byteB

		res := XOR([]byte{byteA}, []byte{byteB})

		assert.Equal(t, int(1), len(res))
		assert.Equal(t, expectedRes, res[0])
	})

	t.Run("should return auto-padded XOR result of two unequal length byte blocks", func(t *testing.T) {
		byteA := byte(gofakeit.Int8())
		byteB := byte(gofakeit.Int8())
		byteC := byte(gofakeit.Int8())
		expectedRes := byteA ^ byteB

		res := XOR([]byte{byteA, byteC}, []byte{byteB})

		assert.Equal(t, int(2), len(res))
		assert.Equal(t, expectedRes, res[0])
		assert.Equal(t, byteC, res[1])

		res = XOR([]byte{byteA}, []byte{byteB, byteC})

		assert.Equal(t, int(2), len(res))
		assert.Equal(t, expectedRes, res[0])
		assert.Equal(t, byteC, res[1])
	})
}
