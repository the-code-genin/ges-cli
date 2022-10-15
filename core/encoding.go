package core

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func EncodeBytes(data []byte, format string) (string, error) {
	switch format {
	case "hex":
		return hex.EncodeToString(data), nil
	case "base64":
		return base64.StdEncoding.EncodeToString(data), nil
	default:
		return "", fmt.Errorf("invalid encoding format")
	}
}

func DecodeBytes(data string, format string) (output []byte, err error) {
	switch format {
	case "hex":
		output, err = hex.DecodeString(data)
	case "base64":
		output, err = base64.StdEncoding.DecodeString(data)
	default:
		err = fmt.Errorf("invalid encoding format")
	}

	return
}
