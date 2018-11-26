package utils

import (
	"crypto/rand"
	"math/big"
)

func RandBytes(l int64) []byte {
	b := make([]byte, l)
	rand.Read(b)

	return b
}

func RandInt(n int64) int64 {
	randInt, _ := rand.Int(rand.Reader, big.NewInt(n))
	return randInt.Int64()
}
