package core

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate random bytes of fixed length
func RandomBytes(length uint64) ([]byte, error) {
	key := make([]byte, length)

	rand.Seed(time.Now().UTC().UnixMilli())
	i, err := rand.Read(key)
	if err != nil {
		return nil, err
	} else if i != int(length) {
		return nil, fmt.Errorf("expecting %v bytes to be generated for random key", i)
	}

	return key, nil
}