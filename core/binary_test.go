package core

import (
	"bytes"
	"testing"
)

func FuzzBinaryXOR(f *testing.F) {
	f.Add([]byte{4, 2, 3, 8}, []byte{1, 9, 3, 1})
	f.Add([]byte{7, 2, 1, 50}, []byte{128, 19, 3, 77})

	binary := Binary{}
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		c, err := binary.RunXOR(a, b)
		if err != nil {
			t.Error(err)
		}

		d, err := binary.RunXOR(a, c)
		if err != nil {
			t.Error(err)
		} else if !bytes.Equal(b, d) {
			t.Errorf("expected %v to equal %v", d, b)
		}

		d, err = binary.RunXOR(b, c)
		if err != nil {
			t.Error(err)
		} else if !bytes.Equal(a, d) {
			t.Errorf("expected %v to equal %v", d, a)
		}
	})
}

func FuzzBitConversion(f *testing.F) {
	f.Add(byte(8))
	f.Add(byte(0))
	f.Add(byte(128))
	f.Add(byte(255))

	binary := Binary{}
	f.Fuzz(func(t *testing.T, a byte) {
		bitArray := binary.ByteToBitArray(a)
		if b, err := binary.BitArrayToByte(bitArray); err != nil {
			t.Error(err)
		} else if b != a {
			t.Errorf("expected %v to equal %v", b, a)
		}
	})
}
