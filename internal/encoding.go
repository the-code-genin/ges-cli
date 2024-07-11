package internal

import (
	"encoding/base64"
	"encoding/hex"
	"strings"
)

type EncodingFormat string

func (e EncodingFormat) String() string {
	return strings.ToLower(strings.TrimSpace(string(e)))
}

const (
	EncodingFormatHex    = EncodingFormat("hex")
	EncodingFormatBase64 = EncodingFormat("base64")
)

func EncodeBytes(format EncodingFormat, data []byte) (string, error) {
	switch format.String() {
	case EncodingFormatHex.String():
		return hex.EncodeToString(data), nil

	case EncodingFormatBase64.String():
		return base64.StdEncoding.EncodeToString(data), nil

	default:
		return "", ErrUnknownEncodingFormat
	}
}

func DecodeBytes(format EncodingFormat, data string) ([]byte, error) {
	switch format.String() {
	case EncodingFormatHex.String():
		return hex.DecodeString(data)

	case EncodingFormatBase64.String():
		return base64.StdEncoding.DecodeString(data)

	default:
		return nil, ErrUnknownEncodingFormat
	}
}
