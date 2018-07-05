package main

import (
	"encoding/hex"
	"fmt"

	"github.com/tpiscitell/cryptopals/utils"
)

func ChallengeThree(in string) ([]byte, byte, float64) {
	b, _ := hex.DecodeString(in)
	return utils.CrackXOR(b)
}

func main() {
	in := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	out, _, _ := ChallengeThree(in)
	fmt.Println(string(out))
}
