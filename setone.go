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
