package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
)

func ChallengeOne(in string) string {
	var out bytes.Buffer

	b, err := hex.DecodeString(in)
	if err != nil {
		println("Error decoding input: ", err)
	}

	encoder := base64.NewEncoder(base64.StdEncoding, &out)

	_, err = encoder.Write(b)
	if err != nil {
		println("Error with base64 encoding: ", err)
	}

	return out.String()
}

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
