package internal

import "errors"

var (
	ErrUnequalBlockLength    = errors.New("unequal block length")
	ErrUnknownEncodingFormat = errors.New("invalid encoding format")
)
