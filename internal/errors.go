package internal

import "errors"

var (
	ErrUnequalBlockLength     = errors.New("unequal block length")
	ErrUnknownEncodingFormat  = errors.New("invalid encoding format")
	ErrInvalidKeyLength       = errors.New("invalid key length")
	ErrInvalidRound           = errors.New("invalid round")
	ErrRequiredOutputFilePath = errors.New("output file path is required")
)
