package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
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

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println(ChallengeOne(input))
}
