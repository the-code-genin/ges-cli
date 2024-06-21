package core

import (
	"bytes"
	"testing"

	"github.com/the-code-genin/ges-cli/internal"
)

func FuzzEncoding(f *testing.F) {
	for _, i := range []uint64{4, 2, 8, 12} {
		randBytes, err := internal.RandomBytes(i)
		if err != nil {
			f.Error(err)
		}

		f.Add(randBytes)
	}
	
	f.Fuzz(func(t *testing.T, a []byte) {
		for _, format := range []string{"hex", "base64"} {
			encodedText, err := EncodeBytes(a, format)
			if err != nil {
				t.Error(err)
			}

			decodedText, err := DecodeBytes(encodedText, format)
			if err != nil {
				t.Error(err)
			}

			if !bytes.Equal(decodedText, a) {
				t.Errorf("expected %v to match %v", decodedText, a)
			}
		}
	})
}
