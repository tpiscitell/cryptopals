package main

import (
	"encoding/hex"
	"fmt"
)

func ChallengeTwo(in, key string) string {
	inBytes, _ := hex.DecodeString(in)
	keyBytes, _ := hex.DecodeString(key)

	n := len(inBytes)

	outBytes := make([]byte, n)

	for i := 0; i < n; i++ {
		outBytes[i] = inBytes[i] ^ keyBytes[i]
	}

	return hex.EncodeToString(outBytes)
}

func main() {
	in := "1c0111001f010100061a024b53535009181c"
	key := "686974207468652062756c6c277320657965"

	fmt.Println(ChallengeTwo(in, key))
}
